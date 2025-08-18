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

package eip

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	CreateEipAction         = "CreateEIP"
	GetEipAction            = "DescribeEIP"
	ListEipsAction          = "DescribeEIP"
	ReleaseEipAction        = "DeleteEIP"
	UpdateEIPAction         = "UpdateEIP"
)



func (c *Client) CreateEip(args *CreateEIPReq) (*CreateEIPResult, error) {
	result := &CreateEIPResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(EipURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateEipAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}



func (c *Client) GetEip(args *GetEipReq) (*GetEipResult, error) {
	result := &GetEipResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(EipURI).
		WithMethod(http.POST).
		WithQueryParam("Action", GetEipAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}



func (c *Client) ListEips(args *ListEipsReq) (*ListEipsResult, error) {
	result := &ListEipsResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(EipURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ListEipsAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) ReleaseEip(args *ReleaseEipReq) (*ReleaseEipResult, error) {
	result := &ReleaseEipResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(EipURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ReleaseEipAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) UpdateEip(args *UpdateEIPReq) (*UpdateEIPResult, error) {
	result := &UpdateEIPResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(EipURI).
		WithMethod(http.POST).
		WithQueryParam("Action", UpdateEIPAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}