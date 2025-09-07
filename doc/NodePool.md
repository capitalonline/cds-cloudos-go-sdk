# EKS NodePool SDK 使用指南

本文档详细介绍如何使用 EKS NodePool 相关的 SDK 方法，包括节点池的创建、查询、扩缩容和删除操作。

## 功能概述

EKS NodePool SDK 提供了完整的 Kubernetes 节点池管理功能：

- **CreateNodePool**: 创建新的节点池，支持 ECS 云主机和 BMS 裸金属两种节点类型
- **ListNodePool**: 查询指定集群中的所有节点池信息
- **ScalingNodePool**: 弹性伸缩节点池中的节点数量
- **DeleteNodePool**: 安全删除指定的节点池

支持多种场景：
- 🖥️ **ECS-CPU** 节点池：适用于通用计算工作负载
- 🚀 **ECS-GPU** 节点池：适用于AI/ML推理任务
- ⚡ **BMS-GPU** 节点池：适用于高性能计算场景
- 💰 **按需付费/包年包月**：灵活的计费方式

## 快速开始

### 初始化客户端

```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"

// 替换为您的实际访问密钥
ak := "your_access_key"
sk := "your_secret_key"

// 创建 EKS 客户端实例
client, err := eks.NewClient(ak, sk)
if err != nil {
    log.Fatalf("初始化 EKS 客户端失败: %v", err)
}
```

## 节点池管理示例

### 1. 创建 ECS-CPU 节点池（按需付费）

```go
func CreateECSNodePool() {
    ak, sk := "your-ak", "your-sk"  // 替换为您的实际密钥
    eksClient, _ := eks.NewClient(ak, sk)
    
    req := &eks.CreateNodePoolReq{
        ClusterId: "cluster-01",  // 集群ID，必须是已存在的集群
        VpcId:     "vpc-01",      // VPC ID，必须与集群在同一VPC
        Config: eks.NodePoolConfiguration{
            PoolName: "ecs-cpu-node-pool",  // 节点池名称，1-26个字符
            NodeType: eks.NodePoolNodeTypeECS,  // 节点类型：ECS 云主机
            SubjectId: 0,  // 测试金项目ID; 如果没有申请默认不传或传0
            NodeConfig: eks.NodePoolNodeConfig{
                BillingSpec: eks.NodePoolBillingSpec{
                    BillingMethod: eks.NodePoolBillingMethodPostPaid,  // 按需付费
                },
                SystemDisk: eks.NodePoolDiskInfo{
                    DiskType: eks.NodePoolDiskTypeSSD,  // 系统盘类型：SSD
                    DiskSize: 40,  // 系统盘为固定值40GB
                },
                DataDisk: []eks.NodePoolDiskInfo{
                    {
                        DiskType: eks.NodePoolDiskTypeSSD,  // 数据盘类型：SSD
                        DiskSize: 80,  // 数据盘大小：最小值80，步长80
                    },
                },
                OsImageName: eks.EcsUbuntu2204K8s13014Cpu,  // 实例镜像名称
                SubnetIds: []string{"subnet-01", "subnet-02"},  // VPC子网ID
                InstanceTypeIds: []string{
                    eks.EcsCpuC11Compute2XLarge,  // CPU计算型C11 8C16G
                },
                Password: "YourPassword123!",  // eks用户登录密码
                Shell: "#!/bin/bash\necho 'ECS CPU Node initialization complete'",
                Labels: map[string]string{
                    "node-type": "ecs-cpu",
                    "team":      "backend",
                },
            },
            Replicas: 3,  // 期望节点数量
        },
    }
    
    response, err := eksClient.CreateNodePool(req)
    if err != nil {
        fmt.Printf("创建节点池失败: %v\n", err)
        return
    }
    fmt.Printf("节点池创建成功: ID=%s, TaskId=%s\n", 
        response.Data.NodePoolId, response.Data.TaskId)
}
```

### 2. 创建 ECS-GPU 节点池（按需付费）

```go
func CreateECSGPUNodePool() {
    ak, sk := "your-ak", "your-sk"
    eksClient, _ := eks.NewClient(ak, sk)
    
    req := &eks.CreateNodePoolReq{
        ClusterId: "cluster-02",
        VpcId:     "vpc-02",
        Config: eks.NodePoolConfiguration{
            PoolName: "ecs-gpu-node-pool",  // GPU 节点池名称
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
                        DiskSize: 80,   // 第一块数据盘
                    },
                    {
                        DiskType: eks.NodePoolDiskTypeSSD,
                        DiskSize: 160,  // 第二块数据盘，可按需添加多块
                    },
                },
                OsImageName: eks.EcsUbuntu2204K8s13014Cpu,
                SubnetIds:   []string{"subnet-01"},
                InstanceTypeIds: []string{
                    eks.EcsGpuGch4XLarge,  // GPU推理型实例
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
        fmt.Printf("创建 GPU 节点池失败: %v\n", err)
        return
    }
    fmt.Printf("GPU 节点池创建成功: ID=%s, TaskId=%s\n", 
        response.Data.NodePoolId, response.Data.TaskId)
}
```

### 3. 创建 BMS-GPU 节点池（按需付费）

```go
func CreateBMSNodePoolPostPaid() {
    ak, sk := "your-ak", "your-sk"
    eksClient, _ := eks.NewClient(ak, sk)
    
    req := &eks.CreateNodePoolReq{
        ClusterId: "cluster-03",
        VpcId:     "vpc-03",
        Config: eks.NodePoolConfiguration{
            PoolName:  "bms-gpu-postpaid-node-pool",
            NodeType:  eks.NodePoolNodeTypeBMS,  // 裸金属节点类型
            SubjectId: 0,
            NodeConfig: eks.NodePoolNodeConfig{
                BillingSpec: eks.NodePoolBillingSpec{
                    BillingMethod: eks.NodePoolBillingMethodPostPaid,
                },
                // 首云裸金属暂不支持挂载云盘，使用裸金属本地盘
                SystemDisk: eks.NodePoolDiskInfo{},
                DataDisk:   []eks.NodePoolDiskInfo{},
                OsImageName: eks.BmsUbuntu2204K8s13014GpuRtx4090,  // 裸金属实例镜像
                SubnetIds: []string{"subnet-03", "subnet-04"},
                InstanceTypeIds: []string{
                    eks.BmsGpuGbm32XLarge,  // 推理型智算云主机igch.c8.nr4 16C64G
                },
                Password: "YourPassword123!",
                Shell:    "#!/bin/bash\necho 'BMS GPU PostPaid Node initialization complete'",
                Labels: map[string]string{
                    "node-type":    "bms-gpu-rtx-4090",
                    "billing-type": "postpaid",
                },
            },
            Replicas: 1,  // 裸金属通常数量较少且昂贵
        },
    }
    
    response, err := eksClient.CreateNodePool(req)
    if err != nil {
        fmt.Printf("创建裸金属节点池失败: %v\n", err)
        return
    }
    fmt.Printf("裸金属节点池创建成功: ID=%s, TaskId=%s\n", 
        response.Data.NodePoolId, response.Data.TaskId)
}
```

### 4. 创建 BMS-GPU 节点池（包年包月）

```go
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
                    BillingMethod: eks.NodePoolBillingMethodPrePaid,  // 包年包月
                    Duration:      3,   // 购买时长3个月
                    IsToMonth:     1,   // 是否购买至当月底
                    AutoRenew:     1,   // 到期开启自动续费
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
            Replicas: 1,
        },
    }
    
    response, err := eksClient.CreateNodePool(req)
    if err != nil {
        fmt.Printf("创建包月裸金属节点池失败: %v\n", err)
        return
    }
    fmt.Printf("包月裸金属节点池创建成功: ID=%s, TaskId=%s\n", 
        response.Data.NodePoolId, response.Data.TaskId)
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
| InstanceTypeIds | []string | 是 | 实例类型ID列表（只能指定一种实例类型） |
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

- [EKS容器服务 总体文档](./EKS.md)

---

如有问题，请参考项目文档或联系技术支持。