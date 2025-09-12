
# 弹性云服务器ECS

## 概述
本文档主要介绍弹性云服务器ECS GO SDK的使用

## 使用指南

### 获取密钥
您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问服务做签名验证。

可前往[**首云-密钥管理**](https://c2.capitalonline.net/portal/webapps/user/securityKeySubuser)获取用户密钥

### 云服务器状态说明

| code             | 说明       |
| ---------------- | ---------- |
| building         | 创建中     |
| running          | 运行中     |
| restarting       | 重启中     |
| shutting_down    | 关机中     |
| shutdown         | 已关机     |
| starting_up      | 开机中     |
| deleting         | 删除中     |
| deleted          | 已删除     |
| destroying       | 销毁中     |
| destroy          | 已销毁     |
| recovering       | 恢复中     |
| updating         | 更新中     |
| error            | 错误       |
| failed           | 创建失败   |
| recycling        | 回收中     |
| cancel_recycling | 取消回收中 |
| creating_image   | 定制镜像中 |

### 代码示例

[查看示例代码](../examples/ecs/ecs.go)

为了方便开发者进行单元测试，我们为客户端每个SDK方法提供了Mock实现，可以完全模拟真实客户端的行为，而无需实际调用远程 API：

[查看Mock单元测试示例代码](../examples/ecs/ecs_test.go)

### 公共参数

#### OpenApiCommonResp

| 参数名    | 类型   | 描述     |
| --------- | ------ | -------- |
| Code      | string | 响应码   |
| Message   | string | 响应消息 |
| RequestId | string | 请求ID   |

#### CommonOpenApiPage

| 参数名     | 类型   | 描述       |
| ---------- | ------ | ---------- |
| TotalCount | string | 总记录数   |
| PageIndex  | string | 当前页码   |
| PageSize   | string | 每页记录数 |

#### EventIdData

| 参数名  | 类型   | 描述   |
| ------- | ------ | ------ |
| EventId | string | 事件ID |

#### billingMethod

| 常量名                | 类型   | 值   | 描述     |
| --------------------- | ------ | ---- | -------- |
| OnDemandBillingMethod | string | "0"  | 按需计费 |
| MonthlyBillingMethod  | string | "1"  | 包月     |

## SDK接口列表

- [x] ECS弹性GPU云服务器

    - [x] 公共
        - [x] [DescribeRegions：获取可用区信息](#DescribeRegions)
        - [x] [DescribeTaskEvent：获取任务事件信息](#DescribeTaskEvent)

    - [X] GPU云服务器相关
        - [X] [DescribeInstanceList：获取云服务器列表](#DescribeInstanceList)
        - [X] [DescribeInstance：获取云服务器详情 ](#DescribeInstance)
        - [X] [OperateInstance：操作云服务器](#OperateInstance)
        - [X] [ModifyInstanceName：修改云服务器名称](#ModifyInstanceName)
        - [X] [DescribeEcsFamilyInfo：获取规格信息](#DescribeEcsFamilyInfo)
        - [X] [ChangeInstanceConfigure：更变云服务器规格](#ChangeInstanceConfigure)
    - [X] 云盘相关
        - [x] [ExtendDisk：云盘扩容](#ExtendDisk)

### DescribeRegions

*获取可用区信息*

#### 参数说明

##### DescribeRegionsResult

| 参数名 | 类型                           | 描述       |
| ------ | ------------------------------ | ---------- |
| Data   | []*[RegionGroup](#RegionGroup) | 区域组列表 |

##### RegionGroup

| 参数名          | 类型                         | 描述       |
| --------------- | ---------------------------- | ---------- |
| RegionGroupId   | string                       | 区域组ID   |
| RegionGroupName | string                       | 区域组名称 |
| RegionList      | []*[RegionInfo](#RegionInfo) | 区域列表   |

##### RegionInfo

| 参数名     | 类型                 | 描述       |
| ---------- | -------------------- | ---------- |
| RegionId   | string               | 区域组ID   |
| RegionName | string               | 区域组名称 |
| RegionCode | string               | 区域列表   |
| AzList     | []*[AzInfo](#AzInfo) | 可用区列表 |

##### AzInfo

| 参数名                | 类型       | 描述           |
| --------------------- | ---------- | -------------- |
| AzId                  | string     | 可用区ID       |
| AzName                | string     | 可用区名称     |
| **AvailableZoneCode** | **string** | **可用区代码** |

### DescribeTaskEvent

*获取任务事件信息*

#### 参数说明



### DescribeInstanceList

*获取云服务器列表*

#### 参数说明

##### DescribeInstanceListReq

| 参数名            | 类型     | 必填 | 描述                                                    |
| ----------------- | -------- | ---- | ------------------------------------------------------- |
| AvailableZoneCode | string   | 否   | 可用区代码筛选，参考[AzInfo](#AzInfo).AvailableZoneCode |
| VpcId             | string   | 否   | VPC ID筛选                                              |
| SearchInfo        | string   | 否   | 搜索信息（实例ID/实例名称/私网IP）                      |
| EcsIds            | []string | 否   | 批量实例ID筛选                                          |

##### DescribeInstanceListResult

| 参数名 | 类型                                   | 描述         |
| ------ | -------------------------------------- | ------------ |
| Data   | *[InstanceListData](#InstanceListData) | 实例列表数据 |

##### InstanceListData

| 参数名  | 类型                                         | 描述             |
| ------- | -------------------------------------------- | ---------------- |
| EcsList | []*[InstanceSimpleInfo](#InstanceSimpleInfo) | 实例简要信息列表 |

##### InstanceSimpleInfo

| 参数                | 类型                                       | 说明                                                       |
| ------------------- | ------------------------------------------ | ---------------------------------------------------------- |
| EcsId               | string                                     | 实例ID                                                     |
| EcsName             | string                                     | 实例名称                                                   |
| VpcId               | string                                     | VPC ID                                                     |
| VpcName             | string                                     | VPC名称                                                    |
| Status              | string                                     | 云服务器状态码，参考 [云服务器状态说明](#云服务器状态说明) |
| StatusDisplay       | string                                     | 状态显示名称                                               |
| AzId                | string                                     | 可用区ID                                                   |
| AzName              | string                                     | 可用区名称                                                 |
| RegionId            | string                                     | 区域ID                                                     |
| RegionName          | string                                     | 区域名称                                                   |
| PrivateNet          | string                                     | 私网ip                                                     |
| PubNet              | string                                     | 默认出网网关ip                                             |
| VirtualNet          | list                                       | 其他线路出网网关ip列表                                     |
| EipInfo             | map[string]*[EipSimpleInfo](EipSimpleInfo) | 公网ip信息                                                 |
| CreateTime          | string                                     | 创建时间                                                   |
| EcsFamilyName       | string                                     | 规格族名称                                                 |
| SpecFamilyId        | string                                     | 规格族ID                                                   |
| CpuSize             | int                                        | CPU核心数                                                  |
| RamSize             | int                                        | 内存大小(GB)                                               |
| IsGpu               | bool                                       | 是否为GPU实例                                              |
| GpuSize             | int                                        | 显卡数量                                                   |
| CardName            | string                                     | 显卡型号                                                   |
| BillingMethod       | string                                     | [计费方式](#billingMethod)                                 |
| BillingMethodName   | string                                     | 计费方式名称                                               |
| EndBillTime         | string                                     | 到期时间                                                   |
| IsAutoRenewal       | string                                     | 到期是否自动续约                                           |
| OsType              | string                                     | 操作系统类型                                               |
| OsVersion           | string                                     | 操作系统版本                                               |
| OsBit               | int                                        | 操作系统位数                                               |
| ImageName           | string                                     | 镜像名称                                                   |
| SystemDiskType      | string                                     | 系统盘类型                                                 |
| SystemDiskSize      | int                                        | 系统盘大小                                                 |
| NoChargeForShutdown | int                                        | 是否关机不计费： 1是， 0否                                 |

##### EipSimpleInfo

| 参数     | 类型   | 说明                               |
| -------- | ------ | ---------------------------------- |
| ConfName | string | 网络带宽运营商，如电信、移动、联通 |
| EipIp    | string | EIP地址                            |

### DescribeInstance

*获取云服务器配置详情* 

#### 参数说明



### OperateInstance

*批量操作云服务器：开机、关机、重启*

#### 参数说明

##### OperateInstanceReq

| 参数名    | 类型                           | 必填 | 描述                                                         |
| --------- | ------------------------------ | ---- | ------------------------------------------------------------ |
| EcsIds    | string                         | 是   | 实例ID列表                                                   |
| OpType    | [operate](#operate-云主机操作) | 是   | [云服务器操作](#operate-云服务器操作)                        |
| DeleteEip | int                            | 否   | 公网释放选项，仅在关机不计费情况生效(0:保留公网IP,1:释放公网IP) |

##### operate-云服务器操作

| 常量名               | 类型   | 值                | 描述                                                         |
| -------------------- | ------ | ----------------- | ------------------------------------------------------------ |
| ShutdownInstance     | string | shutdown_ecs      | 关机操作                                                     |
| StartUpInstance      | string | start_up_ecs      | 开机操作                                                     |
| RestartInstance      | string | restart_ecs       | 重启操作                                                     |
| HardShutdownInstance | string | hard_shutdown_ecs | 强制关机操作                                                 |
| FreeShutdownInstance | string | free_shutdown_ecs | 关机不计费，注：仅按需计费的云盘实例支持关机不计费，再开机公网IP可能会变化； 若批量操作关机：支持关机不收费的实例，关机后停止 CPU、内存、GPU和公网收费； 不支持关机不收费的实例，正常关机，继续收费； 不计费关机期间不支持除开机、删除、定制镜像外的操作。 注意：目前关机不计费白名单开放，若需要请联系商务。 |

##### OperateInstanceResult

| 参数名 | 类型                         | 描述       |
| ------ | ---------------------------- | ---------- |
| Data   | *[EventIdData](#EventIdData) | 事件ID数据 |

### ModifyInstanceName

*修改云服务器名称*

#### 参数说明



### DescribeEcsFamilyInfo

*获取弹性云服务器ECS计算类型规格信息*

#### 参数说明



### ChangeInstanceConfigure

*更变云服务器规格*

#### 参数说明



### ExtendDisk

*云盘扩容，系统盘或者数据盘*

#### 参数说明







