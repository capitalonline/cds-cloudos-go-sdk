package natgateway

type OpenapiCommonPage struct {
	total int `json:"Total,omitempty"`
}

type OpenApiCommonResp struct {
	Code    interface {} `json:"Code"`
	Message interface{}  `json:"Message"`
}

type CommonTask struct {
	TaskId string `json:"TaskId,omitempty"`
}


type CreateNatGatewayReq struct {
	Name  string `json:"Name"`
	Desc string `json:"Desc,omitempty"`
	VPCId string `json:"VPCId"`
	AvailableZoneCode string `json:"AvailableZoneCode"`
	SpecificationId int `json:"SpecificationId"`
	BillingMethod int `json:"BillingMethod,omitempty"`
	IsAutoRenewal int `json:"IsAutoRenewal,omitempty"`
	IsToMonth int `json:"IsToMonth,omitempty"`
	Duration int `json:"Duration,omitempty"`
}


type CreateNatGatewayResultData struct {
	TaskId  string `json:"TaskId,omitempty"`
	NATId string `json:"NATId"`
}

type CreateNatGatewayResult struct {
	Data CreateNatGatewayResultData `json:"data,omitempty"`
	OpenApiCommonResp
}


type NAT struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Status string `json:"Status"`
	Description string `json:"Description"`
	RegionCode string `json:"RegionCode"`
	AvailableZoneCode string `json:"AvailableZoneCode"`
	VPCId string `json:"VPCId"`
	SpecificationId int `json:"SpecificationId"`
	BillingMethod int `json:"BillingMethod"`
}

type GetNatGatewayData struct {
	Total   int    `json:"Total"`
	NATList []NAT  `json:NATList`
}

type GetNatGatewayResult struct {
	Data GetNatGatewayData `json:"data,omitempty"`
	OpenApiCommonResp
}


type ListNatGatewaysData struct {
	Total   int    `json:"Total"`
	NATList []NAT  `json:NATList`
}


type ListNatGatewaysResult struct {
	Data ListNatGatewaysData `json:"data,omitempty"`
	OpenApiCommonResp
}


type UpdateNatGatewayReq struct {
	NATId string `json:"NATId"`
	Name  string `json:"Name,omitempty"`
	Desc string `json:"Desc,omitempty"`
	SpecificationId int `json:"SpecificationId,omitempty"`
	BillingMethod int `json:"BillingMethod,omitempty"`
	IsAutoRenewal int `json:"IsAutoRenewal,omitempty"`
	IsToMonth int `json:"IsToMonth,omitempty"`
	Duration int `json:"Duration,omitempty"`
}

type UpdateNatGatewayResultData struct {
	TaskId  string `json:"TaskId,omitempty"`
}

type UpdateNatGatewayResult struct {
	Data UpdateNatGatewayResultData `json:"data,omitempty"`
	OpenApiCommonResp
}


type DeleteNatGatewayReq struct {
	NATId string `json:"NATId"`
}


type DeleteNatGatewayResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type CreateNatRuleReq struct {
	Name string `json:"Name"`
	NATId string `json:"NATId"`
	EIPId string `json:"EIPId,omitempty"`
	RuleType string `json:"RuleType"`
	SourceAddress string `json:"SourceAddress,omitempty"`
	TranslatedAddress string `json:"TranslatedAddress,omitempty"`
	SubnetId string `json:"SubnetId,omitempty"`
	Protocol string `json:"Protocol,omitempty"`
	DestinationPort int `json:"DestinationPort,omitempty"`
	TranslatedPort int `json:"TranslatedPort,omitempty"`
}


type CreateNatRuleData struct {
	NATRuleId string `json:"NATRuleId"`
}

type CreateNatRuleResult struct {
	Data CreateNatRuleData `json:"data,omitempty"`
	OpenApiCommonResp
	CommonTask
}


type NatRule struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Status string `json:"Status"`
	Type string `json:"Type"`
	SourceAddress string `json:"SourceAddress,omitempty"`
	DestinationAddress string `json:"DestinationAddress,omitempty"`
	TranslatedAddress string `json:"TranslatedAddress,omitempty"`
	Protocol string `json:"Protocol"`
	DestinationPort int `json:"DestinationPort,omitempty"`
	TranslatedPort int `json:"TranslatedPort,omitempty"`
	NATId string `json:"NATId"`
	EIPId string `json:"EIPId,omitempty"`
}


type GetNatRuleData struct {
	NATRuleList []NatRule `json:"NATRuleList"`
	Total int `json:"Total"`
}

type GetNatRuleResult struct {
	Data GetNatRuleData `json:"data,omitempty"`
	OpenApiCommonResp
	CommonTask
}



type ListNatRulesData struct {
	NATRuleList []NatRule `json:"NATRuleList"`
	Total int `json:"Total"`
}

type ListNatRulesResult struct {
	Data ListNatRulesData `json:"data,omitempty"`
	OpenApiCommonResp
}


type UpdateNATRuleReq struct {
	NATRuleId string `json:"NATRuleId"`
	Name string `json:"Name"`
	NATId string `json:"NATId"`
	EIPId string `json:"EIPId,omitempty"`
	RuleType string `json:"RuleType"`
	SourceAddress string `json:"SourceAddress,omitempty"`
	TranslatedAddress string `json:"TranslatedAddress,omitempty"`
	SubnetId string `json:"SubnetId,omitempty"`
	Protocol string `json:"Protocol,omitempty"`
	DestinationPort int `json:"DestinationPort,omitempty"`
	TranslatedPort int `json:"TranslatedPort,omitempty"`
}



type UpdateNATRuleResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type DeleteNATRuleReq struct {
	NATRuleIdList []string   `json:"NATRuleIdList"`
	NATId string   `json:"NATId"`
}


type DeleteNATRuleResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}