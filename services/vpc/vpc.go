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

package vpc

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	CreateVpcAction = "CreateVPC"
	GetVpcAction    = "DescribeVPC"
	ListVpcsAction  = "DescribeVPC"
	DeleteVpcAction = "DeleteVPC"
	DescribeTask    = "DescribeTask"
)


func (c *Client) CreateVPC(args *CreateVpcReq) (*CreateVpcResult, error) {
	result := &CreateVpcResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateVpcAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}


func (c *Client) GetVpc(args *GetVpcReq) (*GetVpcResult, error) {
	result := &GetVpcResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", GetVpcAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListVpcs(args *ListVpcsReq) (*ListVpcsResult, error) {
	result := &ListVpcsResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ListVpcsAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err

}

func (c *Client) DeleteVpc(args *DeleteVpcReq) (*DeleteVpcResult,error) {
	result := &DeleteVpcResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", DeleteVpcAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err

}

func (c *Client) DescribeTask(args map[string]string) (*DescribeTaskResult, error){
	result := &DescribeTaskResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.GET).
		WithQueryParam("Action", DescribeTask).
		WithQueryParams(args).
		WithResult(result).
		Do()
	return result, err
}