package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/capitalonline/cds-cloudos-go-sdk/auth"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/ecs"
)

func makeClient() ecs.Client {
	/*
		默认环境变量名称：CDS_SECRET_ID/CDS_SECRET_KEY
		os.Getenv("CDS_SECRET_ID") // ak
		os.Getenv("CDS_SECRET_KEY") // sk
	*/
	// 从环境变量中获取认证信息
	credentials, err := auth.NewCdsCredentialsByEnv()
	if err != nil {
		log.Fatalln("Failed to get credentials:", err)
	}

	// 创建 ECS 客户端
	client, err := ecs.NewClient(credentials.AccessKeyId, credentials.SecretAccessKey)
	if err != nil {
		log.Fatalln("Failed to create ECS client:", err)
	}

	/*
		// 也可以通过硬编码变量(ak/sk)创建客户端，或自定义环境变量
		ak := "ak-xxxx" // os.Getenv("your_ak_env")
		sk := "sk-xxxx" // os.Getenv("your_sk_env")
		client, err := ecs.NewClient(ak, sk)
		if err != nil {
			log.Fatal("Failed to create ECS client:", err)
		}
	*/

	return client
}

func describeRegions(client ecs.Client) {
	fmt.Println("=== DescribeRegions Example ===")
	// 调用 DescribeRegions 方法获取区域信息
	result, err := client.DescribeRegions()
	if err != nil {
		log.Fatal("Failed to describe regions:", err)
	}

	// 提示：SDK中传参需要的AvailableZoneCode, 可以通过如下接口获取。
	// AvailableZoneCode的值轻易不会更变，可以一次性获取后，将值以固定变量保存在代码中使用。
	// fmt.Println(result.Data[0].RegionList[0].AzList[0].AvailableZoneCode)

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func describeInstanceList(client ecs.Client) {
	fmt.Println("=== DescribeInstanceList Example ===")
	// 调用 DescribeInstanceList 方法获取实例列表
	result, err := client.DescribeInstanceList(&ecs.DescribeInstanceListReq{
		AvailableZoneCode: "",  // 可用区code过滤筛选，可选
		VpcId:             "",  // vpc_id过滤筛选，可选
		SearchInfo:        "",  // 实例id/实例名称/私网ip模糊过滤筛选，可选
		Ids:               nil, // 实例id批量精确匹配过滤筛选，可选
	})
	if err != nil {
		log.Fatal("Failed to describe instance list:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func createInstanceOnDemand(client ecs.Client) {
	fmt.Println("=== CreateInstance OnDemand Example ===")
	// 创建按需计费实例
	i := 1
	req := &ecs.CreateInstanceReq{
		Name:              "test",                    // 云服务器名,不传自动赋予（自动命名规则：ecs-创建日期），可选
		Password:          "1234#qwer",               // 登录密码,密码为8-30个字符，且同时包含三项（大写字母、小写字母、数字、()`~!@#$%^&*-+=_|{}[]:;,.?/中的特殊字符）
		AvailableZoneCode: "CN_DEMO",                 // 可用区代码
		EcsFamilyName:     "优化型M6",                   // 规格族名称
		UtcTime:           0,                         // 是否utc时间，1:是  0:否 默认为0（UTC+8，上海时间）
		Cpu:               2,                         // Cpu
		Ram:               4,                         // 内存
		Gpu:               0,                         // 显卡数量，可选，默认为0
		Number:            1,                         // 购买数量，可选，默认为1（默认批量最大值为100台）
		BillingMethod:     ecs.OnDemandBillingMethod, //按需计费
		ImageId:           "Ubuntu 20.04 64位",        // 替换为实际的镜像id或者镜像名称
		SystemDisk: &ecs.CreateInstanceDiskData{
			DiskFeature: ecs.SsdDiskFeature, // 盘类型，云盘
			Size:        24,                 // 盘大小
		}, // 系统盘信息
		DataDisk: []*ecs.CreateInstanceDiskData{
			{
				DiskFeature:         ecs.SsdDiskFeature, // 盘类型，云盘
				Size:                40,                 // 盘大小
				SnapshotId:          "",
				ReleaseWithInstance: &i, // 是否随实例删除:1:随实例删除,0:不随实例删除.不传默认随实例删除
			},
		}, // 数据盘信息，仅支持云盘，可选
		VpcInfo: &ecs.CreateInstanceVpcInfo{
			VpcId: "vpc-id", // 私有网络id
		}, // vpc信息
		SubnetInfo: &ecs.CreateInstanceSubnetInfo{
			SubnetId: "subnet-id", // 子网id
			// IpAddress: []string{""}, // 指定私网IP列表,列表中的IP个数与创建云主机个数一致
		}, // 私有网络信息
		SecurityGroups: nil, // 安全组列表，安全组优先级按顺序由高到低，可选
		StartNumber:    0,   // 云服务器名称编号起始数字，可选，不需要服务器编号可不传
		DnsList: &[2]string{
			"114.114.114.114",
			"8.8.8.8",
		}, // dns 解析，可选 需要两个元素  [主dns，从dns]，不选采用默认通用DNS
		PubnetInfo: []*ecs.CreateInstancePubnetInfo{
			{
				SubnetId:          "subnet-id",   // 子网id
				BandwidthConfName: "电信",          // 带宽线路名称，例如：电信、联通
				BandwidthType:     ecs.Bandwidth, // 带宽类型，若需新分配公网IP必填,表示绑定公网IP的带宽类型.绑定已有公网IP不填（若实例计费方式为包年包月选择固定带宽时需传"固定带宽包月"）
				Qos:               5,             // 公网带宽值,单位为M;若带宽类型选择”固定带宽”需填写},
			},
		}, // 支持新分配公网IP和绑定已有的公网IP，可选
		TestAccount: nil, // 测试金账户，可不填
	}

	result, err := client.CreateInstance(req)
	if err != nil {
		log.Fatal("Failed to describe instance list:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func createInstanceMonthly(client ecs.Client) {
	fmt.Println("=== CreateInstance Monthly Example ===")
	// 创建包月计费实例
	i := 1
	result, err := client.CreateInstance(&ecs.CreateInstanceReq{
		Name:              "test",                   // 云服务器名,不传自动赋予（自动命名规则：ecs-创建日期），可选
		Password:          "1234#qwer",              // 登录密码,密码为8-30个字符，且同时包含三项（大写字母、小写字母、数字、()`~!@#$%^&*-+=_|{}[]:;,.?/中的特殊字符）
		AvailableZoneCode: "CN_DEMO",                // 可用区代码
		EcsFamilyName:     "优化型M6",                  // 规格族名称
		UtcTime:           0,                        // 是否utc时间，1:是  0:否 默认为0（UTC+8，上海时间）
		Cpu:               1,                        // Cpu
		Ram:               2,                        // 内存
		Gpu:               0,                        // 显卡数量，可选，默认为0
		Number:            1,                        // 购买数量，可选，默认为1（默认批量最大值为100台）
		BillingMethod:     ecs.MonthlyBillingMethod, // 包月计费
		ImageId:           "Ubuntu 20.04 64位",       // 替换为实际的镜像id或者镜像名称
		SystemDisk: &ecs.CreateInstanceDiskData{
			DiskFeature: ecs.LocalDiskFeature, // 盘类型，本地盘
			Size:        20,                   // 盘大小
		}, // 系统盘信息
		DataDisk: []*ecs.CreateInstanceDiskData{
			{
				DiskFeature: ecs.SsdDiskFeature, // 盘类型，云盘
				Size:        20,                 // 盘大小
			},
		}, // 数据盘信息，仅支持云盘，可选
		VpcInfo: &ecs.CreateInstanceVpcInfo{
			VpcId: "vpc-id", // 私有网络id
		}, // vpc信息
		SubnetInfo: &ecs.CreateInstanceSubnetInfo{
			SubnetId:  "subnet-id",          // 子网id
			IpAddress: []string{"10.0.0.1"}, // 指定私网IP列表,列表中的IP个数与创建云主机个数一致
		}, // 私有网络信息
		SecurityGroups: []*ecs.CreateInstanceSecurityGroupData{
			{
				SecurityGroupId: "sg-id", // SecurityGroupId
			},
		}, // 安全组列表，安全组优先级按顺序由高到低，可选
		StartNumber: 0,   // 云服务器名称编号起始数字，可选，不需要服务器编号可不传
		Duration:    3,   // 只在包月算价时有意义，以月份为单位，一年值为12，大于一年要输入12的整数倍，最大值36(3年)，可选
		IsToMonth:   &i,  // 包月是否到月底，可选 1:是  0:否 默认为1
		DnsList:     nil, // dns 解析，可选 需要两个元素  [主dns，从dns]，不选采用默认通用DNS
		PubnetInfo: []*ecs.CreateInstancePubnetInfo{
			{
				SubnetId: "subnet-id",    // 子网id;
				EipIds:   []string{"id"}, // 绑定的eip的id列表;若需新分配公网IP,不填,绑定已有公网IP需填,数量需要和云服务器数量一致

			},
		}, // 支持新分配公网IP和绑定已有的公网IP，可选
		IsAutoRenewal: &i, // 是否自动续约，包月时需传。1:是  0:否 默认为1
	})
	if err != nil {
		log.Fatal("Failed to describe instance list:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func deleteInstance(client ecs.Client) {
	fmt.Println("=== DeleteInstance Example ===")
	// 调用 DeleteInstance 方法修改实例名称
	req := &ecs.DeleteInstanceReq{
		EcsIds:    []string{"ins-test000000000000"}, // 替换为实际的实例ID
		DeleteEip: 0,
	}
	result, err := client.DeleteInstance(req)
	if err != nil {
		log.Fatal("Failed to modify instance name:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))

}

func modifyInstancePassword(client ecs.Client) {
	fmt.Println("=== ModifyInstancePassword Example ===")
	// 调用 ModifyInstancePassword 方法修改实例名称
	req := &ecs.ModifyInstancePasswordReq{
		EcsIds:   []string{"ins-test000000000000"}, // 替换为实际的实例ID
		Password: "",                               // 登录密码,密码为8-30个字符，且同时包含三项（大写字母、小写字母、数字、()`~!@#$%^&*-+=_|{}[]:;,.?/中的特殊字符）
	}
	result, err := client.ModifyInstancePassword(req)
	if err != nil {
		log.Fatal("Failed to modify instance name:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))

}

func operateInstance(client ecs.Client) {
	fmt.Println("=== OperateInstance Example ===")
	// 调用 OperateInstance 方法启动实例
	req := &ecs.OperateInstanceReq{
		EcsIds: []string{"ins-test000000000000"}, // 替换为实际的实例ID
		OpType: ecs.StartUpInstance,              // 开机
	}

	/*
		// 其他可选操作
		req = &ecs.OperateInstanceReq{
			EcsIds: []string{"ins-test000000000000"},
			OpType: ecs.RestartInstance, // 重启

		}
		req = &ecs.OperateInstanceReq{
			EcsIds: []string{"ins-test000000000000"},
			OpType: ecs.ShutdownInstance, // 关机

		}
		req = &ecs.OperateInstanceReq{
			EcsIds: []string{"ins-test000000000000"},
			OpType: ecs.HardShutdownInstance, // 强制关机

		}
	*/
	result, err := client.OperateInstance(req)
	if err != nil {
		log.Fatal("Failed to operate instance:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func modifyInstanceName(client ecs.Client) {
	fmt.Println("=== ModifyInstanceName Example ===")
	// 调用 ModifyInstanceName 方法修改实例名称
	req := &ecs.ModifyInstanceNameReq{
		EcsId: "ins-test000000000000", // 替换为实际的实例ID
		Name:  "new-instance-name",
	}
	result, err := client.ModifyInstanceName(req)
	if err != nil {
		log.Fatal("Failed to modify instance name:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))

}

func describeInstance(client ecs.Client) {
	fmt.Println("=== DescribeInstance Example ===")
	// 调用 DescribeInstance 方法获取实例详细信息
	result, err := client.DescribeInstance(&ecs.DescribeInstanceReq{
		EcsId: "ins-test000000000000", // 替换为实际的实例ID
	})
	if err != nil {
		log.Fatal("Failed to describe instance:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func describeInstanceStatus(client ecs.Client) {
	fmt.Println("=== DescribeInstanceStatus Example ===")
	// 调用 DescribeInstanceStatus 方法修改实例名称
	req := &ecs.DescribeInstanceStatusReq{
		EcsIds: []string{"ins-test000000000000"}, // 替换为实际的实例ID
	}
	result, err := client.DescribeInstanceStatus(req)
	if err != nil {
		log.Fatal("Failed to modify instance name:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func describeTaskEvent(client ecs.Client) {
	fmt.Println("=== DescribeTaskEvent Example ===")
	// 调用 DescribeTaskEvent 方法获取任务事件信息
	result, err := client.DescribeTaskEvent(&ecs.DescribeTaskEventReq{
		EventId: "event-uuid-demo", // 替换为实际的事件ID
	})
	if err != nil {
		log.Fatal("Failed to describe task event:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))

}

func describeEcsFamilyInfo(client ecs.Client) {
	fmt.Println("=== DescribeEcsFamilyInfo Example ===")
	// 调用 DescribeEcsFamilyInfo 方法获取ECS规格族信息
	result, err := client.DescribeEcsFamilyInfo(&ecs.DescribeEcsFamilyInfoReq{
		AvailableZoneCode: "CN_DEMO",                 // 替换为实际的可用区代码
		BillingMethod:     ecs.OnDemandBillingMethod, //按需计费
	})
	if err != nil {
		log.Fatal("Failed to describe ECS family info:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func describeImages(client ecs.Client) {
	fmt.Println("=== DescribeImages Example ===")
	// 调用 DescribeImages 方法获取任务事件信息
	result, err := client.DescribeImages(&ecs.DescribeImagesReq{
		AvailableZoneCode: "CN_DEMO",                                        // 替换为实际的可用区代码
		ImageIds:          []string{"test0000-0000-0000-0000-000000000000"}, // 替换为实际的ImageId
	})
	if err != nil {
		log.Fatal("Failed to describe task event:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))

}

func changeInstanceConfigure(client ecs.Client) {
	fmt.Println("=== ChangeInstanceConfigure Example ===")
	// 调用 ChangeInstanceConfigure 方法修改实例配置
	req := &ecs.ChangeInstanceConfigureReq{
		EcsIds:            []string{"ins-test000000000000"}, // 替换为实际的实例ID
		AvailableZoneCode: "CN_DEMO",                        // 替换为实际的可用区代码

		// EcsFamilyName, 可通过DescribeInstance/DescribeEcsFamilyInfo方法获取实例规格组信息
		EcsFamilyName: "极速渲染型re3", // 替换为实际的实例规格族名称

		// 更变的配置
		Cpu: 16,
		Ram: 64,
		Gpu: 8,
	}
	result, err := client.ChangeInstanceConfigure(req)
	if err != nil {
		log.Fatal("Failed to change instance configuration:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func extendDisk(client ecs.Client) {
	fmt.Println("=== ExtendDisk Example ===")
	// 调用 ExtendDisk 方法扩展磁盘
	req := &ecs.ExtendDiskReq{
		DiskId:       "disk-test000000000000", // 替换为实际的磁盘ID，可通过 DescribeInstance 方法获取实例磁盘信息
		ExtendedSize: 72,                      // 扩展后的大小(GB)，必须是8的倍数，且不能小于当前磁盘大小
	}
	result, err := client.ExtendDisk(req)
	if err != nil {
		log.Fatal("Failed to extend disk:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
}

func main() {
	// Example
	client := makeClient()
	describeRegions(client)
}
