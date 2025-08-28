package main

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func ListNodes() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	params := eks.ListNodesReq{
		ClusterId: "5ff38515-e942-437d-8e16-60f12c382e08",
	}
	response, err := eksClient.ListNodes(&params)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteNodes() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	params := eks.DeleteNodesReq{
		ClusterId: "5ff38515-e942-437d-8e16-60f12c382e08",
		NodeIds:   []string{"eks-4vdp8eu5nq3zppfg"},
	}
	response, err := eksClient.DeleteNodes(&params)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}
