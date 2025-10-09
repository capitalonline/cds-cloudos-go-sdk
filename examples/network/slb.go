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
	//fmt.Printf(">>> response: %+v", response)
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
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func CreateVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbReq{
		RegionCode:        "CN_Qingyang",
		AvailableZoneCode: "CN_Qingyang_A",
		VpcId:             "635957ee-9543-11f0-8037-c2aae808f99f",
		ConfType:          "slb.v1.small",
		NetType:           "wan",
		Name:              "创建测试openapi",
	}
	response, err := slbClient.CreateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbReq{
		SlbId: "d33fca12-a4ec-11f0-9ab9-5aeb8613167f",
	}
	response, err := slbClient.DeleteVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func UpdateVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbReq{
		SlbId:    "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
		ConfType: "slb.v1.small",
		Name:     "lxl-test",
		NetType:  "wan",
	}
	response, err := slbClient.UpdateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbList() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListReq{
		VpcId: "635957ee-9543-11f0-8037-c2aae808f99f",
	}
	response, err := slbClient.DescribeVpcSlbList(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbDetailInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbDetailInfoReq{
		SlbId: "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.DescribeVpcSlbDetailInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbListenCreateInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenCreateInfoReq{
		SlbId: "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.DescribeVpcSlbListenCreateInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func CreateVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbListenReq{
		SlbId:          "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
		ListenName:     "test1",
		Vip:            "38.123.96.8",
		VipId:          "22e20ec2-9778-11f0-8f0f-962fdefe13b1",
		VipType:        "wan_eip",
		ListenProtocol: "TCP",
		ListenPort:     8082,
		AclId:          "",
		ListenTimeout:  10,
		Scheduler:      "rr",
		HealthCheckInfo: slb.HealthCheckInfo{
			Protocol:         "TCP",
			Virtualhost:      "",
			Port:             0,
			Path:             "",
			StatusCode:       200,
			ConnectTimeout:   5,
			DelayLoop:        5,
			Retry:            2,
			DelayBeforeRetry: 10,
		},
	}
	response, err := slbClient.CreateVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbListenReq{
		ListenIds: []string{"8c1be1ea-9ddd-11f0-ac07-428d3fdeacd7"},
	}
	response, err := slbClient.DeleteVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func UpdateVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbListenReq{
		ListenId:      "4751aff2-9916-11f0-b7f3-7a25eed710cc",
		ListenName:    "test1",
		AclId:         "",
		ListenTimeout: 10,
		Scheduler:     "rr",
		HealthCheckInfo: slb.HealthCheckInfo{
			Protocol:         "TCP",
			Port:             0,
			ConnectTimeout:   5,
			DelayLoop:        5,
			Retry:            2,
			DelayBeforeRetry: 10,
		},
	}
	response, err := slbClient.UpdateVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
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
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func QueryVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.QueryVpcSlbListenReq{
		ListenId: "4751aff2-9916-11f0-b7f3-7a25eed710cc",
	}
	response, err := slbClient.QueryVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DescribeVpcSlbListenRsInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenRsInfoReq{
		VmType: "kvm",
		VpcId:  "635957ee-9543-11f0-8037-c2aae808f99f",
	}
	response, err := slbClient.DescribeVpcSlbListenRsInfo(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func CreateVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbRsPortReq{
		ListenId: "4751aff2-9916-11f0-b7f3-7a25eed710cc",
		RsList: []slb.RsPortItem{
			{
				VmId:        "ins-qyvk5dju93nh1cs0",
				VmName:      "ne1aAHzJSr2S1GxC-003",
				VmPublicIp:  "38.123.96.14",
				VmType:      "kvm",
				VmPrivateIp: "10.0.0.6",
				Port:        "8080",
				Weight:      "100",
			},
		},
	}
	response, err := slbClient.CreateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbRsPortReq{
		RsPortIds: []string{"a7dccc28-a4e0-11f0-be74-5aeb8613167f"},
	}
	response, err := slbClient.DeleteVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func UpdateVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbRsPortReq{
		RsPortList: []slb.RsPortItem{
			{
				RsPortId: "8476da80-a4e0-11f0-8979-5aeb8613167f",
				Port:     "80",
				Weight:   "200",
			},
		},
	}
	response, err := slbClient.UpdateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func QueryVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.QueryVpcSlbRsPortReq{
		ListenId: "4751aff2-9916-11f0-b7f3-7a25eed710cc",
		Keyword:  "",
	}
	response, err := slbClient.QueryVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func BandwidthBindResource() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthBindResourceReq{
		BandwidthId: "8bbe3f84-9777-11f0-b451-c2aae808f99f",
		BindType:    "VPCSLB",
		ResourceId:  "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.BandwidthBindResource(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func BandwidthUnbindResource() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthUnbindResourceReq{
		BandwidthId: "8bbe3f84-9777-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.BandwidthUnbindResource(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func main() {
	//ListVPCSlb()
	//GetVPCSlbDetail()

	//CreateVpcSlb()
	DeleteVpcSlb()
	//UpdateVpcSlb()
	//DescribeVpcSlbList()
	//DescribeVpcSlbDetailInfo()
	//DescribeVpcSlbListenCreateInfo()
	//CreateVpcSlbListen()
	//DeleteVpcSlbListen()
	//UpdateVpcSlbListen()
	//DescribeVpcSlbListenList()
	//QueryVpcSlbListen()
	//DescribeVpcSlbListenRsInfo()
	//CreateVpcSlbRsPort()
	//DeleteVpcSlbRsPort()
	//UpdateVpcSlbRsPort()
	//QueryVpcSlbRsPort()
	//BandwidthBindResource()
	//BandwidthUnbindResource()
}
