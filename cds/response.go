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
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

type CdsResponse struct {
	statusCode   int
	statusText   string
	requestId    string
	debugId      string
	response     *http.Response
	serviceError *CdsServiceError
}

func (r *CdsResponse) IsFail() bool {
	return r.response.StatusCode() >= 400
}

func (r *CdsResponse) StatusCode() int {
	return r.statusCode
}

func (r *CdsResponse) StatusText() string {
	return r.statusText
}

func (r *CdsResponse) RequestId() string {
	return r.requestId
}

func (r *CdsResponse) DebugId() string {
	return r.debugId
}

func (r *CdsResponse) Header(key string) string {
	return r.response.GetHeader(key)
}

func (r *CdsResponse) Headers() map[string]string {
	return r.response.GetHeaders()
}

func (r *CdsResponse) Body() io.ReadCloser {
	return r.response.Body()
}

func (r *CdsResponse) SetHttpResponse(response *http.Response) {
	r.response = response
}

func (r *CdsResponse) ElapsedTime() time.Duration {
	return r.response.ElapsedTime()
}

func (r *CdsResponse) ServiceError() *CdsServiceError {
	return r.serviceError
}

func (r *CdsResponse) ParseResponse() {
	r.statusCode = r.response.StatusCode()
	r.statusText = r.response.StatusText()
	r.requestId = r.response.GetHeader(http.CDS_REQUEST_ID)
	if r.IsFail() {
		r.serviceError = NewCdsServiceError("", r.statusText, r.requestId, r.statusCode)

		// First try to read the error `Code' and `Message' from body
		rawBody, _ := ioutil.ReadAll(r.Body())
		defer r.Body().Close()
		if len(rawBody) != 0 {
			jsonDecoder := json.NewDecoder(bytes.NewBuffer(rawBody))
			if err := jsonDecoder.Decode(r.serviceError); err != nil {
				r.serviceError = NewCdsServiceError(
					EMALFORMED_JSON,
					"Service json error message decode failed",
					r.requestId,
					r.statusCode)
			}
			return
		}

		// Then guess the `Message' from by the return status code
		switch r.statusCode {
		case 400:
			r.serviceError.Code = EINVALID_HTTP_REQUEST
		case 403:
			r.serviceError.Code = EACCESS_DENIED
		case 412:
			r.serviceError.Code = EPRECONDITION_FAILED
		case 500:
			r.serviceError.Code = EINTERNAL_ERROR
		default:
			words := strings.Split(r.statusText, " ")
			r.serviceError.Code = strings.Join(words[1:], "")
		}
	}
}

func (r *CdsResponse) ParseJsonBody(result interface{}) error {
	defer r.Body().Close()
	jsonDecoder := json.NewDecoder(r.Body())

	// debug
	//rawBody, _ := io.ReadAll(r.Body())
	//fmt.Printf(">>> httpResponse: %s\n", string(rawBody))

	return jsonDecoder.Decode(result)
}
