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
		默认环境变量：
		ak := os.Getenv("CDS_SECRET_ID")
		sk := os.Getenv("CDS_SECRET_KEY")
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
	fmt.Println(result.Data[0].RegionList[0].AzList[0].AvailableZoneCode)

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
		AvailableZoneCode: "CN_DEMO", // 替换为实际的可用区代码
	})
	if err != nil {
		log.Fatal("Failed to describe ECS family info:", err)
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(data))
	fmt.Println()
}

func changeInstanceConfigure(client ecs.Client) {
	fmt.Println("=== ChangeInstanceConfigure Example ===")
	// 调用 ChangeInstanceConfigure 方法修改实例配置
	req := &ecs.ChangeInstanceConfigureReq{
		EcsIds:            []string{"ins-test000000000000"}, // 替换为实际的实例ID
		AvailableZoneCode: "CN_DEMO",                        // 替换为实际的可用区代码

		// EcsFamilyName, 可通过DescribeInstance/DescribeEcsFamilyInfo方法获取实例规格信息
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
