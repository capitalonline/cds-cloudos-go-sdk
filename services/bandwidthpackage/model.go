package bandwidthpackage

type OpenapiCpmmonPage struct {
	total int `json:"",omitempty`
}

type OpenApiCommonResp struct {
	Code    interface {} `json:"Code"`
	Message interface{}  `json:"Message`
}

type CommonTask struct {
	TaskId string `json:"TaskId,omitempty"`
}


type CreateBandwidthPackageReq struct {
	RegionCode        string `json:"RegionCode"`
    Name              string `json:"Name"`
    AvailableZoneCode string `json:"AvailableZoneCode"`
    BandwidthType     string `json:"BandwidthType"`
    BillScheme        string `json:"BillScheme"`
    Qos               int `json:"Qos"`
    NETID             string `json:"NETID"`
}


type CreateBandwidthPackageResData struct {
	BandwidthId string `json:"BandwidthId"`
}


type CreateBandwidthPackageResult struct {
	Data []CreateBandwidthPackageResData     `json:"data,omitempty"`
	OpenApiCommonResp
}


type UpdateBandwidthPackageReq struct {
	BandwidthId string `json:"BandwidthId"`
	Qos         int    `json:"Qos"`
}


type UpdateBandwidthResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type GetBandwidthPackageReq struct {
	Keyword  string `json:"Keyword"`
}


type Bandwidthpackage struct {
	AvailableZoneCode string `json:"AvailableZoneCode"`
    BandwidthType string `json:"BandwidthType"`
    BillScheme string `json:"BillScheme"`
    CreateTime string `json:"CreateTime"`
    Id string `json:"Id"`
    Name string `json:"Name"`
    Qos int `json:"Qos"`
    RegionCode string `json:"RegionCode"`
    Status string `json:"Status"`
}

type GetBandwidthPackageData struct {
	BandwidthList []Bandwidthpackage `json:"BandwidthList"`
	Total int `json:"Total"`
}


type GetBandwidthPackageResult struct {
	Data GetBandwidthPackageData `json:"data,omitempty"`
	OpenApiCommonResp
}


type ListBandwidthPackagesData struct {
	BandwidthList []Bandwidthpackage `json:"BandwidthList"`
	Total int `json:"Total"`
}

type ListBandwidthPackagesReq struct {
	Keyword           string `json:"PageNumber,omitempty"`
	RegionCode        string `json:"RegionCode"`
	AvailableZoneCode string `json:"AvailableZoneCode,omitempty"`
}


type ListBandwidthPackagesResult struct {
	Data ListBandwidthPackagesData `json:"data,omitempty"`
	OpenApiCommonResp
}


type DeleteBandwidthPackageReq struct {
	BandwidthId  string `json:"BandwidthId"`
}


type DeleteBandwidthPackageResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type AddBandwidthPackageIpReq struct {
	BandwidthId string `json:"BandwidthId"`
	EIPIdList   []string `json:"EIPIdList"`
}


type AddBandwidthPackageIpResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type RemoveBandwidthPackageIpReq struct {
	EIPIdList     []string `json:"EIPIdList"`
    Delete        bool     `json:"Delete"`
	RegionCode    string   `json:"RegionCode,omitempty"`
	AvailableZoneCode string `json:"AvailableZoneCode,omitempty"`
    BandwidthType string   `json:"BandwidthType,omitempty"`
    BillScheme    string   `json:"BillScheme,omitempty"`
    Qos           int      `json:"Qos,omitempty"`
}


type RemoveBandwidthPackageIpResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type BandwidthBindResourceReq struct {
	BandwidthId  string `json:"BandwidthId"`
	BindType string `json:"BindType"`
	ResourceId string `json:"ResourceId"`
}



type BandwidthBindResourceResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}


type BandwidthUnbindResourceReq struct {
	BandwidthId  string `json:"BandwidthId"`
}


type BandwidthUnbindResourceResult struct {
	Data    map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask

}

