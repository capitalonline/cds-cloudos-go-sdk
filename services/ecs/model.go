package ecs

import (
	"fmt"
)

const (
	ActionDescribeRegions         = "DescribeRegions"
	ActionDescribeInstanceList    = "DescribeInstanceList"
	ActionOperateInstance         = "OperateInstance"
	ActionModifyInstanceName      = "ModifyInstanceName"
	ActionDescribeInstance        = "DescribeInstance"
	ActionDescribeTaskEvent       = "DescribeEvent"
	ActionDescribeEcsFamilyInfo   = "DescribeEcsFamilyInfo"
	ActionChangeInstanceConfigure = "ChangeInstanceConfigure"
)

const (
	ActionExtendDisk = "ExtendDisk"
)

type (
	operate       string
	billingMethod string
)

const (
	ShutdownInstance     operate = "shutdown_ecs"
	StartUpInstance      operate = "start_up_ecs"
	RestartInstance      operate = "restart_ecs"
	HardShutdownInstance operate = "hard_shutdown_ecs"
	FreeShutdownInstance operate = "free_shutdown_ecs"
)

const (
	OnDemandBillingMethod billingMethod = "0"
	MonthlyBillingMethod  billingMethod = "1"
)

var (
	actionKey        = "Action"
	azCodeKey        = "AvailableZoneCode"
	vpcIdKey         = "VpcId"
	searchInfoKey    = "SearchInfo"
	idsKey           = "EcsIds"
	idKey            = "EcsId"
	eventKey         = "EventId"
	billingMethodKey = "BillingMethod"
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

type DescribeRegionsResult struct {
	OpenApiCommonResp
	Data []*RegionGroup `json:"Data"`
}

type RegionGroup struct {
	RegionGroupId   string        `json:"RegionGroupId"`
	RegionGroupName string        `json:"RegionGroupName"`
	RegionList      []*RegionInfo `json:"RegionList"`
}

type RegionInfo struct {
	RegionId   string    `json:"RegionId"`
	RegionName string    `json:"RegionName"`
	AzList     []*AzInfo `json:"AzList"`
	RegionCode string    `json:"RegionCode"`
}

type AzInfo struct {
	AzId              string `json:"AzId"`
	AzName            string `json:"AzName"`
	AvailableZoneCode string `json:"AvailableZoneCode"`
}

type DescribeInstanceListReq struct {
	AvailableZoneCode string   `json:"AvailableZoneCode"`
	VpcId             string   `json:"VpcId"`
	SearchInfo        string   `json:"SearchInfo"`
	Ids               []string `json:"EcsIds"`
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
	EcsIds    []string `json:"EcsIds"`
	OpType    operate  `json:"OpType"`
	DeleteEip int      `json:"DeleteEip"`
}

func (req *OperateInstanceReq) check() error {
	if len(req.EcsIds) == 0 {
		return fmt.Errorf("field EcsIds is required")
	}

	switch req.OpType {
	case StartUpInstance, RestartInstance, HardShutdownInstance, ShutdownInstance:
		req.DeleteEip = 0
	case FreeShutdownInstance:
		if req.DeleteEip != 1 {
			req.DeleteEip = 0
		}
	default:
		return fmt.Errorf("invalid OpType: %s", req.OpType)
	}
	return nil
}

type OperateInstanceResult struct {
	OpenApiCommonResp
	Data *OperateInstanceData `json:"Data"`
}

type OperateInstanceData struct {
	EventIdData
}

type ModifyInstanceNameReq struct {
	EcsId string `json:"EcsId"`
	Name  string `json:"Name"`
}

func (req *ModifyInstanceNameReq) check() error {
	if req.EcsId == "" {
		return fmt.Errorf("field EcsId is required")
	}
	if req.Name == "" {
		return fmt.Errorf("field Name is required")
	}
	return nil
}

type ModifyInstanceNameResult struct {
	OpenApiCommonResp
	Data *ModifyInstanceNameData `json:"Data"`
}

type ModifyInstanceNameData struct {
	EcsId string `json:"EcsId"`
	Name  string `json:"Name"`
}

type DescribeInstanceReq struct {
	EcsId string `json:"EcsId"`
}

func (req *DescribeInstanceReq) check() error {
	if req.EcsId == "" {
		return fmt.Errorf("field EcsId is required")
	}
	return nil
}

type DescribeInstanceResult struct {
	OpenApiCommonResp
	Data *InstanceData `json:"Data"`
}

type InstanceData struct {
	EcsId               string            `json:"EcsId"`
	EcsName             string            `json:"EcsName"`
	RegionId            string            `json:"RegionId"`
	RegionName          string            `json:"RegionName"`
	AzId                string            `json:"AzId"`
	AzName              string            `json:"AzName"`
	Status              string            `json:"Status"`
	StatusDisplay       string            `json:"StatusDisplay"`
	CreateTime          string            `json:"CreateTime"`
	Duration            int               `json:"Duration"`
	EndBillTime         string            `json:"EndBillTime"`
	IsAutoRenewal       string            `json:"IsAutoRenewal"`
	TimeZone            string            `json:"TimeZone"`
	IsRam               bool              `json:"IsRam"`
	NoChargeForShutdown int               `json:"NoChargeForShutdown"`
	EcsRule             *InstanceRuleInfo `json:"EcsRule"`
	OsInfo              *OsInfo           `json:"OsInfo"`
	Disk                *DiskInfo         `json:"Disk"`
	Pipe                *PipeInfo         `json:"Pipe"`
	BillingInfo         *BillingInfo      `json:"BillingInfo"`
}

type InstanceRuleInfo struct {
	Name    string `json:"Name"`
	CpuNum  int    `json:"CpuNum"`
	CpuUnit string `json:"CpuUnit"`
	Ram     int    `json:"Ram"`
	Gpu     int    `json:"Gpu"`
	RamUnit string `json:"RamUnit"`
	GpuUnit string `json:"GpuUnit"`
}

type OsInfo struct {
	ImageId   string `json:"ImageId"`
	ImageName string `json:"ImageName"`
	OsType    string `json:"OsType"`
	Bit       int    `json:"Bit"`
	Version   string `json:"Version"`
}

type DiskInfo struct {
	SystemDiskConf *SystemDiskConfInfo `json:"SystemDiskConf"`
	DataDiskConf   []*DataDiskConfInfo `json:"DataDiskConf"`
}
type SystemDiskConfInfo struct {
	ReleaseWithInstance int    `json:"ReleaseWithInstance"`
	DiskType            string `json:"DiskType"`
	Name                string `json:"Name"`
	Size                int    `json:"Size"`
	DiskIops            int    `json:"DiskIops"`
	BandMbps            int    `json:"BandMbps"`
	Unit                string `json:"Unit"`
	DiskId              string `json:"DiskId"`
	DiskFeature         string `json:"DiskFeature"`
}

type DataDiskConfInfo struct {
	ReleaseWithInstance int    `json:"ReleaseWithInstance"`
	DiskType            string `json:"DiskType"`
	Name                string `json:"Name"`
	Size                int    `json:"Size"`
	DiskIops            int    `json:"DiskIops"`
	BandMbps            int    `json:"BandMbps"`
	Unit                string `json:"Unit"`
	Id                  string `json:"Id"`
	DiskFeature         string `json:"DiskFeature"`
}

type PipeInfo struct {
	VpcName    string              `json:"VpcName"`
	VpcId      string              `json:"VpcId"`
	SubnetId   string              `json:"SubnetId"`
	SubnetName string              `json:"SubnetName"`
	CreateTime string              `json:"CreateTime"`
	PrivateNet string              `json:"PrivateNet"`
	PubNet     string              `json:"PubNet"`
	VirtualNet []interface{}       `json:"VirtualNet"`
	EipInfo    map[string]*EipInfo `json:"EipInfo"`
}

type EipInfo struct {
	ConfName      string `json:"ConfName"`
	EipIp         string `json:"EipIp"`
	CurrentUseQos int    `json:"CurrentUseQos"`
}

type BillingInfo struct {
	BillingMethod     string `json:"BillingMethod"`
	BillingMethodName string `json:"BillingMethodName"`
	BillingStatus     string `json:"BillingStatus"`
	BillCycleId       string `json:"BillCycleId"`
}
type DescribeTaskEventReq struct {
	EventId string `json:"EventId"`
}

func (req *DescribeTaskEventReq) check() error {
	if req.EventId == "" {
		return fmt.Errorf("field EventId is required")
	}
	return nil
}

type DescribeTaskEventResult struct {
	OpenApiCommonResp
	Data *EventResultData `json:"Data"`
}

type EventResultData struct {
	EventIdData
	EventStatus        string          `json:"EventStatus"`
	EventStatusDisplay string          `json:"EventStatusDisplay"`
	EventType          string          `json:"EventType"`
	EventTypeDisplay   string          `json:"EventTypeDisplay"`
	CreateTime         string          `json:"CreateTime"`
	TaskList           []*TaskListInfo `json:"TaskList"`
}

type TaskListInfo struct {
	TaskId          string `json:"TaskId"`
	Status          string `json:"Status"`
	StatusDisplay   string `json:"StatusDisplay"`
	ResourceId      string `json:"ResourceId"`
	CreateTime      string `json:"CreateTime"`
	UpdateTime      string `json:"UpdateTime"`
	EndTime         string `json:"EndTime"`
	ResourceType    string `json:"ResourceType"`
	ResourceDisplay string `json:"ResourceDisplay"`
	TaskType        string `json:"TaskType"`
	TaskTypeDisplay string `json:"TaskTypeDisplay"`
}

type DescribeEcsFamilyInfoReq struct {
	AvailableZoneCode string        `json:"AvailableZoneCode"`
	BillingMethod     billingMethod `json:"BillingMethod"`
}

func (req *DescribeEcsFamilyInfoReq) check() error {
	if req.AvailableZoneCode == "" {
		return fmt.Errorf("field AvailableZoneCode is required")
	}
	if req.BillingMethod != MonthlyBillingMethod {
		req.BillingMethod = OnDemandBillingMethod
	}
	return nil
}

type DescribeEcsFamilyInfoResult struct {
	OpenApiCommonResp
	Data FamilyData `json:"Data"`
}

type FamilyData struct {
	EcsFamilyInfo []*FamilyInfo `json:"EcsFamilyInfo"`
}

type FamilyInfo struct {
	EcsFamilyName string          `json:"EcsFamilyName"`
	SpecList      []*SpecListInfo `json:"SpecList"`
}

type SpecListInfo struct {
	GpuShowType      string `json:"GpuShowType"`
	GpuTypeId        string `json:"GpuTypeId"`
	Cpu              int    `json:"Cpu"`
	Ram              int    `json:"Ram"`
	Gpu              int    `json:"Gpu"`
	SupportGpuDriver string `json:"SupportGpuDriver"`
}
type ChangeInstanceConfigureReq struct {
	EcsIds            []string `json:"EcsIds"`
	AvailableZoneCode string   `json:"AvailableZoneCode"`
	EcsFamilyName     string   `json:"EcsFamilyName"`
	Cpu               int      `json:"Cpu"`
	Ram               int      `json:"Ram"`
	Gpu               int      `json:"Gpu"`
}

func (req *ChangeInstanceConfigureReq) check() error {
	if req.AvailableZoneCode == "" {
		return fmt.Errorf("field AvailableZoneCode is required")
	}
	if req.EcsFamilyName == "" {
		return fmt.Errorf("field EcsFamilyName is required")
	}
	if req.Cpu <= 0 {
		return fmt.Errorf("field Cpu is required")
	}
	if req.Ram <= 0 {
		return fmt.Errorf("field Ram is required")
	}
	if len(req.EcsIds) == 0 {
		return fmt.Errorf("field EcsIds is required")
	}

	return nil
}

type ChangeInstanceConfigureResult struct {
	OpenApiCommonResp
	Data *InstanceConfigureData `json:"Data"`
}

type InstanceConfigureData struct {
	EventIdData
}

type ExtendDiskReq struct {
	DiskId       string `json:"DiskId"`
	ExtendedSize int    `json:"ExtendedSize"`
}

func (req *ExtendDiskReq) check() error {
	if req.DiskId == "" {
		return fmt.Errorf("field DiskId is required")
	}

	if req.ExtendedSize == 0 || req.ExtendedSize%8 != 0 {
		return fmt.Errorf("field ExtendedSize must be a multiple of 8")
	}
	return nil
}

type ExtendDiskResult struct {
	OpenApiCommonResp
	Data *EventIdData `json:"Data"`
}

type EventIdData struct {
	EventId string `json:"EventId"`
}
