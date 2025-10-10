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

package subnet

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	CreateSubnetAction = "CreateSubnet"
	GetSubnetAction    = "DescribeSubnet"
	ListSubnetsAction  = "DescribeSubnet"
	DeleteSubnetAction = "DeleteSubnet"
)




func (c *Client) CreateSubnet(args *CreateSubnetReq) (*CreateSubnetResult, error) {
	result := &CreateSubnetResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(subnetURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateSubnetAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}


func (c *Client) GetSubnet(args *GetSubnetReq) (*GetSubnetResult, error) {
	result := &GetSubnetResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(subnetURI).
		WithMethod(http.POST).
		WithQueryParam("Action", GetSubnetAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListSubnets(args *ListSubnetsReq) (*ListSubnetsResult, error) {
	result := &ListSubnetsResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(subnetURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ListSubnetsAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}


func (c *Client) DeleteSubnet(args *DeleteSubnetReq) (*DeleteSubnetResult,error) {
	result := &DeleteSubnetResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(subnetURI).
		WithMethod(http.POST).
		WithQueryParam("Action", DeleteSubnetAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err

}