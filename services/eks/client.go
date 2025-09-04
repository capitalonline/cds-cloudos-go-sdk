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

package eks

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"os"
)

const (
	// 不同产品线 endpoint可能会不一致
	eksEndpoint     = "https://api.capitalonline.net"
	eksURI          = "/eks/v1"
	EnvEKSAPIHost   = "EKS_API_HOST"
	EnvEKSAPISchema = "EKS_API_SCHEMA"
)

// Client of EKS service is a kind of CdsClient, so derived from CdsClient
type Client struct {
	*cds.CdsClient
}

func NewClient(ak, sk string) (*Client, error) {
	var endpoint string

	host := os.Getenv(EnvEKSAPIHost)
	schema := os.Getenv(EnvEKSAPISchema)

	if host == "" || schema == "" {
		endpoint = eksEndpoint
	} else {
		endpoint = schema + "://" + host
	}

	client, err := cds.NewCdsClientWithAkSk(ak, sk, endpoint)
	if err != nil {
		return nil, err
	}
	client.Config.Retry = cds.NewNoRetryPolicy()
	return &Client{client}, nil
}

func NewClientWidthEndpoint(ak, sk string, endpoint string) (*Client, error) {
	if endpoint == "" {
		endpoint = eksEndpoint
	}
	client, err := cds.NewCdsClientWithAkSk(ak, sk, endpoint)
	if err != nil {
		return nil, err
	}
	client.Config.Retry = cds.NewNoRetryPolicy()
	return &Client{client}, nil
}
