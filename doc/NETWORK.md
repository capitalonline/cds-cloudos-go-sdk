# 私有网络 SDK使用指南
本文档详细介绍如何使用 私有网络 相关的 SDK 方法，包括VPC管理、子网管理、eip管理、共享带宽包管理等操作。

## 功能概述

**私有网络管理** 

- **CreateVPC**: 创建VPC信息
- **GetVPC**: 获取指定VPC信息
- **ListVPCs**: 查询VPC信息
- **DeleteVPC**: 删除VPC

**子网管理**
- **CreateSubnet**: 创建子网
- **GetSubnet**: 获取指定子网信息
- **ListSubnets**: 查询子网信息
- **DeleteSubnet**: 删除子网

**弹性公网IP(EIP)管理**
- **CreateEIP**: 创建弹性EIP
- **GetEIP**: 获取指定弹性EIP信息
- **ListEIPs**: 查询弹性EIP
- **ReleaseEIP**: 释放弹性EIP
- **UpdateEIP**: 更改弹性EIP带宽信息

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
- **ListVpcSlb**:  查询VPC下的SLB列表信息
- **GetVpcSlbDetail**: 查询高性能负载均衡详情
- **CreateVpcSlb**: 创建VPC下的SLB
- **DeleteVpcSlb**: 删除VPC下的SLB
- **UpdateVpcSlb**: 更新VPC下的SLB
- **ListVpcSlbDetail**: 查询VPC下的SLB列表
- **GetVpcSlb**: 查询VPC下的SLB信息
- **GetVpcSlbListenCreateInfo**: 查询SLB可用的的VIP、ACL信息（用于监听创建）
- **CreateVpcSlbListen**: 创建VPC下的SLB监听
- **DeleteVpcSlbListen**: 删除VPC下的SLB监听
- **UpdateVpcSlbListen**: 更新VPC下的SLB监听
- **ListVpcSlbListen**: 查询VPC下的SLB监听列表
- **GetVpcSlbListen**: 查询VPC下的SLB监听
- **GetVpcSlbListenRsInfo**: 查询VPC下可绑定监听的主机信息（用于监听后端服务绑定）
- **CreateVpcSlbRsPort**:创建VPC下监听后端服务绑定关系
- **DeleteVpcSlbRsPort**: 删除VPC下监听后端服务绑定关系
- **UpdateVpcSlbRsPort**: 更新VPC下监听后端服务绑定关系
- **GetVpcSlbRsPort**: 查询VPC下监听后端服务绑定关系
- **BandwidthBindResource**: 共享带宽绑定VPC下SLB
- **BandwidthUnbindResource**: 共享带宽解绑VPC下SLB

## 快速开始
### 初始化VPC客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/vpc"
// 替换为您的实际访问密钥
ak, sk := "ak", "sk"
vpcClient, _ := vpc.NewClient(ak, sk)
```
### 初始化子网客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/subnet"
// 替换为您的实际访问密钥
ak, sk :=  "ak", "sk"
subnetClient, _ := subnet.NewClient(ak, sk)
```

### 初始化EIP客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/eip"
// 替换为您的实际访问密钥
ak, sk := "ak", "sk"
EipClient, _ := eip.NewClient(ak, sk)
```
### 初始化共享带宽包客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/bandwidthpackage"
// 替换为您的实际访问密钥
ak, sk := "ak", "sk"
BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
```
### 初始化NAT网关客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/natgateway"
// 替换为您的实际访问密钥
ak, sk := "ak", "sk"
natgatewayClient, _ := natgateway.NewClient(ak, sk)
```
### 初始化SLB客户端
```go
import "github.com/capitalonline/cds-cloudos-go-sdk/services/slb"
// 替换为您的实际访问密钥
ak, sk := "ak", "sk"
slbClient, _ := slb.NewClient(ak, sk)
```

## 代码示例
### VPC管理代码示例
**创建VPC**
```go
func CreateVPC() {
    // 替换为您的实际访问密钥
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
func GetVPC() {
    // 替换为您的实际访问密钥
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
func ListVPCs() {
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
func ListSubnets() {
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
**创建弹性EIP**
```go
func CreateEIP(){
    // 替换为您的实际访问密钥
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
**获取指定弹性EIP信息**
```go
func GetEIP() {
    // 替换为您的实际访问密钥
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
**查询弹性EIP**
```go
func ListEIPs() {
    // 替换为您的实际访问密钥
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
**释放弹性EIP**
```go
func ReleaseEIP(){
    // 替换为您的实际访问密钥
	ak, sk := "ak", "sk"


	EipClient, _ := eip.NewClient(ak, sk)
	ReleaseEipArgs := &eip.ReleaseEipReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",  // eip ID
	}

	response, err := EipClient.ReleaseEip(ReleaseEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}
```
**更改弹性EIP带宽信息**
```go
func UpdateEIP(){
    // 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	UpdateEipArgs := &eip.UpdateEIPReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
        Qos: 10, // 变更的带宽大小
		Description: "go create", // 描述信息
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
}
```
**从指定带宽包移除并删除弹性eip**
```go
func RemoveBandwidthPackageIp() {
    // 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	// 移除删除eip
	RemoveBandwidthPackageIpArgs := &bandwidthpackage.RemoveBandwidthPackageIpReq{
		EIPIdList: []string{"6870eeac-79ac-11f0-8503-6e18e986f14e"}, // EIP的ID
		Delete: true,  // 移除EIP并删除EIP
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
    // 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	DeletebandwidthpackageArgs := &bandwidthpackage.DeleteBandwidthPackageReq{
        BandwidthId: "868bb384-79a4-11f0-adfa-6e18e986f14e",  // 共享带宽的ID
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
    // 替换为您的实际访问密钥
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
    // 替换为您的实际访问密钥
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
### 高性能负载均衡管理代码示例
**查询VPC下的SLB列表信息**

```go
func ListVpcSlb() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	args := &slb.ListVpcSlbReq{
		VpcId: "",      // 需要查询的VPC ID
	}
	response, err := slbClient.ListVpcSlb(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	fmt.Println(response.Data)  // 获取VPC下SLB列表信息
}
```
**查询高性能负载均衡详情**

```go
func GetVpcSlbDetail() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	args := &slb.GetVpcSlbDetailReq{
        SlbId: "",  // 需要查询的SLB的ID
        SlbName: "",    // 需要查询的SLB的名称
    }
    response, err := slbClient.GetVpcSlbDetail(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	fmt.Println(response.Data)  // 获取SLB详情
}
```
> 注意: 对请求参数的内容解释如下
> - SlbId: 此参数允许为空字符串，当此参数为空时会使用SlbName进行实例详情查询，当此参数不为空时将高优先级使用此参数进行实例详情查询，当SlbId和SlbName同时传参时将使用SlbId进行实例详情查询
> - SlbName: 此参数允许为空字符串，仅当SlbId为空时会使用此参数进行实例详情查询，当SlbId不为空时不会使用此参数进行实例匹配查询
**创建VPC下的SLB**

```go
func CreateVpcSlb() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbReq{
		RegionCode:        "CN_Qingyang",				//区域code
		AvailableZoneCode: "CN_Qingyang_A",			//可用区code
		VpcId:             "",									//SLB归属的VPC ID
		ConfType:          "slb.v1.small",			//规格code，slb.v1.small-高阶型、slb.v1.medium-超强型、slb.v1.large-至强型
		NetType:           "wan",								//网络类型: wan-公网、wan_lan-公网和私网
		Name:              "name",		 //SLB名称
	}
	response, err := slbClient.CreateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**删除VPC下的SLB**

```go
func DeleteVpcSlb() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbReq{
		SlbId: "",	//需要删除的SLB实例ID
	}
	response, err := slbClient.DeleteVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**更新VPC下的SLB**

```go
func UpdateVpcSlb() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbReq{
		SlbId:    "",			//需要更新的SLB实例ID
		ConfType: "slb.v1.small",		//规格code，slb.v1.small-高阶型、slb.v1.medium-超强型、slb.v1.large-至强型
		Name:     "test",			//SLB名称
		NetType:  "wan",			//网络类型: wan-公网、wan_lan-公网和私网
	}
	response, err := slbClient.UpdateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**查询VPC下的SLB列表**

```go
func ListVpcSlbDetail() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.ListVpcSlbDetailReq{
		VpcId: "",		//查询指定VPC下的SLB实例列表
	}
	response, err := slbClient.ListVpcSlbDetail(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取SLB实例列表信息
}
```

**查询VPC下的SLB信息**

```go
func GetVpcSlb() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbReq{
		SlbId: "",				//SLB实例ID
	}
	response, err := slbClient.GetVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取SLB实例详细信息
}
```

**查询SLB可用的的VIP、ACL信息（用于监听创建）**

```go
func GetVpcSlbListenCreateInfo() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenCreateInfoReq{
		SlbId: "",		//SLB实例ID
	}
	response, err := slbClient.GetVpcSlbListenCreateInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取SLB可用的的VIP、ACL信息
}
```

**创建VPC下的SLB监听**

```go
func CreateVpcSlbListen() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbListenReq{
		SlbId:          "",			//SLB实例ID
		ListenName:     "test",		//监听名称
		Vip:            "",			//VIP地址
		VipId:          "",			//VIP实例ID
		VipType:        "wan_eip",		//VIP类型：wan_eip-公网VIP、wan_lan-内网VIP
		ListenProtocol: "TCP",		//监听协议：TCP、UDP
		ListenPort:     8080,		//监听端口
		AclId:          "",			//访问控制实例ID
		ListenTimeout:  10,			//会话保持时间，单位秒，取值范围：0-900
		Scheduler:      "rr",		//负载均衡策略：rr-轮询、wrr-加权轮询、conhash-一致性哈希
		HealthCheckInfo: slb.HealthCheckInfo{		//健康检查配置
			Protocol:         "TCP",		//健康检查协议：TCP、HTTP
			Virtualhost:      "",				//健康检查地址，仅当健康检查协议为HTTP时此参数生效
			Port:             0,				//健康检查端口，取值范围：0-65535，0表示使用后端服务端口
			Path:             "",				//健康检查路径，仅当健康检查协议为HTTP时此参数生效
			StatusCode:       200,			//健康检查状态码
			ConnectTimeout:   5,				//监测超时响应时间
			DelayLoop:        5,				//监测间隔时间
			Retry:            2,				//重连次数(不健康阈值)
			DelayBeforeRetry: 10,				//重连时间间隔(不健康重试间隔时间)
		},
	}
	response, err := slbClient.CreateVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

> 注意: 对请求参数的内容解释如下  
> - AclId、Vip、VipId、VipType四个参数内容可通过接口DescribeVpcSlbListenCreateInfo查询
> - HealthCheckInfo: 如果不传递此参数，则监听不会开启健康检查，建议开启

**删除VPC下的SLB监听**

```go
func DeleteVpcSlbListen() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbListenReq{
		ListenIds: []string{""},		//监听ID列表
	}
	response, err := slbClient.DeleteVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**更新VPC下的SLB监听**

```go
func UpdateVpcSlbListen() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbListenReq{
		ListenId:      "",				//监听ID
		ListenName:    "test",		//监听名称
		AclId:         "",			//访问控制实例ID
		ListenTimeout: 10,			//会话保持时间，单位秒，取值范围：0-900
		Scheduler:     "rr",		//负载均衡策略：rr-轮询、wrr-加权轮询、conhash-一致性哈希
		HealthCheckInfo: slb.HealthCheckInfo{		//健康检查配置
			Protocol:         "TCP",		//健康检查协议：TCP、HTTP
			Port:             0,				//健康检查端口，取值范围：0-65535，0表示使用后端服务端口
			ConnectTimeout:   5,				//监测超时响应时间
			DelayLoop:        5,				//监测间隔时间
			Retry:            2,				//重连次数(不健康阈值)
			DelayBeforeRetry: 10,				//重连时间间隔(不健康重试间隔时间)
		},
	}
	response, err := slbClient.UpdateVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**查询VPC下的SLB监听列表**

```go
func ListVpcSlbListen() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenListReq{
		SlbId: "",			//SLB实例ID
	}
	response, err := slbClient.ListVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取监听列表信息
}
```

**查询VPC下的SLB监听**

```go
func GetVpcSlbListen() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenReq{
		ListenId: "",				//监听实例ID
	}
	response, err := slbClient.GetVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取监听详细信息
}
```

**查询VPC下可绑定监听的主机信息（用于监听后端服务绑定）**

```go
func GetVpcSlbListenRsInfo() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenRsInfoReq{
		VmType: "kvm",		//主机类型：kvm
		VpcId:  "",				//VPC实例ID
	}
	response, err := slbClient.GetVpcSlbListenRsInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取可绑定监听的主机信息
}
```

**创建VPC下监听后端服务绑定关系**

```go
func CreateVpcSlbRsPort() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbRsPortReq{
		ListenId: "",				//监听实例ID
		RsList: []slb.RsPortItem{
			{
				VmId:        "",								//服务器ID
				VmName:      "",								//服务器名称
				VmPublicIp:  "",								//服务器公网IP，如果没有公网IP，可以传空字符串：""
				VmType:      "kvm",							//服务器类型
				VmPrivateIp: "10.0.0.6",				//服务器内网IP
				Port:        "8080",						//后端服务端口，取值范围：2-65535
				Weight:      "100",							//权重，取值范围：0-1000
			},
		},
	}
	response, err := slbClient.CreateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

> 注意: 对请求参数的内容解释如下  
> - VmId、VmName、VmPublicIp、VmType、VmPrivateIp五个参数内容可通过接口DescribeVpcSlbListenRsInfo查询

**删除VPC下监听后端服务绑定关系**

```go
func DeleteVpcSlbRsPort() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbRsPortReq{
		RsPortIds: []string{""},		//后端服务绑定实例ID列表
	}
	response, err := slbClient.DeleteVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**更新VPC下监听后端服务绑定关系**

```go
func UpdateVpcSlbRsPort() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbRsPortReq{
		RsPortList: []slb.RsPortItem{
			{
				RsPortId: "",						//后端服务绑定实例ID
				Port:        "8080",						//后端服务端口，取值范围：2-65535
				Weight:      "100",							//权重，取值范围：0-1000
			},
		},
	}
	response, err := slbClient.UpdateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**查询VPC下监听后端服务绑定关系**

```go
func GetVpcSlbRsPort() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbRsPortReq{
		ListenId: "",				//监听实例ID
		Keyword:  "",				//关键字匹配，可以传递空字符串：""
	}
	response, err := slbClient.GetVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取监听后端服务绑定实例列表
}
```

**共享带宽绑定VPC下SLB**

```go
func BandwidthBindResource() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthBindResourceReq{
		BandwidthId: "",		//共享带宽ID
		BindType:    "VPCSLB",	//绑定资源类型：VPCSLB-绑定VPC SLB
		ResourceId:  "",		//SLB实例ID
	}
	response, err := slbClient.BandwidthBindResource(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```

**共享带宽解绑VPC下SLB**

```go
func BandwidthUnbindResource() {
	ak, sk := "ak", "sk"      // 替换为您的实际密钥

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthUnbindResourceReq{
		BandwidthId: "",		//共享带宽ID
	}
	response, err := slbClient.BandwidthUnbindResource(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)   // 获取返回的完整信息
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))  // 获取任务ID
}
```



# 数据结构说明

## 创建VPC请求参数
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

## 创建VPC SLB请求参数


| 名称              | 类型   | 是否必选 | 示例值        | 描述                                                         |
| ----------------- | ------ | -------- | ------------- | ------------------------------------------------------------ |
| RegionCode        | string | 是       | CN_Hongkong   | 区域code, 见附件五                                           |
| AvailableZoneCode | string | 是       | CN_Hongkong_B | 可用区code, 见附件五                                         |
| VpcId             | string | 是       | VPC ID        | VPC ID                                                       |
| Name              | string | 是       | name          | SLB名称                                                      |
| NetType           | string | 是       | wan/wan_lan   | 网络类型: wan-公网、wan_lan-公网和私网                       |
| ConfType          | string | 是       | slb.v1.small  | 配置规格: slb.v1.small-高阶型、slb.v1.medium-超强型、slb.v1.large-至强型 |

## 更新VPC SLB请求参数

| 名称     | 类型   | 是否必选 | 示例值       | 描述                                                         |
| -------- | ------ | -------- | ------------ | ------------------------------------------------------------ |
| SlbId    | string | 是       |              | SLB实例ID                                                    |
| Name     | string | 是       | name         | SLB名称                                                      |
| NetType  | string | 是       | wan/wan_lan  | 网络类型: wan-公网、wan_lan-公网和私网                       |
| ConfType | string | 是       | slb.v1.small | 配置规格: slb.v1.small-高阶型、slb.v1.medium-超强型、slb.v1.large-至强型 |

## 查询VPC下的SLB列表结果参数

| 名称          | 类型   | 示例值  | 描述                                           |
| :------------ | ------ | :------ | :--------------------------------------------- |
| Code          | string | OK      | 错误码                                         |
| Message       | string | success | 信息                                           |
| Data          | dict   | {}      | 返回数据                                       |
| SlbList       | list   |         | SLB实例列表                                    |
| AzId          | string |         | 实例所属可用区ID                               |
| AzName        | string |         | 可用区名称                                     |
| BandwidthId   | string |         | 实例绑定的共享带宽ID                           |
| BandwidthName | string |         | 共享带宽名称                                   |
| BillingMethod | string | 0       | 实例计费类型：0-按需计费                       |
| ConfName      | string |         | 实例规格名称                                   |
| CreateTime    | string |         | 实例创建时间                                   |
| Id            | string |         | 实例ID                                         |
| Name          | string |         | 实例名称                                       |
| NetType       | string | wan     | 网络类型：wan-公网、wan_lan-同时支持公网和私网 |
| RegionId      | string |         | 实例所属大区ID                                 |
| RegionName    | string |         | 实例所属大区名称                               |
| Status        | string | ok      | 实例状态：ok-正常、error-异常                  |
| StatusZh      | string |         | 实例状态中文翻译                               |
| VipList       | list   |         | 实例VIP列表                                    |
| Address       | string |         | VIP地址                                        |
| Type          | string | public  | VIP类型：public-公网、private-内网             |
| VpcId         | string |         | 所属的VPC ID                                   |
| VpcName       | string |         | 所属的VPC名称                                  |

## 查询VPC下的SLB信息结果参数

| 名称                | 类型   | 示例值  | 描述                                           |
| :------------------ | ------ | :------ | :--------------------------------------------- |
| Code                | string | OK      | 错误码                                         |
| Message             | string | success | 信息                                           |
| Data                | dict   | {}      | 返回数据                                       |
| AzId                | string |         | 实例所属可用区ID                               |
| AzName              | string |         | 可用区名称                                     |
| BandwidthInfo       | dict   |         | 实例绑定的共享带宽信息                         |
| BandwidthInfo--Id   | string |         | 共享带宽ID                                     |
| BandwidthInfo--Name | string |         | 共享带宽名称                                   |
| BillingMethod       | string | 0       | 实例计费类型：0-按需计费                       |
| BillingSchemeId     | string |         | 实例计费ID                                     |
| BillingSchemeName   | string |         | 实例计费名称                                   |
| ConfName            | string |         | 实例规格名称                                   |
| CreateTime          | string |         | 实例创建时间                                   |
| Id                  | string |         | 实例ID                                         |
| Name                | string |         | 实例名称                                       |
| NetType             | string | wan     | 网络类型：wan-公网、wan_lan-同时支持公网和私网 |
| NetTypeZh           | string |         | 网络类型中文翻译                               |
| RegionId            | string |         | 实例所属大区ID                                 |
| RegionName          | string |         | 实例所属大区名称                               |
| Status              | string | ok      | 实例状态：ok-正常、error-异常                  |
| StatusZh            | string |         | 实例状态中文翻译                               |
| VipList             | list   |         | 实例VIP列表                                    |
| Address             | string |         | VIP地址                                        |
| Type                | string | public  | VIP类型：public-公网、private-内网             |
| VpcId               | string |         | 所属的VPC ID                                   |
| VpcName             | string |         | 所属的VPC名称                                  |
| VpcSegments         | string |         | 所属VPC的网段信息                              |

## 查询SLB可用的的VIP、ACL信息结果参数

| 名称       | 类型   | 示例值      | 描述                                      |
| :--------- | ------ | :---------- | :---------------------------------------- |
| Code       | string | OK          | 错误码                                    |
| Message    | string | success     | 信息                                      |
| data       | dict   | {}          | 返回数据                                  |
| SlbAclList | list   |             | SLB可用的ACL信息列表                      |
| AclId      | string |             | ACL实例ID                                 |
| AclName    | string |             | ACL名称                                   |
| SlbVipList | list   |             | SLB可用的VIP信息列表                      |
| Vip        | string | 38.123.96.8 | 监听虚拟IP                                |
| VipId      | string |             | VIP实例ID                                 |
| VipType    | string | wan_eip     | VIP类型：wan_eip-公网VIP、wan_lan-内网VIP |

## 创建VPC下的SLB监听请求参数

| 名称             | 类型   | 是否必选 | 示例值   | 描述                                                    |
| ---------------- | ------ | -------- | -------- | ------------------------------------------------------- |
| SlbId            | string | 是       |          | SLB实例ID                                               |
| ListenName       | string | 是       | test     | 监听名称                                                |
| Vip              | string | 是       | 10.0.0.1 | VIP地址                                                 |
| VipId            | string | 是       |          | VIP实例ID                                               |
| VipType          | string | 是       | wan_eip  | VIP类型：wan_eip-公网VIP、wan_lan-内网VIP               |
| ListenProtocol   | string | 是       | TCP      | 监听协议：TCP、UDP                                      |
| ListenPort       | int    | 是       | 8080     | 监听端口，取值范围：1-65535                             |
| AclId            | string | 否       |          | 需要绑定的ACL实例ID                                     |
| ListenTimeout    | int    | 否       | 10       | 会话保持时间，单位秒，取值范围：0-900                   |
| Scheduler        | string | 是       | rr       | 负载均衡策略：rr-轮询、wrr-加权轮询、conhash-一致性哈希 |
| HealthCheckInfo  | string | 否       |          | 健康检查配置                                            |
| Protocol         | string | 否       |          | 健康检查协议：TCP、HTTP                                 |
| Virtualhost      | string | 否       |          | 健康检查地址，仅当健康检查协议为HTTP时此参数生效        |
| Port             | int    | 否       |          | 健康检查端口，取值范围：0-65535，0表示使用后端服务端口  |
| Path             | string | 否       |          | 健康检查路径，仅当健康检查协议为HTTP时此参数生效        |
| StatusCode       | int    | 否       |          | 健康检查状态码                                          |
| ConnectTimeout   | int    | 否       |          | 监测超时响应时间                                        |
| DelayLoop        | int    | 否       |          | 监测间隔时间                                            |
| Retry            | int    | 否       |          | 重连次数(不健康阈值)                                    |
| DelayBeforeRetry | int    | 否       |          | 重连时间间隔(不健康重试间隔时间)                        |

## 更新VPC下的SLB监听请求参数

| 名称             | 类型   | 是否必选 | 示例值 | 描述                                                    |
| ---------------- | ------ | -------- | ------ | ------------------------------------------------------- |
| ListenId         | string | 是       |        | 监听实例ID                                              |
| ListenName       | string | 是       | test   | 监听名称                                                |
| AclId            | string | 否       |        | 需要绑定的ACL实例ID                                     |
| ListenTimeout    | int    | 否       | 10     | 会话保持时间，单位秒，取值范围：0-900                   |
| Scheduler        | string | 是       | rr     | 负载均衡策略：rr-轮询、wrr-加权轮询、conhash-一致性哈希 |
| HealthCheckInfo  | string | 否       |        | 健康检查配置                                            |
| Protocol         | string | 否       |        | 健康检查协议：TCP、HTTP                                 |
| Virtualhost      | string | 否       |        | 健康检查地址，仅当健康检查协议为HTTP时此参数生效        |
| Port             | int    | 否       |        | 健康检查端口，取值范围：1-65535                         |
| Path             | string | 否       |        | 健康检查路径，仅当健康检查协议为HTTP时此参数生效        |
| StatusCode       | int    | 否       |        | 健康检查状态码                                          |
| ConnectTimeout   | int    | 否       |        | 监测超时响应时间                                        |
| DelayLoop        | int    | 否       |        | 监测间隔时间                                            |
| Retry            | int    | 否       |        | 重连次数(不健康阈值)                                    |
| DelayBeforeRetry | int    | 否       |        | 重连时间间隔(不健康重试间隔时间)                        |

## 查询VPC下的SLB监听列表结果参数

| 名称           | 类型   | 示例值  | 描述                     |
| :------------- | ------ | :------ | :----------------------- |
| Code           | string | OK      | 错误码                   |
| Message        | string | success | 信息                     |
| TaskId         | string |         | 任务ID                   |
| Data           | dict   | {}      | 返回数据                 |
| ListenList     | list   |         | SLB中监听列表            |
| CheckInfo      | dict   |         | 后端服务端口健康检查信息 |
| Error          | int    |         | 健康检查失败数量         |
| Ok             | int    |         | 健康检查成功数量         |
| ListenId       | string |         | 监听ID                   |
| ListenName     | string |         | 监听名称                 |
| ListenPort     | string |         | 监听端口                 |
| ListenProtocol | string | TCP     | 监听协议                 |
| ListenVip      | string |         | 监听VIP                  |
| Scheduler      | string | rr      | 监听转发策略             |
| SchedulerCn    | string | 轮询    | 转发策略中文翻译         |
| Status         | string | ok      | 监听状态                 |
| StatusCn       | string | 正常    | 监听状态中文翻译         |

## 查询VPC下的SLB监听结果参数

| 名称             | 类型   | 示例值   | 描述                                                    |
| ---------------- | ------ | -------- | ------------------------------------------------------- |
| Code             | string |          | 错误码                                                  |
| Message          | string | success  | 信息                                                    |
| Data             | dict   | {}       | 返回数据                                                |
| ListenId         | string |          | 监听实例ID                                              |
| ListenName       | string | test     | 监听名称                                                |
| Vip              | string | 10.0.0.1 | VIP地址                                                 |
| VipId            | string |          | VIP实例ID                                               |
| VipType          | string | wan_eip  | VIP类型：wan_eip-公网VIP、wan_lan-内网VIP               |
| ListenProtocol   | string | TCP      | 监听协议：TCP、UDP                                      |
| ListenPort       | int    | 8080     | 监听端口，取值范围：1-65535                             |
| AclId            | string |          | 需要绑定的ACL实例ID                                     |
| ListenTimeout    | int    | 10       | 会话保持时间，单位秒，取值范围：0-900                   |
| Scheduler        | string | rr       | 负载均衡策略：rr-轮询、wrr-加权轮询、conhash-一致性哈希 |
| HealthCheckInfo  | string |          | 健康检查配置                                            |
| Protocol         | string |          | 健康检查协议：TCP、HTTP                                 |
| Virtualhost      | string |          | 健康检查地址，仅当健康检查协议为HTTP时此参数生效        |
| Port             | int    |          | 健康检查端口，取值范围：1-65535                         |
| Path             | string |          | 健康检查路径，仅当健康检查协议为HTTP时此参数生效        |
| StatusCode       | int    |          | 健康检查状态码                                          |
| ConnectTimeout   | int    |          | 监测超时响应时间                                        |
| DelayLoop        | int    |          | 监测间隔时间                                            |
| Retry            | int    |          | 重连次数(不健康阈值)                                    |
| DelayBeforeRetry | int    |          | 重连时间间隔(不健康重试间隔时间)                        |

## 查询VPC下可绑定监听的主机信息结果参数

| 名称          | 类型   | 示例值  | 描述         |
| :------------ | ------ | :------ | :----------- |
| Code          | string | OK      | 错误码       |
| Message       | string | success | 信息         |
| Data          | list   |         | 返回数据     |
| vm_id         | string |         | 服务器ID     |
| vm_name       | string |         | 服务器名称   |
| vm_private_ip | string |         | 服务器内网IP |
| vm_public_ip  | string |         | 服务器公网IP |
| vm_type       | string |         | 服务器类型   |

## 创建VPC下监听后端服务绑定关系请求参数

| 名称        | 类型   | 是否必选 | 示例值 | 描述                                             |
| ----------- | ------ | -------- | ------ | ------------------------------------------------ |
| ListenId    | string | 是       |        | 监听实例ID                                       |
| RsList      | string | 是       |        | 绑定的后端服务器列表                             |
| VmId        | string | 是       |        | 服务器ID                                         |
| VmName      | string | 是       |        | 服务器名称                                       |
| VmPublicIp  | string | 是       |        | 服务器公网IP，如果没有公网IP，可以传空字符串："" |
| VmType      | string | 是       |        | 服务器类型                                       |
| VmPrivateIp | string | 是       |        | 服务器内网IP                                     |
| Port        | string | 是       |        | 后端服务端口，取值范围：2-65535                  |
| Weight      | string | 是       |        | 权重，取值范围：0-1000                           |

## 更新VPC下监听后端服务绑定关系请求参数

| 名称       | 类型   | 是否必选 | 示例值 | 描述                            |
| ---------- | ------ | -------- | ------ | ------------------------------- |
| RsPortList | list   | 是       |        | 需要修改的后端服务列表          |
| RsPortId   | string | 是       |        | 后端服务实例ID                  |
| Port       | string | 是       |        | 后端服务端口，取值范围：2-65535 |
| Weight     | string | 是       |        | 权重，取值范围：0-1000          |

## 查询VPC下监听后端服务绑定关系结果参数

| 名称           | 类型   | 示例值  | 描述                 |
| :------------- | ------ | :------ | :------------------- |
| Code           | string | OK      | 错误码               |
| Message        | string | success | 信息                 |
| Data           | dict   | {}      | 返回数据             |
| RsPortList     |        |         | 后端服务列表         |
| CheckStatus    | string |         | 健康检查状态         |
| CheckStatusStr | string |         | 健康检查状态中文翻译 |
| LanIp          | string |         | 内网IP               |
| Port           | string |         | 服务端口             |
| ResourceId     | string |         | 服务器ID             |
| ResourceName   | string |         | 服务器名称           |
| RsPortId       | string |         | 后端服务实例ID       |
| Status         | string |         | 后端服务状态         |
| StatusZh       | string |         | 后端服务状态中文翻译 |
| VmType         | string |         | 服务器类型           |
| WanIp          | string |         | 服务器公网IP         |
| Weight         | string |         | 权重                 |
| Total          | int    |         | 后端服务数量         |

## 错误处理

所有SDK方法都会返回标准的错误信息：

```go
vpcClient, err := vpc.NewClient(ak, sk)
if err != nil {
    log.Printf("API call failed: %v", err)
    return
}

if result.Code != "Success" {
    log.Printf("API returned error: Code=%s, Message=%s", result.Code, result.Message)
    return
}

// 处理成功的结果
fmt.Printf("VPC created successfully: %+v\\n", result.Data)
```

## 注意事项

1. **VPC私有网络名称规范**: 名称最多输入255个中文，英文，'_'，'-'及数字
2. **子网名称规范**: 名称最多输入255个中文，英文，'_'，'-'及数字

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
    SLB参考 examples/network/slb.go文件查看完整的使用示例。

如有问题，请参考项目文档或联系技术支持。
