package subnet

type OpenapiCommonPage struct {
	total int `json:"Total,omitempty"`
}

type OpenApiCommonResp struct {
	Code    interface {} `json:"Code"`
	Message interface{}  `json:"Message"`
}

type CommonTask struct {
	TaskId string `json:"TaskId"`
}


type CreateSubnetData struct {
	AvailableZoneCode string `json:"AvailableZoneCode"`
	SubnetName        string `json:"SubnetName"`
	SubnetSegment     string `json:"SubnetSegment"`
}


type CreateSubnetReq struct {
	VPCId      string             `json:"VPCId"`
    SubnetList []CreateSubnetData  `json:"SubnetList"`
}

type CreateSubnetResdata struct {
	SubnetIdList []string `json:"SubnetIdList"`
	TaskIdList   []string `json:"TaskIdList"`
}


type CreateSubnetResult struct {
	Data CreateSubnetResdata     `json:"data,omitempty"`
	OpenApiCommonResp
	
}



type ListSubnetsReq struct {
	Keyword           string `json:"PageNumber,omitempty"`
	RegionCode        string `json:"RegionCode"`
	PageNumber        int    `json:"PageNumber,omitempty"`
	VPCId             string `json:"VPCId,omitempty"`
	AvailableZoneCode string `json:"AvailableZoneCode,omitempty"`
}


type Subnet struct {
	
    Status            string `json:"Status"`
    SubnetId          string `json:"SubnetId"`
    SubnetName        string `json:"SubnetName"`
    SubnetSegment     string `json:"SubnetSegment"` 
    UsedIPNum         int    `json:"UsedIPNum"`
    VPCId             string `json:"VPCId"`
    VPCName           string `json:"VPCName"`
	AvailableZoneCode string `json:"AvailableZoneCode"`

}

type ListSubnetsData struct {
	Total      int       `json:"Total"`
	SubnetList []Subnet  `json:SubnetList`
}

type ListSubnetsResult struct {
	Data ListSubnetsData `json:"data,omitempty"`
	OpenApiCommonResp

}


type GetSubnetReq struct {
	Keyword string `json:"Keyword,omitempty"`
}

type GetSubnetData struct {
	Total      int       `json:"Total"`
	SubnetList []Subnet  `json:SubnetList`
}


type GetSubnetResult struct {
	Data GetSubnetData `json:"data,omitempty"`
	OpenApiCommonResp
}


type DeleteSubnetReq struct {
	SubnetId string   `json:"SubnetId"`
}

type DeleteSubnetResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}
