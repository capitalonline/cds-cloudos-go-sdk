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

package natgateway

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	CreateNatGatewayAction         = "CreateNAT"
	GetNatGatewayAction            = "DescribeNAT"
	ListNatGatewaysAction          = "DescribeNAT"
	DeleteNatGatewaysAction         = "DeleteNAT"
	UpdateNatGatewayAction                 ="UpdateNAT"


	CreateNatRuleAction  = "CreateNATRule"
	GetNatRuleAction = "DescribeNATRule"
	ListNatRulesAction = "DescribeNATRule"
	UpdateNatRuleAction = "UpdateNATRule"
	DeleteNatRuleAction = "DeleteNATRule"

)

func (c *Client) CreateNatGateway(args *CreateNatGatewayReq) (*CreateNatGatewayResult, error) {
	result := &CreateNatGatewayResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateNatGatewayAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) GetNatGateway(args map[string]string) (*GetNatGatewayResult, error) {
	result := &GetNatGatewayResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.GET).
		WithQueryParam("Action", GetNatGatewayAction).
		WithQueryParams(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) ListNatGateways(args map[string]string) (*ListNatGatewaysResult, error) {
	result := &ListNatGatewaysResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ListNatGatewaysAction).
		WithQueryParams(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) UpdateNatGateway(args *UpdateNatGatewayReq) (*UpdateNatGatewayResult, error) {
	result := &UpdateNatGatewayResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", UpdateNatGatewayAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) DeleteNatGateway(args *DeleteNatGatewayReq) (*DeleteNatGatewayResult, error) {
	result := &DeleteNatGatewayResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", DeleteNatGatewaysAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) CreateNATRule(args *CreateNatRuleReq) (*CreateNatRuleResult, error) {
	result := &CreateNatRuleResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateNatRuleAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) GetNATRule(args map[string]string) (*GetNatRuleResult, error) {
	result := &GetNatRuleResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", GetNatRuleAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) ListNATRules(args map[string]string) (*ListNatRulesResult, error) {
	result := &ListNatRulesResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ListNatRulesAction).
		WithQueryParams(args).
		WithResult(result).
		Do()
	return result, err
}


func (c *Client) UpdateNATRule(args *UpdateNATRuleReq) (*UpdateNATRuleResult, error) {
	result := &UpdateNATRuleResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", UpdateNatRuleAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}



func (c *Client) DeleteNATRule(args *DeleteNATRuleReq) (*DeleteNATRuleResult, error) {
	result := &DeleteNATRuleResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(NatURI).
		WithMethod(http.POST).
		WithQueryParam("Action", DeleteNatRuleAction).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}
