package ecs

import (
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
	if err = req.check(); err != nil {
		return
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
	if err = req.check(); err != nil {
		return
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
	if err = req.check(); err != nil {
		return
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
	if err = req.check(); err != nil {
		return
	}

	result = new(DescribeTaskEventResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeTaskEvent).
		WithQueryParam(eventKey, req.EventId).Do()

	return
}

func (c *client) DescribeEcsFamilyInfo(req *DescribeEcsFamilyInfoReq) (result *DescribeEcsFamilyInfoResult, err error) {
	if err = req.check(); err != nil {
		return
	}

	result = new(DescribeEcsFamilyInfoResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ecsRoute).
		WithMethod(http.GET).
		WithQueryParam(actionKey, ActionDescribeEcsFamilyInfo).
		WithQueryParam(azCodeKey, req.AvailableZoneCode).
		WithQueryParam(billingMethodKey, string(req.BillingMethod)).Do()

	return
}

func (c *client) ChangeInstanceConfigure(req *ChangeInstanceConfigureReq) (result *ChangeInstanceConfigureResult, err error) {
	if err = req.check(); err != nil {
		return
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
	if err = req.check(); err != nil {
		return
	}

	result = new(ExtendDiskResult)

	err = cds.NewRequestBuilder(c).
		WithURI(c.ebsRoute).
		WithMethod(http.POST).
		WithQueryParam(actionKey, ActionExtendDisk).
		WithBody(req).
		WithResult(result).Do()

	return
}
