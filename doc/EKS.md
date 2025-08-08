
# GPU容器服务

# 概述
本文档主要介绍GPU容器服务 GO SDK的使用

# 使用指南

## 获取密钥
要使用首云GPU容器服务，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问GPU容器服务做签名验证。


## 代码示例

```go
    	// Init a credential with Access Key Id and Secret Access Key
	// You can apply them from the CDS web portal
	credential := common.NewCredential(
		os.Getenv("CDS_SECRET_ID"),
		os.Getenv("CDS_SECRET_KEY"),
	)
	
	...
```
# SDK接口列表

- [X] EKS容器服务
    - [X] 集群管理
        - [X] CreateCluster
        - [X] GetCluster
        - [X] ListClusters
        - [X] DeleteCluster
        - [X] GetClusterEvents
    - [X] 节点管理
        - [X] CreateNode
        - [X] GetNode
        - [X] ListNodes
        - [X] DeleteNode
        - [X] OperateNode
    - [X] 节点池管理(待开发)
        - [X] CreateNodePool
        - [X] GetNodePool
        - [X] AttachNodeToNodePool
        - [X] ScaleDownNodePool
    - [X] 自动伸缩管理
        - [X] CreateScale 
        - [X] UpdateScale 
        - [X] CreateScaleGroup 
        - [X] GetScaleGroup 
        - [X] ListScaleGroups 
        - [X] DeleteScaleGroup 