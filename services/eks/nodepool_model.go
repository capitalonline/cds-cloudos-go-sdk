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

// ======= NodePool专用数据结构定义 =======

// NodePoolBillingSpec 节点池计费配置
type NodePoolBillingSpec struct {
	BillingMethod string `json:"InstanceChargeType"` // 实例计费类型：PostPaid(按需付费) / PrePaid(包年包月)
	Duration      int    `json:"Period"`             // 实例购买时长，单位：月
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

// ======= NodePool API请求和响应结构体 =======

// CreateNodePoolReq 创建节点池请求
type CreateNodePoolReq struct {
	ClusterId string                `json:"ClusterId"` // 集群ID
	VpcId     string                `json:"VpcId"`     // VPC ID
	Config    NodePoolConfiguration `json:"Config"`    // 节点池配置
}

// CreateNodePoolResult 创建节点池响应
type CreateNodePoolResult struct {
	Data CreateNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

// CreateNodePoolData 创建节点池返回数据
type CreateNodePoolData struct {
	NodePoolId string `json:"NodePoolId"` // 节点池ID
	TaskId     string `json:"TaskId"`     // 任务ID
}

// ListNodePoolReq 查询节点池列表请求
type ListNodePoolReq struct {
	ClusterId string `json:"ClusterId"` // 集群ID
}

// ListNodePoolResult 查询节点池列表响应
type ListNodePoolResult struct {
	Data ListNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

// ListNodePoolData 节点池列表数据
type ListNodePoolData struct {
	Total    int64              `json:"Total"`    // 总数
	NodePool []NodePoolInfoData `json:"NodePool"` // 节点池列表
}

// NodePoolInfoData 节点池详细信息
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

// DeleteNodePoolReq 删除节点池请求
type DeleteNodePoolReq struct {
	ClusterId  string `json:"ClusterId"`  // 集群ID
	NodePoolId string `json:"NodePoolId"` // 节点池ID
}

// DeleteNodePoolResult 删除节点池响应
type DeleteNodePoolResult struct {
	Data DeleteNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

// DeleteNodePoolData 删除节点池返回数据
type DeleteNodePoolData struct {
	TaskId string `json:"TaskId"` // 任务ID
}

// ScalingNodePoolReq 伸缩节点池请求
type ScalingNodePoolReq struct {
	ClusterId  string `json:"ClusterId"`  // 集群ID
	NodePoolId string `json:"NodePoolId"` // 节点池ID
	Replicas   int    `json:"Replicas"`   // 目标节点数量
}

// ScalingNodePoolResult 伸缩节点池响应
type ScalingNodePoolResult struct {
	Data ScalingNodePoolData `json:"Data,omitempty"`
	OpenApiCommonResp
}

// ScalingNodePoolData 伸缩节点池返回数据
type ScalingNodePoolData struct {
	TaskId string `json:"TaskId"` // 任务ID
}

// ======= 实例类型和常量定义 =======

// 计费方式常量
const (
	NodePoolBillingMethodPostPaid = "PostPaid" // 按需付费
	NodePoolBillingMethodPrePaid  = "PrePaid"  // 包年包月
)

// NodePoolDiskTypeSSD 磁盘类型常量
const (
	NodePoolDiskTypeSSD = "SSD" // SSD磁盘
)

// 节点类型常量
const (
	NodePoolNodeTypeECS = "ecs" // 云主机
	NodePoolNodeTypeBMS = "bms" // 裸金属
)

// 支持的实例类型常量
const (
	// NodePoolInstanceTypeECSGPU ECS 云主机实例类型
	NodePoolInstanceTypeECSGPU         = "推理型智算云主机igch.c8.nr4.16c64g1gpu"
	NodePoolInstanceTypeECSCPUC11Small = "CPU计算型C11.2c4g"
	NodePoolInstanceTypeECSCPUC11Large = "CPU计算型C11.8c16g"

	// NodePoolInstanceTypeBMSGPU BMS 裸金属实例类型
	NodePoolInstanceTypeBMSGPU = "推理型GPU裸金属igbm.c6.nr44.128c1024g8gpu"
)

// NodePoolOsImageUbuntu2204K8s1_30_14 操作系统镜像常量
const (
	NodePoolOsImageUbuntu2204K8s1_30_14 = "eks-Ubuntu22.04-cpu-k8s1.30.14-v1"
)

// ======= 实例类型辅助函数 =======

// GetNodePoolSupportedInstanceTypes 获取所有支持的实例类型
func GetNodePoolSupportedInstanceTypes() []string {
	return []string{
		NodePoolInstanceTypeBMSGPU,
		NodePoolInstanceTypeECSGPU,
		NodePoolInstanceTypeECSCPUC11Small,
		NodePoolInstanceTypeECSCPUC11Large,
	}
}

// GetNodePoolSupportedInstanceTypesForNodeType 根据节点类型获取支持的实例类型
func GetNodePoolSupportedInstanceTypesForNodeType(nodeType string) []string {
	switch nodeType {
	case NodePoolNodeTypeBMS:
		return []string{NodePoolInstanceTypeBMSGPU}
	case NodePoolNodeTypeECS:
		return []string{
			NodePoolInstanceTypeECSGPU,
			NodePoolInstanceTypeECSCPUC11Small,
			NodePoolInstanceTypeECSCPUC11Large,
		}
	default:
		return GetNodePoolSupportedInstanceTypes()
	}
}

// IsNodePoolInstanceTypeSupported 检查实例类型是否支持
func IsNodePoolInstanceTypeSupported(instanceType string) bool {
	supportedTypes := map[string]bool{
		NodePoolInstanceTypeBMSGPU:         true,
		NodePoolInstanceTypeECSGPU:         true,
		NodePoolInstanceTypeECSCPUC11Small: true,
		NodePoolInstanceTypeECSCPUC11Large: true,
	}
	return supportedTypes[instanceType]
}

// ValidateNodePoolInstanceTypeForNodeType 验证实例类型是否适用于指定的节点类型
func ValidateNodePoolInstanceTypeForNodeType(instanceType string, nodeType string) bool {
	switch nodeType {
	case NodePoolNodeTypeBMS:
		return instanceType == NodePoolInstanceTypeBMSGPU
	case NodePoolNodeTypeECS:
		return instanceType == NodePoolInstanceTypeECSGPU ||
			instanceType == NodePoolInstanceTypeECSCPUC11Small ||
			instanceType == NodePoolInstanceTypeECSCPUC11Large
	default:
		return IsNodePoolInstanceTypeSupported(instanceType)
	}
}
