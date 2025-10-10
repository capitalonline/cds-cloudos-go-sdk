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
	ActionQueryEventStatus         = "QueryEventStatus"
)

const (
	Success      = "Success"
	EventSuccess = "success"
	EventFailed  = "failed"
	EventError   = "error"
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

func (c *Client) IsAttachedECS(netcardId string) (bool, error) {
	result := &IsAttachedECSResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionIsAttachedECS).
		WithQueryParam("NetcardId", netcardId).
		WithResult(result).
		Do()
	if err != nil {
		return false, err
	}
	return result.Data.Attached, ConfirmResult(result.Code, result.Message)
}

func (c *Client) QueryEventStatus(eventId string) (bool, error) {
	result := &QueryTaskStatusResult{}
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ActionQueryEventStatus).
		WithQueryParam("EventId", eventId).
		WithResult(result).
		Do()
	if err != nil {
		return false, err
	}
	eventStatus := result.Data.EventStatus
	if err = ConfirmResult(result.Code, result.Message); err != nil {
		return false, err
	}
	if eventStatus == EventFailed || eventStatus == EventError {
		return false, fmt.Errorf("event %s not sucess, status:%s", eventId, eventStatus)
	}
	return eventStatus == EventSuccess, nil
}

func ConfirmResult(code string, message string) error {
	if code != Success {
		return fmt.Errorf("code： %s， message： %s", code, message)
	}
	return nil
}
