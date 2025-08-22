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
	"github.com/capitalonline/cds-cloudos-go-sdk/services/slb"
)


func ListVpcSlb() {
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	listVpcSlbArgs := &slb.ListVpcSlbReq{
		VpcId: "be4c8d00-7d8e-11f0-8290-2e07174785c2",
	}
	response, err := slbClient.ListVpcSlb(listVpcSlbArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func GetVpcSlb() {
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	getVpcSlbArgs := &slb.GetVpcSlbReq{
		SlbId: "4c1a7638-7d8f-11f0-9ea8-2e07174785c2",
		SlbName: "",
	}
	response, err := slbClient.GetVpcSlb(getVpcSlbArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func main() {
	ListVpcSlb()
// 	GetVpcSlb()
}
