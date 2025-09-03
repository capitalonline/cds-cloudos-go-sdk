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


func CreateNatGateway() {
	ak, sk := "ak", "sk"


	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	CreateNatGatewayArgs := &natgateway.CreateNatGatewayReq{
		Name: "go create band",
		Desc: "go create abnd",
		VPCId: "9197340c-799e-11f0-adfa-6e18e986f14e",
		AvailableZoneCode: "CN_Qingyang_A",
		SpecificationId: 12246,
	}
	response, err := natgatewayClient.CreateNatGateway(CreateNatGatewayArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func DeleteNatGateway() {
	ak, sk := "ak", "sk"


	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	DeleteNatGatewayArgs := &natgateway.DeleteNatGatewayReq{
		NATId: "c3e95ed4-79a9-11f0-a8a0-7a973848a269",
	}
	response, err := natgatewayClient.DeleteNatGateway(DeleteNatGatewayArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func UpdateNatGateway() {
	ak, sk := "ak", "sk"


	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	UpdateNatGatewayArgs := &natgateway.UpdateNatGatewayReq{
		NATId: "c3e95ed4-79a9-11f0-a8a0-7a973848a269",
		Name: "go create1",
		Desc: "go create1",
		SpecificationId: 12246,
	}
	response, err := natgatewayClient.UpdateNatGateway(UpdateNatGatewayArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func ListNatRules() {
	ak, sk := "ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	ListNatRulesArgs := map[string]string{
		"NATId": "c3e95ed4-79a9-11f0-a8a0-7a973848a269",
		"RuleType": "source",
	}
	response, err := natgatewayClient.ListNATRules(ListNatRulesArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func GetNatRule() {
	ak, sk :="ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	GetNatRuleArgs := map[string]string{
		"NATId": "c3e95ed4-79a9-11f0-a8a0-7a973848a269",
		"RuleType": "source",
	}
	response, err := natgatewayClient.GetNATRule(GetNatRuleArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func CreateNatRule() {
	ak, sk := "ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)
	// snat

	// CreateNatRuleArgs := &natgateway.CreateNatRuleReq{
	// 	Name: "snat",
	// 	NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
	// 	RuleType: "source",
	// 	EIPId:"6870eeac-79ac-11f0-8503-6e18e986f14e",
	// 	SourceAddress:"10.21.1.0/24",
	// }

	// 全端口dnat
	// CreateNatRuleArgs := &natgateway.CreateNatRuleReq{
	// 	Name: "dnat全端口",
	// 	NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
	// 	RuleType: "destination",
	// 	EIPId:"6870eeac-79ac-11f0-8503-6e18e986f14e",
	// 	TranslatedAddress:"10.21.1.4",
	// 	SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",
	// 	Protocol:"ALL",
	// }

	// 具体端口
	CreateNatRuleArgs := &natgateway.CreateNatRuleReq{
		Name: "dnat全端口",
		NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
		RuleType: "destination",
		EIPId:"70cf50e2-79a3-11f0-9be8-6e18e986f14e",
		TranslatedAddress:"10.21.1.5",
		SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",
		Protocol:"TCP",
		DestinationPort:80,
		TranslatedPort: 80,
	}
	response, err := natgatewayClient.CreateNATRule(CreateNatRuleArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func DeleteNatRule() {
	ak, sk := "ak", "sk"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)

	DeleteNatRuleArgs := &natgateway.DeleteNATRuleReq{
		NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
		NATRuleIdList:[]string{"ad152430-79af-11f0-8d7c-7a973848a269"},

	}
	response, err := natgatewayClient.DeleteNATRule(DeleteNatRuleArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func UpdateNatRule() {
	ak, sk := "9fa289fa729d11f09c1702852519bf7f", "651b4e9c04c3ce6cb444c54b81169bc4"
	natgatewayClient, _ := natgateway.NewClient(ak, sk)

	// snat

	UpdateNatRuleArgs := &natgateway.UpdateNATRuleReq{
		NATRuleId: "ad152430-79af-11f0-8d7c-7a973848a269",
		Name: "snat",
		NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
		RuleType: "source",
		EIPId:"832d9442-79b2-11f0-a8a0-7a973848a269",
		SourceAddress:"10.21.2.0/24",
	}

	//全端口dnat
	
	// UpdateNatRuleArgs := &natgateway.UpdateNATRuleReq{
	// 	NATRuleId:"2061819e-79b1-11f0-8503-6e18e986f14e",
	// 	Name: "dnat全端口",
	// 	NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
	// 	RuleType: "destination",
	// 	EIPId:"832d9442-79b2-11f0-a8a0-7a973848a269",
	// 	TranslatedAddress:"10.21.1.6",
	// 	SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",
	// 	Protocol:"ALL",
	// }

	// 具体端口

	// UpdateNatRuleArgs := &natgateway.UpdateNATRuleReq{
	// 	NATRuleId:"2061819e-79b1-11f0-8503-6e18e986f14e",
	// 	Name: "dnat全端口",
	// 	NATId:"c3e95ed4-79a9-11f0-a8a0-7a973848a269",
	// 	RuleType: "destination",
	// 	EIPId:"832d9442-79b2-11f0-a8a0-7a973848a269",
	// 	TranslatedAddress:"10.21.1.6",
	// 	SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",
	// 	Protocol:"UDP",
	// 	DestinationPort:81,
	// 	TranslatedPort: 81,
	// }
	
	response, err := natgatewayClient.UpdateNATRule(UpdateNatRuleArgs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

func main(){
	// ListNatGateways()
	// GetNatGateway()
	// CreateNatGateway()
	// UpdateNatGateway()
	DeleteNatGateway()
	// ListNatRules()
	// CreateNatRule()
	// UpdateNatRule()
	// DeleteNatRule()
}
