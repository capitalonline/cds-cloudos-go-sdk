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
	Status  string `json:"Status,omitempty"`
	Version string `json:"Version,omitempty"`
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
	SubnetId  string `json:"SubnetId"`
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
	Data []GetClusterDetail `json:"Data,omitempty"`
	OpenApiCommonResp
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
	CniConfig   CniConfig `json:"CniConfig"`
	CniType     string    `json:"CniType"`
	ProxyConfig string    `json:"ProxyConfig"`
	ServiceCidr string    `json:"ServiceCidr"`
}

type CniConfig struct {
	FlannelBackendType  string `json:"FlannelBackendType"`
	CalicoBlockSize     int    `json:"CalicoBlockSize"`
	CalicoEncapsulation string `json:"CalicoEncapsulation"`
	NodePodsNum         int    `json:"NodePodsNum"`
	PodCidr             string `json:"PodCidr"`
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
	Keyword        int    `json:"Keyword"`
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
