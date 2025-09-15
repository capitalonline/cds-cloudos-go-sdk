package ecs

import (
	"encoding/json"
	"testing"

	"github.com/capitalonline/cds-cloudos-go-sdk/auth"
)

func TestEcsSdk(t *testing.T) {
	credentials, authErr := auth.NewCdsCredentialsByEnv()

	if authErr != nil {
		t.Skip()
	}

	cli, err := NewClient(credentials.AccessKeyId, credentials.SecretAccessKey)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run(ActionDescribeRegions, func(t *testing.T) {
		result, cliErr := cli.DescribeRegions()
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionDescribeInstanceList, func(t *testing.T) {
		result, cliErr := cli.DescribeInstanceList(&DescribeInstanceListReq{})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionOperateInstance, func(t *testing.T) {
		req := &OperateInstanceReq{
			EcsIds:    []string{"ins-test000000000000"}, // 替换为实际的实例ID
			OpType:    StartUpInstance,                  // 开机
			DeleteEip: 0,
		}
		/*
			// 其他可选操作
			req := &OperateInstanceReq{
				EcsIds: []string{"ins-test000000000000"},
				OpType: ecs.RestartInstance, // 重启

			}
			req := &OperateInstanceReq{
				EcsIds: []string{"ins-test000000000000"},
				OpType: ecs.ShutdownInstance, // 关机

			}
			req := &OperateInstanceReq{
				EcsIds: []string{"ins-test000000000000"},
				OpType: ecs.HardShutdownInstance, // 强制关机

			}
		*/
		result, cliErr := cli.OperateInstance(req)
		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionModifyInstanceName, func(t *testing.T) {
		req := &ModifyInstanceNameReq{
			EcsId: "ins-test000000000000", // 替换为实际的实例ID
			Name:  "new-instance-name",
		}

		result, cliErr := cli.ModifyInstanceName(req)
		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionDescribeInstance, func(t *testing.T) {
		result, cliErr := cli.DescribeInstance(&DescribeInstanceReq{
			EcsId: "ins-test000000000000", // 替换为实际的实例ID
		})

		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionDescribeTaskEvent, func(t *testing.T) {
		result, cliErr := cli.DescribeTaskEvent(&DescribeTaskEventReq{
			EventId: "event-uuid-demo", // 替换为实际的事件ID
		})

		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionDescribeEcsFamilyInfo, func(t *testing.T) {
		result, cliErr := cli.DescribeEcsFamilyInfo(&DescribeEcsFamilyInfoReq{
			AvailableZoneCode: "CN_DEMO",             // 替换为实际的可用区代码
			BillingMethod:     OnDemandBillingMethod, //按需计费
		})

		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionExtendDisk, func(t *testing.T) {
		req := &ExtendDiskReq{
			DiskId:       "disk-test000000000000", // 替换为实际的磁盘ID，可通过 DescribeInstance 方法获取实例磁盘信息
			ExtendedSize: 72,                      // 扩展后的大小(GB)，必须是8的倍数，且不能小于当前磁盘大小
		}

		result, cliErr := cli.ExtendDisk(req)
		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run(ActionChangeInstanceConfigure, func(t *testing.T) {
		req := &ChangeInstanceConfigureReq{
			EcsIds:            []string{"ins-test000000000000"}, // 替换为实际的实例ID
			AvailableZoneCode: "CN_DEMO",                        // 替换为实际的可用区代码
			EcsFamilyName:     "优化型M2",                          // 替换为实际的实例规格族名称，可通过DescribeInstance/DescribeEcsFamilyInfo方法获取实例规格组信息
			Cpu:               2,
			Ram:               4,
			Gpu:               0,
		}

		result, cliErr := cli.ChangeInstanceConfigure(req)
		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

}
