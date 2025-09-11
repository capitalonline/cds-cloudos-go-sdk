
# 弹性云服务器ECS

## 概述
本文档主要介绍弹性云服务器ECS GO SDK的使用

## 使用指南

### 获取密钥
您需要拥有一个有效的AK(Access Key ID)和SK(Secret Access Key)用来进行签名认证。AK/SK是由系统分配给用户的，均为字符串，用于标识用户，为访问服务做签名验证。

可前往首云 https://c2.capitalonline.net/portal/webapps/user/securityKeySubuser 获取用户密钥

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

## SDK接口列表

- [x] ECS弹性GPU云服务器

    - [x] 公共
        - [x] [DescribeRegions](#DescribeRegions)
        - [x] [DescribeTaskEvent](#DescribeTaskEvent)

    - [X] GPU云服务器相关
        - [X] [DescribeInstanceList](#DescribeInstanceList)
        - [X] [DescribeInstance](#DescribeInstance)
        - [X] [OperateInstance](#OperateInstance)
        - [X] [ModifyInstanceName](#ModifyInstanceName)
        - [X] [DescribeEcsFamilyInfo](#DescribeEcsFamilyInfo)
        - [X] [ChangeInstanceConfigure](#ChangeInstanceConfigure)
    - [X] 云盘相关
        - [x] [ExtendDisk](#ExtendDisk)

### DescribeRegions

#### 参数说明

#### 代码示例

[查看示例代码](https://github.com/capitalonline/cds-cloudos-go-sdk/blob/dev-ecs/services/ecs/ecs_test.go#L23)

### DescribeTaskEvent

#### 参数说明

#### 代码示例

### DescribeInstanceList

#### 参数说明

#### 代码示例

### DescribeInstance

#### 参数说明

#### 代码示例

### OperateInstance

#### 参数说明

#### 代码示例

### ModifyInstanceName

#### 参数说明

#### 代码示例

### DescribeEcsFamilyInfo

#### 参数说明

#### 代码示例

### ChangeInstanceConfigure

#### 参数说明

#### 代码示例

### ExtendDisk

#### 参数说明

#### 代码示例





