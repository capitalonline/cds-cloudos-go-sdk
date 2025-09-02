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

	"github.com/capitalonline/cds-cloudos-go-sdk/services/bandwidthpackage"
)


func DescribeBandwidth() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	ListBandwidthPackageArgs := &bandwidthpackage.ListBandwidthPackagesReq{
		RegionCode:"CN_Qingyang",
	}
	response, err := BandwidthPackageClient.ListBandwidthPackages(ListBandwidthPackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}


func CreateBandwidth() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	CreatebandwidthpackageArgs := &bandwidthpackage.CreateBandwidthPackageReq{
        RegionCode: "CN_Qingyang",
        Name: "go_create",
        AvailableZoneCode: "CN_Qingyang_A",
        BandwidthType: "Bandwidth_Multi_ISP_BGP",
        BillScheme: "BandwIdth_Shared",
        Qos: 10,
    
	}
	response, err := BandwidthPackageClient.CreateBandwidthPackage(CreatebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func UpdateBandwidth() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	UpdatebandwidthpackageArgs := &bandwidthpackage.UpdateBandwidthPackageReq{
        BandwidthId: "868bb384-79a4-11f0-adfa-6e18e986f14e",
        Qos: 20,
	}
	response, err := BandwidthPackageClient.UpdateBandwidthPackage(UpdatebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func DeleteBandwidth() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	DeletebandwidthpackageArgs := &bandwidthpackage.DeleteBandwidthPackageReq{
        BandwidthId: "868bb384-79a4-11f0-adfa-6e18e986f14e",
	}
	response, err := BandwidthPackageClient.DeleteBandwidthPackage(DeletebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func BandwidthAddEIP() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	AddBandwidthPackageIpArgs := &bandwidthpackage.AddBandwidthPackageIpReq{
		BandwidthId:"868bb384-79a4-11f0-adfa-6e18e986f14e",
		EIPIdList: []string{"6870eeac-79ac-11f0-8503-6e18e986f14e"},
	}
	response, err := BandwidthPackageClient.AddBandwidthPackageIp(AddBandwidthPackageIpArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func BandwidthRemoveEIP() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	// 移除保留eip
	RemoveBandwidthPackageIpArgs := &bandwidthpackage.RemoveBandwidthPackageIpReq{
		EIPIdList: []string{"70cf50e2-79a3-11f0-9be8-6e18e986f14e"},
		// Delete: true,
		Delete: false,
		RegionCode: "CN_Qingyang",
        BandwidthType: "Bandwidth_Multi_ISP_BGP",
        BillScheme: "BandwIdth",
        Qos: 10,
	}
	
	// 移除删除eip
	// RemoveBandwidthPackageIpArgs := &bandwidthpackage.RemoveBandwidthPackageIpReq{
	// 	EIPIdList: []string{"6870eeac-79ac-11f0-8503-6e18e986f14e"},
	// 	Delete: true,
	// }

	response, err := BandwidthPackageClient.RemoveBandwidthPackageIp(RemoveBandwidthPackageIpArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func BandwidthBindResource() {
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	BandwidthPackageBindResourceArgs := &bandwidthpackage.BandwidthBindResourceReq{
		BandwidthId:"868bb384-79a4-11f0-adfa-6e18e986f14e",
		BindType: "NAT",
		ResourceId: "c3e95ed4-79a9-11f0-a8a0-7a973848a269",
	}
	response, err := BandwidthPackageClient.BandwidthBindResource(BandwidthPackageBindResourceArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func BandwidthUnbindResource() {
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	BandwidthPackageUnbindResourceArgs := &bandwidthpackage.BandwidthUnbindResourceReq{
		BandwidthId:"868bb384-79a4-11f0-adfa-6e18e986f14e",
	}
	response, err := BandwidthPackageClient.BandwidthUnbindResource(BandwidthPackageUnbindResourceArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func main(){
}