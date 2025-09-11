package ecs

import (
	"encoding/json"
	"os"
	"testing"
)

const (
	AK = "CDS_SECRET_ID"
	SK = "CDS_SECRET_KEY"
)

func TestEcsSdk(t *testing.T) {
	ak := os.Getenv(AK)
	sk := os.Getenv(SK)

	if ak == "" || sk == "" {
		t.Skip()
	}

	cli, err := NewClient(ak, sk)
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
			EcsIds:    []string{"ins-ti1ugucuxix8uo5j"},
			OpType:    StartUpInstance,
			DeleteEip: 0,
		}

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
			EcsId: "ins-ti1ugucuxix8uo5j",
			Name:  "yh-test",
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
			EcsId: "ins-ti1ugucuxix8uo5j",
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
			EventId: "f7c1dd60-bfff-4bde-a3ef-0d284ace14bc",
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
			AvailableZoneCode: "CN_Qingyang_A",
			BillingMethod:     OnDemandBillingMethod,
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
			DiskId:       "disk-wrssau7uain81ogj",
			ExtendedSize: 72,
		}

		result, cliErr := cli.ExtendDisk(req)
		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

}
