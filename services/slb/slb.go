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

package slb

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	ListSlbAction = "DescribeVpcSlbList"
	GetSlbAction  = "DescribeVpcSlb"

	CreateVpcSlb                   = "CreateVpcSlb"
	DeleteVpcSlb                   = "DeleteVpcSlb"
	UpdateVpcSlb                   = "UpdateVpcSlb"
	DescribeVpcSlbListInfo         = "DescribeVpcSlbListInfo"
	DescribeVpcSlbDetailInfo       = "DescribeVpcSlbDetailInfo"
	DescribeVpcSlbListenCreateInfo = "DescribeVpcSlbListenCreateInfo"
	CreateVpcSlbListen             = "CreateVpcSLBListen"
	DeleteVpcSlbListen             = "DeleteVpcSLBListen"
	UpdateVpcSlbListen             = "UpdateVpcSLBListen"
	DescribeVpcSlbListenList       = "DescribeVpcSlbListenList"
	QueryVpcSlbListen              = "QueryVpcSLBListen"
	DescribeVpcSlbListenRsInfo     = "DescribeVpcSlbListenRsInfo"
	CreateVpcSlbRsPort             = "CreateVpcSLBRsPort"
	DeleteVpcSlbRsPort             = "DeleteVpcSLBRsPort"
	UpdateVpcSlbRsPort             = "UpdateVpcSLBRsPort"
	QueryVpcSlbRsPort              = "QueryVpcSLBRsPort"
	BandwidthBindResource          = "BandwidthBindResource"
	BandwidthUnbindResource        = "BandwidthUnbindResource"
)

func (c *Client) ListVpcSlb(args *ListVpcSlbReq) (*ListVpcSlbResult, error) {
	result := &ListVpcSlbResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ListSlbAction).
		WithQueryParam("VpcId", args.VpcId).
		WithResult(result).
		Do()
	return result, err

}

func (c *Client) GetVpcSlbDetail(args *GetVpcSlbDetailReq) (*GetVpcSlbResult, error) {
	result := &GetVpcSlbResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.GET).
		WithQueryParam("Action", GetSlbAction).
		WithQueryParam("SlbId", args.SlbId).
		WithQueryParam("SlbName", args.SlbName).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) CreateVpcSlb(args *CreateVpcSlbReq) (*CreateVpcSlbResult, error) {
	result := &CreateVpcSlbResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "CreateVpcSlb").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DeleteVpcSlb(args *DeleteVpcSlbReq) (*DeleteVpcSlbResult, error) {
	result := &DeleteVpcSlbResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DeleteVpcSlb").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) UpdateVpcSlb(args *UpdateVpcSlbReq) (*UpdateVpcSlbResult, error) {
	result := &UpdateVpcSlbResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "UpdateVpcSlb").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DescribeVpcSlbList(args *DescribeVpcSlbListReq) (*DescribeVpcSlbListResult, error) {
	result := &DescribeVpcSlbListResult{}
	builder := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.GET).
		WithQueryParam("Action", "DescribeVpcSlbListInfo")
	if args != nil {
		if args.VpcId != "" {
			builder = builder.WithQueryParam("VpcId", args.VpcId)
		}
		if args.Keyword != "" {
			builder = builder.WithQueryParam("Keyword", args.Keyword)
		}
	}
	err := builder.WithResult(result).Do()
	return result, err
}

func (c *Client) DescribeVpcSlbDetailInfo(args *DescribeVpcSlbDetailInfoReq) (*DescribeVpcSlbDetailInfoResult, error) {
	result := &DescribeVpcSlbDetailInfoResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DescribeVpcSlbDetailInfo").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DescribeVpcSlbListenCreateInfo(args *DescribeVpcSlbListenCreateInfoReq) (*DescribeVpcSlbListenCreateInfoResult, error) {
	result := &DescribeVpcSlbListenCreateInfoResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DescribeVpcSlbListenCreateInfo").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) CreateVpcSlbListen(args *CreateVpcSlbListenReq) (*CreateVpcSlbListenResult, error) {
	result := &CreateVpcSlbListenResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "CreateVpcSlbListen").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DeleteVpcSlbListen(args *DeleteVpcSlbListenReq) (*DeleteVpcSlbListenResult, error) {
	result := &DeleteVpcSlbListenResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DeleteVpcSlbListen").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) UpdateVpcSlbListen(args *UpdateVpcSlbListenReq) (*UpdateVpcSlbListenResult, error) {
	result := &UpdateVpcSlbListenResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "UpdateVpcSlbListen").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DescribeVpcSlbListenList(args *DescribeVpcSlbListenListReq) (*DescribeVpcSlbListenListResult, error) {
	result := &DescribeVpcSlbListenListResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DescribeVpcSlbListenList").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) QueryVpcSlbListen(args *QueryVpcSlbListenReq) (*QueryVpcSlbListenResult, error) {
	result := &QueryVpcSlbListenResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "QueryVpcSlbListen").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DescribeVpcSlbListenRsInfo(args *DescribeVpcSlbListenRsInfoReq) (*DescribeVpcSlbListenRsInfoResult, error) {
	result := &DescribeVpcSlbListenRsInfoResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DescribeVpcSlbListenRsInfo").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) CreateVpcSlbRsPort(args *CreateVpcSlbRsPortReq) (*CreateVpcSlbRsPortResult, error) {
	result := &CreateVpcSlbRsPortResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "CreateVpcSLBRsPort").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DeleteVpcSlbRsPort(args *DeleteVpcSlbRsPortReq) (*DeleteVpcSlbRsPortResult, error) {
	result := &DeleteVpcSlbRsPortResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "DeleteVpcSLBRsPort").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) UpdateVpcSlbRsPort(args *UpdateVpcSlbRsPortReq) (*UpdateVpcSlbRsPortResult, error) {
	result := &UpdateVpcSlbRsPortResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "UpdateVpcSLBRsPort").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) QueryVpcSlbRsPort(args *QueryVpcSlbRsPortReq) (*QueryVpcSlbRsPortResult, error) {
	result := &QueryVpcSlbRsPortResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "QueryVpcSLBRsPort").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) BandwidthBindResource(args *BandwidthBindResourceReq) (*BandwidthBindResourceResult, error) {
	result := &BandwidthBindResourceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "BandwidthBindResource").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) BandwidthUnbindResource(args *BandwidthUnbindResourceReq) (*BandwidthUnbindResourceResult, error) {
	result := &BandwidthUnbindResourceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(vpcURI).
		WithMethod(http.POST).
		WithQueryParam("Action", "BandwidthUnbindResource").
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}
