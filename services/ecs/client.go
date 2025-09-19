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

package ecs

import (
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
)

const (
	ecsV1Route  = "/ecs/v1"
	ecsV1Route1 = "/ecs/v1.1"
	ebsV1Route  = "/ebs/v1"
)

type Client interface {
	DescribeRegions() (*DescribeRegionsResult, error)
	DescribeTaskEvent(*DescribeTaskEventReq) (*DescribeTaskEventResult, error)

	CreateInstance(*CreateInstanceReq) (*CreateInstanceResult, error)
	DeleteInstance(*DeleteInstanceReq) (*DeleteInstanceResult, error)
	ModifyInstancePassword(*ModifyInstancePasswordReq) (*ModifyInstancePasswordResult, error)
	DescribeInstanceStatus(*DescribeInstanceStatusReq) (*DescribeInstanceStatusResult, error)
	DescribeInstanceList(*DescribeInstanceListReq) (*DescribeInstanceListResult, error)
	DescribeInstance(*DescribeInstanceReq) (*DescribeInstanceResult, error)
	OperateInstance(*OperateInstanceReq) (*OperateInstanceResult, error)
	ModifyInstanceName(*ModifyInstanceNameReq) (*ModifyInstanceNameResult, error)
	DescribeEcsFamilyInfo(*DescribeEcsFamilyInfoReq) (*DescribeEcsFamilyInfoResult, error)
	ChangeInstanceConfigure(*ChangeInstanceConfigureReq) (*ChangeInstanceConfigureResult, error)
	DescribeImages(*DescribeImagesReq) (*DescribeImagesResult, error)

	ExtendDisk(*ExtendDiskReq) (*ExtendDiskResult, error)
}

type client struct {
	*cds.CdsClient
}

func NewClient(ak, sk string) (Client, error) {
	cli, err := cds.NewCdsClientWithAkSk(ak, sk)
	if err != nil {
		return nil, err
	}

	cli.Config.Retry = cds.NewNoRetryPolicy()
	return &client{
		CdsClient: cli,
	}, nil
}
