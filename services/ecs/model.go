package ecs

const (
	ActionDescribeInstanceList     = "DescribeInstanceList"
	ActionOperateInstance          = "OperateInstance"
	ActionModifyInstanceName       = "ModifyInstanceName"
	ActionDescribeInstance         = "DescribeInstance"
	ActionDescribeTaskEvent        = "DescribeEvent"
	ActionDescribeZoneInstanceType = "DescribeZoneInstanceType"
	ActionChangeInstanceConfigure  = "ChangeInstanceConfigure"
)

const (
	ActionExtendDisk = "ExtendDisk"
)

var (
	actionKey     = "Action"
	azCodeKey     = "AvailableZoneCode"
	vpcIdKey      = "VpcId"
	searchInfoKey = "SearchInfo"
)

type OpenApiCommonResp struct {
	Code      string `json:"Code"`
	Message   string `json:"Msg"`
	RequestId string `json:"RequestId"`
	CommonOpenApiPage
}

type CommonOpenApiPage struct {
	TotalCount int `json:"TotalCount,omitempty"`
	PageIndex  int `json:"PageIndex,omitempty"`
	PageSize   int `json:"PageSize,omitempty"`
}

type DescribeInstanceListReq struct {
	AvailableZoneCode string `json:"AvailableZoneCode"`
	VpcId             string `json:"VpcId"`
	SearchInfo        string `json:"SearchInfo"`
}

type DescribeInstanceListResult struct {
	OpenApiCommonResp
	Data *InstanceListData `json:"Data"`
}

type InstanceListData struct {
	EcsList []*InstanceSimpleInfo `json:"EcsList"`
}

type InstanceSimpleInfo struct {
	EcsId               string                    `json:"EcsId"`
	EcsName             string                    `json:"EcsName"`
	VpcId               string                    `json:"VpcId"`
	VpcName             string                    `json:"VpcName"`
	Status              string                    `json:"Status"`
	StatusDisplay       string                    `json:"StatusDisplay"`
	AzId                string                    `json:"AzId"`
	AzName              string                    `json:"AzName"`
	RegionId            string                    `json:"RegionId"`
	RegionName          string                    `json:"RegionName"`
	PrivateNet          string                    `json:"PrivateNet"`
	PubNet              string                    `json:"PubNet"`
	VirtualNet          []interface{}             `json:"VirtualNet"`
	EipInfo             map[string]*EipSimpleInfo `json:"EipInfo"`
	CreateTime          string                    `json:"CreateTime"`
	EndBillTime         string                    `json:"EndBillTime"`
	IsAutoRenewal       string                    `json:"IsAutoRenewal"`
	IsGpu               bool                      `json:"IsGpu"`
	CardName            string                    `json:"CardName"`
	CpuSize             int                       `json:"CpuSize"`
	RamSize             int                       `json:"RamSize"`
	GpuSize             int                       `json:"GpuSize"`
	BillingMethodName   string                    `json:"BillingMethodName"`
	BillingMethod       string                    `json:"BillingMethod"`
	OsType              string                    `json:"OsType"`
	OsVersion           string                    `json:"OsVersion"`
	ImageName           string                    `json:"ImageName"`
	OsBit               int                       `json:"OsBit"`
	CustomerId          string                    `json:"CustomerId"`
	SystemDiskType      string                    `json:"SystemDiskType"`
	SystemDiskFeature   string                    `json:"SystemDiskFeature"`
	SystemDiskSize      int                       `json:"SystemDiskSize"`
	SupportGpuDriver    string                    `json:"SupportGpuDriver"`
	SpecFamilyId        string                    `json:"SpecFamilyId"`
	EcsFamilyName       string                    `json:"EcsFamilyName"`
	NoChargeForShutdown int                       `json:"NoChargeForShutdown"`
}

type EipSimpleInfo struct {
	ConfName string `json:"ConfName"`
	EipIp    string `json:"EipIp"`
}

type OperateInstanceReq struct {
}

type OperateInstanceResult struct {
}

type ModifyInstanceNameReq struct {
}

type ModifyInstanceNameResult struct {
}

type DescribeInstanceReq struct {
}

type DescribeInstanceResult struct {
}

type DescribeTaskEventReq struct {
}

type DescribeTaskEventResult struct {
}

type DescribeZoneInstanceTypeReq struct {
}

type DescribeZoneInstanceTypeResult struct {
}

type ChangeInstanceConfigureReq struct {
}

type ChangeInstanceConfigureResult struct {
}

type ExtendDiskReq struct {
}

type ExtendDiskResult struct {
}
