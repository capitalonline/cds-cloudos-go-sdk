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

// credentials.go - the credentials data structure definition

// Package auth implements the authorization functionality for CDS.
// It use the CDS access key ID and secret access key with the specific sign algorithm to generate the authorization string.

package auth

import (
	"errors"
	"os"
)

const (
	AK = "CDS_SECRET_ID"
	SK = "CDS_SECRET_KEY"
)

type CdsCredentials struct {
	AccessKeyId     string // access key id to the service
	SecretAccessKey string // secret access key to the service
}

func (c *CdsCredentials) String() string {
	str := "ak: " + c.AccessKeyId + ", sk: " + c.SecretAccessKey
	return str
}

func NewCdsCredentials(ak, sk string) (*CdsCredentials, error) {
	if len(ak) == 0 {
		return nil, errors.New("accessKeyId should not be empty")
	}
	if len(sk) == 0 {
		return nil, errors.New("secretKey should not be empty")
	}

	return &CdsCredentials{ak, sk}, nil
}

func NewCdsCredentialsByEnv() (*CdsCredentials, error) {
	ak := os.Getenv(AK)
	sk := os.Getenv(SK)
	if len(ak) == 0 {
		return nil, errors.New("accessKeyId should not be empty")
	}
	if len(sk) == 0 {
		return nil, errors.New("secretKey should not be empty")
	}

	return &CdsCredentials{ak, sk}, nil
}
