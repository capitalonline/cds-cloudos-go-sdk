# 私有网络 SDK使用指南
本文档详细介绍如何使用 私有网络 相关的 SDK 方法，包括VPC管理、子网管理、eip管理、共享带宽包管理等操作。

## 功能概述

**私有网络管理** 

- **CreateVpc**: 创建VPC信息
- **GetVpc**: 获取指定VPC信息
- **ListVpcs**: 查询VPC信息
- **DeleteVpc**: 删除VPC

**子网管理**
- **CreateSubnet**: 创建子网
- **GetSubnet**: 获取指定子网信息
- **ListSubnets**: 查询子网信息
- **DeleteSubnet**: 删除子网

**弹性公网IP(EIp)管理**
- **CreateEip**: 创建弹性eip
- **GetEip**: 获取指定弹性eip信息
- **ListEips**: 查询弹性eip
- **ReleaseEip**: 释放弹性eip
- **UpdateEip**: 更改弹性eip带宽信息

**共享带宽包管理**
- **CreateBandwidthPackage**: 创建共享带宽包
- **GetBandwidthPackage**: 获取指定共享带宽包信息
- **ListBandwidthPackage**: 查询共享带宽包
- **UpdateBandwidthPackage**: 更新共享带宽包
- **AddBandwidthPackageIp**: 向指定共享带宽包添加eip
- **RemoveBandwidthPackageIp**: 从指定带宽包移除弹性eip
- **DeleteBandwidthPackage**:  删除共享带宽包


**NAT网关管理**
- **GetNatGateway**: 获取指定nat网关信息
- **ListNatGateways**: 查询nat网关

**高性能负载均衡管理**
- **GetSlb**: 获取指定高性能负载均衡信息
- **ListSlb**:  查询高性能负载均衡信息

## 快速开始
### 初始化VPC客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/vpc"
ak, sk := "ak", "sk"
vpcClient, _ := vpc.NewClient(ak, sk)
```
### 初始化子网客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/subnet"
ak, sk :=  "ak", "sk"
subnetClient, _ := subnet.NewClient(ak, sk)
```

### 初始化EIP客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/eip"
ak, sk := "ak", "sk"
EipClient, _ := eip.NewClient(ak, sk)
```
### 初始化共享带宽包客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/bandwidthpackage"
ak, sk := "ak", "sk"
BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
```
### 初始化NAT网关客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/natgateway"
ak, sk := "ak", "sk"
natgatewayClient, _ := natgateway.NewClient(ak, sk)
```

## 代码示例
### VPC管理代码示例
**创建VPC**
```go
func CreateVpc() {
	ak, sk := "ak", "sk"

	vpcClient, _ := vpc.NewClient(ak, sk)
	CreateVpcArgs := &vpc.CreateVpcReq{
        RegionCode:     "CN_Qingyang",  // 区域code
        VPCName:        "go_create_bandtype",  // vpc名称
        VPCSegment:     "10.21.0.0/16", // vpc网段
        VPCSegmentType: "auto", // 网段类型
		BandwidthType: "Bandwidth_Multi_ISP_BGP",  // 线路类型
        SubnetList: []vpc.Subnet{
            {
                AvailableZoneCode: "CN_Qingyang_A",  // 可用区 code
                SubnetName:        "gocreate子网1",
                SubnetSegment:     "10.21.1.0/24", // 子网网段
            },
        },
    }

	response, err := vpcClient.CreateVPC(CreateVpcArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```

**获取指定VPC信息**
```go
func Getvpc() {
	ak, sk := "ak", "sk"

	vpcClient, _ := vpc.NewClient(ak, sk)
	GetVpcArgs := &vpc.GetVpcReq{
		Keyword: "9197340c-799e-11f0-adfa-6e18e986f14e",  // vpc id
	}
	response, err := vpcClient.GetVpc(GetVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.VPCList)
}
```
**查询VPC信息**
```go
func ListVpcs() {
	ak, sk := "ak", "sk"

	vpcClient, _ := vpc.NewClient(ak, sk)
	ListVpcArgs := &vpc.ListVpcsReq{
		RegionCode:"CN_Qingyang",  // 区域code
        Keyword: "", //   查询关键字
        PageNumber:10, // 列表页码。起始值：1, 默认值：1
	}
	response, err := vpcClient.ListVpcs(ListVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.VPCList)
}

```
**删除VPC**
```go
func DeleteVPC() {
	ak, sk := "ak", "sk"
	vpcClient, _ := vpc.NewClient(ak, sk)
	DeleteVpcArgs := &vpc.DeleteVpcReq{
		VPCId: "9197340c-799e-11f0-adfa-6e18e986f14e",  // 删除私有网络VPC
	}
	response, err := vpcClient.DeleteVpc(DeleteVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
### 子网管理代码示例
**创建子网**
```go
func CreateSubnet() {
	ak, sk :=  "ak", "sk"
	subnetClient, _ := subnet.NewClient(ak, sk)
	CreateSubnetArgs := &subnet.CreateSubnetReq{
        VPCId:        "9197340c-799e-11f0-adfa-6e18e986f14e",  // 私有网络VPC ID
        SubnetList: []subnet.CreateSubnetData{
            {
                AvailableZoneCode: "CN_Qingyang_A",  // 可用区code
                SubnetName:        "子网2",
                SubnetSegment:     "10.21.2.0/24", // 子网网段
            },
        },
    }
	response, err := subnetClient.CreateSubnet(CreateSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

```

**获取指定子网信息**
```go
func GetSubnet() {
	ak, sk :=  "ak", "sk"
	subnetClient, _ := subnet.NewClient(ak, sk)
	GetSubnetArgs := &subnet.GetSubnetReq{
		Keyword: "919a5290-799e-11f0-adfa-6e18e986f14e",  // 子网id或者名称
	}
	response, err := subnetClient.GetSubnet(GetSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
**查询子网信息**
```go
func ListSubnet() {
	ak, sk :=  "ak", "sk"

	subnetClient, _ := subnet.NewClient(ak, sk)
	ListSubnetArgs := &subnet.ListSubnetsReq{
		RegionCode: "CN_Qingyang",  // 区域code
        AvailableZoneCode: "", // 可用区code
        VPCId: "", // 私有网络vpc id
        PageNumber: "",// 页码
	}
	response, err := subnetClient.ListSubnets(ListSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```

**删除子网**
```go
func DeleteSubnet() {
	ak, sk :=  "ak", "sk"

	subnetClient, _ := subnet.NewClient(ak, sk)
	DeleteSubnetArgs := &subnet.DeleteSubnetReq{
		SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",   // 子网id
    }

	response, err := subnetClient.DeleteSubnet(DeleteSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

```
### EIP管理代码示例
**创建弹性eip**
```go
func CreateEip(){
	ak, sk := "ak", "sk"
	EipClient, _ := eip.NewClient(ak, sk)
	CreateEipArgs := &eip.CreateEIPReq{
		RegionCode: "CN_Qingyang",  // 区域code
        BandwidthType: "Bandwidth_Multi_ISP_BGP",  // 带宽类型
        BillScheme: "BandwIdth",  // 计费方案
        Qos: 5,  // 带宽大小
        Size: 1,  // 数量
        Description: "go create",
        SubjectId: 111, //测试金id
        ProjectId: "0-0", //项目组id, 默认为0-0，默认项目组
	}
	response, err := EipClient.CreateEip(CreateEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
**获取指定弹性eip信息**
```go
func GetEip() {
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	GetEipArgs := &eip.GetEipReq{
		Keyword: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",  // eip地址或者id
	}
	response, err := EipClient.GetEip(GetEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
```
**查询弹性eip**
```go
func ListEips() {
	ak, sk := "ak", "sk"
	EipClient, _ := eip.NewClient(ak, sk)
	ListEipArgs := &eip.ListEipsReq{
		RegionCode:"CN_Qingyang",
        Keyword:"127.0.0.1", // eip地址
	}
	response, err := EipClient.ListEips(ListEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
```
**释放弹性eip**
```go
func ReleaseEip(){
	ak, sk := "ak", "sk"


	EipClient, _ := eip.NewClient(ak, sk)
	ReleaseEipArgs := &eip.ReleaseEipReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
	}

	response, err := EipClient.ReleaseEip(ReleaseEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}
```
**更改弹性eip带宽信息**
```go
func UpdateEip(){
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	UpdateEipArgs := &eip.UpdateEIPReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
        Qos: 10,
		Description: "go create",
	}

	response, err := EipClient.UpdateEip(UpdateEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
### 共享带宽包管理代码示例
**创建共享带宽包**
```go
func CreateBandwidthPackage() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	CreatebandwidthpackageArgs := &bandwidthpackage.CreateBandwidthPackageReq{
        RegionCode: "CN_Qingyang",  // 	VPC区域code
        Name: "go_create",
        AvailableZoneCode: "CN_Qingyang_A",  // VPC可用区code
        BandwidthType: "Bandwidth_Multi_ISP_BGP", // 带宽类型
        BillScheme: "BandwIdth_Shared", // 计费方案
        Qos: 10,
        SubjectId: 111   //测试金id
    
	}
	response, err := BandwidthPackageClient.CreateBandwidthPackage(CreatebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```

**获取指定共享带宽包信息**
```go
func GetBandwidthPackage() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	GetBandwidthPackageArgs := &bandwidthpackage.GetBandwidthPackageReq{
		Keyword: "868bb384-79a4-11f0-adfa-6e18e986f14e", //查询关键字: id或名称
	}
	response, err := BandwidthPackageClient.GetBandwidthPackage(GetBandwidthPackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

```
**查询共享带宽包**
```go
func ListBandwidthPackage() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	ListBandwidthPackageArgs := &bandwidthpackage.ListBandwidthPackagesReq{
		RegionCode:"CN_Qingyang", // VPC区域code
        Keyword: "", // 共享带宽包id或名称
	}
	response, err := BandwidthPackageClient.ListBandwidthPackages(ListBandwidthPackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
```
**更新共享带宽包**
```go
func UpdateBandwidthPackage() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	UpdatebandwidthpackageArgs := &bandwidthpackage.UpdateBandwidthPackageReq{
        BandwidthId: "868bb384-79a4-11f0-adfa-6e18e986f14e",
        Qos: 20,
	}
	response, err := BandwidthPackageClient.UpdateBandwidthPackage(UpdatebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
**向指定共享带宽包添加eip**
```go
func AddBandwidthPackageIp() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	AddBandwidthPackageIpArgs := &bandwidthpackage.AddBandwidthPackageIpReq{
		BandwidthId:"868bb384-79a4-11f0-adfa-6e18e986f14e",  // 共享带宽包id
		EIPIdList: []string{"6870eeac-79ac-11f0-8503-6e18e986f14e"},  // eip id
	}
	response, err := BandwidthPackageClient.AddBandwidthPackageIp(AddBandwidthPackageIpArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

```
**从指定带宽包移除并保留弹性eip**
```go
func RemoveBandwidthPackageIp() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	// 移除保留eip
	RemoveBandwidthPackageIpArgs := &bandwidthpackage.RemoveBandwidthPackageIpReq{
		EIPIdList: []string{"70cf50e2-79a3-11f0-9be8-6e18e986f14e"},
		Delete: false,  // 保留弹性eip, 为true时，下方参数可不传
		RegionCode: "CN_Qingyang",  // 区域code
        BandwidthType: "Bandwidth_Multi_ISP_BGP",  // 带宽类型
        BillScheme: "BandwIdth",  // 计费方案
        Qos: 10,
	}
```
**从指定带宽包移除并删除弹性eip**
```go
func RemoveBandwidthPackageIp() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	// 移除删除eip
	RemoveBandwidthPackageIpArgs := &bandwidthpackage.RemoveBandwidthPackageIpReq{
		EIPIdList: []string{"6870eeac-79ac-11f0-8503-6e18e986f14e"},
		Delete: true,
	}

	response, err := BandwidthPackageClient.RemoveBandwidthPackageIp(RemoveBandwidthPackageIpArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

```
**删除共享带宽包**
```go
func DeleteBandwidthPackage() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	DeletebandwidthpackageArgs := &bandwidthpackage.DeleteBandwidthPackageReq{
        BandwidthId: "868bb384-79a4-11f0-adfa-6e18e986f14e",
	}
	response, err := BandwidthPackageClient.DeleteBandwidthPackage(DeletebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
### NAT网关管理代码示例
**获取指定nat网关信息**
```go
func GetNatGateway() {
	ak, sk := "ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	GetNatGatewayArgs := map[string]string{
		"Keyword": "c3e95ed4-79a9-11f0-a8a0-7a973848a269", // nat网关id或者名称

	}
	response, err := natgatewayClient.GetNatGateway(GetNatGatewayArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
**查询nat网关**
```go
func ListNatGateways() {
	ak, sk := "ak", "sk"

	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	ListNatGatewaysArgs := map[string]string{
		"RegionCode": "CN_Qingyang",  // 区域code
        "Keyword": "c3e95ed4-79a9-11f0-a8a0-7a973848a269", // nat网关id或者名称
	}
	response, err := natgatewayClient.ListNatGateways(ListNatGatewaysArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
```
# 数据结构说明
## 创建vpc请求参数
| 名称              | 类型   | 是否必选 | 示例值                  | 描述                                |
| ----------------- | ------ | -------- | ----------------------- | ----------------------------------- |
| RegionCode        | string | 是       | CN_Hongkong             | VPC区域code               |
| VPCName           | string | 是       | name                    | VPC名称                             |
| VPCSegment        | string | 是       | 10.15.0.0/16            | VPC网段   推荐网段          |
| VPCSegmentType    | string | 是       | auto/manual             | 使用推荐网段或手动分配，              |
| BandwidthType     | string | 否       | Bandwidth_China_Telecom | VPC带宽类型，边缘节点必传 |
| SubnetList        | list   | 是       | []                      | 创建VPC必须创建一个子网             |
| AvailableZoneCode | string | 是       | CN_Hongkong_B           | VPC可用区code           |
| SubnetName        | string | 是       | 子网1                   | 子网名称                            |
| SubnetSegment     | string | 是       | 10.15.1.0/24            | 子网网段                            |

## 查询VPC信息响应数据
| 名称       | 类型     | 示例值                               | 描述        |
| :--------- | -------- | :----------------------------------- | :---------- |
| VPCList    | list     | []                                   | VPC信息列表 |
| VPCId      | string   | 9304b130-e25b-11ec-a12c-823092f7a5bd | VPC ID      |
| VPCNmae    | string   | test                                 | VPC名称     |
| VPCSegment | string   | 10.15.0.0/16                         | VPC网段     |
| Status     | string   | ok                                   | VPC状态     |
| RegionCode | string   | CN_Hongkong                          | VPC区域code |
| CreateTime | datetime | 2022-06-02 18:05:47                  | 创建时间    |

## 创建子网请求参数
| 名称              | 类型   | 是否必选 | 示例值                               | 描述                    |
| ----------------- | ------ | -------- | ------------------------------------ | ----------------------- |
| VPCId             | string | 是       | 0266c864-e573-11ec-a09a-cabfed3cc5e1 | VPC ID                  |
| SubnetList        | list   | 是       | []                                   | 创建VPC必须创建一个子网 |
| AvailableZoneCode | string | 是       | CN_Hongkong_B                        | VPC可用区code           |
| SubnetName        | string | 是       | 子网1                                | 子网名称                |
| SubnetSegment     | string | 是       | 10.15.1.0/24                         | 子网网段                |

## 查询子网信息响应数据
| 名称              | 类型   | 示例值                               | 描述       |
| :---------------- | ------ | :----------------------------------- | :--------- |
| SubnetList        | list   | []                                   | 子网列表   |
| SubnetId          | string | 026f4386-e573-11ec-a09a-cabfed3cc5e1 | 子网ID     |
| SubnetName        | string | 子网1                                | 子网名称   |
| SubnetSegment     | string | 026f4386-e573-11ec-a09a-cabfed3cc5e1 | 子网网段   |
| UsedIPNum         | int    | 026f4386-e573-11ec-a09a-cabfed3cc5e1 | 已用IP数量 |
| VPCId             | string | 026f4386-e573-11ec-a09a-cabfed3cc5e1 | VPC ID     |
| VPCName           | string | 026f4386-e573-11ec-a09a-cabfed3cc5e1 | VPC名称    |
| AvailableZoneCode | string | 026f4386-e573-11ec-a09a-cabfed3cc5e1 | 可用区code |

## 创建弹性EIP请求参数
| 名称              | 类型   | 是否必选                      | 示例值                  | 描述                                                         |
| ----------------- | ------ | ----------------------------- | ----------------------- | --------------------- |
| RegionCode        | string | 是                            | CN_Hongkong             | VPC区域code           |
| AvailableZoneCode | string | 边缘节点：是 / 云平台节点: 否     |                         | VPC可用区code|
| BandwidthType     | string | 是                            | Bandwidth_China_Telecom | 带宽类型                  |
| BillScheme        | string | 是                            | BandwIdth               | EIP计费方案 |
| Qos               | int    | 是                            | 5                       | 带宽大小    |
| Size              | int    | 是                            | 1                       | 创建个数        |
| Description       | string | 否                            | test                    | EIP描述             |
| SubjectId      | int | 否       | 123                                                  | 测试金Id     |

## 获取EIP信息响应结果参数
| 名称              | 类型     | 示例值                               | 描述                                                    |
| :---------------- | -------- | :----------------------------------- | :---------------------------------- |
| EIPList           | list     | []                                   | EIP信息列表             |
| Id                | string   | 22597c56-e646-11ec-97e2-7687d6f44ced | EIP ID                   |
| IP                | string   | 118.186.70.138                       | EIP 地址                      |
| RegionCode        | string   | CN_Hongkong                          | VPC区域code                |
| AvailableZoneCode | string   | ""                                   | VPC可用区code |
| Status            | string   | ok                                   | EIP 状态       |
| Description       | string   | test                                 | EIP 描述信息                     |
| IsBind            | bool     | True                                 | EIP是否绑定了子网ip               |
| CreateTime        | datetime | 2022-06-02 18:05:47                  | 创建时间                  |
| BandwidthInfo     | dict     | {}                                   | EIP带宽信息               |
| Id                | string   | 2232dede-e646-11ec-97e2-7687d6f44ced | 带宽ID                 |
| Name              | string   | ""                                   | 共享带宽名称 |
| Qos               | int      | 10                                   | 带宽大小（Mbps）               |
| AvailableZoneCode | string   | ""                                   | VPC可用区code |
| BandwidthType     | string ｜Bandwidth_China_Telecom      |         |
| BillScheme        | string   | BandwIdth        |                                                         |
| Status            | string   | ok                                   | 带宽状态             |
| CreateTime        | string   | 2022-06-02 18:05:47                  | 带宽创建时间                     |
| BindResourceInfo  | dict     | {}                                   | EIP绑定的子网ip信息               |
| SubnetId          | string   | d9b88a1e-aa54-11ec-a512-06d5ff412043 | 子网ID（边缘节点，此处为虚拟出网网关ID）      |
| SubnetIP          | string   | 10.3.2.4                             | 子网IP（边缘节点，此处为虚拟出网网关IP）       |
## 创建共享带宽包请求参数
| 名称              | 类型   | 是否必选 | 示例值                               | 描述                     |
| ----------------- | ------ | -------- | ----------------------- | -------------------------- |
| RegionCode        | string | 是       | CN_Hongkong                          | VPC区域code               |
| AvailableZoneCode | string | 是       | CN_Hongkong_B                        | VPC可用区code     |
| Name              | string | 是       | 香港共享带宽                         | 共享带宽名称                 |
| BandwidthType     | string | 是       | Bandwidth_China_Telecom| 带宽类型              |
| BillScheme        | string | 是       | BandwIdth                            | 共享带宽计费方案 |
| Qos               | int    | 是       | 5                                    | 带宽大小              |
| NETID             | string | 否       | ce11eb1e-e6fa-11ec-8b50-bafaaf87d540 | 子网ID，边缘节点必传          |
| SubjectId         | int | 否       | 123                                                  | 测试金Id |
## 获取共享带宽响应结果参数
| 名称              | 类型   | 示例值                        | 描述                      |
| :---------------- | ------ | :--------------------------- | :------------------------ |
| BandwidthList     | list   | []                                   | 共享带宽信息列表                  |
| BandwidthInfo     | dict   | {}                                   | EIP带宽信息          |
| Id                | string | 2232dede-e646-11ec-97e2-7687d6f44ced | 带宽ID                 |
| Name              | string | ""                                   | 共享带宽名称       |
| Qos               | int    | 10                                   | 带宽大小（Mbps）          |
| RegionCode        | string | “”                                   | VPC区域code       |
| AvailableZoneCode | string | ""                                   | VPC可用区code |
| BandwidthType     | string | Bandwidth_China_Telecom            | 带宽类型           |
| BillScheme        | string | BandwIdth                            | 计费方案          |
| Status            | string | ok                                   | 带宽状态                |
| CreateTime        | string | 2022-06-02 18:05:47                  | 带宽创建时间                  |


# 常用常量说明
## 私有网络区域名称
| 区域名称 | RegionCode    | 区域类型   | 所在大区 |
| -------- | ------------- | ---------- | -------- |
|  庆阳    | CN_Qingyang | 边缘节点   | 中国大陆 |
|  达拉斯   | US_Dallas   | 边缘节点   | 中国大陆 |
## 私有网络可用区名称

| 可用区名称 | AvailableZoneCode | 站点类型   | 所在区域 |
| ---------- | ----------------- | ---------- | -------- |
| 庆阳A      | CN_Qingyang_A     | 边缘节点   | 庆阳    |
| 达拉斯N    | US_Dallas_N        | 边缘节点   | 达拉斯     |

## VPC推荐网段
|网段|
| -------- |
|10.0.0.0/16|
|172.16.0.0/16|
|192.168.0.0/16|

## vpc带宽类型
| 名称 | BandwidthType    | 
| -------- | ------------- |
| 移动    | Bandwidth_CMCC |
| 联通    | Bandwidth_China_Unicom |
| 电信    | Bandwidth_China_Telecom |
| BGP多线 | Bandwidth_Multi_ISP_BGP |

## EIP计费方案
| 名称 | BillScheme    | 
| -------- | ------------- |
|电信-固定带宽|BandwIdth|
|移动-固定带宽|BandwIdth|
|联通-固定带宽|BandwIdth|
|BGP(多线)-固定带宽|BandwIdth|
|电信-固定带宽EIP电信包月|BandwIdthMonth|
|移动-固定带宽EIP移动包月|BandwIdthMonth|
|联通-固定带宽EIP联通包月|BandwIdthMonth|
|BGP（多线）-固定带宽EIPBGP包月|BandwIdthMonth|
|电信-流量按需|Traffic|
|移动-流量按需|Traffic|
|联通-流量按需|Traffic|
|BGP(多线)-流量按需|Traffic|

## 共享带宽计费类型
| 名称 | BillScheme    | 
| -------- | ------------- |
|电信-固定带宽(共享)|BandwIdth_Shared|
|移动-固定带宽(共享)|BandwIdth_Shared|
|联通-固定带宽(共享)|BandwIdth_Shared|
|BGP(多线)-固定带宽(共享)|BandwIdth_Shared|
|电信-固定带宽(电信包月共享)|BandwIdthMonth_Shared|
|移动-固定带宽(移动包月共享)|BandwIdthMonth_Shared|
|联通-固定带宽(联通包月共享)|BandwIdthMonth_Shared|
|BGP(多线)-固定带宽(BGP包月共享)|BandwIdthMonth_Shared|
|电信-流量按需(共享)|Traffic_Shared|
|移动-流量按需(共享)|Traffic_Shared|
|联通-流量按需(共享)|Traffic_Shared|
|BGP(多线)-流量按需(共享)|Traffic_Shared|

# 完整示例

    私有网络VPC参考 examples/network/vpc.go文件查看完整的使用示例。
    子网参考 examples/network/subnet.go文件查看完整的使用示例。
    EIP参考 examples/network/eip.go文件查看完整的使用示例。
    共享带宽包参考 examples/network/bandwidthpackage.go文件查看完整的使用示例。
    NAT网关参考 examples/network/natgateway.go文件查看完整的使用示例。

如有问题，请参考项目文档或联系技术支持。