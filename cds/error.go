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

// error.go - define the error types for CDS

package cds

const (
	EACCESS_DENIED            = "AccessDenied"
	EINAPPROPRIATE_JSON       = "InappropriateJSON"
	EINTERNAL_ERROR           = "InternalError"
	EINVALID_ACCESS_KEY_ID    = "InvalidAccessKeyId"
	EINVALID_HTTP_AUTH_HEADER = "InvalidHTTPAuthHeader"
	EINVALID_HTTP_REQUEST     = "InvalidHTTPRequest"
	EINVALID_URI              = "InvalidURI"
	EMALFORMED_JSON           = "MalformedJSON"
	EINVALID_VERSION          = "InvalidVersion"
	EOPT_IN_REQUIRED          = "OptInRequired"
	EPRECONDITION_FAILED      = "PreconditionFailed"
	EREQUEST_EXPIRED          = "RequestExpired"
	ESIGNATURE_DOES_NOT_MATCH = "SignatureDoesNotMatch"
)

// CdsError abstracts the error for CDS
type CdsError interface {
	error
}

// CdsClientError defines the error struct for the client when making request
type CdsClientError struct{ Message string }

func (b *CdsClientError) Error() string { return b.Message }

func NewCdsClientError(msg string) *CdsClientError { return &CdsClientError{msg} }

// CdsServiceError defines the error struct for the CDS service when receiving response
type CdsServiceError struct {
	Code       string `json:"Code"`
	Message    string `json:"Message"`
	Msg        string `json:"Msg"` // EKS service uses "Msg" instead of "Message"
	RequestId  string `json:"RequestId"`
	StatusCode int    `json:"-"`
}

func (b *CdsServiceError) Error() string {
	ret := "[Code: " + b.Code
	// Prefer Msg over Message for EKS service compatibility
	if b.Msg != "" {
		ret += "; Message: " + b.Msg
	} else {
		ret += "; Message: " + b.Message
	}
	ret += "; RequestId: " + b.RequestId + "]"
	return ret
}

func NewCdsServiceError(code, msg, reqId string, status int) *CdsServiceError {
	return &CdsServiceError{
		Code:       code,
		Message:    msg, // For backward compatibility
		Msg:        "",  // Will be populated when parsing JSON
		RequestId:  reqId,
		StatusCode: status,
	}
}
