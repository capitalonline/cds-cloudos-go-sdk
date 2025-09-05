# 网络服务

# 概述
本文档主要介绍网络 GO SDK的使用

# 使用指南

## 获取密钥
要使用网络相关服务，您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问网络相关服务做签名验证。


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
- [X] 私有网络（VPC）
    - [X] 私有网络管理
        - [X] CreateVpc
        - [X] GetVpc
        - [X] ListVpcs
        - [X] DeleteVpc
    - [X] 子网管理
        - [X] CreateSubnet
        - [X] GetSubnet
        - [X] ListSubnets
        - [X] DeleteSubnet

- [X] 弹性公网IP(EIP)
    - [X] 弹性公网IP(EIp)管理
        - [X] CreateEip
        - [X] GetEip
        - [X] ListEips
        - [X] ReleaseEip
        - [X] UpdateEIP

        
- [X] 共享带宽包
    - [X] 共享带宽包管理
        - [X] CreateBandwidthPackage
        - [X] GetBandwidthPackage
        - [X] ListBandwidthPackages
        - [X] UpdateBandwidthPackage
        - [X] AddBandwidthPackageIp
        - [X] RemoveBandwidthPackageIp
        - [X] DeleteBandwidthPackage
        - [X] BandwidthPackageBindResource
        - [X] BandwidthPackageUnbindResource


- [X] NAT网关
    - [X] NAT网关管理
        - [X] CreateNatGateway
        - [X] GetNatGateway
        - [X] ListNatGateways
        - [X] UpdateNatGateway
        - [X] DeleteNatGateway
    - [X] NAT规则管理
        - [X] CreateNATRule
        - [X] GetNATRule
        - [X] ListNATRules
        - [X] UpdateNATRule
        - [X] DeleteNATRule

## 高性能负载均衡管理

### 查询VPC下的SLB列表
查询指定VPC下的SLB列表信息
```go
// import github.com/capitalonline/cds-cloudos-go-sdk/services/slb
ak, sk := "", ""

slbClient, _ := slb.NewClient(ak, sk)
args := &slb.ListVpcSlbReq{
    // 需要查询的VPC ID
    VpcId: "",
}
response, err := slbClient.ListVpcSlb(args)
if err != nil {
    fmt.Println(err)
}
// 获取返回的完整信息
fmt.Printf(">>> response: %+v", response)
// 获取VPC下SLB列表信息
fmt.Println(response.Data)
```

### 查询指定SLB
```go
// import github.com/capitalonline/cds-cloudos-go-sdk/services/slb
ak, sk := "", ""

slbClient, _ := slb.NewClient(ak, sk)
args := &slb.GetVpcSlbDetailReq{
    // 需要查询的SLB的ID
    SlbId: "",
    // 需要查询的SLB的名称
    SlbName: "",
}
response, err := slbClient.GetVpcSlbDetail(args)
if err != nil {
    fmt.Println(err)
}
// 获取返回的完整信息
fmt.Printf(">>> response: %+v", response)
// 获取SLB详细信息
fmt.Println(response.Data)
```
> 注意: 对请求参数的内容解释如下
> - SlbId: 此参数允许为空字符串，当此参数为空时会使用SlbName进行实例详情查询，当此参数不为空时将高优先级使用此参数进行实例详情查询，当SlbId和SlbName同时传参时将使用SlbId进行实例详情查询
> - SlbName: 此参数允许为空字符串，仅当SlbId为空时会使用此参数进行实例详情查询，当SlbId不为空时不会使用此参数进行实例匹配查询