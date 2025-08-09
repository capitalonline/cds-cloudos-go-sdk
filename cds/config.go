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
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/auth"
	"reflect"
	"time"
)

// Constants and default values for the package cds
const (
	DEFAULT_DOMAIN                       = "cdsapi.capitalonline.net"
	DEFAULT_PROTOCOL                     = "http"
	HTTPS_PROTOCAL                       = "https"
	DEFAULT_CONTENT_TYPE                 = "application/json;charset=utf-8"
	DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS = 1200 * 1000
	DEFAULT_WARN_LOG_TIMEOUT_IN_MILLS    = 5 * 1000
)

var (
	DEFAULT_RETRY_POLICY = NewBackOffRetryPolicy(3, 20000, 300)
)

// CdsClientConfiguration defines the config components structure.
type CdsClientConfiguration struct {
	Endpoint                  string
	ProxyUrl                  string
	Region                    string
	UserAgent                 string
	Credentials               *auth.CdsCredentials
	SignOption                *auth.SignOptions
	Retry                     RetryPolicy
	ConnectionTimeoutInMillis int
	// CnameEnabled should be true when use custom domain as endpoint to visit bos resource
	CnameEnabled          bool
	BackupEndpoint        string
	RedirectDisabled      bool
	DisableKeepAlives     bool
	DialTimeout           *time.Duration // timeout of building a connection
	KeepAlive             *time.Duration // the interval between keep-alive probes for an active connection
	ReadTimeout           *time.Duration // read timeout of net.Conn
	WriteTimeOut          *time.Duration // write timeout of net.Conn
	TLSHandshakeTimeout   *time.Duration // http.Transport.TLSHandshakeTimeout
	IdleConnectionTimeout *time.Duration // http.Transport.IdleConnTimeout
	ResponseHeaderTimeout *time.Duration // http.Transport.ResponseHeaderTimeout
	HTTPClientTimeout     *time.Duration // http.Client.Timeout
}

func (c *CdsClientConfiguration) String() string {
	return fmt.Sprintf(`CdsClientConfiguration [
        Endpoint=%s;
        ProxyUrl=%s;
        Region=%s;
        UserAgent=%s;
        Credentials=%v;
        SignOption=%v;
        RetryPolicy=%v;
        ConnectionTimeoutInMillis=%v;
		RedirectDisabled=%v
    ]`, c.Endpoint, c.ProxyUrl, c.Region, c.UserAgent, c.Credentials,
		c.SignOption, reflect.TypeOf(c.Retry).Name(), c.ConnectionTimeoutInMillis, c.RedirectDisabled)
}
