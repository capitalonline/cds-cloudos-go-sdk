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

	t.Run("DescribeRegions", func(t *testing.T) {
		result, cliErr := cli.DescribeRegions()
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run("DescribeInstanceList", func(t *testing.T) {
		result, cliErr := cli.DescribeInstanceList(&DescribeInstanceListReq{})
		if cliErr != nil {
			t.Error(cliErr)
			return
		}
		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run("OperateInstance", func(t *testing.T) {
		req := &OperateInstanceReq{
			EcsIds:    []string{"ins-ti1ugucuxix8uo5j"},
			OpType:    "shutdown_ecs",
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

	t.Run("ModifyInstanceName", func(t *testing.T) {
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

	t.Run("DescribeInstance", func(t *testing.T) {
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

	t.Run("DescribeTaskEvent", func(t *testing.T) {
		result, cliErr := cli.DescribeTaskEvent(&DescribeTaskEventReq{
			EventId: "1c004e06-8e14-11f0-9162-9601b14a8b77",
		})

		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run("DescribeEcsFamilyInfo", func(t *testing.T) {
		result, cliErr := cli.DescribeEcsFamilyInfo(&DescribeEcsFamilyInfoReq{
			AvailableZoneCode: "CN_Qingyang_A",
			BillingMethod:     "0",
		})

		if cliErr != nil {
			t.Error(cliErr)
			return
		}

		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run("ExtendDisk", func(t *testing.T) {
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
