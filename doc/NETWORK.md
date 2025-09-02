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
        - [X] CreateEIP
        - [X] DeleteEIP
        - [X] UpdateEIP

        
- [X] 共享带宽包
    - [X] 共享带宽包管理
        - [X] CreateBandwidth
        - [X] DescribeBandwidth
        - [X] UpdateBandwidth
        - [X] BandwidthAddEIP
        - [X] BandwidthRemoveEIP
        - [X] DeleteBandwidth
        - [X] BandwidthBindResource
        - [X] BandwidthUnbindResource


- [X] NAT网关
    - [X] NAT网关管理
        - [X] GetNatGateway
        - [X] ListNatGateways

- [X] 高性能负载均衡管理
    - [X] SLB实例管理
        - [X] GetSlb
        - [X] ListSlb

   