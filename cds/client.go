/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cds

import (
	"bytes"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/auth"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
	"io"
	"io/ioutil"
	"time"
)

// Client is the general interface which can perform sending request. Different service
// will define its own client in case of specific extension.
type Client interface {
	SendRequest(*CdsRequest, *CdsResponse) error
	GetCdsClientConfig() *CdsClientConfiguration
}

// CdsClient defines the general client to access the CDS services.
type CdsClient struct {
	Config *CdsClientConfiguration
	Signer auth.Signer // the sign algorithm
}

// BuildHttpRequest - the helper method for the client to build http request
//
// PARAMS:
//   - request: the input request object to be built
func (c *CdsClient) buildHttpRequest(request *CdsRequest) {
	// Construct the http request instance for the special fields
	request.BuildHttpRequest()

	// Set the client specific configurations
	if request.Endpoint() == "" {
		request.SetEndpoint(c.Config.Endpoint)
	}
	if request.Protocol() == "" {
		request.SetProtocol(DEFAULT_PROTOCOL)
	}
	if len(c.Config.ProxyUrl) != 0 {
		request.SetProxyUrl(c.Config.ProxyUrl)
	}
	request.SetTimeout(c.Config.ConnectionTimeoutInMillis / 1000)

	// Set the CDS request headers
	request.SetHeader(http.HOST, request.Host())
	request.SetHeader(http.USER_AGENT, c.Config.UserAgent)

	//set default content-type if null
	if request.Header(http.CONTENT_TYPE) == "" {
		request.SetHeader(http.CONTENT_TYPE, DEFAULT_CONTENT_TYPE)
	}

	// Generate the auth string if needed
	if c.Config.Credentials != nil {
		c.Signer.Sign(&request.Request, c.Config.Credentials, c.Config.SignOption)
	}
}

// SendRequest - the client performs sending the http request with retry policy and receive the
// response from the CDS services.
//
// PARAMS:
//   - req: the request object to be sent to the CDS service
//   - resp: the response object to receive the content from CDS service
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *CdsClient) SendRequest(req *CdsRequest, resp *CdsResponse) error {
	// Return client error if it is not nil
	if req.ClientError() != nil {
		return req.ClientError()
	}

	// Build the http request and prepare to send
	c.buildHttpRequest(req)
	// log.Infof("send http request: %v", req)

	// Send request with the given retry policy
	retries := 0
	if req.Body() != nil {
		defer req.Body().Close() // Manually close the ReadCloser body for retry
	}
	for {
		// The request body should be temporarily saved if retry to send the http request
		var retryBuf bytes.Buffer
		var teeReader io.Reader
		if c.Config.Retry.ShouldRetry(nil, 0) && req.Body() != nil {
			teeReader = io.TeeReader(req.Body(), &retryBuf)
			req.Request.SetBody(ioutil.NopCloser(teeReader))
		}
		httpResp, err := http.Execute(&req.Request)

		if err != nil {
			if c.Config.Retry.ShouldRetry(err, retries) {
				delayInMills := c.Config.Retry.GetDelayBeforeNextRetryInMillis(err, retries)
				time.Sleep(delayInMills)
			} else {
				return &CdsClientError{
					fmt.Sprintf("execute http request failed! Retried %d times, error: %v",
						retries, err)}
			}
			retries++
			// todo unImplement log pkg
			// log.Warnf("send request failed: %v, retry for %d time(s)", err, retries)
			if req.Body() != nil {
				ioutil.ReadAll(teeReader)
				req.Request.SetBody(ioutil.NopCloser(&retryBuf))
			}
			continue
		}
		resp.SetHttpResponse(httpResp)
		resp.ParseResponse()

		//log.Infof("receive http response: status: %s, debugId: %s, requestId: %s, elapsed: %v",
		//	resp.StatusText(), resp.DebugId(), resp.RequestId(), resp.ElapsedTime())
		//
		//if resp.ElapsedTime().Milliseconds() > DEFAULT_WARN_LOG_TIMEOUT_IN_MILLS {
		//	log.Warnf("request time more than 5 second, debugId: %s, requestId: %s, elapsed: %v",
		//		resp.DebugId(), resp.RequestId(), resp.ElapsedTime())
		//}
		//for k, v := range resp.Headers() {
		//	log.Debugf("%s=%s", k, v)
		//}

		if resp.IsFail() {
			err := resp.ServiceError()
			if c.Config.Retry.ShouldRetry(err, retries) {
				delayInMills := c.Config.Retry.GetDelayBeforeNextRetryInMillis(err, retries)
				time.Sleep(delayInMills)
			} else {
				return err
			}
			retries++
			// log.Warnf("send request failed, retry for %d time(s)", retries)
			if req.Body() != nil {
				ioutil.ReadAll(teeReader)
				req.Request.SetBody(ioutil.NopCloser(&retryBuf))
			}
			continue
		}
		return nil
	}
}

func (c *CdsClient) GetCdsClientConfig() *CdsClientConfiguration {
	return c.Config
}

func NewCdsClient(conf *CdsClientConfiguration, sign auth.Signer) *CdsClient {
	clientConfig := http.ClientConfig{
		RedirectDisabled:  conf.RedirectDisabled,
		DisableKeepAlives: conf.DisableKeepAlives,
	}
	http.InitClient(clientConfig)
	return &CdsClient{conf, sign}
}

func NewCdsClientWithAkSk(ak, sk, endPoint string) (*CdsClient, error) {
	credentials, err := auth.NewCdsCredentials(ak, sk)
	if err != nil {
		return nil, err
	}

	defaultConf := &CdsClientConfiguration{
		Endpoint:                  endPoint,
		Credentials:               credentials,
		Retry:                     DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
		RedirectDisabled:          false,
	}
	v1Signer := &auth.CdsSigner{}

	return NewCdsClient(defaultConf, v1Signer), nil
}
