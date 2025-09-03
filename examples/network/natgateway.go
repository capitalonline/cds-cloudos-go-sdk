package main

import (
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/natgateway"
)

func ListNatGateways() {
	ak, sk := "ak", "sk"

	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	ListNatGatewaysArgs := map[string]string{
		"RegionCode": "CN_Qingyang",
	}
	response, err := natgatewayClient.ListNatGateways(ListNatGatewaysArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func GetNatGateway() {
	ak, sk := "ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	GetNatGatewayArgs := map[string]string{
		// RegionCode: "CN_SJZ",
		"Keyword": "c3e95ed4-79a9-11f0-a8a0-7a973848a269",

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
