package vpc

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



type Subnet struct {
	AvailableZoneCode string `json:"AvailableZoneCode"`
	SubnetName        string `json:"SubnetName"`
	SubnetSegment     string `json:"SubnetSegment"`
}

type CreateVpcReq struct {
	RegionCode     string   `json:"RegionCode"`
	VPCName        string   `json:"VPCName"`
	VPCSegment     string   `json:"VPCSegment"`
	VPCSegmentType string   `json:"VPCSegmentType"`
	BandwidthType  string   `json:"BandwidthType`
	SubnetList     []Subnet `json:"SubnetList"`
	ProjectId      string   `json:"ProjectId,omitempty"`
}



type CreateVpcResultData struct {
	SubnetIdList []string `json:"SubnetIdList"`
	SubnetTaskId []string `json:"SubnetTaskId"`
	VPCId 		 string   `json:"VPCId"`
} 

type CreateVpcResult struct {
	Data CreateVpcResultData `json:"data,omitempty"`
	OpenApiCommonResp
	CommonTask
}


type GetVpcReq struct {
	Keyword string `json:"Keyword,omitempty"`
}


type VPC struct {
	CreateTime    string  `json:"CreateTime"`
	RegionCode    string  `json:"RegionCode"`
	Status 		  string  `json:"Status"`
	VPCId 		  string  `json:"VPCId"`
	VPCName 	  string  `json:"VPCNmae"`
	BandwidthType string `json:"BandwidthType,omitempty"`
	VPCSegment    string  `json:"VPCSegment"`
}


type GetVpcResultData struct {
	Total   int    `json:"Total"`
	VPCList []VPC  `json:"VPCList"`

}


type GetVpcResult struct {
	Data GetVpcResultData `json:"data,omitempty"`
	OpenApiCommonResp
}


type ListVpcsReq struct {
	Keyword    string `json:"PageNumber,omitempty"`
	RegionCode string `json:"RegionCode,omitempty"`
	PageNumber int    `json:"PageNumber,omitempty"`
}


type ListVpcsData struct {
	Total   int `json:"Total"`
	VPCList []VPC  `json:VPCList"`
}
type ListVpcsResult struct {
	Data ListVpcsData `json:"data,omitempty"`
	OpenApiCommonResp

}

type DeleteVpcReq struct {
	VPCId string `json:"VPCId"`
}

type DeleteVpcResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type DescribeTaskReq struct {
	TaskId string `json:"TaskId"`
}
type DescribeTaskData struct {
	TaskStatus string `json:"TaskStatus,omitempty"`
	ResourceId string `json:"ResourceId,omitempty"`
}
type DescribeTaskResult struct {
	Data    DescribeTaskData `json:"data,omitempty"`
	OpenApiCommonResp
	CommonTask
}
