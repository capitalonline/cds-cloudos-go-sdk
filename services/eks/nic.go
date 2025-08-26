package eks

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	ActionAttachNetworkInterface   = "AttachNetworkInterface"
	ActionDetachNetworkInterface   = "DetachNetworkInterface"
	ActionDescribeNetworkInterface = "DescribeNetworkInterface"
	ActionIsAttachedECS            = "IsAttachedECS"
	ActionQueryTaskStatus          = "QueryTaskStatus"
)

func (c *Client) AttachNetworkInterface(args *AttachNetworkInterfaceReq) (*AttachNetworkInterfaceResult, error) {
	result := &AttachNetworkInterfaceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionAttachNetworkInterface).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DetachNetworkInterface(args *DetachNetworkInterfaceReq) (*DetachNetworkInterfaceResult, error) {
	result := &DetachNetworkInterfaceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionDetachNetworkInterface).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DescribeNetworkInterface(args *DescribeNetworkInterfaceReq) (*DescribeNetworkInterfaceResult, error) {
	result := &DescribeNetworkInterfaceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionDescribeNetworkInterface).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) IsAttachedECS(netcardId string) (*DescribeNetworkInterfaceResult, error) {
	result := &DescribeNetworkInterfaceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionIsAttachedECS).
		WithQueryParam("NetcardId", netcardId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) QueryTaskStatus(taskId string) (*QueryTaskStatusResult, error) {
	result := &QueryTaskStatusResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionQueryTaskStatus).
		WithQueryParam("TaskId", taskId).
		WithResult(result).
		Do()

	return result, err
}
