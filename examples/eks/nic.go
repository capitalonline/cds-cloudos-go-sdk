package main

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func AttachNetworkInterface() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)
	req := eks.AttachNetworkInterfaceReq{
		EcsId:     "",
		SubnetId:  "",
		NetcardId: "",
		VlanId:    "",
	}
	response, err := eksClient.AttachNetworkInterface(&req)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DetachNetworkInterface() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)
	req := eks.DetachNetworkInterfaceReq{}
	response, err := eksClient.DetachNetworkInterface(&req)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeNetworkInterface() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)
	req := eks.DescribeNetworkInterfaceReq{
		NetcardIds: []string{},
	}
	response, err := eksClient.DescribeNetworkInterface(&req)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func IsAttachedECS() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.IsAttachedECS("")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func QueryTaskStatus() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.QueryTaskStatus("")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}
