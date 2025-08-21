/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/vpc"
)


func Getvpc() {
	ak, sk := "ak", "sk"

	vpcClient, _ := vpc.NewClient(ak, sk)
	GetVpcArgs := &vpc.GetVpcReq{
		Keyword: "9197340c-799e-11f0-adfa-6e18e986f14e",
	}
	response, err := vpcClient.GetVpc(GetVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.VPCList)
}

func CreateVpc() {
	ak, sk := "ak", "sk"

	vpcClient, _ := vpc.NewClient(ak, sk)
	CreateVpcArgs := &vpc.CreateVpcReq{
        RegionCode:     "CN_Qingyang",
        VPCName:        "go_create_bandtype",
        VPCSegment:     "10.21.0.0/16",
        VPCSegmentType: "auto",
		BandwidthType: "Bandwidth_Multi_ISP_BGP",
        SubnetList: []vpc.Subnet{
            {
                AvailableZoneCode: "CN_Qingyang_A",
                SubnetName:        "gocreate子网1",
                SubnetSegment:     "10.21.1.0/24",
            },
        },
    }

	response, err := vpcClient.CreateVPC(CreateVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func ListVpcs() {
	ak, sk := "ak", "sk"

	vpcClient, _ := vpc.NewClient(ak, sk)
	ListVpcArgs := &vpc.ListVpcsReq{
		RegionCode:"CN_Qingyang",
	}
	response, err := vpcClient.ListVpcs(ListVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.VPCList)
}

func DeleteVPC() {
	ak, sk := "ak", "sk"
	vpcClient, _ := vpc.NewClient(ak, sk)
	DeleteVpcArgs := &vpc.DeleteVpcReq{
		VPCId: "9197340c-799e-11f0-adfa-6e18e986f14e",
	}
	response, err := vpcClient.DeleteVpc(DeleteVpcArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func main() {
	CreateVpc()
	// Getvpc()
	// ListVpcs()
	// DeleteVPC()
}
