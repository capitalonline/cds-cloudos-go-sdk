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

	"github.com/capitalonline/cds-cloudos-go-sdk/services/eip"
)

// ListEIPs 查询eip列表
func ListEIPs() {
	ak, sk := "ak", "sk"
	EipClient, _ := eip.NewClient(ak, sk)
	ListEipArgs := &eip.ListEipsReq{
		RegionCode:"CN_Qingyang",
	}
	response, err := EipClient.ListEips(ListEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
// GetEIP 获取某个eip信息
func GetEIP() {
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	GetEipArgs := &eip.GetEipReq{
		Keyword: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",  // eip地址或者id
	}
	response, err := EipClient.GetEip(GetEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
// CreateEIP 创建eip
func CreateEIP(){
	ak, sk := "ak", "sk"
	EipClient, _ := eip.NewClient(ak, sk)
	CreateEipArgs := &eip.CreateEIPReq{
		RegionCode: "CN_Qingyang",  // 区域code
        BandwidthType: "Bandwidth_Multi_ISP_BGP",  // 带宽类型
        BillScheme: "BandwIdth",  // 计费方案
        Qos: 5,  // 带宽大小
        Size: 1,  // 数量
        Description: "go create",
	}
	response, err := EipClient.CreateEip(CreateEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
// UpdateEIP 更新eip带宽或者描述
func UpdateEIP(){
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	UpdateEipArgs := &eip.UpdateEIPReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
        Qos: 10,
		Description: "go create",
	}

	response, err := EipClient.UpdateEip(UpdateEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

// ReleaseEIP删除弹性eip
func ReleaseEIP(){
	ak, sk := "ak", "sk"


	EipClient, _ := eip.NewClient(ak, sk)
	ReleaseEipArgs := &eip.ReleaseEipReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
	}

	response, err := EipClient.ReleaseEip(ReleaseEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func main() {
	// ListEIPs()
	// GetEIP()
	// CreateEIP()
	// UpdateEIP()
	ReleaseEIP()
}