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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.ListVpcSlbReq{
		VpcId: "",
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
	ak, sk := "ak", "sk"

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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbReq{
		RegionCode:        "CN_Qingyang",
		AvailableZoneCode: "CN_Qingyang_A",
		VpcId:             "",
		ConfType:          "slb.v1.small",
		NetType:           "wan",
		Name:              "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbReq{
		SlbId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbReq{
		SlbId:    "",
		ConfType: "slb.v1.small",
		Name:     "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.ListVpcSlbDetailReq{
		VpcId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbReq{
		SlbId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenCreateInfoReq{
		SlbId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbListenReq{
		SlbId:          "",
		ListenName:     "",
		Vip:            "",
		VipId:          "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbListenReq{
		ListenIds: []string{""},
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbListenReq{
		ListenId:      "",
		ListenName:    "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DescribeVpcSlbListenListReq{
		SlbId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenReq{
		ListenId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbListenRsInfoReq{
		VmType: "kvm",
		VpcId:  "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.CreateVpcSlbRsPortReq{
		ListenId: "",
		RsList: []slb.RsPortItem{
			{
				VmId:        "",
				VmName:      "",
				VmPublicIp:  "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.DeleteVpcSlbRsPortReq{
		RsPortIds: []string{""},
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.UpdateVpcSlbRsPortReq{
		RsPortList: []slb.RsPortItem{
			{
				RsPortId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.GetVpcSlbRsPortReq{
		ListenId: "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthBindResourceReq{
		BandwidthId: "",
		BindType:    "VPCSLB",
		ResourceId:  "",
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
	ak, sk := "ak", "sk"

	slbClient, _ := slb.NewClient(ak, sk)
	req := &slb.BandwidthUnbindResourceReq{
		BandwidthId: "",
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
