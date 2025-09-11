package main

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func AttachNetworkInterface() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)
	req := eks.AttachNetworkInterfaceReq{
		EcsId:     "",
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
	ak, sk := "Your Ak", "Your Sk"

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
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.DescribeNetworkInterface("")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func IsAttachedECS() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.IsAttachedECS("")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func QueryTaskStatus() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.QueryEventStatus("")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response)
}
