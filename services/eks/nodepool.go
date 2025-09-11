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
	"encoding/json"
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	ActionCreateNodePool  = "CreateNodePool"
	ActionListNodePool    = "ListNodePool"
	ActionDeleteNodePool  = "DeleteNodePool"
	ActionScalingNodePool = "ScalingNodePool"
)

// CreateNodePool creates a new node pool in the cluster
func (c *Client) CreateNodePool(args *CreateNodePoolReq) (*CreateNodePoolResult, error) {
	result := &CreateNodePoolResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionCreateNodePool).
		WithBody(args).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}

// ListNodePool retrieves a list of node pools in a cluster
func (c *Client) ListNodePool(args *ListNodePoolReq) (*ListNodePoolResult, error) {
	result := &ListNodePoolResult{}
	bytes, _ := json.Marshal(args)
	tmp := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &tmp)
	builder := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionListNodePool)
	params := make(map[string]string)
	for key, value := range tmp {
		params[key] = fmt.Sprintf("%v", value)
	}
	err := builder.WithQueryParams(params).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}

// DeleteNodePool deletes a node pool from the cluster
func (c *Client) DeleteNodePool(args *DeleteNodePoolReq) (*DeleteNodePoolResult, error) {
	result := &DeleteNodePoolResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionDeleteNodePool).
		WithBody(args).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}

// ScalingNodePool scales the number of nodes in a node pool
func (c *Client) ScalingNodePool(args *ScalingNodePoolReq) (*ScalingNodePoolResult, error) {
	result := &ScalingNodePoolResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionScalingNodePool).
		WithBody(args).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}
