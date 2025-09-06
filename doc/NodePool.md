# EKS NodePool SDK 使用指南

本文档介绍如何使用EKS NodePool相关的4个SDK方法：CreateNodePool、ListNodePool、DeleteNodePool、ScalingNodePool。

## 功能概述

这些SDK方法提供了完整的节点池管理功能：

- **CreateNodePool**: 创建新的节点池
- **ListNodePool**: 查询集群中的节点池列表
- **DeleteNodePool**: 删除指定的节点池
- **ScalingNodePool**: 伸缩节点池中的节点数量

## 快速开始

### 1. 初始化客户端

```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"

ak := "your_access_key"
sk := "your_secret_key"

client, err := eks.NewClient(ak, sk)
if err != nil {
    log.Fatalf("Failed to create client: %v", err)
}
```

### 2. 创建节点池

#### ECS GPU节点池示例

```go
createReq := &eks.CreateNodePoolReq{
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
            Shell:    "#!/bin/bash\\necho 'GPU Node initialization complete'",
            Labels: map[string]string{
                "env":       "production",
                "node-type": "gpu",
                "team":      "ai",
            },
        },
        Replicas: 1,
    },
}

result, err := client.CreateNodePool(createReq)
if err != nil {
    log.Printf("Failed to create node pool: %v", err)
} else {
    fmt.Printf("NodePool created: ID=%s, TaskId=%s\\n", 
        result.Data.NodePoolId, result.Data.TaskId)
}
```

#### ECS CPU节点池示例

```go
createReq := &eks.CreateNodePoolReq{
    ClusterId: "cluster-xxxxx",
    VpcId:     "vpc-xxxxx",
    Config: eks.NodePoolConfiguration{
        PoolName:  "cpu-nodepool",
        NodeType:  eks.NodePoolNodeTypeECS,
        SubjectId: 1,
        NodeConfig: eks.NodePoolNodeConfig{
            // ... 其他配置相同 ...
            // 使用CPU规格
            InstanceTypeIds: []string{
                eks.NodePoolInstanceTypeECSCPUC11Small, // CPU计算型C11.2c4g
                eks.NodePoolInstanceTypeECSCPUC11Large, // CPU计算型C11.8c16g
            },
            // ... 其他配置 ...
        },
        Replicas: 3,
    },
}
```

#### BMS GPU节点池示例

```go
createReq := &eks.CreateNodePoolReq{
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
            // ... 其他配置 ...
            // BMS只支持GPU规格
            InstanceTypeIds: []string{
                eks.NodePoolInstanceTypeBMSGPU, // 推理型GPU裸金属
            },
            // ... 其他配置 ...
        },
        Replicas: 1,
    },
}
```

### 3. 查询节点池列表

```go
listReq := &eks.ListNodePoolReq{
    ClusterId: "cluster-xxxxx",
}

result, err := client.ListNodePool(listReq)
if err != nil {
    log.Printf("Failed to list node pools: %v", err)
} else {
    fmt.Printf("Found %d node pools\\n", result.Data.Total)
    for _, pool := range result.Data.NodePool {
        fmt.Printf("- ID: %s, Name: %s, Status: %s, Replicas: %d\\n",
            pool.Id, pool.Name, pool.Status, pool.Replicas)
    }
}
```

### 4. 伸缩节点池

```go
scaleReq := &eks.ScalingNodePoolReq{
    ClusterId:  "cluster-xxxxx",
    NodePoolId: "nodepool-xxxxx",
    Replicas:   5,  // 新的节点数量
}

result, err := client.ScalingNodePool(scaleReq)
if err != nil {
    log.Printf("Failed to scale node pool: %v", err)
} else {
    fmt.Printf("Scaling initiated, TaskId: %s\\n", result.Data.TaskId)
}
```

### 5. 删除节点池

```go
deleteReq := &eks.DeleteNodePoolReq{
    ClusterId:  "cluster-xxxxx",
    NodePoolId: "nodepool-xxxxx",
}

result, err := client.DeleteNodePool(deleteReq)
if err != nil {
    log.Printf("Failed to delete node pool: %v", err)
} else {
    fmt.Printf("Deletion initiated, TaskId: %s\\n", result.Data.TaskId)
}
```

## 数据结构说明

### CreateNodePoolReq

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|------|
| ClusterId | string | 是 | 集群ID |
| VpcId | string | 是 | VPC ID |
| Config | NodePoolConfiguration | 是 | 节点池配置 |

### NodePoolConfiguration

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|------|
| PoolName | string | 是 | 节点池名称 |
| NodeType | string | 是 | 节点类型 (ecs/bms) |
| SubjectId | int | 是 | 主体ID |
| NodeConfig | NodePoolNodeConfig | 是 | 节点配置 |
| Replicas | int | 是 | 节点数量 |

### NodePoolNodeConfig

| 字段 | 类型 | 必填 | 描述 |
|------|------|------|------|
| BillingSpec | NodePoolBillingSpec | 是 | 计费配置 |
| SystemDisk | NodePoolDiskInfo | 是 | 系统卷配置 |
| DataDisk | []NodePoolDiskInfo | 否 | 数据卷配置 |
| OsImageName | string | 是 | 操作系统镜像 |
| SubnetIds | []string | 是 | 子网ID列表 |
| InstanceTypeIds | []string | 是 | 实例类型ID列表 |
| Password | string | 是 | 节点密码 |
| Shell | string | 否 | 初始化脚本 |
| Labels | map[string]string | 否 | 节点标签 |

## 实例类型常量

### ECS 云主机实例类型

- `eks.NodePoolInstanceTypeECSGPU`: "推理型智算云主机igch.c8.nr4.16c64g1gpu" - GPU计算实例
- `eks.NodePoolInstanceTypeECSCPUC11Small`: "CPU计算型C11.2c4g" - 2核4G CPU实例
- `eks.NodePoolInstanceTypeECSCPUC11Large`: "CPU计算型C11.8c16g" - 8核16G CPU实例

### BMS 裸金属实例类型

- `eks.NodePoolInstanceTypeBMSGPU`: "推理型GPU裸金属igbm.c6.nr44.128c1024g8gpu" - 128核1024G 8GPU

### 实例类型辅助方法

```go
// 获取所有支持的实例类型
supportedTypes := eks.GetNodePoolSupportedInstanceTypes()

// 根据节点类型获取支持的实例类型
ecsTypes := eks.GetNodePoolSupportedInstanceTypesForNodeType(eks.NodePoolNodeTypeECS)
bmsTypes := eks.GetNodePoolSupportedInstanceTypesForNodeType(eks.NodePoolNodeTypeBMS)

// 检查实例类型是否支持
isSupported := eks.IsNodePoolInstanceTypeSupported("推理型智算云主机igch.c8.nr4.16c64g1gpu")

// 验证实例类型是否适用于指定节点类型
isValid := eks.ValidateNodePoolInstanceTypeForNodeType("推理型GPU裸金属igbm.c6.nr44.128c1024g8gpu", eks.NodePoolNodeTypeBMS)
```

### 其他常用常量

#### 节点类型
- `eks.NodePoolNodeTypeECS`: 云主机节点
- `eks.NodePoolNodeTypeBMS`: 裸金属节点

#### 计费方式
- `eks.NodePoolBillingMethodPostPaid`: 按需付费
- `eks.NodePoolBillingMethodPrePaid`: 包年包月

#### 磁盘类型
- `eks.NodePoolDiskTypeSSD`: SSD磁盘

#### 操作系统镜像
- `eks.NodePoolOsImageUbuntu2204K8s1_30_14`: "eks-Ubuntu22.04-cpu-k8s1.30.14-v1"

## 重要变更说明

### 🔄 全新的NodePool SDK

本SDK采用了全新的数据结构设计，完全按照eks-service中的openapi_body.CreateNodePoolReq结构体生成：

- 使用 `InstanceTypeIds` 字段指定实例类型，支持最新的实例规格
- 计费配置中字段名与OpenAPI项目完全一致（如 `InstanceChargeType`、`Period` 等）
- 磁盘配置使用 `SystemVolume` 和 `DataVolumes` 字段名
- 所有结构体都有 `NodePool` 前缀以避免与原有集群创建API的命名冲突

### 与集群创建API的区别

NodePool专用API与集群创建时的NodePool配置在以下方面有所不同：

| 特性 | 集群创建API | NodePool专用API |
|------|------------|----------------|
| 结构体前缀 | 无前缀（如 `NodeConfig`） | `NodePool` 前缀（如 `NodePoolNodeConfig`） |
| 实例类型指定 | `Specifics` 字段（已废弃） | `InstanceTypeIds` 字段 |
| 字段命名风格 | 驼峰命名 | 与OpenAPI项目一致 |
| 自动伸缩 | 包含 `AutoScaling` 配置 | 不包含，通过独立API管理 |

## 错误处理

所有SDK方法都会返回标准的错误信息：

```go
result, err := client.CreateNodePool(createReq)
if err != nil {
    log.Printf("API call failed: %v", err)
    return
}

if result.Code != "Success" {
    log.Printf("API returned error: Code=%s, Message=%s", result.Code, result.Message)
    return
}

// 处理成功的结果
fmt.Printf("NodePool created successfully: %+v\\n", result.Data)
```

## 注意事项

1. **节点池名称规范**: 1-26个字符，只能包含字母、数字和连字符，必须以字母或数字开头和结尾
2. **密码要求**: 必须包含大小写字母、数字和特殊字符，长度8-30位
3. **磁盘配置**: 系统卷必须是SSD且大小为40GB，数据卷大小必须是80GB的倍数
4. **标签限制**: 最多可以添加5个标签
5. **节点数量**: 伸缩时节点数量范围为0-99
6. **实例类型选择**:
   - 仅支持以下实例类型：
     - ECS: 推理型智算云主机、CPU计算型C11.2c4g、CPU计算型C11.8c16g
     - BMS: 推理型GPU裸金属
7. **节点类型与实例类型匹配**:
   - BMS 节点只能使用 GPU 规格
   - ECS 节点可以使用 GPU 或 CPU 规格
8. **计费方式建议**:
   - ECS 节点推荐使用按需付费（灵活性更好）
   - BMS 节点推荐使用包月（成本更优）

## 完整示例

参考 `examples/eks/nodepool.go` 文件查看完整的使用示例。

## 相关文档

- [EKS 总体文档](./EKS.md)

---

如有问题，请参考项目文档或联系技术支持。