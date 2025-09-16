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
	ShutdownInstance     operate = "shutdown_ecs"      // 关机操作
	StartUpInstance      operate = "start_up_ecs"      // 开机操作
	RestartInstance      operate = "restart_ecs"       // 重启操作
	HardShutdownInstance operate = "hard_shutdown_ecs" // 强制关机操作

	/*
		FreeShutdownInstance 参数说明：
		仅按需计费的云盘实例支持关机不计费，再开机公网IP可能会变化；
		若批量操作关机：支持关机不收费的实例，关机后停止 CPU、内存、GPU和公网收费；
		不支持关机不收费的实例，正常关机，继续收费；
		不计费关机期间不支持除开机、删除、定制镜像外的操作。
		注意：目前关机不计费白名单开放，若需要请联系商务。
	*/
	FreeShutdownInstance operate = "free_shutdown_ecs" // 关机不计费
)

const (
	OnDemandBillingMethod billingMethod = "0" // 按需计费
	MonthlyBillingMethod  billingMethod = "1" // 包月
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

// OpenApiCommonResp 开放API通用响应结构
type OpenApiCommonResp struct {
	Code      string `json:"Code"`      // 响应码
	Message   string `json:"Msg"`       // 响应消息
	RequestId string `json:"RequestId"` // 请求ID
	CommonOpenApiPage
}

// CommonOpenApiPage 通用分页信息
type CommonOpenApiPage struct {
	TotalCount int `json:"TotalCount,omitempty"` // 总记录数
	PageIndex  int `json:"PageIndex,omitempty"`  // 当前页码
	PageSize   int `json:"PageSize,omitempty"`   // 每页记录数
}

// DescribeRegionsResult DescribeRegions接口响应结果
type DescribeRegionsResult struct {
	OpenApiCommonResp
	Data []*RegionGroup `json:"Data"` // 区域组列表
}

// RegionGroup 区域组信息
type RegionGroup struct {
	RegionGroupId   string        `json:"RegionGroupId"`   // 区域组ID
	RegionGroupName string        `json:"RegionGroupName"` // 区域组名称
	RegionList      []*RegionInfo `json:"RegionList"`      // 区域列表
}

// RegionInfo 区域信息
type RegionInfo struct {
	RegionId   string    `json:"RegionId"`   // 区域ID
	RegionName string    `json:"RegionName"` // 区域名称
	AzList     []*AzInfo `json:"AzList"`     // 可用区列表
	RegionCode string    `json:"RegionCode"` // 区域代码
}

// AzInfo 可用区信息
type AzInfo struct {
	AzId              string `json:"AzId"`              // 可用区ID
	AzName            string `json:"AzName"`            // 可用区名称
	AvailableZoneCode string `json:"AvailableZoneCode"` // 可用区代码
}

// EventIdData 事件ID数据
type EventIdData struct {
	EventId string `json:"EventId"` // 事件ID
}

// DescribeInstanceListReq DescribeInstanceList接口请求参数
type DescribeInstanceListReq struct {
	AvailableZoneCode string   `json:"AvailableZoneCode"` // 可用区代码
	VpcId             string   `json:"VpcId"`             // VPC ID
	SearchInfo        string   `json:"SearchInfo"`        // 搜索信息（实例ID/实例名称/私网IP）
	Ids               []string `json:"EcsIds"`            // 实例ID列表
}

// DescribeInstanceListResult DescribeInstanceList接口响应结果
type DescribeInstanceListResult struct {
	OpenApiCommonResp
	Data *InstanceListData `json:"Data"` // 实例列表数据
}

// InstanceListData 实例列表数据
type InstanceListData struct {
	EcsList []*InstanceSimpleInfo `json:"EcsList"` // 实例简基础息列表
}

// InstanceSimpleInfo 实例简要信息
type InstanceSimpleInfo struct {
	EcsId                string                    `json:"EcsId"`                // 实例ID
	EcsName              string                    `json:"EcsName"`              // 实例名称
	VpcId                string                    `json:"VpcId"`                // VPC ID
	VpcName              string                    `json:"VpcName"`              // VPC名称
	Status               string                    `json:"Status"`               // 云服务器状态码
	StatusDisplay        string                    `json:"StatusDisplay"`        // 状态显示名称
	AzId                 string                    `json:"AzId"`                 // 可用区ID
	AzName               string                    `json:"AzName"`               // 可用区名称
	RegionId             string                    `json:"RegionId"`             // 区域ID
	RegionName           string                    `json:"RegionName"`           // 区域名称
	PrivateNet           string                    `json:"PrivateNet"`           // 私网IP
	PubNet               string                    `json:"PubNet"`               // 默认出网网关ip
	VirtualNet           []interface{}             `json:"VirtualNet"`           // 虚拟网络信息
	EipInfo              map[string]*EipSimpleInfo `json:"EipInfo"`              // EIP信息
	CreateTime           string                    `json:"CreateTime"`           // 创建时间
	SpecFamilyId         string                    `json:"SpecFamilyId"`         // 规格族ID
	EcsFamilyName        string                    `json:"EcsFamilyName"`        // 实例规格族名称
	CpuSize              int                       `json:"CpuSize"`              // CPU核心数
	RamSize              int                       `json:"RamSize"`              // 内存大小(GB)
	EndBillTime          string                    `json:"EndBillTime"`          // 计费结束时间
	IsAutoRenewal        string                    `json:"IsAutoRenewal"`        // 是否自动续费
	IsGpu                bool                      `json:"IsGpu"`                // 是否为GPU实例
	CardName             string                    `json:"CardName"`             // 显卡型号
	GpuSize              int                       `json:"GpuSize"`              // GPU数量
	BillingMethodName    string                    `json:"BillingMethodName"`    // 计费方式名称
	BillingMethod        string                    `json:"BillingMethod"`        // 计费方式
	OsType               string                    `json:"OsType"`               // 操作系统类型
	OsVersion            string                    `json:"OsVersion"`            // 操作系统版本
	ImageName            string                    `json:"ImageName"`            // 镜像名称
	OsBit                int                       `json:"OsBit"`                // 操作系统位数
	SystemDiskType       string                    `json:"SystemDiskType"`       // 系统盘类型
	SystemDiskSize       int                       `json:"SystemDiskSize"`       // 系统盘大小(GB)
	NoChargeForShutdown  int                       `json:"NoChargeForShutdown"`  // 关机是否收费
	CustomerId           string                    `json:"CustomerId"`           // 客户ID
	SystemDiskFeature    string                    `json:"SystemDiskFeature"`    // 系统盘特性
	SupportGpuDriver     string                    `json:"SupportGpuDriver"`     // 支持的GPU驱动
	BindingPubnetIp      bool                      `json:"BindingPubnetIp"`      // 是否绑定公网 IP
	EcsGoodsId           string                    `json:"EcsGoodsId"`           // 云主机商品/套餐 ID
	GpuCardInfo          *GpuCardData              `json:"GpuCardInfo"`          // GPU 卡信息
	GpuDriver            string                    `json:"GpuDriver"`            // GPU 驱动版本信息
	IsHaveSnapshot       bool                      `json:"IsHaveSnapshot"`       // 是否有快照
	ProductSource        string                    `json:"ProductSource"`        // 产品来源标识
	ProductSourceDisplay string                    `json:"ProductSourceDisplay"` // 产品来源显示名
	SecurityGroup        []interface{}             `json:"SecurityGroup"`        // 安全组
	SubnetId             string                    `json:"SubnetId"`             // 子网 ID
	Tag                  []interface{}             `json:"Tag"`                  // 标签
}

// GpuCardData GPU 卡数据信息
type GpuCardData struct {
	GicBusinessName string   `json:"GicBusinessName"` // GPU 型号的名称
	GpuDriver       string   `json:"GpuDriver"`       // GPU 驱动版本
	GpuTypeId       string   `json:"GpuTypeId"`       // GPU 类型 ID
	RealName        []string `json:"RealName"`        // GPU 型号的真实名称
}

// EipSimpleInfo EIP简要信息
type EipSimpleInfo struct {
	ConfName string `json:"ConfName"` // 网络带宽运营商
	EipIp    string `json:"EipIp"`    // EIP地址
}

// OperateInstanceReq OperateInstance接口请求参数
type OperateInstanceReq struct {
	EcsIds    []string `json:"EcsIds"`    // 实例ID列表
	OpType    operate  `json:"OpType"`    // 操作类型
	DeleteEip int      `json:"DeleteEip"` // 公网释放选项，仅在关机不计费情况生效(0:保留公网IP,1:释放公网IP)
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

// OperateInstanceResult OperateInstance接口响应结果
type OperateInstanceResult struct {
	OpenApiCommonResp
	Data *EventIdData `json:"Data"` // 事件ID数据
}

// ModifyInstanceNameReq ModifyInstanceName接口请求参数
type ModifyInstanceNameReq struct {
	EcsId string `json:"EcsId"` // 实例ID
	Name  string `json:"Name"`  // 新实例名称
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

// ModifyInstanceNameResult ModifyInstanceName接口响应结果
type ModifyInstanceNameResult struct {
	OpenApiCommonResp
	Data *ModifyInstanceNameData `json:"Data"` // 修改实例名称数据
}

// ModifyInstanceNameData 修改实例名称数据
type ModifyInstanceNameData struct {
	EcsId string `json:"EcsId"` // 实例ID
	Name  string `json:"Name"`  // 实例名称
}

// DescribeInstanceReq DescribeInstance接口请求参数
type DescribeInstanceReq struct {
	EcsId string `json:"EcsId"` // 实例ID
}

func (req *DescribeInstanceReq) check() error {
	if req.EcsId == "" {
		return fmt.Errorf("field EcsId is required")
	}
	return nil
}

// DescribeInstanceResult DescribeInstance接口响应结果
type DescribeInstanceResult struct {
	OpenApiCommonResp
	Data *InstanceData `json:"Data"` // 实例详细数据
}

// InstanceData 实例详细数据
type InstanceData struct {
	EcsId         string `json:"EcsId"`         // 实例ID
	EcsName       string `json:"EcsName"`       // 实例名称
	RegionId      string `json:"RegionId"`      // 区域ID
	RegionName    string `json:"RegionName"`    // 区域名称
	AzId          string `json:"AzId"`          // 可用区ID
	AzName        string `json:"AzName"`        // 可用区名称
	Status        string `json:"Status"`        // 实例状态
	StatusDisplay string `json:"StatusDisplay"` // 实例状态显示名称
	CreateTime    string `json:"CreateTime"`    // 创建时间
	Duration      int    `json:"Duration"`      // 运行时长
	EndBillTime   string `json:"EndBillTime"`   // 计费结束时间
	IsAutoRenewal string `json:"IsAutoRenewal"` // 是否自动续费
	TimeZone      string `json:"TimeZone"`      // 时区
	//IsRam               bool              `json:"IsRam"`               // 是否为内存优化型
	IsGpu               bool                 `json:"IsGpu"`               // 是否为 GPU 实例
	IsToMonth           int                  `json:"IsToMonth"`           // 是否按月计费
	RealCardName        string               `json:"RealCardName"`        // GPU 实际型号
	SecurityGroup       []*SecurityGroupInfo `json:"SecurityGroup"`       // 安全组
	StockRelease        bool                 `json:"StockRelease"`        // 库存释放标志
	Supplier            string               `json:"Supplier"`            // 实例供应商
	SystemDiskFeature   string               `json:"SystemDiskFeature"`   // 系统盘类型
	Tag                 []interface{}        `json:"Tag"`                 // 标签
	NoChargeForShutdown int                  `json:"NoChargeForShutdown"` // 关机是否收费
	EcsRule             *InstanceRuleInfo    `json:"EcsRule"`             // 实例规格信息
	OsInfo              *OsInfo              `json:"OsInfo"`              // 操作系统信息
	Disk                *DiskInfo            `json:"Disk"`                // 磁盘信息
	Pipe                *PipeInfo            `json:"Pipe"`                // 网络信息
	BillingInfo         *BillingInfo         `json:"BillingInfo"`         // 计费信息
}

// SecurityGroupInfo 安全组信息
type SecurityGroupInfo struct {
	GroupInterconnected bool   `json:"GroupInterconnected"` // 安全组内是否可以实例互通
	Priority            int    `json:"Priority"`            // 优先级
	SecurityGroupId     string `json:"SecurityGroupId"`     // 安全组ID
	SecurityGroupName   string `json:"SecurityGroupName"`   // 安全组名称
	SecurityGroupType   string `json:"SecurityGroupType"`   // 安全组类型
}

// InstanceRuleInfo 实例规格信息
type InstanceRuleInfo struct {
	Name       string `json:"Name"`       // 规格名称
	CpuNum     int    `json:"CpuNum"`     // CPU核心数
	CpuUnit    string `json:"CpuUnit"`    // CPU单位
	Ram        int    `json:"Ram"`        // 内存大小
	Gpu        int    `json:"Gpu"`        // GPU数量
	RamUnit    string `json:"RamUnit"`    // 内存单位
	GpuUnit    string `json:"GpuUnit"`    // GPU单位
	CategoryId string `json:"CategoryId"` // 规格类别ID
	ConfName   string `json:"ConfName"`   // 配置名称
	EcsGoodsId string `json:"EcsGoodsId"` // 商品ID
}

// OsInfo 操作系统信息
type OsInfo struct {
	ImageId   string `json:"ImageId"`   // 镜像ID
	ImageName string `json:"ImageName"` // 镜像名称
	OsType    string `json:"OsType"`    // 操作系统类型
	Bit       int    `json:"Bit"`       // 操作系统位数
	Version   string `json:"Version"`   // 操作系统版本
	Unit      string `json:"Unit"`      // 单位
}

// DiskInfo 磁盘信息
type DiskInfo struct {
	SystemDiskConf *SystemDiskConfInfo `json:"SystemDiskConf"` // 系统盘配置信息
	DataDiskConf   []*DataDiskConfInfo `json:"DataDiskConf"`   // 数据盘配置信息列表
}

// SystemDiskConfInfo 系统盘配置信息
type SystemDiskConfInfo struct {
	ReleaseWithInstance int    `json:"ReleaseWithInstance"` // 是否随实例释放(0:否,1:是)
	DiskType            string `json:"DiskType"`            // 磁盘类型(system: 系统盘, data: 数据盘)
	Name                string `json:"Name"`                // 磁盘名称(用户自定义或默认名称)
	Size                int    `json:"Size"`                // 磁盘大小(GB)
	DiskIops            int    `json:"DiskIops"`            // 磁盘IOPS
	BandMbps            int    `json:"BandMbps"`            // 磁盘带宽(Mbps)
	Unit                string `json:"Unit"`                // 单位
	DiskId              string `json:"DiskId"`              // 磁盘ID
	DiskFeature         string `json:"DiskFeature"`         // 磁盘特性
	DiskName            string `json:"DiskName"`            // 磁盘显示名称(例如“SSD云盘”)
	EbsGoodsId          string `json:"EbsGoodsId"`          // 磁盘商品 ID
	EcsGoodsId          string `json:"EcsGoodsId"`          // 云主机商品 ID
	IsFollowDelete      int    `json:"IsFollowDelete"`      // 是否随实例删除(0: 否, 1: 是)
}

// DataDiskConfInfo 数据盘配置信息
type DataDiskConfInfo struct {
	ReleaseWithInstance int    `json:"ReleaseWithInstance"` // 是否随实例释放(0:否,1:是)
	DiskType            string `json:"DiskType"`            // 磁盘类型(system: 系统盘, data: 数据盘)
	Name                string `json:"Name"`                // 磁盘名称(用户自定义或默认名称)
	Size                int    `json:"Size"`                // 磁盘大小(GB)
	DiskIops            int    `json:"DiskIops"`            // 磁盘IOPS
	BandMbps            int    `json:"BandMbps"`            // 磁盘带宽(Mbps)
	Unit                string `json:"Unit"`                // 单位
	Id                  string `json:"Id"`                  // 磁盘ID
	DiskFeature         string `json:"DiskFeature"`         // 磁盘特性
	DiskName            string `json:"DiskName"`            // 磁盘显示名称(例如“SSD云盘”)
	EbsGoodsId          string `json:"EbsGoodsId"`          // 磁盘商品 ID
	EcsGoodsId          string `json:"EcsGoodsId"`          // 云主机商品 ID
	IsFollowDelete      int    `json:"IsFollowDelete"`      // 是否随实例删除(0: 否, 1: 是)
}

// PipeInfo 网络信息
type PipeInfo struct {
	VpcName    string              `json:"VpcName"`    // VPC名称
	VpcId      string              `json:"VpcId"`      // VPC ID
	SubnetId   string              `json:"SubnetId"`   // 子网ID
	SubnetName string              `json:"SubnetName"` // 子网名称
	CreateTime string              `json:"CreateTime"` // 创建时间
	PrivateNet string              `json:"PrivateNet"` // 私网IP
	PubNet     string              `json:"PubNet"`     // 公网IP
	VirtualNet []interface{}       `json:"VirtualNet"` // 虚拟网络信息
	EipInfo    map[string]*EipInfo `json:"EipInfo"`    // EIP信息
}

// EipInfo EIP信息
type EipInfo struct {
	ConfName      string `json:"ConfName"`      // 配置名称
	EipIp         string `json:"EipIp"`         // EIP地址
	CurrentUseQos int    `json:"CurrentUseQos"` // 当前使用带宽
	BandwidthName string `json:"BandwidthName"` // 带宽名称
}

// BillingInfo 计费信息
type BillingInfo struct {
	BillingMethod       string `json:"BillingMethod"`       // 计费方式
	BillingMethodName   string `json:"BillingMethodName"`   // 计费方式名称
	BillingStatus       string `json:"BillingStatus"`       // 计费状态
	BillCycleId         string `json:"BillCycleId"`         // 计费周期ID
	BillingMethodStatus string `json:"BillingMethodStatus"` // 计费方式状态
}

// DescribeTaskEventReq DescribeTaskEvent接口请求参数
type DescribeTaskEventReq struct {
	EventId string `json:"EventId"` // 事件ID
}

func (req *DescribeTaskEventReq) check() error {
	if req.EventId == "" {
		return fmt.Errorf("field EventId is required")
	}
	return nil
}

// DescribeTaskEventResult DescribeTaskEvent接口响应结果
type DescribeTaskEventResult struct {
	OpenApiCommonResp
	Data *EventResultData `json:"Data"` // 事件结果数据
}

// EventResultData 事件结果数据
type EventResultData struct {
	EventId            string          `json:"EventId"`            // 事件ID
	EventStatus        string          `json:"EventStatus"`        // 事件状态
	EventStatusDisplay string          `json:"EventStatusDisplay"` // 事件状态显示名称
	EventType          string          `json:"EventType"`          // 事件类型
	EventTypeDisplay   string          `json:"EventTypeDisplay"`   // 事件类型显示名称
	CreateTime         string          `json:"CreateTime"`         // 创建时间
	TaskList           []*TaskListInfo `json:"TaskList"`           // 任务列表
}

// TaskListInfo 任务列表信息
type TaskListInfo struct {
	TaskId          string `json:"TaskId"`          // 任务ID
	Status          string `json:"Status"`          // 任务状态
	StatusDisplay   string `json:"StatusDisplay"`   // 任务状态显示名称
	ResourceId      string `json:"ResourceId"`      // 资源ID
	CreateTime      string `json:"CreateTime"`      // 创建时间
	UpdateTime      string `json:"UpdateTime"`      // 更新时间
	EndTime         string `json:"EndTime"`         // 结束时间
	ResourceType    string `json:"ResourceType"`    // 资源类型
	ResourceDisplay string `json:"ResourceDisplay"` // 资源显示名称
	TaskType        string `json:"TaskType"`        // 任务类型
	TaskTypeDisplay string `json:"TaskTypeDisplay"` // 任务类型显示名称
}

// DescribeEcsFamilyInfoReq DescribeEcsFamilyInfo接口请求参数
type DescribeEcsFamilyInfoReq struct {
	AvailableZoneCode string        `json:"AvailableZoneCode"` // 可用区代码
	BillingMethod     billingMethod `json:"BillingMethod"`     // 计费方式
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

// DescribeEcsFamilyInfoResult DescribeEcsFamilyInfo接口响应结果
type DescribeEcsFamilyInfoResult struct {
	OpenApiCommonResp
	Data FamilyData `json:"Data"` // 规格族数据
}

// FamilyData 规格族数据
type FamilyData struct {
	EcsFamilyInfo []*FamilyInfo `json:"EcsFamilyInfo"` // ECS规格族信息列表
}

// FamilyInfo 规格族信息
type FamilyInfo struct {
	EcsFamilyName string          `json:"EcsFamilyName"` // ECS规格族名称
	SpecList      []*SpecListInfo `json:"SpecList"`      // 规格列表
}

// SpecListInfo 规格列表信息
type SpecListInfo struct {
	GpuShowType      string `json:"GpuShowType"`      // GPU显示类型
	GpuTypeId        string `json:"GpuTypeId"`        // GPU类型ID
	Cpu              int    `json:"Cpu"`              // CPU核心数
	Ram              int    `json:"Ram"`              // 内存大小(GB)
	Gpu              int    `json:"Gpu"`              // GPU数量
	SupportGpuDriver string `json:"SupportGpuDriver"` // 支持的GPU驱动
}

// ChangeInstanceConfigureReq ChangeInstanceConfigure接口请求参数
type ChangeInstanceConfigureReq struct {
	EcsIds            []string `json:"EcsIds"`            // 实例ID列表
	AvailableZoneCode string   `json:"AvailableZoneCode"` // 可用区代码
	EcsFamilyName     string   `json:"EcsFamilyName"`     // 实例规格族名称
	Cpu               int      `json:"Cpu"`               // CPU核心数
	Ram               int      `json:"Ram"`               // 内存大小(GB)
	Gpu               int      `json:"Gpu"`               // GPU数量
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

// ChangeInstanceConfigureResult ChangeInstanceConfigure接口响应结果
type ChangeInstanceConfigureResult struct {
	OpenApiCommonResp
	Data *EventIdData `json:"Data"` // 事件ID数据
}

// ExtendDiskReq ExtendDisk接口请求参数
type ExtendDiskReq struct {
	DiskId       string `json:"DiskId"`       // 磁盘ID
	ExtendedSize int    `json:"ExtendedSize"` // 扩展后的大小(GB)
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

// ExtendDiskResult ExtendDisk接口响应结果
type ExtendDiskResult struct {
	OpenApiCommonResp
	Data *EventIdData `json:"Data"` // 事件ID数据
}
