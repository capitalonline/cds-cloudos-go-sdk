package main

import (
	"fmt"
	"log"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func NodePoolExamples() {
	ak := "your_access_key"
	sk := "your_secret_key"

	client, err := eks.NewClient(ak, sk)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: Create ECS GPU Node Pool
	fmt.Println("===== Create ECS GPU Node Pool =====")
	createECSGPUReq := &eks.CreateNodePoolReq{
		ClusterId: "cluster-xxxxx",
		VpcId:     "vpc-xxxxx",
		Config: eks.NodePoolConfiguration{
			PoolName:  "gpu-nodepool",
			NodeType:  eks.NodePoolNodeTypeECS,
			SubjectId: 1,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPostPaid, // 按需付费
					Duration:      1,
					IsToMonth:     0,
					AutoRenew:     0,
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
				},
				OsImageName: eks.NodePoolOsImageUbuntu2204K8s1_30_14,
				SubnetIds:   []string{"subnet-xxxxx"},
				InstanceTypeIds: []string{
					eks.NodePoolInstanceTypeECSGPU, // 推理型智算云主机
				},
				Password: "YourPassword123!",
				Shell:    "#!/bin/bash\necho 'GPU Node initialization complete'",
				Labels: map[string]string{
					"env":       "production",
					"node-type": "gpu",
					"team":      "ai",
				},
			},
			Replicas: 1, // GPU节点通常数量较少
		},
	}

	createResult, err := client.CreateNodePool(createECSGPUReq)
	if err != nil {
		log.Printf("Failed to create ECS GPU node pool: %v", err)
	} else {
		fmt.Printf("ECS GPU node pool created successfully: %+v\n", createResult.Data)
	}

	// Example: Create ECS CPU Node Pool
	fmt.Println("\n===== Create ECS CPU Node Pool =====")
	createECSCPUReq := &eks.CreateNodePoolReq{
		ClusterId: "cluster-xxxxx",
		VpcId:     "vpc-xxxxx",
		Config: eks.NodePoolConfiguration{
			PoolName:  "cpu-nodepool",
			NodeType:  eks.NodePoolNodeTypeECS,
			SubjectId: 1,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPostPaid,
					Duration:      1,
					IsToMonth:     0,
					AutoRenew:     0,
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
				},
				OsImageName: eks.NodePoolOsImageUbuntu2204K8s1_30_14,
				SubnetIds:   []string{"subnet-xxxxx"},
				// 使用多种CPU规格
				InstanceTypeIds: []string{
					eks.NodePoolInstanceTypeECSCPUC11Small, // CPU计算型C11.2c4g
					eks.NodePoolInstanceTypeECSCPUC11Large, // CPU计算型C11.8c16g
				},
				Password: "YourPassword123!",
				Shell:    "#!/bin/bash\necho 'CPU Node initialization complete'",
				Labels: map[string]string{
					"env":       "production",
					"node-type": "cpu",
					"team":      "backend",
				},
			},
			Replicas: 3, // CPU节点可以多一些
		},
	}

	createResult2, err := client.CreateNodePool(createECSCPUReq)
	if err != nil {
		log.Printf("Failed to create ECS CPU node pool: %v", err)
	} else {
		fmt.Printf("ECS CPU node pool created successfully: %+v\n", createResult2.Data)
	}

	// Example: Create BMS GPU Node Pool
	fmt.Println("\n===== Create BMS GPU Node Pool =====")
	createBMSReq := &eks.CreateNodePoolReq{
		ClusterId: "cluster-xxxxx",
		VpcId:     "vpc-xxxxx",
		Config: eks.NodePoolConfiguration{
			PoolName:  "bms-gpu-nodepool",
			NodeType:  eks.NodePoolNodeTypeBMS,
			SubjectId: 1,
			NodeConfig: eks.NodePoolNodeConfig{
				BillingSpec: eks.NodePoolBillingSpec{
					BillingMethod: eks.NodePoolBillingMethodPrePaid, // 裸金属通常用包月
					Duration:      1,
					IsToMonth:     0,
					AutoRenew:     0,
				},
				SystemDisk: eks.NodePoolDiskInfo{
					DiskType: eks.NodePoolDiskTypeSSD,
					DiskSize: 40,
				},
				DataDisk: []eks.NodePoolDiskInfo{
					{
						DiskType: eks.NodePoolDiskTypeSSD,
						DiskSize: 160, // 裸金属可以用更大的磁盘
					},
				},
				OsImageName: eks.NodePoolOsImageUbuntu2204K8s1_30_14,
				SubnetIds:   []string{"subnet-xxxxx"},
				// BMS只支持GPU规格
				InstanceTypeIds: []string{
					eks.NodePoolInstanceTypeBMSGPU, // 推理型GPU裸金属
				},
				Password: "YourPassword123!",
				Shell:    "#!/bin/bash\necho 'BMS GPU Node initialization complete'",
				Labels: map[string]string{
					"env":       "production",
					"node-type": "bms-gpu",
					"team":      "ml",
				},
			},
			Replicas: 1, // 裸金属通常数量较少且昂贵
		},
	}

	createResult3, err := client.CreateNodePool(createBMSReq)
	if err != nil {
		log.Printf("Failed to create BMS GPU node pool: %v", err)
	} else {
		fmt.Printf("BMS GPU node pool created successfully: %+v\n", createResult3.Data)
	}

	// Example: List Node Pools
	fmt.Println("\n===== List Node Pools =====")
	listReq := &eks.ListNodePoolReq{
		ClusterId: "cluster-xxxxx",
	}

	listResult, err := client.ListNodePool(listReq)
	if err != nil {
		log.Printf("Failed to list node pools: %v", err)
	} else {
		fmt.Printf("Found %d node pools:\n", listResult.Data.Total)
		for _, pool := range listResult.Data.NodePool {
			fmt.Printf("- ID: %s, Name: %s, Status: %s, NodeType: %s, Replicas: %d\n",
				pool.Id, pool.Name, pool.Status, pool.NodeType, pool.Replicas)
			// 显示实例类型
			if len(pool.NodeConfig.InstanceTypeIds) > 0 {
				fmt.Printf("  Instance Types: %v\n", pool.NodeConfig.InstanceTypeIds)
			}
		}
	}

	// 演示如何获取支持的实例类型
	fmt.Println("\n===== Supported Instance Types =====")
	fmt.Println("所有支持的实例类型:")
	for _, instanceType := range eks.GetNodePoolSupportedInstanceTypes() {
		fmt.Printf("- %s\n", instanceType)
	}

	fmt.Println("\nECS支持的实例类型:")
	for _, instanceType := range eks.GetNodePoolSupportedInstanceTypesForNodeType(eks.NodePoolNodeTypeECS) {
		fmt.Printf("- %s\n", instanceType)
	}

	fmt.Println("\nBMS支持的实例类型:")
	for _, instanceType := range eks.GetNodePoolSupportedInstanceTypesForNodeType(eks.NodePoolNodeTypeBMS) {
		fmt.Printf("- %s\n", instanceType)
	}

	// Example: Scale Node Pool
	fmt.Println("\n===== Scale Node Pool =====")
	scaleReq := &eks.ScalingNodePoolReq{
		ClusterId:  "cluster-xxxxx",
		NodePoolId: "nodepool-xxxxx",
		Replicas:   5, // 扩容到5个节点
	}

	scaleResult, err := client.ScalingNodePool(scaleReq)
	if err != nil {
		log.Printf("Failed to scale node pool: %v", err)
	} else {
		fmt.Printf("Node pool scaling initiated successfully, TaskId: %s\n", scaleResult.Data.TaskId)
	}

	// Example: Delete Node Pool
	fmt.Println("\n===== Delete Node Pool =====")
	deleteReq := &eks.DeleteNodePoolReq{
		ClusterId:  "cluster-xxxxx",
		NodePoolId: "nodepool-xxxxx",
	}

	deleteResult, err := client.DeleteNodePool(deleteReq)
	if err != nil {
		log.Printf("Failed to delete node pool: %v", err)
	} else {
		fmt.Printf("Node pool deletion initiated successfully, TaskId: %s\n", deleteResult.Data.TaskId)
	}

	// 实例类型验证示例
	fmt.Println("\n===== Instance Type Validation =====")
	testInstanceTypes := []string{
		eks.NodePoolInstanceTypeECSGPU,
		eks.NodePoolInstanceTypeECSCPUC11Small,
		"不支持的规格",
	}

	for _, instanceType := range testInstanceTypes {
		isSupported := eks.IsNodePoolInstanceTypeSupported(instanceType)
		isValidForECS := eks.ValidateNodePoolInstanceTypeForNodeType(instanceType, eks.NodePoolNodeTypeECS)
		isValidForBMS := eks.ValidateNodePoolInstanceTypeForNodeType(instanceType, eks.NodePoolNodeTypeBMS)
		fmt.Printf("实例类型: %s\n", instanceType)
		fmt.Printf("  是否支持: %t\n", isSupported)
		fmt.Printf("  适用于ECS: %t\n", isValidForECS)
		fmt.Printf("  适用于BMS: %t\n", isValidForBMS)
		fmt.Println()
	}
}
