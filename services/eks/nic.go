package eks

import (
	"fmt"
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

const Success = "Success"

func (c *Client) AttachNetworkInterface(args *AttachNetworkInterfaceReq) (*AttachNetworkInterfaceResult, error) {
	result := &AttachNetworkInterfaceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", ActionAttachNetworkInterface).
		WithBody(args).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
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
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}

func (c *Client) DescribeNetworkInterface(netcardId string) (*DescribeNetworkInterfaceResult, error) {
	result := &DescribeNetworkInterfaceResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionDescribeNetworkInterface).
		WithQueryParam("NetcardId", netcardId).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
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
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}

func (c *Client) QueryEventStatus(taskId string) (*QueryTaskStatusResult, error) {
	result := &QueryTaskStatusResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionQueryTaskStatus).
		WithQueryParam("TaskId", taskId).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, ConfirmResult(result.Code, result.Message)
}

func ConfirmResult(code string, message string) error {
	if code != Success {
		return fmt.Errorf("request not success,with code:%s,with message:%s", code, message)
	}
	return nil
}
