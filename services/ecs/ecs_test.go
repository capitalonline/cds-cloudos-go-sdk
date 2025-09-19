package ecs

import (
	"encoding/json"
	"testing"

	"github.com/capitalonline/cds-cloudos-go-sdk/auth"
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
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

func TestTryEcsApi(t *testing.T) {
	credentials, authErr := auth.NewCdsCredentialsByEnv()

	if authErr != nil {
		t.Skip()
	}

	cdsCli, err := cds.NewCdsClientWithAkSk(credentials.AccessKeyId, credentials.SecretAccessKey)
	if err != nil {
		t.Skip()
	}

	cdsCli.Config.Retry = cds.NewNoRetryPolicy()

	cli := &client{
		CdsClient: cdsCli,
	}

	t.Run(ActionDescribeImages, func(t *testing.T) {
		result, cliErr := cli.DescribeImages(&DescribeImagesReq{
			AvailableZoneCode: "CN_DEMO", // 替换为实际的可用区代码
			// ImageIds:          []string{"test0000-0000-0000-0000-000000000000", "test0000-0000-0000-0000-000000000001"}, // 替换为实际的镜像id
		})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("%s", string(data))
	})

	t.Run(ActionCreateInstance, func(t *testing.T) {
		i := 1
		k := 0
		result, cliErr := cli.CreateInstance(&CreateInstanceReq{
			AvailableZoneCode: "CN_DEMO", // 替换为实际的可用区代码
			EcsFamilyName:     "优化型M6",
			Cpu:               2,
			Ram:               2,
			Gpu:               0,
			Number:            1,
			BillingMethod:     MonthlyBillingMethod,
			Password:          "？",
			ImageId:           "test0000-0000-0000-0000-000000000000", // 替换为实际的镜像id
			SystemDisk: &CreateInstanceDiskData{
				DiskFeature: SsdDiskFeature,
				Size:        40,
			},
			VpcInfo: &CreateInstanceVpcInfo{VpcId: "test0000-0000-0000-0000-000000000000"}, // 替换为实际的Vpc id
			SubnetInfo: &CreateInstanceSubnetInfo{
				SubnetId: "test0000-0000-0000-0000-000000000000", // 替换为实际的子网id
			},
			PubnetInfo: []*CreateInstancePubnetInfo{
				{
					SubnetId: "test0000-0000-0000-0000-000000000000",           // 替换为实际的子网id
					EipIds:   []string{"test0000-0000-0000-0000-000000000000"}, // 替换为实际的弹性公网id

					//BandwidthType:     Bandwidth,
					//Qos:               5,
					//BandwidthConfName: "电信",
				},
			},
			Name:          "yh",
			StartNumber:   0,
			Duration:      24,
			IsToMonth:     &i,
			IsAutoRenewal: &i,
			UtcTime:       0,
			DataDisk: []*CreateInstanceDiskData{
				{
					DiskFeature:         SsdDiskFeature,
					Size:                24,
					ReleaseWithInstance: &k,
				},
			},
			DnsList: &[2]string{
				"114.114.114.114",
				"8.8.8.8",
			},
			// SecurityGroups: []*CreateInstanceSecurityGroupData{
			// 	{
			// 		SecurityGroupId: "sg-0000000000000000",  // 替换为实际的安全组id
			// 	},
			// },
		})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("%s", string(data))
	})

	t.Run(ActionDeleteInstance, func(t *testing.T) {
		result, cliErr := cli.DeleteInstance(&DeleteInstanceReq{
			EcsIds:    []string{"ins-0000000000000000"}, // 替换为实际的实例id
			DeleteEip: 0,
		})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("%s", string(data))
	})

	t.Run(ActionModifyInstancePassword, func(t *testing.T) {
		result, cliErr := cli.ModifyInstancePassword(&ModifyInstancePasswordReq{
			EcsIds:   []string{"ins-0000000000000000"}, // 替换为实际的实例id
			Password: "？",
		})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("%s", string(data))
	})

	t.Run(ActionDescribeInstanceStatus, func(t *testing.T) {
		result, cliErr := cli.DescribeInstanceStatus(&DescribeInstanceStatusReq{
			EcsIds: []string{"ins-0000000000000000"}, // 替换为实际的实例id
		})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		data, _ := json.MarshalIndent(result, "", "  ")
		t.Logf("%s", string(data))
	})

}
