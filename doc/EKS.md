
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
        - [X] [CreateCluster](#CreateCluster)
        - [X] GetCluster
        - [X] ListClusters
        - [X] DeleteCluster
        - [X] GetClusterEvents
        - [X] [GetTaskStatus](#GetTaskStatus)
    - [X] 节点管理
        - [X] ListNodes
        - [X] DeleteNode
    - [X] 节点池管理
        - [X] CreateNodePool
        - [X] UpdateNodePoolConfig
        - [X] ScalingNodePoolNum
        - [X] ListNodePools
        - [X] DeleteNodePool
    - [X] 弹性网卡ENI管理
        - [X] CreateNetworkInterface 
        - [X] DeleteNetworkInterface 
        - [X] AttachNetworkInterface 
        - [X] DetachNetworkInterface 
        - [X] DescribeNetworkInterfaces 
## CreateCluster
使用以下代码可以创建一个EKS Cluster。
```
// import "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
req := eks.CreateClusterReq{
        // 集群名称
        ClusterName: "zbh-sdk-test",
        VpcId:       "cc32706e-a09a-11ef-8de7-06aabd6a5066",
        Cni: eks.CniInfo{
            CniType: eks.CniTypeFlannel,
            CniConfig: eks.CniConfig{
                PodCidr:            "",
                FlannelBackendType: eks.FlannelBackendTypeVxlan,
                NodePodsNum:        eks.NodePodsNum64,
                SubnetList: []eks.SubnetInfo{
                    {
                        SubnetId: "SubnetId",
                    },
                },
            },
            ProxyConfig: eks.ProxyConfigIPVS,
            ServiceCidr: eks.ServiceCidr192_16,
        },
        K8sVersion:       eks.K8sVersion1_30_14,
        NatId:            "ea4f7218-a09a-11ef-8de7-06aabd6a5066",
        SourceEipId:      "491dd24e-83c1-11f0-a1c5-268f674581cc",
        DestinationEipId: "491dd24e-83c1-11f0-a1c5-268f674581cc",
        SlbId:            "",
        Ak:               "77a166e8f32611ef8e00d67a82aa41e6",
        Sk:               "76387540dc16f17c0b9c85867dcefce0",
        MasterNumber:     eks.MasterNumber3,
        MasterConfig: eks.NodeConfig{
            BillingSpec: eks.BillingSpec{
                AutoRenew:     eks.AutoRenew0,
                BillingMethod: eks.BillingMethod0,
                Duration:      eks.Duration1,
                IsToMonth:     eks.IsToMonth0,
                },
            SystemDisk: eks.DiskInfo{
                DiskFeature: eks.DiskFeatureSsd,
                DiskSize:    40,
                },
            DataDisk: []eks.DiskInfo{
                {
                    DiskFeature: eks.DiskFeatureSsd,
                    DiskSize:    80,
                    Number:      1,
                },
            },
            OsImageName: eks.OsImageNameUbuntu2204_1_30_14,
            SubnetList: []eks.SubnetInfo{
				{
					SubnetId: "c0ba8808-048f-11f0-a617-3ed2fcc2525f",
				},
			},
			Shell: "",
			Specifics: []eks.Specifics{
				{
					ProductName: eks.ProductName1,
				},
			},
			Password:    "***********",
			Labels:      map[string]string{},
			Annotations: map[string]string{},
		},
		NodePoolConfig: eks.NodePoolConfig{
			PoolName:  "node-pool-1",
			NodeType:  eks.NodeTypeEcs,
			SubjectId: 0,
			NodeConfig: eks.NodeConfig{
				BillingSpec: eks.BillingSpec{
					AutoRenew:     eks.AutoRenew0,
					BillingMethod: eks.BillingMethod0,
					Duration:      eks.Duration1,
					IsToMonth:     eks.IsToMonth1,
				},
				SystemDisk: eks.DiskInfo{
					DiskFeature: eks.DiskFeatureSsd,
					DiskSize:    40,
				},
				DataDisk: []eks.DiskInfo{
					{
						DiskFeature: eks.DiskFeatureSsd,
						DiskSize:    80,
						Number:      1,
					},
				},
				OsImageName: eks.OsImageNameUbuntu2204_1_30_14,
				SubnetList: []eks.SubnetInfo{
					{
						SubnetId: "c0ba8808-048f-11f0-a617-3ed2fcc2525f",
					},
				},
				Specifics: []eks.Specifics{
					{
						ProductName: eks.ProductName1,
					},
				},
				Shell:       "",
				Password:    "***********",
				Labels:      map[string]string{},
				Annotations: map[string]string{},
			},
			AutoScaling: eks.AutoScaling{
				Enable:         false,
				MaxReplicas:    10,
				MinReplicas:    1,
				Priority:       1,
				SubnetPolicy:   "random",
				ScalingGroupId: "",
			},
			Replicas: 1,
		},
	}
	response, err := eksClient.CreateCluster(&req)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
 ```
## 创建集群参数说明
|请求参数|类型|必填| 描述                                         |
|------|----|----|--------------------------------------------|
|ClusterName|string|是| 集群名称                                       |
|VpcId|string|是| 集群所在VPC的ID                                 |
|Cni|[CniInfo](#CniInfo)|是| 集群网络配置信息                                   |
|K8sVersion|string|是| 集群版本                                       |
|NatId|string|是| 集群所在VPC的NAT网关ID                            |
|SourceEipId|string|是| 集群nat的公网EIP的ID                             |
|DestinationEipId|string|是| 集群nat的公网EIP的ID，如果SlbId不为空，则需要slb所挂载的EIP id |
|SlbId|string|否| 集群使用的slb实例ID                               |
|Ak|string|是| 集群CCM所使用的ak                                |                                  |
|Sk|string|是| 集群CCM所使用的sk                                     |
|MasterNumber|int|是| 集群主节点数量                                     |
|MasterConfig|[NodeConfig](#NodeConfig)|是| 集群主节点配置信息                                 |
|NodePoolConfig|[NodePoolConfig](#NodePoolConfig)|是| 集群节点池配置信息                               |
## CniInfo
|请求参数|类型|必填| 描述                              |
|------|----|----|---------------------------------|
|CniType|string|是| 集群内部网络类型，flannel，calico，vpc-cni |
|CniConfig|[CniConfig](#CniConfig)|是| 集群内部网络配置信息                  |
|ProxyConfig|string|是| 集群内部网络代理配置，ipvs，iptables |
|ServiceCidr|string|是| 集群服务CIDR                        |
## CniConfig
|请求参数| 类型                        | 必填 | 描述                                   
|------|---------------------------|----|--------------------------------------|
|PodCidr| string                    | 是  | 集群pod网络，使用vpc-cni时无需给空值，其他网络类型需要明确指定 |
|FlannelBackendType| string                    | 否  | 使用flannel时指定backend类型，可选值vxlan       |
|NodePodsNum| int                       | 否  | 每个节点pod数量，默认64                       |
|SubnetList| [SubnetInfo](#SubnetInfo) | 否  | 使用vpc-cni时指定子网信息                     |
|CalicoBlockSize| int                       | 否  | 使用calico时指定blockSize                   |
|CalicoEncapsulation| string                    | 否  | 使用calico时指定是否使用ipip               |
|
## SubnetInfo
|请求参数| 类型     | 必填 | 描述
|------|--------|----|--------------------------------------|
|SubnetId| string | 是  | 子网ID
## NodeConfig
|请求参数|类型|必填| 描述                                         |
|------|--------|----|--------------------------------------|
|BillingSpec|[BillingSpec](#BillingSpec)|是| 集群节点计费信息                                   |
|SystemDisk|[DiskInfo](#DiskInfo)|是| 集群节点系统盘信息                                   |
|DataDisk|[DiskInfo](#DiskInfo)|是| 集群节点数据盘信息                                   |
|OsImageName|string|是| 集群节点镜像名称                                   |
|SubnetList|[SubnetInfo](#SubnetInfo)|是| 集群节点所在子网信息                                |
|Specifics|[Specifics](#Specifics)|是| 集群节点规格信息                                   |
|Shell|string|否| 集群节点自定义shell                                |
|Password|string|是| 集群节点密码                                       |
|Labels|map[string]string|否| 集群节点标签                                       |
|Annotations|map[string]string|否| 集群节点注释                                       |
## BillingSpec
|请求参数|类型|必填| 描述                                         |
|------|--------|----|--------------------------------------|
|AutoRenew|int|是| 集群节点自动续费 , 0: 不自动续费, 1: 自动续费                                  |
|BillingMethod|string|是| 集群节点计费方式, 0: 按量计费, 1: 包年包月                                  |
|Duration|int|是| 集群节点包年包月时长, 1: 1个月, 2: 2个月, 3: 3个月, 4: 4个月, 5: 5个月, 6: 6个月, 7: 7个月, 8: 8个月, 9: 9个月, 10: 10个月, 11: 11个月, 12: 1年, 24: 2年, 36: 3年                                  |
|IsToMonth|int|是| 是否购买到月底, 0: 购买整月, 1: 购买到月底                                  |
## DiskInfo
|请求参数|类型|必填| 描述                                         |
|------|--------|----|--------------------------------------|
|DiskFeature|string|是| 集群节点磁盘特性, ssd: ssd磁盘                                  |
|DiskSize|int|是| 集群节点磁盘大小, 单位G                                  |
## Specifics
|请求参数|类型|必填| 描述                                         |
|------|--------|----|--------------------------------------|
|ProductName|string|是| 集群节点规格名称, e.g. Optimized M6,2,4,0                                  |
## NodePoolConfig
| 请求参数     | 类型     |必填| 描述                |
|----------|--------|----|-------------------|
| PoolName | string |是| 集群节点池名称           |
| NodeType | string |是| 集群节点池节资源类型        |
|SubjectId|int|否| 测试金编号，默认为0：不使用测试金 |
|NodeConfig|[NodeConfig](#NodeConfig)|是| 集群worker节点配置信息    |
|AutoScaling|[AutoScaling](#AutoScaling)|是| 集群节点池自动伸缩配置信息 |
## AutoScaling
| 请求参数     | 类型     | 必填 | 描述                |
|----------|--------|----|-------------------|
|Enable| bool| 是  | 集群节点池是否开启自动伸缩功能 |
|MaxReplicas| int| 是  | 集群节点池最大节点数,最大值为99|
|MinReplicas| int| 是  | 集群节点池最小节点数,最小值为1|
|Priority| int| 否  | 集群节点池cpu目标值,最小值为1,最大值为100|
|SubnetPolicy| string| 否  | 集群节点池子网分配策略,可选值: random|
|ScalingGroupId| string| 否  | 集群节点池伸缩组id|

## 返回参数
| 参数名称      | 类型                         |描述|
|-----------|----------------------------|----|
| RequestId | string                     |唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求ID。|
| Code      | string                     |集群ID|
|Msg| string                     |返回信息|
|Data| [map[string]string](#Data) |返回数据|
## Data
|参数名称|类型|描述|
|------|--------|----|
|ClusterId|string|集群ID|
|TaskId|string|任务ID|

## GetTaskStatus
使用以下代码可以获取任务状态
```
ak, sk := "ak", "sk"
endpoint := "http://gateway.gic.test"
eksClient, _ := eks.NewClientWidthEndpoint(ak, sk, endpoint)

response, err := eksClient.GetTaskStatus("cf8a2f1e-8941-11f0-8366-9ea42c0f2d3c")
if err != nil {
	fmt.Println(err)
}
fmt.Printf(">>> response: %+v\n", response)

fmt.Println(response.RequestId)
bytes, _ := json.Marshal(response)
fmt.Println(string(bytes))
```
## GetTaskStatus参数说明
|参数名称|类型|描述|
|------|--------|----|
|TaskId|string|任务ID|
## GetTaskStatus返回参数
|参数名称|类型|描述|
|------|--------|----|
|RequestId|string|唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求ID。|
|Code|string|返回码|
|Msg|string|返回信息|
|Data|[map[string]string](#Data)|返回数据|
## Data
|参数名称|类型|描述|
|------|--------|----|
|TaskId|string|任务ID|
|TaskStatus|string|任务状态|
|TaskMsg|string|任务信息|







