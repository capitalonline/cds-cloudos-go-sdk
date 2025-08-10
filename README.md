[English](./README.md) | 简体中文

<p align="center">
<a href=" https://www.alibabacloud.com"><img src="https://www.capitalonline.net/templets/default/icon/logo_header.png"></a>
</p>

<h1 align="center">CapitalOnline Cloud SDK for Go</h1>

欢迎使用`CapitalOnline Cloud SDK for Go`，它可以管理[首都在线](https://www.capitalonline.net)多个全球服务，如`首云全球云服务器`、`EKS容器服务`、`GPU裸金属实例`、`首云私有网络`等，帮您轻松管理所有在线资源。基于首云官方[OpenAPI文档](https://github.com/capitalonline/openapi/blob/master/README.md)。

## 功能特点

你可以通过SDK管理各种资源，[这里](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md)可以看到所有可用接口清单。下面列出部分接口列表：

- [X] EKS容器服务
    - [X] 集群管理
        - [X] CreateCluster
        - [X] GetCluster
        - [X] ListClusters
        - [X] DeleteCluster
        - [X] GetClusterEvents
    - [X] 节点管理
        - [X] CreateNode
        - [X] DeleteNode
        - [X] OperateNode
    - [X] 节点池管理(待开发)
        - [X] CreateNodePool
        - [X] GetNodePool
        - [X] AttachNodeToNodePool
        - [X] ScaleDownNodePool
    - [X] 自动伸缩管理
        - [X] 创建自动伸缩
        - [X] 修改自动伸缩
        - [X] 创建伸缩组
        - [X] 查看伸缩组列表
        - [X] 修改伸缩组
        - [X] 删除伸缩组


- [X] 私有网络管理
    - [X] DescribeVpc

- [X] GPU裸金属管理
    - [X] CreateBms

- [X] 还有更多

## 使用指南

使用 `go get` 下载安装 SDK

```sh
$ go get -u github.com/capitalonline/cds-cloudos-go-sdk
```

## 使用示例

```go
    	// Init a credential with Access Key Id and Secret Access Key
	// You can apply them from the CDS web portal
	credential := common.NewCredential(
		os.Getenv("CDS_SECRET_ID"),
		os.Getenv("CDS_SECRET_KEY"),
	)
	
	...
```

更多示例参见 [example](./example) 目录。

## 如何贡献

欢迎提交 Issue 或 Pull Request。

## 相关参考

- [CDS OpenAPI Explorer](https://github.com/capitalonline/openapi)

## 许可证

[Apache License v2.0](./LICENSE)
