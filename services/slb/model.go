package slb

type OpenapiCpmmonPage struct {
	total int `json:"",omitempty`
}

type OpenApiCommonResp struct {
	Code    interface {} `json:"Code"`
	Message interface{}  `json:"Message`
}

type CommonTask struct {
	TaskId string `json:"TaskId"`
}

type ListVpcSlbReq struct {
	VpcId string `json:"VpcId"`
}

type Slb struct {
	SlbId string `json:"SlbId"`
	SlbName string `json:"SlbName"`
	BandwidthId string `json:"BandwidthId,omitempty"`
    EipList []string `json:"EipList,omitempty"`
	LanVipList []string `json:"LanVipList,omitempty"`
}

type ListVpcSlbResult struct {
	// Data []Slb  `json:Data`
	Data []map[string]interface{} `json:"Data,omitempty"`
	OpenApiCommonResp
	CommonTask
}

type GetVpcSlbReq struct {
	SlbId string `json:"SlbId"`
	SlbName string `json:"SlbName"`
}


type SlbResultData struct {
	SlbId string `json:"SlbId"`
	SlbName string `json:"SlbName"`
	SlbStatus string `json:"SlbStatus"`
	BandwidthId string `json:"BandwidthId"`
	BandwidthName string `json:"BandwidthName"`
    VipList []map[string]interface{} `json:"VipList,omitempty"`
}

type GetVpcSlbResult struct {
	Data SlbResultData `json:"data,omitempty"`
	OpenApiCommonResp
	CommonTask
}
