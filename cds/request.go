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
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
	"github.com/capitalonline/cds-cloudos-go-sdk/util"
	"hash"
	"io"
	"io/ioutil"
)

// Body defines the data structure used in CDS request.
// Every CDS request that sets the body field must set its content-length and content-md5 headers
// to ensure the correctness of the body content force, and users can also set the content-sha256
// header to strengthen the correctness with the "SetHeader" method.
type Body struct {
	stream     io.ReadCloser
	writer     io.Writer
	size       int64
	contentMD5 string
}

func (b *Body) Stream() io.ReadCloser { return b.stream }

func (b *Body) SetStream(stream io.ReadCloser) { b.stream = stream }

func (b *Body) Writer() io.Writer { return b.writer }

func (b *Body) SetWriter(w io.Writer) { b.writer = w }

func (b *Body) Size() int64 { return b.size }

func (b *Body) SetSize(size int64) { b.size = size }

func (b *Body) ContentMD5() string { return b.contentMD5 }

func (b *Body) SetContentMD5(md5Str string) { b.contentMD5 = md5Str }

func (b *Body) Read(p []byte) (int, error) {
	n, err := b.stream.Read(p)
	if n > 0 && b.writer != nil {
		if n, err := b.writer.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return n, err
}

func (b *Body) Close() error {
	if rclose, ok := b.stream.(io.ReadCloser); ok {
		rclose.Close()
	}
	if b.writer != nil {
		if wclose, ok := b.writer.(io.WriteCloser); ok {
			wclose.Close()
		}
	}
	return nil
}

func (b *Body) Crc32() uint32 {
	if b.writer != nil {
		if hc32, ok := b.writer.(hash.Hash32); ok {
			return hc32.Sum32()
		}
	}
	return 0
}

func NewBodyFromBytes(stream []byte) (*Body, error) {
	buf := bytes.NewBuffer(stream)
	size := int64(buf.Len())
	contentMD5, err := util.CalculateContentMD5(buf, size)
	if err != nil {
		return nil, err
	}
	buf = bytes.NewBuffer(stream)
	return &Body{stream: ioutil.NopCloser(buf), size: size, contentMD5: contentMD5}, nil
}

// CdsRequest defines the request structure for accessing CDS services
type CdsRequest struct {
	http.Request
	requestId   string
	clientError *CdsClientError
}

func (b *CdsRequest) RequestId() string { return b.requestId }

func (b *CdsRequest) SetRequestId(val string) { b.requestId = val }

func (b *CdsRequest) ClientError() *CdsClientError { return b.clientError }

func (b *CdsRequest) SetClientError(err *CdsClientError) { b.clientError = err }

func (b *CdsRequest) SetBody(body *Body) { // override SetBody derived from http.Request
	b.Request.SetBody(body)
	b.SetLength(body.Size()) // set field of "net/http.Request.ContentLength"
	if body.ContentMD5() != "" {
		b.SetHeader(http.CONTENT_MD5, body.ContentMD5())
	}
	if body.Size() > 0 {
		b.SetHeader(http.CONTENT_LENGTH, fmt.Sprintf("%d", body.Size()))
	}
}

func (b *CdsRequest) BuildHttpRequest() {
	// Only need to build the specific `requestId` field for CDS, other fields are same as the
	// `http.Request` as well as its methods.
	if len(b.requestId) == 0 {
		// Construct the request ID with UUID
		b.requestId = util.NewRequestId()
	}
	b.SetHeader(http.CDS_REQUEST_ID, b.requestId)
}

func (b *CdsRequest) String() string {
	requestIdStr := "requestId=" + b.requestId
	if b.clientError != nil {
		return requestIdStr + ", client error: " + b.ClientError().Error()
	}
	return requestIdStr + "\n" + b.Request.String()
}
