package ecs

import (
	"encoding/json"
	"testing"

	"github.com/capitalonline/cds-cloudos-go-sdk/auth"
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
)

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
