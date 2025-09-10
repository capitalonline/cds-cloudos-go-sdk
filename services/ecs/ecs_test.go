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
			t.Error(err)
			return
		}
		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

	t.Run("DescribeInstanceList", func(t *testing.T) {
		result, cliErr := cli.DescribeInstanceList(&DescribeInstanceListReq{
			AvailableZoneCode: "",
			VpcId:             "",
			SearchInfo:        "",
		})

		if cliErr != nil {
			t.Error(err)
			return
		}
		b, _ := json.Marshal(result)
		t.Logf("%s", string(b))
	})

}
