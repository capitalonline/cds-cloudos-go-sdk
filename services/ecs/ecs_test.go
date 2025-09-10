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

func TestDescribeInstanceList(t *testing.T) {
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

	result, err := cli.DescribeInstanceList(&DescribeInstanceListReq{
		AvailableZoneCode: "",
		VpcId:             "",
		SearchInfo:        "",
	})

	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(result)
	t.Logf("%s", string(b))
}
