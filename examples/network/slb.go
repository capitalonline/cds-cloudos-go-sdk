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

func ListVpcSlb() {
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

func GetVpcSlbDetail() {
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
	fmt.Printf(">>> response: %+v", response)
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
	fmt.Printf(">>> response: %+v", response)
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
		NetType:  "wan_lan",
	}
	response, err := slbClient.UpdateVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func ListVpcSlbDetail() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.ListVpcSlbDetailReq{
		VpcId: "635957ee-9543-11f0-8037-c2aae808f99f",
	}
	response, err := slbClient.ListVpcSlbDetail(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetVpcSlb() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbReq{
		SlbId: "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.GetVpcSlb(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetVpcSlbListenCreateInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenCreateInfoReq{
		SlbId: "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.GetVpcSlbListenCreateInfo(req)
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
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbListenReq{
		ListenIds: []string{"37f895ec-a588-11f0-a90a-5aeb8613167f"},
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
		ListenId:      "7f208a88-a588-11f0-8979-5aeb8613167f",
		ListenName:    "test2",
		AclId:         "",
		ListenTimeout: 12,
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
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func ListVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenListReq{
		SlbId: "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
	}
	response, err := slbClient.ListVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetVpcSlbListen() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenReq{
		ListenId: "7f208a88-a588-11f0-8979-5aeb8613167f",
	}
	response, err := slbClient.GetVpcSlbListen(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetVpcSlbListenRsInfo() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenRsInfoReq{
		VmType: "kvm",
		VpcId:  "635957ee-9543-11f0-8037-c2aae808f99f",
	}
	response, err := slbClient.GetVpcSlbListenRsInfo(req)
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
		ListenId: "7f208a88-a588-11f0-8979-5aeb8613167f",
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
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbRsPortReq{
		RsPortIds: []string{"b4a85c84-a589-11f0-bf17-5aeb8613167f"},
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
		RsPortList: []slb.RsPortItem{
			{
				RsPortId: "b4a85c84-a589-11f0-bf17-5aeb8613167f",
				Port:     "80",
				Weight:   "200",
			},
		},
	}
	response, err := slbClient.UpdateVpcSlbRsPort(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetVpcSlbRsPort() {
	// 替换为您的实际访问密钥
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbRsPortReq{
		ListenId: "7f208a88-a588-11f0-8979-5aeb8613167f",
		Keyword:  "",
	}
	response, err := slbClient.GetVpcSlbRsPort(req)
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
		BandwidthId: "8bbe3f84-9777-11f0-b451-c2aae808f99f",
		BindType:    "VPCSLB",
		ResourceId:  "0f9c3a1e-9551-11f0-b451-c2aae808f99f",
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
		BandwidthId: "8bbe3f84-9777-11f0-b451-c2aae808f99f",
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
	//ListVpcSlb()
	//GetVpcSlbDetail()

	//CreateVpcSlb()
	//DeleteVpcSlb()
	//UpdateVpcSlb()
	//ListVpcSlbDetail()
	//GetVpcSlb()
	//GetVpcSlbListenCreateInfo()
	//CreateVpcSlbListen()
	//DeleteVpcSlbListen()
	//UpdateVpcSlbListen()
	//ListVpcSlbListen()
	//GetVpcSlbListen()
	//GetVpcSlbListenRsInfo()
	//CreateVpcSlbRsPort()
	//DeleteVpcSlbRsPort()
	//UpdateVpcSlbRsPort()
	//GetVpcSlbRsPort()
	//BandwidthBindResource()
	//BandwidthUnbindResource()
}
