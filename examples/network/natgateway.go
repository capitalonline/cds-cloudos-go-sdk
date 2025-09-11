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

	"github.com/capitalonline/cds-cloudos-go-sdk/services/natgateway"
)

// ListNatGateways 查询nat网关
func ListNatGateways() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"

	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	ListNatGatewaysArgs := map[string]string{
		"RegionCode": "CN_Qingyang",  // 区域code
	}
	response, err := natgatewayClient.ListNatGateways(ListNatGatewaysArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
// GetNatGateway 获取nat网关数据
func GetNatGateway() {
	// 替换为您的实际访问密钥
	ak, sk := "ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	GetNatGatewayArgs := map[string]string{
		// RegionCode: "CN_SJZ",
		"Keyword": "c3e95ed4-79a9-11f0-a8a0-7a973848a269", // nat网关id或者名称

	}
	response, err := natgatewayClient.GetNatGateway(GetNatGatewayArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}




func main(){
	// ListNatGateways()
	GetNatGateway()

}
