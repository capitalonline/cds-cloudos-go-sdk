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

// NodePoolBillingSpec 节点池计费配置
type NodePoolBillingSpec struct {
	BillingMethod string `json:"InstanceChargeType"` // 实例计费类型：PostPaid(按需付费) / PrePaid(包年包月)
	Duration      int    `json:"Duration"`           // 实例购买时长，单位：月
	IsToMonth     int    `json:"IsToMonth"`          // 是否购买至当前月底：0(购买整月) / 1(购买至月底)
	AutoRenew     int    `json:"AutoRenew"`          // 是否自动续费：0(不自动续费) / 1(自动续费)
}

// NodePoolDiskInfo 节点池磁盘配置
type NodePoolDiskInfo struct {
	DiskType string `json:"Type"` // 磁盘类型：SSD
	DiskSize int    `json:"Size"` // 磁盘大小，单位：GB
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
	ClusterId string `json:"ClusterId"` // 集群ID
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

	// BmsGpuGbm32XLarge 20250907 改英文
	BmsGpuGbm32XLarge = "推理型GPU裸金属igbm.c6.nr44.128c1024g8gpu"
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
