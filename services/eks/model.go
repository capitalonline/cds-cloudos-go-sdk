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
	StatusStr     string `json:"StatusStr"`
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
	ErrorInfo          string `json:"ErrorInfo"`
	Frontend           string `json:"Frontend"`
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
	ClusterId       string     `json:"ClusterId"`
	Ak              string     `json:"Ak"`
	ClusterIp       string     `json:"ClusterIp"`
	VpcId           string     `json:"VpcId"`
	ClusterName     string     `json:"ClusterName"`
	ClusterStatus   string     `json:"ClusterStatus"`
	CreateTime      string     `json:"CreateTime"`
	CustomerId      string     `json:"CustomerId"`
	DashboardAddr   string     `json:"DashboardAddr"`
	K8sCaExpiryData int        `json:"K8sCaExpiryData"`
	K8sVersion      string     `json:"K8sVersion"`
	MasterSum       int        `json:"MasterSum"`
	NatId           string     `json:"NatId"`
	NatName         string     `json:"NatName"`
	NodeSum         int        `json:"NodeSum"`
	OsName          string     `json:"OsName"`
	RegionId        string     `json:"RegionId"`
	RegionName      string     `json:"RegionName"`
	RuntimeType     string     `json:"RuntimeType"`
	SlbId           string     `json:"SlbId"`
	SshPort         int        `json:"SshPort"`
	StatusStr       string     `json:"StatusStr"`
	SubDomain       string     `json:"SubDomain"`
	UpdateTime      string     `json:"UpdateTime"`
	UserId          string     `json:"UserId"`
	Vip             string     `json:"Vip"`
	VipId           string     `json:"VipId"`
	VpcName         string     `json:"VpcName"`
	WorkerSum       int        `json:"WorkerSum"`
	CniInfo         CniInfo    `json:"CniInfo"`
	DestinationEip  EipInfo    `json:"DestinationEip"`
	SourceEip       EipInfo    `json:"SourceEip"`
	MasterSubnet    SubnetInfo `json:"MasterSubnet"`
	WorkerSubnet    SubnetInfo `json:"WorkerSubnet"`
}

type CniInfo struct {
	CniConfig   CniConfig   `json:"CniConfig"`
	CniType     CniType     `json:"CniType"`
	ProxyConfig ProxyConfig `json:"ProxyConfig"`
	ServiceCidr ServiceCidr `json:"ServiceCidr"`
}

type CniConfig struct {
	FlannelBackendType  FlannelBackendType  `json:"FlannelBackendType"`
	CalicoBlockSize     int                 `json:"CalicoBlockSize"`
	CalicoEncapsulation CalicoEncapsulation `json:"CalicoEncapsulation"`
	NodePodsNum         NodePodsNum         `json:"NodePodsNum"`
	PodCidr             string              `json:"PodCidr"`
	SubnetList          []SubnetInfo        `json:"SubnetList"`
}

type EipInfo struct {
	EipId string `json:"EipId"`
	Ip    string `json:"Ip"`
}

type SubnetInfo struct {
	AvailableZoneId string `json:"AvailableZoneId"`
	SubnetId        string `json:"SubnetId"`
	Segment         string `json:"Segment"`
	SelectSegment   string `json:"SelectSegment"`
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
	NodeId            string `json:"NodeId"`
	NodeName          string `json:"NodeName"`
	NodeResourceId    string `json:"NodeResourceId"`
	StorageResourceId string `json:"StorageResourceId"`
	NodeNumber        int    `json:"NodeNumber"`
	NodeStatus        string `json:"NodeStatus"`
	StatusStr         string `json:"StatusStr"`
	EcsStatus         string `json:"EcsStatus"`
	IsValid           int    `json:"IsValid"`
	NodeType          string `json:"NodeType"`
	ClusterId         string `json:"ClusterId"`
	ClusterName       string `json:"ClusterName"`
	CustomerId        string `json:"CustomerId"`
	UserId            string `json:"UserId"`
	Hostname          string `json:"Hostname"`
	VpcId             string `json:"VpcId"`
	SubnetId          string `json:"SubnetId"`
	PrivateIp         string `json:"PrivateIp"`
	CpuSize           int    `json:"CpuSize"`
	RamSize           int    `json:"RamSize"`
	GpuSum            int    `json:"GpuSum"`
	GpuShowType       string `json:"GpuShowType"`
	BmsHostId         string `json:"BmsHostId"`
	OsImageId         string `json:"OsImageId"`
	FamilyId          string `json:"FamilyId"`
	FamilyName        string `json:"FamilyName"`
	Schedulable       int    `json:"Schedulable"`
	K8sStatus         string `json:"K8SStatus"`
	RegionName        string `json:"RegionName"`
	RegionId          string `json:"RegionId"`
	AzName            string `json:"AzName"`
	AzId              string `json:"AzId"`
	SystemDisk        string `json:"SystemDisk"`
	DataDisk          string `json:"DataDisk"`
	SuborderId        string `json:"SuborderId"`
	SourceType        string `json:"SourceType"`
	ScalingGroupId    string `json:"ScalingGroupId"`
	Labels            string `json:"Labels"`
	Annotations       string `json:"Annotations"`
	BillingMethod     string `json:"BillingMethod"`
	Duration          int    `json:"Duration"`
	IsToMonth         int    `json:"IsToMonth"`
	AutoRenew         int    `json:"AutoRenew"`
	CreateTime        string `json:"CreateTime"`
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
	ClusterName      string         `json:"ClusterName"`
	VpcId            string         `json:"VpcId"`
	Cni              CniInfo        `json:"Cni"`
	K8sVersion       K8sVersion     `json:"K8SVersion"`
	RuntimeType      string         `json:"RuntimeType"`
	NatId            string         `json:"NatId"`
	SourceEipId      string         `json:"SourceEipId"`
	DestinationEipId string         `json:"DestinationEipId"`
	SlbId            string         `json:"SlbId"`
	Ak               string         `json:"Ak"`
	Sk               string         `json:"Sk"`
	MasterNumber     MasterNumber   `json:"MasterNumber"`
	MasterConfig     NodeConfig     `json:"MasterConfig"`
	NodePoolConfig   NodePoolConfig `json:"NodePoolConfig"`
}

type NodeConfig struct {
	BillingSpec BillingSpec       `json:"BillingSpec"`
	SystemDisk  DiskInfo          `json:"SystemDisk"`
	DataDisk    []DiskInfo        `json:"DataDisk"`
	OsImageName OsImageName       `json:"OsImageName"`
	SubnetList  []SubnetInfo      `json:"SubnetList"`
	Specifics   []Specifics       `json:"Specifics"`
	Password    string            `json:"Password"`
	Shell       string            `json:"Shell"`
	Labels      map[string]string `json:"Labels"`
	Annotations map[string]string `json:"Annotations"`
}

type BillingSpec struct {
	BillingMethod BillingMethod `json:"BillingMethod"`
	Duration      Duration      `json:"Duration"`
	IsToMonth     IsToMonth     `json:"IsToMonth"`
	AutoRenew     AutoRenew     `json:"AutoRenew"`
}

type DiskInfo struct {
	DiskName     string      `json:"DiskName,omitempty"`
	DiskId       string      `json:"DiskId,omitempty"`
	DeviceScsiId string      `json:"DeviceScsiId,omitempty"`
	DiskFeature  DiskFeature `json:"DiskFeature"`
	Snapshot     string      `json:"Snapshot,omitempty"`
	DiskSize     int         `json:"DiskSize"`
	DiskStorage  []string    `json:"DiskStorage,omitempty"`
	Number       int         `json:"Number"`
}

type Specifics struct {
	ProductName ProductName `json:"ProductName"`
}

type NodePoolConfig struct {
	PoolName    string      `json:"PoolName"`
	NodeType    NodeType    `json:"NodeType"`
	SubjectId   int         `json:"SubjectId"`
	NodeConfig  NodeConfig  `json:"NodeConfig"`
	AutoScaling AutoScaling `json:"AutoScaling,omitempty"`
	Replicas    int         `json:"Replicas"`
}

type AutoScaling struct {
	Enable         bool   `json:"Enable"`
	MaxReplicas    int    `json:"MaxReplicas"`
	MinReplicas    int    `json:"MinReplicas"`
	Priority       int    `json:"Priority"`
	SubnetPolicy   string `json:"SubnetPolicy"`
	ScalingGroupId string `json:"ScalingGroupId"`
}

type CreateClusterData struct {
	ClusterId string `json:"ClusterId"`
	TaskId    string `json:"TaskId"`
}

type CreateClusterResult struct {
	Data CreateClusterData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type CniType string

// cni类型
const (
	CniTypeFlannel CniType = "flannel"
	CniTypeCalico  CniType = "calico"
	CniTypeVpcCni  CniType = "vpc-cni"
)

type K8sVersion string

// k8s版本
const (
	K8sVersion1_30_14 K8sVersion = "v1.30.14"
	K8sVersion1_26_5  K8sVersion = "v1.26.5"
)

type FlannelBackendType string

// flannel后端类型
const (
	FlannelBackendTypeVxlan FlannelBackendType = "vxlan"
)

type ProxyConfig string

// 服务转发模式
const (
	ProxyConfigIptables ProxyConfig = "iptables"
	ProxyConfigIPVS     ProxyConfig = "ipvs"
)

type MasterNumber int

// master节点数
const (
	MasterNumber3 MasterNumber = 3
	MasterNumber5 MasterNumber = 5
	MasterNumber7 MasterNumber = 7
)

type DiskFeature string

// 云盘类型，ecs适用
const (
	DiskFeatureSsd DiskFeature = "ssd"
)

type ProductName string

// 产品名称及CPU个数,内存大小，GPU数量
const (
	ProductName1 ProductName = "Optimized M6,2,4,0"
	ProductName2 ProductName = "Optimized M6,4,8,0"
	ProductName3 ProductName = "Optimized M6,8,16,0"
	ProductName4 ProductName = "Optimized M6,16,32,0"
)

type OsImageName string

// 镜像名称
const (
	OsImageNameUbuntu2204_1_30_14 OsImageName = "eks-Ubuntu22.04-cpu-k8s1.30.14-v1"
)

type AutoRenew int

// 自动续费,0:关闭自动续费,1:开启自动续费,包年包月适用
const (
	AutoRenew1 AutoRenew = 1 // 开启自动续费,包年包月适用
	AutoRenew0 AutoRenew = 0 // 关闭自动续费,包年包月适用
)

type BillingMethod string

// 计费方式,0:按需付费,1:包年包月
const (
	BillingMethod0 BillingMethod = "0" // 按需付费
	BillingMethod1 BillingMethod = "1" // 包年包月
)

type NodeType string

// 节点类型,ecs:云主机，bms:裸金属
const (
	NodeTypeEcs NodeType = "ecs"
	NodeTypeBms NodeType = "bms"
)

type NodePodsNum int

// 节点的Pods数量
const (
	NodePodsNum16  NodePodsNum = 16
	NodePodsNum32  NodePodsNum = 32
	NodePodsNum64  NodePodsNum = 64
	NodePodsNum128 NodePodsNum = 128
	NodePodsNum256 NodePodsNum = 256
)

type CalicoEncapsulation string

// calico使用的网络封装模式
const (
	CalicoEncapsulationIPIP CalicoEncapsulation = "IPIP"
)

type ServiceCidr string

// 集群服务CIDR ip地址段
const (
	ServiceCidr10_16  ServiceCidr = "10.0.0.0/16"
	ServiceCidr10_24  ServiceCidr = "10.0.0.0/24"
	ServiceCidr192_16 ServiceCidr = "192.168.0.0/16"
	ServiceCidr192_24 ServiceCidr = "192.168.0.0/24"
	ServiceCidr172_16 ServiceCidr = "172.16.0.0/16"
	ServiceCidr172_24 ServiceCidr = "172.16.0.0/24"
)

type IsToMonth int

// 是否购买至月底
const (
	IsToMonth1 IsToMonth = 1 // 购买至月底,包年包月适用
	IsToMonth0 IsToMonth = 0 // 购买整月,包年包月适用
)

type Duration int

// 购买的月份数，购买前请确保账户余额充足，包年包月适用
const (
	Duration1  Duration = 1
	Duration2  Duration = 2
	Duration3  Duration = 3
	Duration4  Duration = 4
	Duration5  Duration = 5
	Duration6  Duration = 6
	Duration7  Duration = 7
	Duration8  Duration = 8
	Duration9  Duration = 9
	Duration10 Duration = 10
	Duration11 Duration = 11
	Duration12 Duration = 12
	Duration13 Duration = 24
	Duration14 Duration = 36
)
