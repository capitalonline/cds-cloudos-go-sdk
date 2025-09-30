package slb

type OpenapiCommonPage struct {
	Total int `json:"Total,omitempty"`
}

type OpenApiCommonResp struct {
	Code    interface{} `json:"Code"`
	Message interface{} `json:"Message"`
}

type CommonTask struct {
	TaskId string `json:"TaskId"`
}

type ListVpcSlbReq struct {
	VpcId string `json:"VpcId"`
}

type Slb struct {
	SlbId       string   `json:"SlbId"`
	SlbName     string   `json:"SlbName"`
	BandwidthId string   `json:"BandwidthId,omitempty"`
	EipList     []string `json:"EipList,omitempty"`
	LanVipList  []string `json:"LanVipList,omitempty"`
}

type ListVpcSlbResult struct {
	// Data []Slb  `json:Data`
	Data []map[string]interface{} `json:"Data,omitempty"`
	OpenApiCommonResp
	CommonTask
}

type GetVpcSlbDetailReq struct {
	SlbId   string `json:"SlbId"`
	SlbName string `json:"SlbName"`
}

type SlbResultData struct {
	SlbId         string                   `json:"SlbId"`
	SlbName       string                   `json:"SlbName"`
	SlbStatus     string                   `json:"SlbStatus"`
	BandwidthId   string                   `json:"BandwidthId"`
	BandwidthName string                   `json:"BandwidthName"`
	VipList       []map[string]interface{} `json:"VipList,omitempty"`
}

type GetVpcSlbResult struct {
	Data SlbResultData `json:"data,omitempty"`
	OpenApiCommonResp
	CommonTask
}

type CreateVpcSlbReq struct {
	RegionCode        string `json:"RegionCode"`
	AvailableZoneCode string `json:"AvailableZoneCode"`
	VpcId             string `json:"VpcId"`
	Name              string `json:"Name"`
	NetType           string `json:"NetType"`
	ConfType          string `json:"ConfType"`
}

type CreateVpcSlbResult struct {
	Code    string             `json:"Code"`
	Message string             `json:"Message"`
	TaskId  string             `json:"TaskId,omitempty"`
	Data    CreateVpcSlbData   `json:"Data,omitempty"`
}

type CreateVpcSlbData struct {
	SlbId string `json:"SlbId"`
}

type DeleteVpcSlbReq struct {
	SlbId string `json:"SlbId"`
}

type DeleteVpcSlbResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type UpdateVpcSlbReq struct {
	SlbId   string `json:"SlbId"`
	Name    string `json:"Name,omitempty"`
	NetType string `json:"NetType,omitempty"`
	ConfType string `json:"ConfType,omitempty"`
}

type UpdateVpcSlbResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type DescribeVpcSlbListReq struct {
	VpcId      string `json:"VpcId,omitempty"`
	Keyword    string `json:"Keyword,omitempty"`
}

type DescribeVpcSlbListResult struct {
	Code    string                 `json:"Code"`
	Message string                 `json:"Message"`
	Data    DescribeVpcSlbListData `json:"Data,omitempty"`
}

type DescribeVpcSlbListData struct {
	SlbList []DescribeVpcSlbEntry `json:"SlbList,omitempty"`
	Total   int                   `json:"Total,omitempty"`
}

type DescribeVpcSlbEntry struct {
	AzId         string      `json:"AzId,omitempty"`
	AzName       string      `json:"AzName,omitempty"`
	BandwidthId  string      `json:"BandwidthId,omitempty"`
	BandwidthName string     `json:"BandwidthName,omitempty"`
	BillingMethod string     `json:"BillingMethod,omitempty"`
	ConfName      string     `json:"ConfName,omitempty"`
	SlbId        string      `json:"SlbId,omitempty"`
	SlbName      string      `json:"SlbName,omitempty"`
	NetType      string      `json:"NetType,omitempty"`
	NetTypeZh    string      `json:"NetTypeZh,omitempty"`
	RegionId     string      `json:"RegionId,omitempty"`
	RegionName   string      `json:"RegionName,omitempty"`
	Status       string      `json:"Status,omitempty"`
	StatusZh     string      `json:"StatusZh,omitempty"`
	VipList      []VipEntry  `json:"VipList,omitempty"`
	VpcId        string      `json:"VpcId,omitempty"`
	VpcName      string      `json:"VpcName,omitempty"`
	VpcSegments  string      `json:"VpcSegments,omitempty"`
}

type VipEntry struct {
	Address string `json:"Address,omitempty"`
	Type    string `json:"Type,omitempty"`
}

type DescribeVpcSlbDetailInfoReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbDetailInfoResult struct {
	Code    string                        `json:"Code"`
	Message string                        `json:"Message"`
	Data    DescribeVpcSlbDetailInfoData  `json:"Data,omitempty"`
}

type DescribeVpcSlbDetailInfoData struct {
	SlbId   string    `json:"SlbId,omitempty"`
	SlbName string    `json:"SlbName,omitempty"`
	NetType string    `json:"NetType,omitempty"`
	ConfName string   `json:"ConfName,omitempty"`
	Status  string    `json:"Status,omitempty"`
	VipList []VipEntry `json:"VipList,omitempty"`
}

type DescribeVpcSlbListenCreateInfoReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbListenCreateInfoResult struct {
	Code    string                                `json:"Code"`
	Message string                                `json:"Message"`
	Data    DescribeVpcSlbListenCreateInfoData    `json:"Data,omitempty"`
}

type DescribeVpcSlbListenCreateInfoData struct {
	SlbAclList []AclEntry `json:"SlbAclList,omitempty"`
}

type AclEntry struct {
	AclId   string `json:"AclId,omitempty"`
	AclName string `json:"AclName,omitempty"`
}

type CreateVpcSlbListenReq struct {
	SlbId          string `json:"SlbId"`
	ListenName     string `json:"ListenName,omitempty"`
	Vip            string `json:"Vip,omitempty"`
	VipId          string `json:"VipId,omitempty"`
	VipType        string `json:"VipType,omitempty"`
	ListenProtocol string `json:"ListenProtocol,omitempty"`
	ListenPort     int    `json:"ListenPort,omitempty"`
	AclId          string `json:"AclId,omitempty"`
	Scheduler      string `json:"Scheduler,omitempty"`
	HealthCheckInfo *HealthCheckInfo `json:"HealthCheckInfo,omitempty"`
}

type HealthCheckInfo struct {
	Protocol           string `json:"Protocol,omitempty"`
	Port               string `json:"Port,omitempty"`
	ConnectTimeout     string `json:"ConnectTimeout,omitempty"`
	DelayLoop          int    `json:"DelayLoop,omitempty"`
	Retry              int    `json:"Retry,omitempty"`
	Rise               int    `json:"Rise,omitempty"`
	DelayBeforeRetry   int    `json:"DelayBeforeRetry,omitempty"`
}

type CreateVpcSlbListenResult struct {
	Code    string                      `json:"Code"`
	Message string                      `json:"Message"`
	Data    CreateVpcSlbListenData      `json:"Data,omitempty"`
}

type CreateVpcSlbListenData struct {
	ListenId string `json:"ListenId,omitempty"`
}

type DeleteVpcSlbListenReq struct {
	ListenIds []string `json:"ListenIds"`
}

type DeleteVpcSlbListenResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type UpdateVpcSlbListenReq struct {
	ListenId       string           `json:"ListenId"`
	ListenName     string           `json:"ListenName,omitempty"`
	Scheduler      string           `json:"Scheduler,omitempty"`
	HealthCheckInfo *HealthCheckInfo `json:"HealthCheckInfo,omitempty"`
	Virtualhost    string           `json:"Virtualhost,omitempty"`
}

type UpdateVpcSlbListenResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type DescribeVpcSlbListenListReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbListenListResult struct {
	Code    string                        `json:"Code"`
	Message string                        `json:"Message"`
	Data    DescribeVpcSlbListenListData  `json:"Data,omitempty"`
}

type DescribeVpcSlbListenListData struct {
	ListenList []DescribeVpcListenEntry `json:"ListenList,omitempty"`
}

type DescribeVpcListenEntry struct {
	ListenId       string `json:"ListenId,omitempty"`
	ListenName     string `json:"ListenName,omitempty"`
	ListenProtocol string `json:"ListenProtocol,omitempty"`
	ListenPort     int    `json:"ListenPort,omitempty"`
	Scheduler      string `json:"Scheduler,omitempty"`
}

type QueryVpcSlbListenReq struct {
	SlbId   string `json:"SlbId,omitempty"`
	ListenId string `json:"ListenId,omitempty"`
}

type QueryVpcSlbListenResult struct {
	Code    string               `json:"Code"`
	Message string               `json:"Message"`
	Data    []QueryListenEntry   `json:"Data,omitempty"`
}

type QueryListenEntry struct {
	ListenId string `json:"ListenId,omitempty"`
	Port     int    `json:"Port,omitempty"`
	Protocol string `json:"Protocol,omitempty"`
}

type DescribeVpcSlbListenRsInfoReq struct {
	ListenId string `json:"ListenId"`
}

type DescribeVpcSlbListenRsInfoResult struct {
	Code    string                        `json:"Code"`
	Message string                        `json:"Message"`
	Data    DescribeVpcSlbListenRsInfoData `json:"Data,omitempty"`
}

type DescribeVpcSlbListenRsInfoData struct {
	ListenId string        `json:"ListenId,omitempty"`
	RsList   []BackendRs    `json:"RsList,omitempty"`
}

type BackendRs struct {
	RsId string `json:"RsId,omitempty"`
	Address string `json:"Address,omitempty"`
	Port int `json:"Port,omitempty"`
	Weight int `json:"Weight,omitempty"`
	Status string `json:"Status,omitempty"`
}

type CreateVpcSlbRsPortReq struct {
	ListenId string `json:"ListenId"`
	RsList   []RsPort `json:"RsList"`
}

type RsPort struct {
	RsId string `json:"RsId"`
	Port int    `json:"Port"`
	Weight int `json:"Weight,omitempty"`
}

type CreateVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type DeleteVpcSlbRsPortReq struct {
	ListenId string `json:"ListenId,omitempty"`
	RsList   []RsPort `json:"RsList,omitempty"`
}

type DeleteVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type UpdateVpcSlbRsPortReq struct {
	ListenId string `json:"ListenId"`
	RsList   []RsPort `json:"RsList"`
}

type UpdateVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type QueryVpcSlbRsPortReq struct {
	SlbId   string `json:"SlbId,omitempty"`
	ListenId string `json:"ListenId,omitempty"`
}

type QueryVpcSlbRsPortResult struct {
	Code    string        `json:"Code"`
	Message string        `json:"Message"`
	Data    []RsPortEntry `json:"Data,omitempty"`
}

type RsPortEntry struct {
	RsId  string `json:"RsId,omitempty"`
	Port  int    `json:"Port,omitempty"`
	Weight int   `json:"Weight,omitempty"`
	Status string `json:"Status,omitempty"`
}

type BandwidthBindResourceReq struct {
	BandwidthId string `json:"BandwidthId"`
	ResourceId  string `json:"ResourceId"`
	ResourceType string `json:"ResourceType,omitempty"`
}

type BandwidthBindResourceResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type BandwidthUnbindResourceReq struct {
	BandwidthId string `json:"BandwidthId"`
	ResourceId  string `json:"ResourceId"`
}

type BandwidthUnbindResourceResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}
