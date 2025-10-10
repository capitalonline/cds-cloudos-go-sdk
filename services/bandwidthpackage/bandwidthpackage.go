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

package bandwidthpackage

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	CreateBandwidthPackageAction         = "CreateBandwidth"
	GetBandwidthPackageAction            = "DescribeBandwidth"
	ListBandwidthPackagesAction          = "DescribeBandwidth"
	DeleteBandwidthPackageAction         = "DeleteBandwidth"
	UpdateBandwidthPackageAction         = "UpdateBandwidth"
	AddBandwidthPackageIpAction          = "BandwidthAddEIP"
	RemoveBandwidthPackageIpAction       = "BandwidthRemoveEIP"
	BandwidthPackageBindResourceAction   = "BandwidthBindResource"
	BandwidthPackageUnbindResourceAction = "BandwidthUnbindResource"
)



func (c *Client) CreateBandwidthPackage(args *CreateBandwidthPackageReq) (*CreateBandwidthPackageResult, error) {
	result := &CreateBandwidthPackageResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateBandwidthPackageAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) GetBandwidthPackage(args *GetBandwidthPackageReq) (*GetBandwidthPackageResult, error) {
	result := &GetBandwidthPackageResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", GetBandwidthPackageAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}


func (c *Client) ListBandwidthPackages(args *ListBandwidthPackagesReq) (*ListBandwidthPackagesResult, error) {
	result := &ListBandwidthPackagesResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ListBandwidthPackagesAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) UpdateBandwidthPackage(args *UpdateBandwidthPackageReq) (*UpdateBandwidthResult, error) {
	result := &UpdateBandwidthResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", UpdateBandwidthPackageAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) DeleteBandwidthPackage(args *DeleteBandwidthPackageReq) (*DeleteBandwidthPackageResult,error) {
	result := &DeleteBandwidthPackageResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", DeleteBandwidthPackageAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err

}


func (c *Client) AddBandwidthPackageIp(args *AddBandwidthPackageIpReq) (*AddBandwidthPackageIpResult,error) {
	result := &AddBandwidthPackageIpResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", AddBandwidthPackageIpAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err

}

func (c *Client) RemoveBandwidthPackageIp(args *RemoveBandwidthPackageIpReq) (*RemoveBandwidthPackageIpResult,error) {
	result := &RemoveBandwidthPackageIpResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", RemoveBandwidthPackageIpAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) BandwidthBindResource(args *BandwidthBindResourceReq) (*BandwidthBindResourceResult,error) {
	result := &BandwidthBindResourceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", BandwidthPackageBindResourceAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err

}

func (c *Client) BandwidthUnbindResource(args *BandwidthUnbindResourceReq) (*BandwidthUnbindResourceResult,error) {
	result := &BandwidthUnbindResourceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(BandwidthpackageURI).
		WithMethod(http.POST).
		WithQueryParam("Action", BandwidthPackageUnbindResourceAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}