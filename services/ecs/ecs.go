package ecs

import (
	"fmt"
	"strings"

	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

func (c *client) DescribeRegions() (result *DescribeRegionsResult, err error) {
	result = new(DescribeRegionsResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeRegions).
		WithResult(result).Do()

	return
}

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
		op = op.WithQueryParam(searchInfoKey, req.SearchInfo)
	}

	if len(req.Ids) != 0 {
		op = op.WithQueryParamFilter(idsKey, strings.Join(req.Ids, ","))
	}

	err = op.WithResult(result).Do()

	return
}

func (c *client) OperateInstance(req *OperateInstanceReq) (result *OperateInstanceResult, err error) {
	if req.EcsIds == nil || len(req.EcsIds) == 0 {
		return nil, fmt.Errorf("field EcsIds is required")
	}
	if req.OpType == "" {
		return nil, fmt.Errorf("field OpType is required")
	}
	switch req.OpType {
	case StartUpInstance, RestartInstance, HardShutdownInstance, ShutdownInstance:
		req.DeleteEip = 0
	case FreeShutdownInstance:
		if req.DeleteEip != 1 {
			req.DeleteEip = 0
		}
	default:
		return nil, fmt.Errorf("invalid OpType: %s", req.OpType)
	}
	result = new(OperateInstanceResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.POST).
		WithQueryParam(actionKey, ActionOperateInstance).
		WithBody(req).
		WithResult(result).Do()

	return
}

func (c *client) ModifyInstanceName(req *ModifyInstanceNameReq) (result *ModifyInstanceNameResult, err error) {
	if req.EcsId == "" {
		return nil, fmt.Errorf("field EcsId is required")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("field Name is required")
	}
	result = new(ModifyInstanceNameResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.POST).
		WithQueryParam(actionKey, ActionModifyInstanceName).
		WithBody(req).
		WithResult(result).Do()

	return
}

func (c *client) DescribeInstance(req *DescribeInstanceReq) (result *DescribeInstanceResult, err error) {
	if req.EcsId == "" {
		return nil, fmt.Errorf("field EcsId is required")
	}
	result = new(DescribeInstanceResult)

	op := cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeInstance).
		WithQueryParam(idKey, req.EcsId)

	err = op.WithResult(result).Do()

	return
}

func (c *client) DescribeTaskEvent(req *DescribeTaskEventReq) (result *DescribeTaskEventResult, err error) {
	if req.EventId == "" {
		return nil, fmt.Errorf("field EventId is required")
	}
	result = new(DescribeTaskEventResult)

	op := cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeTaskEvent).
		WithQueryParam(eventKey, req.EventId)

	err = op.WithResult(result).Do()

	return
}

func (c *client) DescribeEcsFamilyInfo(req *DescribeEcsFamilyInfoReq) (result *DescribeEcsFamilyInfoResult, err error) {
	if req.AvailableZoneCode == "" {
		return nil, fmt.Errorf("field AvailableZoneCode is required")
	}
	if req.BillingMethod == "" {
		return nil, fmt.Errorf("field BillingMethod is required")
	}
	result = new(DescribeEcsFamilyInfoResult)

	op := cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeEcsFamilyInfo).
		WithQueryParam(azCodeKey, req.AvailableZoneCode)

	if req.BillingMethod != "" {
		op = op.WithQueryParam(billingMethodKey, string(req.BillingMethod))
	} else {
		op = op.WithQueryParam(billingMethodKey, string(OnDemandBillingMethod))
	}

	err = op.WithResult(result).Do()

	return
}

func (c *client) ChangeInstanceConfigure(req *ChangeInstanceConfigureReq) (result *ChangeInstanceConfigureResult, err error) {
	if req.AvailableZoneCode == "" {
		return nil, fmt.Errorf("field AvailableZoneCode is required")
	}
	if req.EcsFamilyName == "" {
		return nil, fmt.Errorf("field EcsFamilyName is required")
	}
	if req.Cpu <= 0 {
		return nil, fmt.Errorf("field Cpu is required")
	}
	if req.Ram <= 0 {
		return nil, fmt.Errorf("field Ram is required")
	}
	if req.EcsIds == nil || len(req.EcsIds) == 0 {
		return nil, fmt.Errorf("field EcsIds is required")
	}
	result = new(ChangeInstanceConfigureResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.POST).
		WithQueryParam(actionKey, ActionChangeInstanceConfigure).
		WithBody(req).
		WithResult(result).Do()

	return
}

func (c *client) ExtendDisk(req *ExtendDiskReq) (result *ExtendDiskResult, err error) {
	if req.DiskId == "" {
		return nil, fmt.Errorf("field DiskId is required")
	}
	result = new(ExtendDiskResult)

	if req.ExtendedSize%8 != 0 {
		return nil, fmt.Errorf("disk size must be a multiple of 8")
	}

	err = cds.NewRequestBuilder(c).
		WithURI(c.ebsRoute).
		WithMethod(http.POST).
		WithQueryParam(actionKey, ActionExtendDisk).
		WithBody(req).
		WithResult(result).Do()

	return
}
