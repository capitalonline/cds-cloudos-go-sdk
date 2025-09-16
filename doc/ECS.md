
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

| 参数名            | 类型   | 描述       |
| :---------------- | ------ | ---------- |
| AzId              | string | 可用区ID   |
| AzName            | string | 可用区名称 |
| AvailableZoneCode | string | 可用区代码 |

### DescribeTaskEvent

*获取任务事件信息*

#### 参数说明

##### DescribeTaskEventReq

| 参数名  | 类型   | 必填 | 描述   |
| ------- | ------ | ---- | ------ |
| EventId | string | 是   | 事件id |

##### DescribeTaskEventResult

| 参数名 | 类型                                   | 描述         |
| ------ | -------------------------------------- | ------------ |
| Data   | []*[EventResultData](#EventResultData) | 事件结果数据 |

##### EventResultData

| 参数名             | 类型                             | 描述             |
| ------------------ | -------------------------------- | ---------------- |
| EventId            | string                           | 事件ID           |
| EventStatus        | string                           | 事件状态         |
| EventStatusDisplay | string                           | 事件状态显示名称 |
| EventType          | string                           | 事件类型         |
| EventTypeDisplay   | string                           | 事件类型显示名称 |
| CreateTime         | string                           | 创建时间         |
| TaskList           | []*[TaskListInfo](#TaskListInfo) | 任务列表         |

##### TaskListInfo

| 参数名          | 类型 | 描述             |
| --------------- | ---- | ---------------- |
| TaskId          |      | 任务ID           |
| Status          |      | 任务状态         |
| StatusDisplay   |      | 任务状态显示名称 |
| ResourceId      |      | 资源ID           |
| CreateTime      |      | 创建时间         |
| UpdateTime      |      | 更新时间         |
| EndTime         |      | 结束时间         |
| ResourceType    |      | 资源类型         |
| ResourceDisplay |      | 资源显示名称     |
| TaskType        |      | 任务类型         |
| TaskTypeDisplay |      | 任务类型显示名称 |

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

| 参数                 | 类型                                       | 说明                                                       |
| -------------------- | ------------------------------------------ | ---------------------------------------------------------- |
| EcsId                | string                                     | 实例ID                                                     |
| EcsName              | string                                     | 实例名称                                                   |
| VpcId                | string                                     | VPC ID                                                     |
| VpcName              | string                                     | VPC名称                                                    |
| Status               | string                                     | 云服务器状态码，参考 [云服务器状态说明](#云服务器状态说明) |
| StatusDisplay        | string                                     | 状态显示名称                                               |
| AzId                 | string                                     | 可用区ID                                                   |
| AzName               | string                                     | 可用区名称                                                 |
| RegionId             | string                                     | 区域ID                                                     |
| RegionName           | string                                     | 区域名称                                                   |
| PrivateNet           | string                                     | 私网ip                                                     |
| PubNet               | string                                     | 默认出网网关ip                                             |
| VirtualNet           | list                                       | 其他线路出网网关ip列表                                     |
| EipInfo              | map[string]*[EipSimpleInfo](EipSimpleInfo) | 公网ip信息                                                 |
| CreateTime           | string                                     | 创建时间                                                   |
| EcsFamilyName        | string                                     | 规格族名称                                                 |
| SpecFamilyId         | string                                     | 规格族ID                                                   |
| CpuSize              | int                                        | CPU核心数                                                  |
| RamSize              | int                                        | 内存大小(GB)                                               |
| IsGpu                | bool                                       | 是否为GPU实例                                              |
| GpuSize              | int                                        | 显卡数量                                                   |
| CardName             | string                                     | 显卡型号                                                   |
| BillingMethod        | string                                     | [计费方式](#billingMethod)                                 |
| BillingMethodName    | string                                     | 计费方式名称                                               |
| EndBillTime          | string                                     | 到期时间                                                   |
| IsAutoRenewal        | string                                     | 到期是否自动续约                                           |
| OsType               | string                                     | 操作系统类型                                               |
| OsVersion            | string                                     | 操作系统版本                                               |
| OsBit                | int                                        | 操作系统位数                                               |
| ImageName            | string                                     | 镜像名称                                                   |
| SystemDiskType       | string                                     | 系统盘类型                                                 |
| SystemDiskSize       | int                                        | 系统盘大小                                                 |
| NoChargeForShutdown  | int                                        | 是否关机不计费： 1是， 0否                                 |
| BindingPubnetIp      | bool                                       | 是否绑定公网 IP                                            |
| CustomerId           | string                                     | 客户ID                                                     |
| EcsGoodsId           | string                                     | 云主机商品/套餐 ID                                         |
| GpuCardInfo          | *[GpuCardData](#GpuCardData)               | GPU 卡信息                                                 |
| GpuDriver            | string                                     | GPU 驱动版本信息                                           |
| IsHaveSnapshot       | bool                                       | 是否有快照                                                 |
| ProductSource        | string                                     | 产品来源标识                                               |
| ProductSourceDisplay | string                                     | 产品来源显示名                                             |
| SecurityGroup        | []*[SecurityGroupInfo](#SecurityGroupInfo) | 安全组                                                     |
| SubnetId             | string                                     | 子网 ID                                                    |
| SupportGpuDriver     | string                                     | 支持的GPU驱动                                              |
| SystemDiskFeature    | string                                     | 系统盘特性                                                 |
| Tag                  | list                                       | 标签                                                       |

##### EipSimpleInfo

| 参数     | 类型   | 说明                               |
| -------- | ------ | ---------------------------------- |
| ConfName | string | 网络带宽运营商，如电信、移动、联通 |
| EipIp    | string | EIP地址                            |

##### SecurityGroupInfo

| 参数                | 类型   | 说明                     |
| ------------------- | ------ | ------------------------ |
| GroupInterconnected | bool   | 安全组内是否可以实例互通 |
| Priority            | int    | 优先级                   |
| SecurityGroupId     | string | 安全组ID                 |
| SecurityGroupName   | string | 安全组名称               |
| SecurityGroupType   | string | 安全组类型               |

##### GpuCardData

| 参数            | 类型   | 说明               |
| --------------- | ------ | ------------------ |
| GicBusinessName | string | GPU 型号的名称     |
| GpuDriver       | string | GPU 驱动版本       |
| GpuTypeId       | string | GPU 类型 ID        |
| RealName        | list   | GPU 型号的真实名称 |

### DescribeInstance

*获取云服务器配置详情* 

#### 参数说明

##### DescribeInstanceReq

| 参数名 | 类型   | 必填 | 描述       |
| ------ | ------ | ---- | ---------- |
| EcsId  | string | 是   | 云服务器id |

##### DescribeInstanceResult

| 参数 | 类型                           | 说明         |
| ---- | ------------------------------ | ------------ |
| Data | *[InstanceData](#InstanceData) | 实例详细数据 |

##### InstanceData

| 参数                | 类型                                   | 说明             |
| ------------------- | -------------------------------------- | ---------------- |
| EcsId               | string                                 | 实例ID           |
| EcsName             | string                                 | 实例名称         |
| RegionId            | string                                 | 区域ID           |
| RegionName          | string                                 | 区域名称         |
| AzId                | string                                 | 可用区ID         |
| AzName              | string                                 | 可用区名称       |
| Status              | string                                 | 实例状态         |
| StatusDisplay       | string                                 | 实例状态显示名称 |
| CreateTime          | string                                 | 创建时间         |
| Duration            | int                                    | 运行时长         |
| EndBillTime         | string                                 | 计费结束时间     |
| IsAutoRenewal       | string                                 | 是否自动续费     |
| TimeZone            | string                                 | 时区             |
| IsGpu               | bool                                   | 是否为 GPU 实例  |
| IsToMonth           | int                                    | 是否按月计费     |
| RealCardName        | string                                 | GPU 实际型号     |
| SecurityGroup       | list                                   | 安全组           |
| StockRelease        | bool                                   | 库存释放标志     |
| Supplier            | string                                 | 实例供应商       |
| SystemDiskFeature   | string                                 | 系统盘类型       |
| Tag                 | list                                   | 标签             |
| NoChargeForShutdown | int                                    | 关机是否收费     |
| EcsRule             | *[InstanceRuleInfo](#InstanceRuleInfo) | 实例规格信息     |
| OsInfo              | *[OsInfo](#OsInfo)                     | 操作系统信息     |
| Disk                | *[DiskInfo](#DiskInfo)                 | 磁盘信息         |
| Pipe                | *[PipeInfo](#PipeInfo)                 | 网络信息         |
| BillingInfo         | *[BillingInfo](#BillingInfo)           | 计费信息         |

##### InstanceRuleInfo

| 参数       | 类型   | 说明       |
| ---------- | ------ | ---------- |
| Name       | string | 规格名称   |
| CpuNum     | int    | CPU核心数  |
| CpuUnit    | string | CPU单位    |
| Ram        | int    | 内存大小   |
| Gpu        | int    | GPU数量    |
| RamUnit    | string | 内存单位   |
| GpuUnit    | string | GPU单位    |
| CategoryId | string | 规格类别ID |
| ConfName   | string | 配置名称   |
| EcsGoodsId | string | 商品ID     |

##### OsInfo

| 参数      | 类型   | 说明         |
| --------- | ------ | ------------ |
| ImageId   | string | 镜像ID       |
| ImageName | string | 镜像名称     |
| OsType    | string | 操作系统类型 |
| Bit       | int    | 操作系统位数 |
| Version   | string | 操作系统版本 |
| Unit      | string | 单位         |

##### DiskInfo

| 参数           | 类型                                       | 说明               |
| -------------- | ------------------------------------------ | ------------------ |
| SystemDiskConf | *[SystemDiskConfInfo](#SystemDiskConfInfo) | 系统盘配置信息     |
| DataDiskConf   | []*[DataDiskConfInfo](#DataDiskConfInfo)   | 数据盘配置信息列表 |

##### SystemDiskConfInfo

| 参数                | 类型   | 说明                                   |
| ------------------- | ------ | -------------------------------------- |
| ReleaseWithInstance | int    | 是否随实例释放(0:否,1:是)              |
| DiskType            | string | 磁盘类型(system: 系统盘, data: 数据盘) |
| Name                | string | 磁盘名称(用户自定义或默认名称)         |
| Size                | int    | 磁盘大小(GB)                           |
| DiskIops            | int    | 磁盘IOPS                               |
| BandMbps            | int    | 磁盘带宽(Mbps)                         |
| Unit                | string | 单位                                   |
| DiskId              | string | 磁盘ID                                 |
| DiskFeature         | string | 磁盘特性                               |
| DiskName            | string | 磁盘显示名称(例如 SSD云盘)             |
| EbsGoodsId          | string | 磁盘商品 ID                            |
| EcsGoodsId          | string | 云主机商品 ID                          |
| IsFollowDelete      | int    | 是否随实例删除(0: 否, 1: 是)           |

##### DataDiskConfInfo

| 参数                | 类型   | 说明                                   |
| ------------------- | ------ | -------------------------------------- |
| ReleaseWithInstance | int    | 是否随实例释放(0:否,1:是)              |
| DiskType            | string | 磁盘类型(system: 系统盘, data: 数据盘) |
| Name                | string | 磁盘名称(用户自定义或默认名称)         |
| Size                | int    | 磁盘大小(GB)                           |
| DiskIops            | int    | 磁盘IOPS                               |
| BandMbps            | int    | 磁盘带宽(Mbps)                         |
| Unit                | string | 单位                                   |
| Id                  | string | 磁盘ID                                 |
| DiskFeature         | string | 磁盘特性                               |
| DiskName            | string | 磁盘显示名称(例如 SSD云盘)             |
| EbsGoodsId          | string | 磁盘商品 ID                            |
| EcsGoodsId          | string | 云主机商品 ID                          |
| IsFollowDelete      | int    | 是否随实例删除(0: 否, 1: 是)           |

##### PipeInfo

| 参数       | 类型                           | 说明         |
| ---------- | ------------------------------ | ------------ |
| VpcName    | string                         | VPC名称      |
| VpcId      | string                         | VPC ID       |
| SubnetId   | string                         | 子网ID       |
| SubnetName | string                         | 子网名称     |
| CreateTime | string                         | 创建时间     |
| PrivateNet | string                         | 私网IP       |
| PubNet     | string                         | 公网IP       |
| VirtualNet | list                           | 虚拟网络信息 |
| EipInfo    | map[string]*[EipInfo](EipInfo) | EIP信息      |

##### EipInfo

| 参数          | 类型   | 说明         |
| ------------- | ------ | ------------ |
| ConfName      | string | 配置名称     |
| EipIp         | string | EIP地址      |
| CurrentUseQos | int    | 当前使用带宽 |
| BandwidthName | string | 带宽名称     |

##### BillingInfo

| 参数                | 类型   | 说明                       |
| ------------------- | ------ | -------------------------- |
| BillingMethod       | string | [计费方式](#billingMethod) |
| BillingMethodName   | string | 计费方式名称               |
| BillingStatus       | string | 计费状态                   |
| BillCycleId         | string | 计费周期ID                 |
| BillingMethodStatus | string | 计费方式状态               |

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

##### ModifyInstanceNameReq

| 参数名 | 类型   | 必填 | 描述       |
| ------ | ------ | ---- | ---------- |
| EcsId  | string | 是   | 实例ID     |
| Name   | string | 是   | 新实例名称 |

##### ModifyInstanceNameResult

| 参数名 | 类型                                               | 必填             |
| ------ | -------------------------------------------------- | ---------------- |
| Data   | *[ModifyInstanceNameData](#ModifyInstanceNameData) | 修改实例名称数据 |

##### ModifyInstanceNameData

| 参数名 | 类型   | 必填     |
| ------ | ------ | -------- |
| EcsId  | string | 实例ID   |
| Name   | string | 实例名称 |

### DescribeEcsFamilyInfo

*获取弹性云服务器ECS计算类型规格信息*

#### 参数说明

##### DescribeEcsFamilyInfoReq

| 参数名            | 类型   | 必填 | 描述                                                    |
| ----------------- | ------ | ---- | ------------------------------------------------------- |
| AvailableZoneCode | string | 是   | 可用区代码筛选，参考[AzInfo](#AzInfo).AvailableZoneCode |
| BillingMethod     | string | 是   | [计费方式](#billingMethod)                              |

##### DescribeEcsFamilyInfoResult

| 参数名 | 类型                       | 描述       |
| ------ | -------------------------- | ---------- |
| Data   | *[FamilyData](#FamilyData) | 规格族数据 |

##### FamilyData

| 参数名        | 类型                         | 描述              |
| ------------- | ---------------------------- | ----------------- |
| EcsFamilyInfo | []*[FamilyInfo](#FamilyInfo) | ECS规格族信息列表 |

##### FamilyInfo

| 参数名        | 类型                              | 描述          |
| ------------- | --------------------------------- | ------------- |
| EcsFamilyName | string                            | ECS规格族名称 |
| SpecList      | []*[SpecListInfo](#[SpecListInfo) | 规格列表      |

##### SpecListInfo

| 参数名           | 类型   | 描述          |
| ---------------- | ------ | ------------- |
| GpuShowType      | string | GPU显示类型   |
| GpuTypeId        | string | GPU类型ID     |
| Cpu              | int    | CPU核心数     |
| Ram              | int    | 内存大小(GB)  |
| Gpu              | int    | GPU数量       |
| SupportGpuDriver | string | 支持的GPU驱动 |

### ChangeInstanceConfigure

*更变云服务器规格*

***限制：***

- ***已关机状态下才可以操作***
- ***不支持跨类型修改实例规格***
- ***GPU型不支持跨驱动***
- ***按需计费支持升降配置***
-  ***包年包月只支持升配置***
-  ***云盘实例支持跨规格族，本地盘实例不支持跨规格族***
- ***批量操作要具有一致性***

#### 参数说明

##### ChangeInstanceConfigureReq

| 参数名            | 类型   | 必填 | 描述           |
| ----------------- | ------ | ---- | -------------- |
| EcsIds            | list   | 是   | 实例ID列表     |
| AvailableZoneCode | string | 是   | 可用区代码     |
| EcsFamilyName     | string | 是   | 实例规格族名称 |
| Cpu               | int    | 是   | CPU核心数      |
| Ram               | int    | 是   | 内存大小(GB)   |
| Gpu               | int    | 否   | GPU数量        |

##### ChangeInstanceConfigureResult

| 参数名 | 类型                         | 描述       |
| ------ | ---------------------------- | ---------- |
| Data   | *[EventIdData](#EventIdData) | 事件ID数据 |

### ExtendDisk

*云盘扩容，系统盘或者数据盘*

#### 参数说明

##### ExtendDiskReq

| 参数名       | 类型   | 必填 | 描述             |
| ------------ | ------ | ---- | ---------------- |
| DiskId       | string | 是   | 磁盘ID           |
| ExtendedSize | int    | 是   | 扩展后的大小(GB) |

##### ExtendDiskResult

| 参数名 | 类型                         | 描述       |
| ------ | ---------------------------- | ---------- |
| Data   | *[EventIdData](#EventIdData) | 事件ID数据 |



