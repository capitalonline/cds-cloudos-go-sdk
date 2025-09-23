package ecs

import (
	"fmt"
)

const (
	ActionDescribeRegions         = "DescribeRegions"
	ActionCreateInstance          = "CreateInstance"
	ActionDeleteInstance          = "DeleteInstance"
	ActionModifyInstancePassword  = "ModifyInstancePassword"
	ActionDescribeInstanceStatus  = "DescribeInstanceStatus"
	ActionDescribeInstanceList    = "DescribeInstanceList"
	ActionOperateInstance         = "OperateInstance"
	ActionModifyInstanceName      = "ModifyInstanceName"
	ActionDescribeInstance        = "DescribeInstance"
	ActionDescribeTaskEvent       = "DescribeEvent"
	ActionDescribeEcsFamilyInfo   = "DescribeEcsFamilyInfo"
	ActionChangeInstanceConfigure = "ChangeInstanceConfigure"
	ActionDescribeImages          = "DescribeImages"
)

const (
	ActionExtendDisk = "ExtendDisk"
)

type (
	operate       string
	billingMethod string

	DiskFeature   string
	bandwidthType string
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

const (
	LocalDiskFeature DiskFeature = "local"
	SsdDiskFeature   DiskFeature = "ssd"
)

const (
	Bandwidth      bandwidthType = "Bandwidth"      // 固定带宽
	BandwidthMonth bandwidthType = "BandwidthMonth" // 固定带宽包月
	Traffic        bandwidthType = "Traffic"        // 流量按需
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
	imageIdsKey      = "ImageIds"
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

// CreateInstanceReq CreateInstance接口请求参数
type CreateInstanceReq struct {
	Name              string                             `json:"Name"`                     // 云服务器名,不传自动赋予（自动命名规则：ecs-创建日期）
	Password          string                             `json:"Password"`                 // 登录密码
	AvailableZoneCode string                             `json:"AvailableZoneCode"`        // 可用区代码
	EcsFamilyName     string                             `json:"EcsFamilyName"`            // 规格族名称
	UtcTime           int                                `json:"UtcTime"`                  // 是否utc时间，1:是  0:否 默认为0（UTC+8，上海时间）
	Cpu               int                                `json:"Cpu"`                      // Cpu
	Ram               int                                `json:"Ram"`                      // 内存
	Gpu               int                                `json:"Gpu,omitempty"`            // 显卡数量，默认为0
	Number            int                                `json:"Number"`                   // 购买数量，默认为1（默认批量最大值为100台）
	BillingMethod     billingMethod                      `json:"BillingMethod"`            // 计费方式："0": 按需  "1":包年包月
	ImageId           string                             `json:"ImageId"`                  // 镜像id或者镜像名称
	SystemDisk        *CreateInstanceDiskData            `json:"SystemDisk"`               // 系统盘信息
	DataDisk          []*CreateInstanceDiskData          `json:"DataDisk,omitempty"`       // 数据盘信息
	VpcInfo           *CreateInstanceVpcInfo             `json:"VpcInfo"`                  // vpc信息
	SubnetInfo        *CreateInstanceSubnetInfo          `json:"SubnetInfo"`               // 私有网络信息
	SecurityGroups    []*CreateInstanceSecurityGroupData `json:"SecurityGroups,omitempty"` // 安全组列表，安全组优先级按顺序由高到低
	StartNumber       int                                `json:"StartNumber,omitempty"`    // 云服务器名称编号起始数字，不需要服务器编号可不传
	Duration          int                                `json:"Duration,omitempty"`       // 只在包月算价时有意义，以月份为单位，一年值为12，大于一年要输入12的整数倍，最大值36(3年)
	IsToMonth         *int                               `json:"IsToMonth,omitempty"`      // 包月是否到月底 1:是  0:否 默认为1。
	DnsList           *[2]string                         `json:"DnsList,omitempty"`        // dns 解析 需要两个元素  [主dns，从dns]，不选采用默认通用DNS
	PubnetInfo        []*CreateInstancePubnetInfo        `json:"PubnetInfo,omitempty"`     // 公有网络信息,支持新分配公网IP和绑定已有的公网IP
	TestAccount       *string                            `json:"TestAccount,omitempty"`    // 测试账户名称
	IsAutoRenewal     *int                               `json:"IsAutoRenewal,omitempty"`  // 是否自动续约，包月时需传。1:是  0:否 默认为1
}

func (req *CreateInstanceReq) check() error {
	// 必填参数检查
	if req.AvailableZoneCode == "" {
		return fmt.Errorf("field AvailableZoneCode is required")
	}
	if req.EcsFamilyName == "" {
		return fmt.Errorf("field EcsFamilyName is required")
	}
	if req.Cpu <= 0 {
		return fmt.Errorf("field Cpu must be > 0")
	}
	if req.Ram <= 0 {
		return fmt.Errorf("field Ram must be > 0")
	}
	if req.BillingMethod != MonthlyBillingMethod && req.BillingMethod != OnDemandBillingMethod {
		return fmt.Errorf("field BillingMethod must be '0' or '1'")
	}
	if req.Password == "" {
		return fmt.Errorf("field Password is required")
	}
	if req.ImageId == "" {
		return fmt.Errorf("field ImageId is required")
	}
	if req.SystemDisk == nil {
		return fmt.Errorf("field SystemDisk is required")
	}
	if err := req.SystemDisk.check(true, req.BillingMethod); err != nil {
		return fmt.Errorf("invalid SystemDisk: %v", err)
	}
	if req.VpcInfo == nil {
		return fmt.Errorf("field VpcInfo is required")
	}
	if req.SubnetInfo == nil {
		return fmt.Errorf("field SubnetInfo is required")
	}
	// 可选字段检查
	if req.Gpu < 0 {
		req.Gpu = 0
	}
	if req.Number <= 0 {
		req.Number = 1 // 默认值 1
	}
	if req.Number > 100 {
		return fmt.Errorf("field Number maximum value is 100")
	}
	if req.SubnetInfo == nil || req.SubnetInfo.SubnetId == "" {
		return fmt.Errorf("field SubnetInfo.SubnetId is required")
	}
	for i, p := range req.PubnetInfo {
		if err := p.check(req.Number, req.BillingMethod); err != nil {
			return fmt.Errorf("invalid PubnetInfo[%d]: %v", i, err)
		}
	}
	if req.BillingMethod == "1" { // 包年包月
		if req.Duration > 36 {
			return fmt.Errorf("field Duration maximum value is 36")
		}
		if req.Duration > 12 && req.Duration%12 != 0 {
			return fmt.Errorf("field Duration must be a multiple of 8 when greater than 12")
		}
	}
	for i, d := range req.DataDisk {
		if err := d.check(false, req.BillingMethod); err != nil {
			return fmt.Errorf("invalid DataDisk[%d]: %v", i, err)
		}
	}
	return nil
}

// CreateInstanceDiskData 系统盘信息
type CreateInstanceDiskData struct {
	DiskFeature         DiskFeature `json:"DiskFeature"`                   // 盘类型:本地盘:"local", 云盘:"ssd"
	Size                int         `json:"Size"`                          // 磁盘大小，单位为 GB
	SnapshotId          string      `json:"SnapshotId,omitempty"`          // 创建磁盘时使用的快照 ID
	ReleaseWithInstance *int        `json:"ReleaseWithInstance,omitempty"` // 磁盘是否随实例释放
}

func (d *CreateInstanceDiskData) check(isSystemDisk bool, BillingMethod billingMethod) error {
	if d.DiskFeature == "" {
		return fmt.Errorf("field DiskFeature is required")
	}

	if d.Size <= 0 {
		return fmt.Errorf("field Size must be > 0")
	}

	if isSystemDisk {
		// 系统盘不能传 SnapshotId 和 ReleaseWithInstance
		if d.SnapshotId != "" {
			return fmt.Errorf("field SnapshotId is not allowed in SystemDisk")
		}
	} else {
		if BillingMethod == MonthlyBillingMethod {
			if d.ReleaseWithInstance != nil && *d.ReleaseWithInstance != 0 {
				return fmt.Errorf("MonthlyBillingMethod do not support ReleaseWithInstance")
			}
		}
	}
	return nil
}

// CreateInstanceVpcInfo vpc信息
type CreateInstanceVpcInfo struct {
	VpcId string `json:"VpcId"` // 私有网络id
}

// CreateInstanceSubnetInfo 私有网络信息
type CreateInstanceSubnetInfo struct {
	SubnetId  string   `json:"SubnetId"`            // 子网id
	IpAddress []string `json:"IpAddress,omitempty"` // 指定私网IP列表,列表中的IP个数与创建云主机个数一致
}

// CreateInstanceSecurityGroupData 安全组列表
type CreateInstanceSecurityGroupData struct {
	SecurityGroupId string `json:"SecurityGroupId"` // 安全组id
}

// CreateInstancePubnetInfo 公有网络信息
type CreateInstancePubnetInfo struct {
	SubnetId          string        `json:"SubnetId"`                // 子网id
	BandwidthConfName string        `json:"BandwidthConfName"`       // 带宽线路名称.使用新创建的vpp网络需要指定线路名称.例如：电信、联通. 带宽线路名称参考VPCBandWidthBillingScheme获取
	IpType            string        `json:"IpType"`                  // 历史遗留参数，可不写。若使用虚拟出网网关必填。默认出网网关:"default_gateway",虚拟网关：”virtual”
	EipIds            []string      `json:"EipIds,omitempty"`        // 选填,绑定的eip的id列表;若需新分配公网IP,不填,绑定已有公网IP需填,数量需要和云服务器数量一致
	BandwidthType     bandwidthType `json:"BandwidthType,omitempty"` // 带宽类型;若需新分配公网IP必填,表示绑定公网IP的带宽类型.绑定已有公网IP不填.固定带宽:”Bandwidth”,固定带宽包月:”BandwidthMonth”,流量按需: “Traffic”（若实例计费方式为包年包月选择固定带宽时需传"固定带宽包月"）
	Qos               int           `json:"Qos,omitempty"`           // 公网带宽值,单位为M;若带宽类型选择”固定带宽”需填写
}

func (p *CreateInstancePubnetInfo) check(instanceCount int, billingMethod billingMethod) error {
	if p.SubnetId == "" {
		return fmt.Errorf("field SubnetId is required")
	}

	if len(p.EipIds) > 0 {
		// 绑定已有公网IP
		if len(p.EipIds) != instanceCount {
			return fmt.Errorf("field EipIds must contain %d elements", instanceCount)
		}

	} else {
		// 新分配公网IP
		if p.BandwidthType == "" {
			return fmt.Errorf("field BandwidthType is required when assigning new public IP")
		}

		switch p.BandwidthType {
		case Bandwidth:
			if p.Qos <= 0 {
				return fmt.Errorf("field Qos must be > 0 when BandwidthType=Bandwidth")
			}
		case BandwidthMonth:
			if billingMethod != MonthlyBillingMethod { // 包年包月
				return fmt.Errorf("field BandwidthType=BandwidthMonth is only valid for yearly/monthly billing")
			}
		case Traffic:
		default:
			return fmt.Errorf("field BandwidthType has invalid value: %s", p.BandwidthType)
		}
	}

	return nil
}

// CreateInstanceResult CreateInstance接口响应结果
type CreateInstanceResult struct {
	OpenApiCommonResp
	Data *CreateInstanceEvent `json:"Data"` // 创建实例id信息
}

// CreateInstanceEvent 创建实例id信息
type CreateInstanceEvent struct {
	EventId  string   `json:"EventId"`  // 事件id
	EcsIdSet []string `json:"EcsIdSet"` // 创建的资源id列表
}

// DeleteInstanceReq DeleteInstance接口请求参数
type DeleteInstanceReq struct {
	EcsIds    []string `json:"EcsIds"`    // 云服务器id列表
	DeleteEip int      `json:"DeleteEip"` // 1:解绑并删除服务器绑定的EIP，0:解绑EIP  默认为0
}

func (req *DeleteInstanceReq) check() error {
	if len(req.EcsIds) == 0 {
		return fmt.Errorf("field EcsIds is required")
	}

	if req.DeleteEip != 0 && req.DeleteEip != 1 {
		return fmt.Errorf("field DeleteEip must be 0 or 1")
	}

	return nil
}

// DeleteInstanceResult DeleteInstance接口响应结果
type DeleteInstanceResult struct {
	OpenApiCommonResp
	Data *EventIdData `json:"Data"` // 事件ID数据
}

// ModifyInstancePasswordReq ModifyInstancePassword接口请求参数
type ModifyInstancePasswordReq struct {
	EcsIds   []string `json:"EcsIds"`   // 云服务器id列表
	Password string   `json:"Password"` // 新密码
}

func (req *ModifyInstancePasswordReq) check() error {
	if len(req.EcsIds) == 0 {
		return fmt.Errorf("field EcsIds is required")
	}

	if req.Password == "" {
		return fmt.Errorf("field Password is required")
	}

	return nil
}

// ModifyInstancePasswordResult ModifyInstancePassword接口响应结果
type ModifyInstancePasswordResult struct {
	OpenApiCommonResp
	Data *EventIdData `json:"Data"` // 事件ID数据
}

// DescribeInstanceStatusReq DescribeInstanceStatus接口请求参数
type DescribeInstanceStatusReq struct {
	EcsIds []string `json:"EcsIds"` // 云服务器列表
}

func (req *DescribeInstanceStatusReq) check() error {
	if len(req.EcsIds) == 0 {
		return fmt.Errorf("field EcsIds is required")
	}

	return nil
}

// DescribeInstanceStatusResult DescribeInstanceStatus接口响应结果
type DescribeInstanceStatusResult struct {
	OpenApiCommonResp
	Data *InstanceEcsStatusInfo `json:"Data"` // 云服务器状态信息
}

// InstanceEcsStatusInfo 云服务器状态字典
type InstanceEcsStatusInfo struct {
	EcsStatus map[string]*InstanceEcsStatusData `json:"EcsStatus"` // 云服务器状态字典，key为实例id
}

// InstanceEcsStatusData 云服务器状态信息
type InstanceEcsStatusData struct {
	Status        string `json:"Status"`        // 状态码
	StatusDisplay string `json:"StatusDisplay"` // 状态
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
	SystemDiskFeature    string                    `json:"SystemDiskFeature"`    // 系统盘类型，本地盘:"local", 云盘:"ssd"
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
	SystemDiskFeature   string               `json:"SystemDiskFeature"`   // 系统盘类型，本地盘:"local", 云盘:"ssd"
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
	DiskFeature         string `json:"DiskFeature"`         // 磁盘类型，本地盘:"local", 云盘:"ssd"
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
	DiskFeature         string `json:"DiskFeature"`         // 磁盘类型，本地盘:"local", 云盘:"ssd"
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

// DescribeTaskEventReq DescribeTaskEvent 接口请求参数
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
	if req.BillingMethod == "" {
		return fmt.Errorf("field BillingMethod is required")
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

// DescribeImagesReq DescribeImages接口请求参数
type DescribeImagesReq struct {
	AvailableZoneCode string   `json:"AvailableZoneCode"` // 可用区代码
	ImageIds          []string `json:"ImageIds"`          // 镜像id列表
}

// DescribeImagesResult DescribeImages接口响应结果
type DescribeImagesResult struct {
	OpenApiCommonResp
	Data *ImageList `json:"Data"` // 镜像列表
}

// ImageList 镜像列表
type ImageList struct {
	ImageList []*ImageInfo `json:"ImageList"` // 镜像信息
}

// ImageInfo 镜像信息
type ImageInfo struct {
	AvailableZoneCode string   `json:"AvailableZoneCode"` // 可用区code
	AzId              string   `json:"AzId"`              // 可用区ID
	AzName            string   `json:"AzName"`            // 可用区名称
	CreateTime        string   `json:"CreateTime"`        // 创建时间
	ImageId           string   `json:"ImageId"`           // 镜像id
	ImageName         string   `json:"ImageName"`         // 镜像名称
	IsOptimized       int      `json:"IsOptimized"`       // 镜像是否开启优化选项，1为开启，0为不开启
	OsBit             int      `json:"OsBit"`             // 系统位数
	OsSize            int      `json:"OsSize"`            // 镜像容量(GB)
	OsType            string   `json:"OsType"`            // 镜像类型
	OsVersion         string   `json:"OsVersion"`         // 镜像版本
	Status            string   `json:"Status"`            // 镜像状态code
	StatusDisplay     string   `json:"StatusDisplay"`     // 镜像状态中文
	SupportGpuDriver  string   `json:"SupportGpuDriver"`  // 支持的GPU驱动类型
	SupportType       []string `json:"SupportType"`       // 支持类型
	TemplateType      string   `json:"TemplateType"`      // 公共镜像为public，私有镜像为private
	Username          string   `json:"Username"`          // 用户名称
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
