package eks

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	ActionListNodes   = "ListNodes"
	ActionDeleteNodes = "DeleteNodes"
)

func (c *Client) ListNodes(args *ListNodesReq) (*ListNodesResult, error) {
	result := &ListNodesResult{}
	bytes, _ := json.Marshal(args)
	tmp := make(map[string]interface{})
	_ = json.Unmarshal(bytes, &tmp)
	builder := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionListNodes)
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

func (c *Client) DeleteNodes(args *DeleteNodesReq) (*DeleteNodesResult, error) {
	result := &DeleteNodesResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionDeleteNodes).
		WithBody(args).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}
