package main

import (
	"encoding/json"
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

// CreateECSNodePool 创建 ECS-CPU云主机 节点池（按需付费）
func CreateECSNodePool() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.CreateNodePoolReq{
		ClusterId: "cluster-01",
		VpcId:     "vpc-01",
		Config: eks.NodePoolConfiguration{
			PoolName: "ecs-cpu-node-pool",
			NodeType: eks.NodePoolNodeTypeECS,
			// SubjectId 测试金项目ID; 如果没有申请默认不传或传0; 如需申请请联系销售
			SubjectId: 0,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPostPaid, // 按需付费
				},
				SystemDisk: eks.NodePoolDiskInfo{
					DiskType: eks.NodePoolDiskTypeSSD,
					// DiskSize 系统盘为固定值40GB
					DiskSize: 40,
				},
				DataDisk: []eks.NodePoolDiskInfo{
					{
						DiskType: eks.NodePoolDiskTypeSSD,
						// DiskSize: 最小值80,步长80
						DiskSize: 80,
					},
				},
				//OsImageName 实例镜像名称，根据您的k8s版本，CPU/GPU实例类型按需选择
				OsImageName: eks.EcsUbuntu2204K8s13014Cpu,

				// SubnetIds VPC子网ID,支持多选，必须是同一可用区
				SubnetIds: []string{"subnet-01", "subnet-02"},

				// InstanceTypeIds 云主机实例类型，目前首云仅支持一个，后续开放多实例规格
				InstanceTypeIds: []string{
					// CPU计算型C11 8C16G
					eks.EcsCpuC11Compute2XLarge,
				},

				// Password eks用户登录密码，节点初始化完毕后自动创建eks用户
				Password: "YourPassword123!",

				// Shell 节点初始化完成后执行脚本命令
				Shell: "#!/bin/bash\necho 'ECS CPU Node initialization complete'",

				// Labels 节点k8s标签
				Labels: map[string]string{
					"node-type": "ecs-cpu",
					"team":      "backend",
				},
			},
			// Replicas期望节点数量
			Replicas: 3,
		},
	}

	response, err := eksClient.CreateNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// CreateECSGPUNodePool 创建 ECS-GPU云主机 节点池（按需付费）
func CreateECSGPUNodePool() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.CreateNodePoolReq{
		ClusterId: "cluster-02",
		VpcId:     "vpc-02",
		Config: eks.NodePoolConfiguration{
			PoolName: "ecs-gpu-node-pool",
			NodeType: eks.NodePoolNodeTypeECS,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPostPaid,
				},
				SystemDisk: eks.NodePoolDiskInfo{
					DiskType: eks.NodePoolDiskTypeSSD,
					DiskSize: 40,
				},
				DataDisk: []eks.NodePoolDiskInfo{
					{
						DiskType: eks.NodePoolDiskTypeSSD,
						DiskSize: 80,
					},
					{
						DiskType: eks.NodePoolDiskTypeSSD,
						DiskSize: 160,
					},
				},
				OsImageName: eks.EcsUbuntu2204K8s13014Cpu,
				SubnetIds:   []string{"subnet-01"},
				InstanceTypeIds: []string{
					eks.EcsGpuGch4XLarge,
				},
				Password: "YourPassword123!",
				Shell:    "#!/bin/bash\necho 'ECS GPU Node initialization complete'",
				Labels: map[string]string{
					"env":       "production",
					"node-type": "ecs-gpu",
					"team":      "ai",
				},
			},
			Replicas: 2,
		},
	}

	response, err := eksClient.CreateNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// CreateBMSNodePoolPostPaid 创建 BMS-GPU裸金属 节点池（按需付费）
func CreateBMSNodePoolPostPaid() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.CreateNodePoolReq{
		ClusterId: "cluster-03",
		VpcId:     "vpc-03",
		Config: eks.NodePoolConfiguration{
			PoolName:  "bms-gpu-postpaid-node-pool",
			NodeType:  eks.NodePoolNodeTypeBMS,
			SubjectId: 0,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPostPaid, // 按需付费
				},
				// 首云裸金属暂不支持挂载云盘，使用裸金属本地盘
				SystemDisk: eks.NodePoolDiskInfo{},
				DataDisk:   []eks.NodePoolDiskInfo{},

				// OsImageName 裸金属实例镜像
				OsImageName: eks.BmsUbuntu2204K8s13014GpuRtx4090,

				// SubnetIds VPC子网ID,支持多选，必须是同一可用区
				SubnetIds: []string{"subnet-03", "subnet-04"},

				// InstanceTypeIds 裸金属实例类型，目前首云仅支持一个，后续开放多实例规格
				InstanceTypeIds: []string{
					// 推理型智算云主机igch.c8.nr4 16C64G
					eks.BmsGpuGbm32XLarge,
				},

				// Password eks用户登录密码，节点初始化完毕后自动创建eks用户
				Password: "YourPassword123!",

				// Shell 节点初始化完成后执行脚本命令
				Shell: "#!/bin/bash\necho 'BMS GPU PostPaid Node initialization complete'",

				// Labels 节点k8s标签
				Labels: map[string]string{
					"node-type":    "bms-gpu-rtx-4090",
					"billing-type": "postpaid",
				},
			},
			Replicas: 1, // 裸金属通常数量较少且昂贵
		},
	}

	response, err := eksClient.CreateNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// CreateBMSNodePoolPrePaid 创建 BMS-GPU裸金属 节点池（包年包月）
func CreateBMSNodePoolPrePaid() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.CreateNodePoolReq{
		ClusterId: "cluster-04",
		VpcId:     "vpc-04",
		Config: eks.NodePoolConfiguration{
			PoolName:  "bms-gpu-prepaid-node-pool",
			NodeType:  eks.NodePoolNodeTypeBMS,
			SubjectId: 0,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPrePaid, // 包年包月
					Duration:      3,                                // 购买市场3个月
					IsToMonth:     1,                                // 是否购买至当月底
					AutoRenew:     1,                                // 到期开启自动续费
				},
				SystemDisk:  eks.NodePoolDiskInfo{},
				DataDisk:    []eks.NodePoolDiskInfo{},
				OsImageName: eks.BmsUbuntu2204K8s13014GpuRtx4090,
				SubnetIds:   []string{"subnet-01"},
				InstanceTypeIds: []string{
					eks.BmsGpuGbm32XLarge,
				},
				Password: "YourPassword123!",
				Shell:    "#!/bin/bash\necho 'BMS GPU PrePaid Node initialization complete'",
				Labels: map[string]string{
					"env":          "production",
					"node-type":    "bms-gpu",
					"billing-type": "prepaid",
				},
			},
			Replicas: 1, // 裸金属通常数量较少且昂贵
		},
	}

	response, err := eksClient.CreateNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// ListNodePools 查询节点池列表
func ListNodePools() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.ListNodePoolReq{
		ClusterId: "cluster-xxxxx",
	}

	response, err := eksClient.ListNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// ScaleNodePool 扩缩容节点池
func ScaleNodePool() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.ScalingNodePoolReq{
		ClusterId:  "cluster-xxxxx",
		NodePoolId: "nodepool-xxxxx",
		Replicas:   5, // 扩容到5个节点
	}

	response, err := eksClient.ScalingNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

// DeleteNodePool 删除节点池
func DeleteNodePool() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := &eks.DeleteNodePoolReq{
		ClusterId:  "cluster-xxxxx",
		NodePoolId: "nodepool-xxxxx",
	}

	response, err := eksClient.DeleteNodePool(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}
