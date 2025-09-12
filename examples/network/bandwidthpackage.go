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

// ListBandwidthPackage 查询共享带宽包
func ListBandwidthPackage() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	ListBandwidthPackageArgs := &bandwidthpackage.ListBandwidthPackagesReq{
		RegionCode:"CN_Qingyang", // VPC区域code
	}
	response, err := BandwidthPackageClient.ListBandwidthPackages(ListBandwidthPackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
//获取共享带宽包
func GetBandwidthPackage() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	GetBandwidthPackageArgs := &bandwidthpackage.GetBandwidthPackageReq{
		Keyword: "868bb384-79a4-11f0-adfa-6e18e986f14e", //查询关键字: id或名称
	}
	response, err := BandwidthPackageClient.GetBandwidthPackage(GetBandwidthPackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

// CreateBandwidthPackage 创建共享带宽包
func CreateBandwidthPackage() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	CreatebandwidthpackageArgs := &bandwidthpackage.CreateBandwidthPackageReq{
        RegionCode: "CN_Qingyang",  // 	VPC区域code
        Name: "go_create",
        AvailableZoneCode: "CN_Qingyang_A",  // VPC可用区code
        BandwidthType: "Bandwidth_Multi_ISP_BGP", // 带宽类型
        BillScheme: "BandwIdth_Shared", // 计费方案
        Qos: 10,
    
	}
	response, err := BandwidthPackageClient.CreateBandwidthPackage(CreatebandwidthpackageArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

//UpdateBandwidthPackage 更新带宽包带宽
func UpdateBandwidthPackage() {
	// 替换为您的实际访问密钥
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

//DeleteBandwidthPackage 删除共享带宽包
func DeleteBandwidthPackage() {
	// 替换为您的实际访问密钥
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

//AddBandwidthPackageIp 为共享带宽包添加EIP
func AddBandwidthPackageIp() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	AddBandwidthPackageIpArgs := &bandwidthpackage.AddBandwidthPackageIpReq{
		BandwidthId:"868bb384-79a4-11f0-adfa-6e18e986f14e",
		EIPIdList: []string{"6870eeac-79ac-11f0-8503-6e18e986f14e"},  // eip id
	}
	response, err := BandwidthPackageClient.AddBandwidthPackageIp(AddBandwidthPackageIpArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

//RemoveBandwidthPackageIp 从共享带宽包移除eip
func RemoveBandwidthPackageIp() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	BandwidthPackageClient, _ := bandwidthpackage.NewClient(ak, sk)
	// 移除保留eip
	RemoveBandwidthPackageIpArgs := &bandwidthpackage.RemoveBandwidthPackageIpReq{
		EIPIdList: []string{"70cf50e2-79a3-11f0-9be8-6e18e986f14e"},
		// Delete: true,
		Delete: false,  // 保留弹性eip, 为true时，下方参数可不传
		RegionCode: "CN_Qingyang",  // 区域code
        BandwidthType: "Bandwidth_Multi_ISP_BGP",  // 带宽类型
        BillScheme: "BandwIdth",  // 计费方案
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

func main(){
	// ListBandwidthPackage()
	// GetBandwidthPackage()
	// CreateBandwidthPackage()
	// UpdateBandwidthPackage()
	// DeleteBandwidthPackage()
	// AddBandwidthPackageIp()
	// RemoveBandwidthPackageIp()
}