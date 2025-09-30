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
	"encoding/json"
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/slb"
)

func ListVPCSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.ListVpcSlbReq{
		VpcId: "635957ee-9543-11f0-8037-c2aae808f99f",
	}
	response, err := slbClient.ListVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetVPCSlbDetail() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	args := &slb.GetVpcSlbDetailReq{
		SlbId:   "",
		SlbName: "",
	}
	response, err := slbClient.GetVpcSlbDetail(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func CreateVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.CreateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DeleteVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func UpdateVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.UpdateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbList() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DescribeVpcSlbList(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbDetailInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbDetailInfoReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DescribeVpcSlbDetailInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbListenCreateInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenCreateInfoReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DescribeVpcSlbListenCreateInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func CreateVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbListenReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.CreateVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbListenReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DeleteVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func UpdateVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbListenReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.UpdateVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbListenList() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenListReq{
		SlbId: "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.DescribeVpcSlbListenList(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func QueryVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.QueryVpcSlbListenReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.QueryVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbListenRsInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenRsInfoReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DescribeVpcSlbListenRsInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func CreateVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbRsPortReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.CreateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbRsPortReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.DeleteVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func UpdateVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbRsPortReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.UpdateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func QueryVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.QueryVpcSlbRsPortReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.QueryVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func BandwidthBindResource() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthBindResourceReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.BandwidthBindResource(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func BandwidthUnbindResource() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthUnbindResourceReq{
		// TODO: 填充请求参数
	}
	response, err := slbClient.BandwidthUnbindResource(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func main() {
	ListVPCSlb()
	// GetVPCSlbDetail()
	//DescribeVpcSlbList()
	//DescribeVpcSlbListenList()
}
