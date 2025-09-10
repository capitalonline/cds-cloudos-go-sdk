package ecs

type Client interface {
	DescribeRegions() (*DescribeRegionsResult, error)
	DescribeInstanceList(*DescribeInstanceListReq) (*DescribeInstanceListResult, error)
	OperateInstance(*OperateInstanceReq) (*OperateInstanceResult, error)
	ModifyInstanceName(*ModifyInstanceNameReq) (*ModifyInstanceNameResult, error)
	DescribeInstance(*DescribeInstanceReq) (*DescribeInstanceResult, error)
	DescribeTaskEvent(*DescribeTaskEventReq) (*DescribeTaskEventResult, error)
	DescribeZoneInstanceType(*DescribeZoneInstanceTypeReq) (*DescribeZoneInstanceTypeResult, error)
	ChangeInstanceConfigure(*ChangeInstanceConfigureReq) (*ChangeInstanceConfigureResult, error)

	ExtendDisk(*ExtendDiskReq) (*ExtendDiskResult, error)
}
