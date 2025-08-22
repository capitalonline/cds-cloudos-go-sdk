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
	ListSlbAction  = "DescribeVpcSlbList"
	GetSlbAction    = "DescribeVpcSlb"
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

func (c *Client) GetVpcSlb(args *GetVpcSlbReq) (*GetVpcSlbResult, error) {
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
