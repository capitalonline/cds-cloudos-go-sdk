package ecs

import (
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

func (c *client) DescribeInstanceList(req *DescribeInstanceListReq) (result *DescribeInstanceListResult, err error) {
	result = new(DescribeInstanceListResult)

	op := cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeInstanceList)

	if req.AvailableZoneCode != "" {
		op = op.WithQueryParam(azCodeKey, req.AvailableZoneCode)
	}

	if req.VpcId != "" {
		op = op.WithQueryParam(vpcIdKey, req.VpcId)
	}

	if req.SearchInfo != "" {
		op = op.WithQueryParam(azCodeKey, req.SearchInfo)
	}

	err = op.WithResult(result).Do()

	return
}

func (c *client) OperateInstance(req *OperateInstanceReq) (*OperateInstanceResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}

func (c *client) ModifyInstanceName(req *ModifyInstanceNameReq) (*ModifyInstanceNameResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}

func (c *client) DescribeInstance(req *DescribeInstanceReq) (*DescribeInstanceResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}

func (c *client) DescribeTaskEvent(req *DescribeTaskEventReq) (*DescribeTaskEventResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}

func (c *client) DescribeZoneInstanceType(req *DescribeZoneInstanceTypeReq) (*DescribeZoneInstanceTypeResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}

func (c *client) ChangeInstanceConfigure(req *ChangeInstanceConfigureReq) (*ChangeInstanceConfigureResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}

func (c *client) ExtendDisk(req *ExtendDiskReq) (*ExtendDiskResult, error) {
	//TODO implement me
	return nil, fmt.Errorf("implement me")
}
