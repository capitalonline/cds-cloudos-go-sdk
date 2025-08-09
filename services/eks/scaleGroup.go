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

// DescribeScalingGroup - Describe Scale Group
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - reqBody: http request body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeScalingGroup(cli cds.Client, reqBody *cds.Body) error {

	// Build the request
	req := &cds.CdsRequest{}
	req.SetUri(UriPrefix)
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &cds.CdsResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}
