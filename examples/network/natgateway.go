package main

import (
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/natgateway"
)

func DescribeNAT() {
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


func main(){
	DescribeNAT()
}
