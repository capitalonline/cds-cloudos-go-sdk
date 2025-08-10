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

package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
	"github.com/capitalonline/cds-cloudos-go-sdk/util"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	version          = "2019-08-08"
	signatureVersion = "1.0"
	signatureMethod  = "HMAC-SHA1"
	timeStampFormat  = "2006-01-02T15:04:05Z"
)

// Signer abstracts the entity that implements the `Sign` method
type Signer interface {
	// Sign the given Request with the Credentials and SignOptions
	Sign(*http.Request, *CdsCredentials)
}

// SignOptions defines the data structure used by Signer
type SignOptions struct {
	HeadersToSign map[string]struct{}
	Timestamp     int64
	ExpireSeconds int
}

func (opt *SignOptions) String() string {
	return fmt.Sprintf(`SignOptions [
        HeadersToSign=%s;
        Timestamp=%d;
        ExpireSeconds=%d
    ]`, opt.HeadersToSign, opt.Timestamp, opt.ExpireSeconds)
}

// CdsSigner implements the sign algorithm
type CdsSigner struct{}

func (b *CdsSigner) Sign(req *http.Request, cred *CdsCredentials) {
	if req == nil {
		log.Fatal("request should not be null for sign")
		return
	}
	if cred == nil {
		log.Fatal("credentials should not be null for sign")
		return
	}

	// Prepare parameters
	accessKeyId := cred.AccessKeyId
	secretAccessKey := cred.SecretAccessKey

	urlParams := map[string]string{
		"Action":           req.Action(),
		"AccessKeyId":      accessKeyId,
		"SignatureMethod":  signatureMethod,
		"SignatureNonce":   util.NewUUID(),
		"SignatureVersion": signatureVersion,
		"Timestamp":        time.Now().UTC().Format(timeStampFormat),
		"Version":          version,
	}

	if req.Params() != nil {
		for k, v := range req.Params() {
			urlParams[k] = v
		}
	}

	var paramSortKeys sort.StringSlice
	for k, _ := range urlParams {
		paramSortKeys = append(paramSortKeys, k)
	}
	sort.Sort(paramSortKeys)

	var urlStr string
	for _, k := range paramSortKeys {
		urlStr += "&" + percentEncode(k) + "=" + percentEncode(urlParams[k])
	}
	urlStr = req.Method() + "&%2F&" + percentEncode(urlStr[1:])

	// Generate signature
	h := hmac.New(sha1.New, []byte(secretAccessKey))
	h.Write([]byte(urlStr))
	signStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	urlParams["Signature"] = signStr

	// generate http request URL.RawQuery
	urlVal := url.Values{}
	for k, v := range urlParams {
		urlVal.Add(k, v)
	}
	urlValStr := urlVal.Encode()

	req.SetQueryString(urlValStr)
}

func percentEncode(str string) string {
	str = url.QueryEscape(str)
	strings.Replace(str, "+", "%20", -1)
	strings.Replace(str, "*", "%2A", -1)
	strings.Replace(str, "%7E", "~", -1)
	return str
}
