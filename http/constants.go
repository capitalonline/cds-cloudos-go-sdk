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

// constants.go - defines constants of the CDS http package including headers and methods

package http

// Constants of the supported HTTP methods for CDS
const (
	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
)

// Constants of the HTTP headers for CDS
const (
	// Standard HTTP Headers
	AUTHORIZATION       = "Authorization"
	CACHE_CONTROL       = "Cache-Control"
	CONTENT_DISPOSITION = "Content-Disposition"
	CONTENT_ENCODING    = "Content-Encoding"
	CONTENT_LANGUAGE    = "Content-Language"
	CONTENT_LENGTH      = "Content-Length"
	CONTENT_MD5         = "Content-Md5"
	CONTENT_RANGE       = "Content-Range"
	CONTENT_TYPE        = "Content-Type"
	DATE                = "Date"
	ETAG                = "Etag"
	EXPIRES             = "Expires"
	HOST                = "Host"
	LAST_MODIFIED       = "Last-Modified"
	LOCATION            = "Location"
	RANGE               = "Range"
	SERVER              = "Server"
	TRANSFER_ENCODING   = "Transfer-Encoding"
	USER_AGENT          = "User-Agent"

	// CDS Common HTTP Heahder
	CDS_REQUEST_ID = "cds-request-id"
)
