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

package eks

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

type CommonScalingGroupTask struct {
	TaskId string `json:"TaskId"`
}

type GetScalingGroupDetailResult struct {
	Data ScalingGroupDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ScalingGroupDetail struct {
	Cpu              int    `json:"Cpu"`
	Ram              int    `json:"Ram"`
	Gpu              int    `json:"Gpu"`
	GpuShowType      string `json:"GpuShowType"`
	MaxSize          int    `json:"MaxSize"`
	MinSize          int    `json:"MinSize"`
	ActivitySize     int    `json:"ActivitySize"`
	PendingSize      int    `json:"PendingSize"`
	RemovingSize     int    `json:"RovingSize"`
	TotalScalingSize int    `json:"TotalScalingSize"`
	ScalingGroupId   string `json:"ScalingGroupId"`
	Status           string `json:"Status"`
	StatusStr        string `json:"StatusStr"`
}

type AddScalingGroupNodeReq struct {
	ScalingGroupId string `json:"ScalingGroupId"`
	AddNum         int    `json:"AddNum"`
}

type AddScalingGroupNodeResult struct {
	Data CommonScalingGroupTask `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ListClustersReq struct {
	Keyword string `json:"Keyword,omitempty"`
	VpcId   string `json:"VpcId,omitempty"`
	//Status  string `json:"Status,omitempty"`
	//Version string `json:"Version,omitempty"`
}

type ListClustersResult struct {
	Data []ListClustersDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ListClustersDetail struct {
	ClusterId     string `json:"ClusterId"`
	ClusterIp     string `json:"ClusterIp"`
	ClusterName   string `json:"ClusterName"`
	ClusterStatus string `json:"ClusterStatus"`
	CreateTime    string `json:"CreateTime"`
	K8SVersion    string `json:"K8sVersion"`
	NodeSum       int    `json:"NodeSum"`
	RegionId      string `json:"RegionId"`
	RegionName    string `json:"RegionName"`
	SlbId         string `json:"SlbId"`
	SshPort       int    `json:"SshPort"`
	SubDomain     string `json:"SubDomain"`
	UpdateTime    string `json:"UpdateTime"`
	Vip           string `json:"Vip"`
	VpcId         string `json:"VpcId"`
	VpcName       string `json:"VpcName"`
}

type GetClusterEventsResult struct {
	Data []GetClusterEventsDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type GetClusterEventsDetail struct {
	Status             string `json:"Status"`
	SubtaskElapsedTime string `json:"SubtaskElapsedTime"`
	SubtaskName        string `json:"SubtaskName"`
	SubtaskStartTime   string `json:"SubtaskStartTime"`
	SubtaskType        string `json:"SubtaskType"`
}

type DeleteClusterResult struct {
	Data DeleteClusterData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type DeleteClusterData struct {
	ClusterId string `json:"ClusterId"`
	TaskId    string `json:"TaskId"`
}

type AttachNetworkInterfaceReq struct {
	EcsId     string `json:"EcsId"`
	NetcardId string `json:"NetcardId"`
	VlanId    string `json:"VlanId"`
}

type AttachNetworkInterfaceResult struct {
	Data AttachNetworkInterfaceData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type AttachNetworkInterfaceData struct {
	EventId string `json:"EventId"`
}

type DetachNetworkInterfaceReq struct {
	EcsId     string `json:"EcsId"`
	NetcardId string `json:"NetcardId"`
}

type DetachNetworkInterfaceResult struct {
	Data DetachNetworkInterfaceData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type DetachNetworkInterfaceData struct {
	EventId string `json:"EventId"`
}

type DescribeNetworkInterfaceResult struct {
	Data DescribeNetworkInterfaceData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type DescribeNetworkInterfaceData struct {
	NetcardId   string `json:"NetcardId"`
	MacAddress  string `json:"MacAddress"`
	SubnetId    string `json:"SubnetId"`
	EcsId       string `json:"EcsId"`
	IsValid     bool   `json:"IsValid"`
	IpAddress   string `json:"IpAddress"`
	Mask        int    `json:"Mask"`
	NetcardName string `json:"NetcardName"`
}

type IsAttachedECSReq struct {
	NetcardId string `json:"NetcardId"`
}

type IsAttachedECSResult struct {
	Data IsAttachedECSData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type IsAttachedECSData struct {
	Attached bool `json:"Attached"`
}

type QueryTaskStatusResult struct {
	Data QueryTaskStatusData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type QueryTaskStatusData struct {
	EventId            string                       `json:"EventId"`
	EventStatus        string                       `json:"EventStatus"`
	EventStatusDisplay string                       `json:"EventStatusDisplay"`
	EventType          string                       `json:"EventType"`
	EventTypeDisplay   string                       `json:"EventTypeDisplay"`
	CreateTime         string                       `json:"CreateTime"`
	TaskList           []QueryTaskStatusDataSubtask `json:"TaskList"`
}

type QueryTaskStatusDataSubtask struct {
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

type GetClusterResult struct {
	Data GetClusterDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type TaskStatusResult struct {
	Data TaskStatusDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type TaskStatusDetail struct {
	TaskId     string `json:"TaskId"`
	TaskMsg    string `json:"TaskMsg"`
	TaskStatus string `json:"TaskStatus"`
}

type GetClusterDetail struct {
	ClusterId       string       `json:"ClusterId"`
	Ak              string       `json:"Ak"`
	ClusterIp       string       `json:"ClusterIp"`
	VpcId           string       `json:"VpcId"`
	ClusterName     string       `json:"ClusterName"`
	ClusterStatus   string       `json:"ClusterStatus"`
	CreateTime      string       `json:"CreateTime"`
	CustomerId      string       `json:"CustomerId"`
	DashboardAddr   string       `json:"DashboardAddr"`
	K8sCaExpiryData int          `json:"K8sCaExpiryData"`
	K8sVersion      string       `json:"K8sVersion"`
	MasterSum       int          `json:"MasterSum"`
	NatId           string       `json:"NatId"`
	NatName         string       `json:"NatName"`
	NodeSum         int          `json:"NodeSum"`
	OsName          string       `json:"OsName"`
	RegionId        string       `json:"RegionId"`
	RegionName      string       `json:"RegionName"`
	RuntimeType     string       `json:"RuntimeType"`
	SlbId           string       `json:"SlbId"`
	SshPort         int          `json:"SshPort"`
	StatusStr       string       `json:"StatusStr"`
	SubDomain       string       `json:"SubDomain"`
	UpdateTime      string       `json:"UpdateTime"`
	UserId          string       `json:"UserId"`
	Vip             string       `json:"Vip"`
	VipId           string       `json:"VipId"`
	VpcName         string       `json:"VpcName"`
	WorkerSum       int          `json:"WorkerSum"`
	CniInfo         CniDetail    `json:"CniInfo"`
	DestinationEip  EipInfo      `json:"DestinationEip"`
	SourceEip       EipInfo      `json:"SourceEip"`
	MasterSubnet    []SubnetInfo `json:"MasterSubnet"`
	WorkerSubnet    []SubnetInfo `json:"WorkerSubnet"`
	SubnetList      []SubnetInfo `json:"SubnetList"`
}

type CniDetail struct {
	CniConfig   CniConfigDetail `json:"CniConfig"`
	CniType     string          `json:"CniType"`
	ProxyConfig string          `json:"ProxyConfig"`
	ServiceCidr string          `json:"ServiceCidr"`
}

type CniConfigDetail struct {
	NodePodsNum int          `json:"NodePodsNum"`
	PodCidr     string       `json:"PodCidr"`
	SubnetList  []SubnetInfo `json:"SubnetList"`
}

type CniInfo struct {
	CniConfig   CniConfig `json:"CniConfig"`
	CniType     string    `json:"CniType"`
	ProxyConfig string    `json:"ProxyConfig"`
	ServiceCidr string    `json:"ServiceCidr"`
}

type CniConfig struct {
	NodePodsNum int      `json:"NodePodsNum"`
	PodCidr     string   `json:"PodCidr"`
	SubnetIds   []string `json:"SubnetIds"`
}

type EipInfo struct {
	EipId string `json:"EipId"`
	Ip    string `json:"Ip"`
}

type SubnetInfo struct {
	AvailableZoneId string `json:"AvailableZoneId"`
	SubnetId        string `json:"SubnetId"`
	Segment         string `json:"Segment"`
}

type ListNodesReq struct {
	ClusterId      string `json:"ClusterId"`
	NodeStatus     string `json:"NodeStatus"`
	NodeType       string `json:"NodeType"`
	SchedulableStr string `json:"SchedulableStr"`
	Status         string `json:"Status"`
	PageSize       int    `json:"PageSize"`
	PageIndex      int    `json:"PageIndex"`
	Keyword        string `json:"Keyword"`
}

type ListNodesResult struct {
	Data []ListNodesData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ListNodesData struct {
	NodeId        string `json:"NodeId"`
	NodeName      string `json:"NodeName"`
	NodeStatus    string `json:"NodeStatus"`
	NodeType      string `json:"NodeType"`
	ClusterId     string `json:"ClusterId"`
	ClusterName   string `json:"ClusterName"`
	CustomerId    string `json:"CustomerId"`
	UserId        string `json:"UserId"`
	VpcId         string `json:"VpcId"`
	SubnetId      string `json:"SubnetId"`
	PrivateIp     string `json:"PrivateIp"`
	Cpu           int    `json:"Cpu"`
	Ram           int    `json:"Ram"`
	Gpu           int    `json:"Gpu"`
	GpuType       string `json:"GpuType"`
	FamilyName    string `json:"FamilyName"`
	Schedulable   int    `json:"Schedulable"`
	K8sStatus     string `json:"K8sStatus"`
	RegionName    string `json:"RegionName"`
	RegionId      string `json:"RegionId"`
	AzName        string `json:"AzName"`
	AzId          string `json:"AzId"`
	SourceType    string `json:"SourceType"`
	BillingMethod string `json:"BillingMethod"`
	Duration      int    `json:"Duration"`
	IsToMonth     int    `json:"IsToMonth"`
	AutoRenew     int    `json:"AutoRenew"`
	CreateTime    string `json:"CreateTime"`
}

type DeleteNodesReq struct {
	ClusterId string   `json:"ClusterId"`
	NodeIds   []string `json:"NodeIds"`
}

type DeleteNodesResult struct {
	Data DeleteNodesData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type DeleteNodesData struct {
	TaskId string `json:"TaskId"`
}

type CreateClusterReq struct {
	ClusterName    string                `json:"ClusterName"`
	VpcId          string                `json:"VpcId"`
	Cni            CniInfo               `json:"Cni"`
	K8sVersion     string                `json:"K8sVersion"`
	RuntimeType    string                `json:"RuntimeType"`
	NatId          string                `json:"NatId"`
	SourceEip      string                `json:"SourceEip"`
	DestinationEip string                `json:"DestinationEip"`
	Password       string                `json:"Password"`
	SlbId          string                `json:"SlbId"`
	Ak             string                `json:"Ak"`
	Sk             string                `json:"Sk"`
	MasterNumber   int                   `json:"MasterNumber"`
	MasterConfig   NodePoolNodeConfig    `json:"MasterConfig"`
	NodePoolConfig NodePoolConfiguration `json:"NodePoolConfig"`
}

type NodePoolDiskInfo struct {
	DiskType string `json:"Type,omitempty"` // 磁盘类型：SSD
	DiskSize int    `json:"Size,omitempty"` // 磁盘大小，单位：GB
}

type CreateClusterData struct {
	ClusterId string `json:"ClusterId"`
	TaskId    string `json:"TaskId"`
}

type CreateClusterResult struct {
	Data CreateClusterData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type AddClusterSubnetReq struct {
	ClusterId  string          `json:"ClusterId"`
	SubnetList []ClusterSubnet `json:"SubnetList"`
}

type ClusterSubnet struct {
	SubnetId string `json:"SubnetId"`
}

type AddClusterSubnetResult struct {
	OpenApiCommonResp
}

// cni类型
const (
	CniTypeFlannel = "flannel"
	CniTypeVpcCni  = "vpc-cni"
)

// k8s版本
const (
	K8sVersion1_30_14 = "v1.30.14"
	K8sVersion1_26_5  = "v1.26.5"
)

// 服务转发模式
const (
	ProxyConfigIptables = "iptables"
	ProxyConfigIPVS     = "ipvs"
)

// master节点数
const (
	MasterNumber3 = 3
	MasterNumber5 = 5
	MasterNumber7 = 7
)

// 集群服务CIDR ip地址段
const (
	ServiceCidr10_16  = "10.0.0.0/16"
	ServiceCidr10_24  = "10.0.0.0/24"
	ServiceCidr192_16 = "192.168.0.0/16"
	ServiceCidr192_24 = "192.168.0.0/24"
	ServiceCidr172_16 = "172.16.0.0/16"
	ServiceCidr172_24 = "172.16.0.0/24"
)

// NodePoolBillingSpec 节点池计费配置
type NodePoolBillingSpec struct {
	BillingMethod string `json:"InstanceChargeType"` // 实例计费类型：PostPaid(按需付费) / PrePaid(包年包月)
	Duration      int    `json:"Period"`             // 实例购买时长，单位：月
	IsToMonth     int    `json:"IsToMonth"`          // 是否购买至当前月底：0(购买整月) / 1(购买至月底)
	AutoRenew     int    `json:"AutoRenew"`          // 是否自动续费：0(不自动续费) / 1(自动续费)
}

// NodePoolNodeConfig 节点池节点配置
type NodePoolNodeConfig struct {
	BillingSpec     NodePoolBillingSpec `json:"BillingSpec"`     // 计费配置
	SystemDisk      NodePoolDiskInfo    `json:"SystemVolume"`    // 系统卷配置
	DataDisk        []NodePoolDiskInfo  `json:"DataVolumes"`     // 数据卷配置
	OsImageName     string              `json:"OsImageName"`     // 操作系统镜像名称
	SubnetIds       []string            `json:"SubnetIds"`       // 子网ID列表
	InstanceTypeIds []string            `json:"InstanceTypeIds"` // 实例类型ID列表
	Password        string              `json:"Password"`        // 节点密码
	Shell           string              `json:"ExecCommand"`     // 节点初始化脚本
	Labels          map[string]string   `json:"Labels"`          // 节点标签
}

// NodePoolConfiguration 节点池配置
type NodePoolConfiguration struct {
	Id         string             `json:"Id,omitempty"` // 节点池ID（更新时使用）
	PoolName   string             `json:"Name"`         // 节点池名称
	NodeType   string             `json:"NodeType"`     // 节点类型：ecs / bms
	SubjectId  int                `json:"SubjectId"`    // 主体ID
	NodeConfig NodePoolNodeConfig `json:"NodeConfig"`   // 节点配置
	Replicas   int                `json:"Replicas"`     // 节点数量
}

type CreateNodePoolReq struct {
	ClusterId string                `json:"ClusterId"` // 集群ID
	VpcId     string                `json:"VpcId"`     // VPC ID
	Config    NodePoolConfiguration `json:"Config"`    // 节点池配置
}

type CreateNodePoolResult struct {
	Data CreateNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type CreateNodePoolData struct {
	NodePoolId string `json:"NodePoolId"` // 节点池ID
	TaskId     string `json:"TaskId"`     // 任务ID
}

type ListNodePoolReq struct {
	ClusterId  string `json:"ClusterId"`  // 集群ID
	NodePoolId string `json:"NodePoolId"` // 节点池ID
	Filter     string `json:"Filter"`     // 节点池名称、ID模糊搜索
}

type ListNodePoolResult struct {
	Data ListNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ListNodePoolData struct {
	Total    int64              `json:"Total"`    // 总数
	NodePool []NodePoolInfoData `json:"NodePool"` // 节点池列表
}

type NodePoolInfoData struct {
	AvailableNodeCount int                `json:"AvailableNodeCount"` // 可用节点数
	TotalNodeCount     int                `json:"TotalNodeCount"`     // 总节点数
	Replicas           int                `json:"Replicas"`           // 期望节点数
	Id                 string             `json:"Id"`                 // 节点池ID
	Name               string             `json:"Name"`               // 节点池名称
	NodeType           string             `json:"NodeType"`           // 节点类型
	NodeConfig         NodePoolNodeConfig `json:"NodeConfig"`         // 节点配置
	Status             string             `json:"Status"`             // 状态
	ClusterId          string             `json:"ClusterId"`          // 集群ID
	CustomerId         string             `json:"CustomerId"`         // 客户ID
	UserId             string             `json:"UserId"`             // 用户ID
	AzId               string             `json:"AvailableZoneId"`    // 可用区ID
	AzName             string             `json:"AvailableZoneName"`  // 可用区名称
	CreateTime         string             `json:"CreateTime"`         // 创建时间
	UpdateTime         string             `json:"UpdateTime"`         // 更新时间
}

type DeleteNodePoolReq struct {
	ClusterId  string `json:"ClusterId"`  // 集群ID
	NodePoolId string `json:"NodePoolId"` // 节点池ID
}

type DeleteNodePoolResult struct {
	Data DeleteNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type DeleteNodePoolData struct {
	TaskId string `json:"TaskId"` // 任务ID
}

type ScalingNodePoolReq struct {
	ClusterId  string `json:"ClusterId"`  // 集群ID
	NodePoolId string `json:"NodePoolId"` // 节点池ID
	Replicas   int    `json:"Replicas"`   // 目标节点数量
}

type ScalingNodePoolResult struct {
	Data ScalingNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ScalingNodePoolData struct {
	TaskId string `json:"TaskId"` // 任务ID
}

// 节点池计费方式
const (
	NodePoolBillingMethodPostPaid = "PostPaid" // 按需付费
	NodePoolBillingMethodPrePaid  = "PrePaid"  // 包年包月
)

// NodePoolDiskTypeSSD 节点池支持磁盘类型
const (
	NodePoolDiskTypeSSD = "SSD" // SSD磁盘
)

// 节点可选类型
const (
	NodePoolNodeTypeECS = "ecs" // 云主机
	NodePoolNodeTypeBMS = "bms" // 裸金属
)

// 节点池ECS、BMS实例类型 (持续更新中)
const (
	EcsGpuGch4XLarge        = "Inference gch c8 nr4.16c64g1gpu"
	EcsCpuC11ComputeLarge   = "CPU Compute C11.2c4g"
	EcsCpuC11Compute2XLarge = "CPU Compute C11.8c16g"
	EcsCpuC11Compute4XLarge = "CPU Compute C11.16c32g"
	EcsCpuC11Compute8XLarge = "CPU Compute C11.32c64g"

	BmsGpuGbm32XLarge = "Inference gbm c6 nr44.128c1024g8gpu"
)

// ECS、BMS支持的操作系统镜像
const (
	EcsUbuntu2204K8s13014Cpu         = "Ubuntu22.04-CPU-1.30.14"
	EcsUbuntu2204K8s12615GpuRtx4090  = "Ubuntu22.04-GPU-Geforce-1.26"
	EcsUbuntu2204K8s12615GpuRtxA5000 = "Ubuntu22.04-GPU-Datacenter-1.26"

	BmsUbuntu2004K8s12615Cpu         = "Ubuntu20.04-CPU-1.26-bms"
	BmsUbuntu2004K8s12615GpuRtx4090  = "Ubuntu20.04-GPU_CUDA12.8-bms"
	BmsUbuntu2204K8s12615GpuRtx4090  = "Ubuntu22.04-GPU-Geforce-1.26-bms"
	BmsUbuntu2004K8s12615GpuRtxA5000 = "Ubuntu20.04-GPU-1.26-bms"
	BmsUbuntu2204K8s13014GpuRtx4090  = "Ubuntu22.04-GPU-1.30.14-bms"
)
