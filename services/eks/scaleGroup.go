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
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	DescribeScalingGroupAction = "DescribeScalingGroup"
)

// GetScalingGroupDetail - Get Scale Group Detail
//
// PARAMS:
//   - scaleGroupId: the Scale Group id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) GetScalingGroupDetail(ScalingGroupId string) (*GetScalingGroupDetailResult, error) {
	result := &GetScalingGroupDetailResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", DescribeScalingGroupAction).
		WithQueryParam("ScalingGroupId", ScalingGroupId).
		WithResult(result).
		Do()

	return result, err
}
