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

	"github.com/capitalonline/cds-cloudos-go-sdk/services/subnet"
)

func DescribeSubnet() {
	ak, sk :=  "ak", "sk"

	subnetClient, _ := subnet.NewClient(ak, sk)
	ListSubnetArgs := &subnet.ListSubnetsReq{
		RegionCode: "CN_Qingyang",
	}
	response, err := subnetClient.ListSubnets(ListSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func CreateSubnet() {
	ak, sk :=  "ak", "sk"
	subnetClient, _ := subnet.NewClient(ak, sk)
	CreateSubnetArgs := &subnet.CreateSubnetReq{
        VPCId:        "9197340c-799e-11f0-adfa-6e18e986f14e",
        SubnetList: []subnet.CreateSubnetData{
            {
                AvailableZoneCode: "CN_Qingyang_A",
                SubnetName:        "子网2",
                SubnetSegment:     "10.21.2.0/24",
            },
        },
    }
	response, err := subnetClient.CreateSubnet(CreateSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func DeleteSubnet() {
	ak, sk :=  "ak", "sk"

	subnetClient, _ := subnet.NewClient(ak, sk)
	DeleteSubnetArgs := &subnet.DeleteSubnetReq{
		SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",
    }

	response, err := subnetClient.DeleteSubnet(DeleteSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func main(){
	// ListSubnet()
//    CreateSubnet()
    DeleteSubnet()
}