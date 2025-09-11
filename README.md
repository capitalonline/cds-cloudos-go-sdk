# CDS-CloudOS Go SDK 文档

<p align="center">
<a href="https://www.capitalonline.net"><img src="https://www.capitalonline.net/templets/default/icon/logo_header.png"></a>
</p>

<h1 align="center">首都在线 Go SDK</h1>

# 概述

本文档主要介绍首都在线（CDS）Go语言版的开发者工具包（SDK），用户可基于该SDK使用Go语言接入首都在线的各项云服务产品（详见支持产品列表）。SDK封装了便捷的调用接口，保持了统一的认证方式、调用接口，提供了标准的错误码和返回格式，方便开发者调试和使用。

## 支持的服务

- **EKS容器服务** - 企业级Kubernetes容器编排服务
- **GPU裸金属服务** - 高性能GPU计算实例
- **私有网络VPC** - 云上私有网络环境
- **弹性公网IP** - 灵活的公网IP资源
- **负载均衡SLB** - 高可用负载均衡服务
- **NAT网关** - 网络地址转换服务
- **带宽包管理** - 网络带宽资源管理
- **子网管理** - VPC内子网资源管理

# 安装SDK工具包

## 运行环境

CDS Go SDK可以在Go 1.18及以上环境下运行。我们建议使用最新的Go版本以获得最佳性能和安全性。

## 安装SDK

**使用go mod安装（推荐）**

在您的项目中使用`go get`工具进行安装：

```shell
go get github.com/capitalonline/cds-cloudos-go-sdk
```

**从源码编译**

```shell
git clone https://github.com/capitalonline/cds-cloudos-go-sdk.git
cd cds-cloudos-go-sdk
go mod tidy
go build ./...
```

## SDK目录结构

```text
cds-cloudos-go-sdk
|--auth                   //CDS签名和权限认证
|--cds                    //CDS公用基础组件  
|--http                   //CDS的HTTP通信模块
|--services               //CDS相关服务目录
|  |--eks                 //EKS容器服务
|  |  |--client.go        //EKS客户端入口
|  |  |--cluster.go       //集群管理
|  |  |--node.go          //节点管理
|  |  |--nodepool.go      //节点池管理
|  |  |--nodepool_model.go//节点池数据模型
|  |  |--model.go         //EKS数据模型
|  |  |--network_interface.go //弹性网卡
|  |--vpc                 //私有网络服务
|  |  |--client.go        //VPC客户端入口
|  |  |--api.go           //VPC相关
|  |  |--model.go         //VPC数据模型
|  |--eip                 //弹性公网IP服务
|  |--slb                 //负载均衡服务
|  |--natgateway          //NAT网关服务
|  |--subnet              //子网服务
|  |--bandwidthpackage    //带宽包服务
|--examples               //使用示例目录
|  |--eks                 //EKS服务使用示例
|  |--vpc                 //VPC服务使用示例
|--doc                    //详细文档目录
|  |--EKS.md              //EKS服务文档
|  |--NodePool.md         //节点池专用文档
|  |--NETWORK.md          //网络服务文档
|--util                   //CDS公用的工具实现
```

## 卸载SDK

如需卸载SDK，删除下载的源码并从go.mod中移除依赖即可：

```shell
go mod edit -droprequire github.com/capitalonline/cds-cloudos-go-sdk
go mod tidy
```

# 使用步骤

## 获取访问凭证

在使用SDK之前，您需要在首都在线控制台获取Access Key ID和Secret Access Key。

1. 登录[首都在线控制台](https://console.capitalonline.net)
2. 进入"访问控制" > "API密钥管理"
3. 创建或查看您的Access Key ID和Secret Access Key

## 确认Endpoint

根据您使用的服务和所在地域，确认相应的服务端点（Endpoint）。CDS服务通常使用统一的API端点，SDK会自动处理不同服务的路由。

## 创建Client对象

每种具体的服务都有一个`Client`对象，为开发者与对应的服务进行交互封装了一系列易用的方法。

```go
import (
    "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

// 创建EKS服务客户端
eksClient, err := eks.NewClient("your-access-key-id", "your-secret-access-key")
if err != nil {
    panic(err)
}
```

## 调用功能接口

基于创建的对应服务的`Client`对象，即可调用相应的功能接口，使用CDS云产品的功能。

## 示例

以下以EKS容器服务为例，演示基本的使用方式：

### EKS集群管理示例

```go
package main

import (
    "fmt"
    "os"
    
    "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func main() {
    // 从环境变量获取认证信息
    ak := os.Getenv("CDS_SECRET_ID")
    sk := os.Getenv("CDS_SECRET_KEY")
    
    // 创建EKS客户端
    client, err := eks.NewClient(ak, sk)
    if err != nil {
        panic(err)
    }
    
    // 查询集群列表
    listReq := &eks.ListClustersReq{
        VpcId: "your-vpc-id",
    }
    
    result, err := client.ListClusters(listReq)
    if err != nil {
        fmt.Printf("查询集群列表失败: %v\\n", err)
        return
    }
    
    fmt.Printf("找到 %d 个集群\\n", result.Data.Total)
    for _, cluster := range result.Data.ClusterList {
        fmt.Printf("集群: %s, 状态: %s\\n", cluster.ClusterName, cluster.ClusterStatus)
    }
}
```

### NodePool节点池管理示例

```go
package main

import (
  "fmt"
  "os"

  "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func main() {
  // 创建客户端
  client, err := eks.NewClient(os.Getenv("CDS_SECRET_ID"), os.Getenv("CDS_SECRET_KEY"))
  if err != nil {
    panic(err)
  }

  // 创建GPU裸金属节点池
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

  result, err := client.CreateNodePool(createReq)
  if err != nil {
    fmt.Printf("创建节点池失败: %v\\n", err)
    return
  }

  fmt.Printf("节点池创建成功: ID=%s, TaskId=%s\\n",
    result.Data.NodePoolId, result.Data.TaskId)
}

```

# 配置

## 使用HTTPS协议

CDS Go SDK默认使用HTTPS协议确保数据传输安全。所有API调用都通过加密通道进行。

## 详细配置

开发者使用CDS Go SDK时，可以对客户端进行详细配置：

### 环境变量配置

推荐使用环境变量配置敏感信息：

```bash
export CDS_SECRET_ID="your-access-key-id"
export CDS_SECRET_KEY="your-secret-access-key"
```

### 代码配置示例

```go
// 基础配置
client, err := eks.NewClient(ak, sk)

// 自定义HTTP客户端配置
client.Config.ConnectionTimeoutInMillis = 30 * 1000  // 30秒超时
```

### 配置选项说明

| 配置项名称 | 类型 | 默认值 | 说明 |
|-----------|------|--------|------|
| ConnectionTimeoutInMillis | int | 120000 | 连接超时时间（毫秒） |
| UserAgent | string | cds-go-sdk | HTTP请求的User-Agent |
| MaxRetries | int | 3 | 最大重试次数 |

## 重试策略

SDK内置了智能重试机制：

- **指数退避算法**: 重试间隔按指数级增长
- **最大重试次数**: 默认3次，可配置
- **重试条件**: 网络错误、5xx服务器错误等临时性错误

```go
// 配置重试策略
client.Config.MaxRetries = 5                    // 最大重试5次
client.Config.RetryBaseDelay = 1000             // 基础延迟1秒
```

# 错误处理

CDS Go SDK定义了完善的错误处理机制，包含以下错误类型：

## 错误类型

| 错误类型 | 说明 | 示例 |
|---------|------|------|
| ClientError | 客户端错误，如网络连接问题 | 网络超时、连接失败 |
| ServerError | 服务端错误，如API调用失败 | 参数错误、权限不足 |
| ValidationError | 参数验证错误 | 必填参数缺失、格式不正确 |

## 错误处理示例

```go
result, err := client.CreateNodePool(createReq)
if err != nil {
    switch realErr := err.(type) {
    case *cds.ClientError:
        fmt.Printf("客户端错误: %s\\n", realErr.Error())
        fmt.Printf("错误码: %s\\n", realErr.GetErrorCode())
    case *cds.ServerError:
        fmt.Printf("服务端错误: %s\\n", realErr.Error())
        fmt.Printf("错误码: %s\\n", realErr.GetErrorCode())
        fmt.Printf("请求ID: %s\\n", realErr.GetRequestId())
    default:
        fmt.Printf("未知错误: %s\\n", err.Error())
    }
    return
}

// 检查业务错误码
if result.Code != "Success" {
    fmt.Printf("API调用失败: %s - %s\\n", result.Code, result.Message)
    return
}
```

## 常见错误码

| 错误码 | 说明 | 解决方案 |
|-------|------|---------|
| ParamError | 参数错误 | 检查请求参数是否正确 |
| AuthFailure | 认证失败 | 检查Access Key和Secret Key |
| ResourceNotFound | 资源不存在 | 检查资源ID是否正确 |
| InsufficientBalance | 余额不足 | 充值账户余额 |
| QuotaExceeded | 配额超限 | 申请提升配额或清理资源 |

# 服务详细文档

## EKS容器服务

EKS（Elastic Kubernetes Service）是首都在线提供的企业级Kubernetes容器编排服务。

### 支持的功能

- ✅ **集群管理**
    - 创建/删除/查询集群
    - 集群事件查询
    - 集群配置更新

- ✅ **节点管理**
    - 查询节点列表
    - 删除节点
    - 节点状态监控

- ✅ **节点池管理**
    - 创建/删除节点池
    - 节点池列表查询
    - 节点池伸缩

### 实例类型支持

| 节点类型 | 实例规格 | 适用场景 |
|---------|---------|---------|
| ECS云主机 | 推理型智算云主机 (GPU) | AI推理、机器学习 |
| ECS云主机 | CPU计算型C11.2c4g | 通用计算负载 |
| ECS云主机 | CPU计算型C11.8c16g | 高性能计算 |
| BMS裸金属 | 推理型GPU裸金属 | 高性能GPU计算 |

详细使用说明请参考：[EKS服务文档](./doc/EKS.md) | [NodePool专用文档](./doc/NodePool.md)

## VPC私有网络

VPC为您提供了一个在云上的私有网络环境。

### 支持的功能

- ✅ VPC查询和管理
- ✅ 子网管理
- ✅ SLB负载均衡


详细SDK文档请参考doc目录各服务的说明文档。

# 最佳实践

## 认证信息安全

1. **使用环境变量**: 避免在代码中硬编码Access Key
2. **定期轮换密钥**: 定期更新Access Key和Secret Key
3. **最小权限原则**: 为SDK分配最小必要权限

```go
// ❌ 不推荐：硬编码
client, _ := eks.NewClient("AKID...", "SecretKey...")

// ✅ 推荐：使用环境变量
client, _ := eks.NewClient(os.Getenv("CDS_SECRET_ID"), os.Getenv("CDS_SECRET_KEY"))
```

## 错误处理

1. **完整的错误处理**: 对所有API调用进行错误检查
2. **记录错误信息**: 保存RequestId用于问题排查
3. **优雅降级**: 设计合理的降级策略

```go
result, err := client.CreateNodePool(req)
if err != nil {
    log.Printf("创建节点池失败: %v, RequestId: %s", err, getRequestId(err))
    // 实现降级逻辑
    return handleFallback()
}
```

## 性能优化

1. **连接复用**: 重用Client对象避免频繁创建
2. **并发控制**: 合理控制并发API调用数量
3. **分页查询**: 对大量数据使用分页接口

```go
// ✅ 重用客户端
var eksClient *eks.Client

func init() {
    eksClient, _ = eks.NewClient(os.Getenv("CDS_SECRET_ID"), os.Getenv("CDS_SECRET_KEY"))
}

// ✅ 并发控制
sem := make(chan struct{}, 10) // 限制并发数为10
for _, task := range tasks {
    sem <- struct{}{}
    go func(t Task) {
        defer func() { <-sem }()
        processTask(t)
    }(task)
}
```

# 示例代码

完整的使用示例请参考 [examples](./examples) 目录：

- [EKS集群管理示例](./examples/eks/cluster.go)
- [NodePool节点池示例](./examples/eks/nodepool.go)
- [VPC网络管理示例](./examples/vpc/vpc.go)

# 版本历史

## v0.0.24 (最新)
- ✨ 新增NodePool节点池管理功能
- ✨ 支持GPU裸金属实例类型
- 🐛 修复网络超时问题
- 📚 完善文档和示例


## v0.0.1
- 🎉 首次发布
- ✨ 支持VPC、EKS、SLB等基础服务

# 社区支持

## 获取帮助

- 📖 [技术文档](./doc/)
- 🐛 [问题反馈](https://github.com/capitalonline/cds-cloudos-go-sdk/issues)
- 💬 [讨论区](https://github.com/capitalonline/cds-cloudos-go-sdk/discussions)


## 开发规范

- 遵循Go语言编码规范
- 添加必要的单元测试
- 更新相关文档
- 确保向后兼容性

# 相关链接

- 🌐 [首都在线官网](https://www.capitalonline.net)
- 📋 [OpenAPI文档](https://github.com/capitalonline/openapi)
- 🔧 [控制台](https://console.capitalonline.net)
- 📞 [技术支持](https://www.capitalonline.net/support)

# 许可证

本项目采用 [Apache License 2.0](./LICENSE) 许可证。

---

**注意**: 使用本SDK前，请确保您已经注册并开通了相应的首都在线云服务。