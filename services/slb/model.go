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
	Code    string           `json:"Code"`
	Message string           `json:"Message"`
	TaskId  string           `json:"TaskId,omitempty"`
	Data    CreateVpcSlbData `json:"Data,omitempty"`
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
	SlbId    string `json:"SlbId"`
	Name     string `json:"Name"`
	NetType  string `json:"NetType"`
	ConfType string `json:"ConfType"`
}

type UpdateVpcSlbResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type ListVpcSlbDetailReq struct {
	VpcId   string `json:"VpcId,omitempty"`
	Keyword string `json:"Keyword,omitempty"`
}

type DescribeVpcSlbListResult struct {
	Code    string                 `json:"Code"`
	Message string                 `json:"Message"`
	Data    DescribeVpcSlbListData `json:"Data,omitempty"`
}

type DescribeVpcSlbListData struct {
	SlbList []DescribeVpcSlbEntry `json:"SlbList,omitempty"`
	//Total   int                   `json:"Total,omitempty"`
}

type DescribeVpcSlbEntry struct {
	AzId          string `json:"AzId,omitempty"`
	AzName        string `json:"AzName,omitempty"`
	BandwidthId   string `json:"BandwidthId,omitempty"`
	BandwidthName string `json:"BandwidthName,omitempty"`
	BillingMethod string `json:"BillingMethod,omitempty"`
	ConfName      string `json:"ConfName,omitempty"`
	CreateTime    string `json:"CreateTime,omitempty"`
	Id            string `json:"Id,omitempty"`
	Name          string `json:"Name,omitempty"`
	NetType       string `json:"NetType,omitempty"`
	//NetTypeZh     string     `json:"NetTypeZh,omitempty"`
	RegionId   string     `json:"RegionId,omitempty"`
	RegionName string     `json:"RegionName,omitempty"`
	Status     string     `json:"Status,omitempty"`
	StatusZh   string     `json:"StatusZh,omitempty"`
	VipList    []VipEntry `json:"VipList,omitempty"`
	VpcId      string     `json:"VpcId,omitempty"`
	VpcName    string     `json:"VpcName,omitempty"`
	//VpcSegments   string     `json:"VpcSegments,omitempty"`
}

type VipEntry struct {
	Address string `json:"Address,omitempty"`
	Type    string `json:"Type,omitempty"`
	Vip     string `json:"Vip,omitempty"`
	VipId   string `json:"VipId,omitempty"`
	VipType string `json:"VipType,omitempty"`
}

type GetVpcSlbReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbDetailInfoResult struct {
	Code    string                       `json:"Code"`
	Message string                       `json:"Message"`
	Data    DescribeVpcSlbDetailInfoData `json:"Data,omitempty"`
}

type DescribeVpcSlbDetailInfoData struct {
	AzId              string        `json:"AzId,omitempty"`
	AzName            string        `json:"AzName,omitempty"`
	BandwidthInfo     BandwidthInfo `json:"BandwidthInfo,omitempty"`
	BillingMethod     string        `json:"BillingMethod,omitempty"`
	BillingSchemeId   string        `json:"BillingSchemeId,omitempty"`
	BillingSchemeName string        `json:"BillingSchemeName,omitempty"`
	ConfName          string        `json:"ConfName,omitempty"`
	CreateTime        string        `json:"CreateTime,omitempty"`
	Id                string        `json:"Id,omitempty"`
	Name              string        `json:"Name,omitempty"`
	NetType           string        `json:"NetType,omitempty"`
	NetTypeZh         string        `json:"NetTypeZh,omitempty"`
	RegionId          string        `json:"RegionId,omitempty"`
	RegionName        string        `json:"RegionName,omitempty"`
	Status            string        `json:"Status,omitempty"`
	StatusZh          string        `json:"StatusZh,omitempty"`
	VipList           []VipEntry    `json:"VipList,omitempty"`
	VpcId             string        `json:"VpcId,omitempty"`
	VpcName           string        `json:"VpcName,omitempty"`
	VpcSegments       string        `json:"VpcSegments,omitempty"`
}

type BandwidthInfo struct {
	Id   string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type GetVpcSlbListenCreateInfoReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbListenCreateInfoResult struct {
	Code    string                             `json:"Code"`
	Message string                             `json:"Message"`
	Data    DescribeVpcSlbListenCreateInfoData `json:"Data,omitempty"`
}

type DescribeVpcSlbListenCreateInfoData struct {
	SlbAclList []AclEntry `json:"SlbAclList,omitempty"`
	SlbVipList []VipEntry `json:"SlbVipList,omitempty"`
}

type AclEntry struct {
	AclId   string `json:"AclId,omitempty"`
	AclName string `json:"AclName,omitempty"`
}

type CreateVpcSlbListenReq struct {
	SlbId           string          `json:"SlbId"`
	ListenName      string          `json:"ListenName,omitempty"`
	Vip             string          `json:"Vip,omitempty"`
	VipId           string          `json:"VipId,omitempty"`
	VipType         string          `json:"VipType,omitempty"`
	ListenProtocol  string          `json:"ListenProtocol,omitempty"`
	ListenPort      int             `json:"ListenPort,omitempty"`
	ListenTimeout   int             `json:"ListenTimeout,omitempty"`
	AclId           string          `json:"AclId,omitempty"`
	Scheduler       string          `json:"Scheduler,omitempty"`
	HealthCheckInfo HealthCheckInfo `json:"HealthCheckInfo,omitempty"`
}

type HealthCheckInfo struct {
	Protocol         string `json:"Protocol,omitempty"`
	Port             int    `json:"Port,omitempty"`
	ConnectTimeout   int    `json:"ConnectTimeout,omitempty"`
	Virtualhost      string `json:"Virtualhost,omitempty"`
	Path             string `json:"Path,omitempty"`
	DelayLoop        int    `json:"DelayLoop,omitempty"`
	Retry            int    `json:"Retry,omitempty"`
	Rise             int    `json:"Rise,omitempty"`
	StatusCode       int    `json:"StatusCode,omitempty"`
	DelayBeforeRetry int    `json:"DelayBeforeRetry,omitempty"`
}

type CreateVpcSlbListenResult struct {
	Code    string                 `json:"Code"`
	Message string                 `json:"Message"`
	Data    CreateVpcSlbListenData `json:"Data,omitempty"`
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
	ListenId        string          `json:"ListenId"`
	ListenName      string          `json:"ListenName,omitempty"`
	AclId           string          `json:"AclId,omitempty"`
	ListenTimeout   int             `json:"ListenTimeout,omitempty"`
	Scheduler       string          `json:"Scheduler,omitempty"`
	Virtualhost     string          `json:"Virtualhost,omitempty"`
	HealthCheckInfo HealthCheckInfo `json:"HealthCheckInfo,omitempty"`
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
	Code    string                       `json:"Code"`
	Message string                       `json:"Message"`
	Data    DescribeVpcSlbListenListData `json:"Data,omitempty"`
}

type DescribeVpcSlbListenListData struct {
	ListenList []DescribeVpcListenEntry `json:"ListenList,omitempty"`
	Total      int                      `json:"Total"`
}

type DescribeVpcListenEntry struct {
	ListenId       string        `json:"ListenId,omitempty"`
	ListenName     string        `json:"ListenName,omitempty"`
	ListenPort     string        `json:"ListenPort,omitempty"`
	ListenProtocol string        `json:"ListenProtocol,omitempty"`
	ListenVip      string        `json:"ListenVip,omitempty"`
	Scheduler      string        `json:"Scheduler,omitempty"`
	SchedulerCn    string        `json:"SchedulerCn,omitempty"`
	Status         string        `json:"Status,omitempty"`
	StatusCn       string        `json:"StatusCn,omitempty"`
	CheckInfo      CheckInfoData `json:"CheckInfo,omitempty"`
}

type CheckInfoData struct {
	Error int `json:"Error"`
	Ok    int `json:"Ok"`
}

type GetVpcSlbListenReq struct {
	SlbId    string `json:"SlbId,omitempty"`
	ListenId string `json:"ListenId,omitempty"`
}

type QueryVpcSlbListenResult struct {
	Code    string           `json:"Code"`
	Message string           `json:"Message"`
	Data    QueryListenEntry `json:"Data,omitempty"`
}

type QueryListenEntry struct {
	ListenId        string          `json:"ListenId,omitempty"`
	ListenName      string          `json:"ListenName,omitempty"`
	ListenProtocol  string          `json:"ListenProtocol,omitempty"`
	ListenVip       string          `json:"ListenVip,omitempty"`
	Vip             string          `json:"Vip,omitempty"`
	VipId           string          `json:"VipId,omitempty"`
	VipType         string          `json:"VipType,omitempty"`
	ListenPort      int             `json:"ListenPort,omitempty"`
	ListenTimeout   int             `json:"ListenTimeout"`
	AclId           string          `json:"AclId"`
	Scheduler       string          `json:"Scheduler,omitempty"`
	SchedulerCn     string          `json:"SchedulerCn,omitempty"`
	Status          string          `json:"Status,omitempty"`
	StatusCn        string          `json:"StatusCn,omitempty"`
	HealthCheckInfo HealthCheckInfo `json:"HealthCheckInfo,omitempty"`
}

type GetVpcSlbListenRsInfoReq struct {
	VmType string `json:"VmType"`
	VpcId  string `json:"VpcId"`
}

type DescribeVpcSlbListenRsInfoResult struct {
	Code    string                           `json:"Code"`
	Message string                           `json:"Message"`
	Data    []DescribeVpcSlbListenRsInfoData `json:"Data,omitempty"`
}

type DescribeVpcSlbListenRsInfoData struct {
	VmId        string `json:"vm_id,omitempty"`
	VmName      string `json:"vm_name,omitempty"`
	VmPrivateIp string `json:"vm_private_ip,omitempty"`
	VmPublicIp  string `json:"vm_public_ip,omitempty"`
	VmType      string `json:"vm_type,omitempty"`
}

type BackendRs struct {
	RsId    string `json:"RsId,omitempty"`
	Address string `json:"Address,omitempty"`
	Port    int    `json:"Port,omitempty"`
	Weight  int    `json:"Weight,omitempty"`
	Status  string `json:"Status,omitempty"`
}

type CreateVpcSlbRsPortReq struct {
	ListenId string       `json:"ListenId"`
	RsList   []RsPortItem `json:"RsList"`
}

type RsPortItem struct {
	RsPortId    string `json:"RsPortId,omitempty"`
	VmId        string `json:"VmId,omitempty"`
	VmName      string `json:"VmName,omitempty"`
	VmPublicIp  string `json:"VmPublicIp,omitempty"`
	VmType      string `json:"VmType,omitempty"`
	VmPrivateIp string `json:"VmPrivateIp,omitempty"`
	Port        string `json:"Port,omitempty"`
	Weight      string `json:"Weight,omitempty"`
}

type RsPort struct {
	RsId   string `json:"RsId"`
	Port   int    `json:"Port"`
	Weight int    `json:"Weight,omitempty"`
}

type CreateVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type DeleteVpcSlbRsPortReq struct {
	RsPortIds []string `json:"RsPortIds"`
}

type DeleteVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type UpdateVpcSlbRsPortReq struct {
	RsPortList []RsPortItem `json:"RsPortList"`
}

type UpdateVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type GetVpcSlbRsPortReq struct {
	Keyword  string `json:"Keyword"`
	ListenId string `json:"ListenId,omitempty"`
}

type QueryVpcSlbRsPortResult struct {
	Code    string                       `json:"Code"`
	Message string                       `json:"Message"`
	Data    QueryVpcSlbRsPortResultEntry `json:"Data,omitempty"`
}

type QueryVpcSlbRsPortResultEntry struct {
	RsPortList []RsPortEntry `json:"RsPortList,omitempty"`
	Total      int           `json:"Total"`
}

type RsPortEntry struct {
	CheckStatus    string `json:"CheckStatus,omitempty"`
	CheckStatusStr string `json:"CheckStatusStr,omitempty"`
	LanIp          string `json:"LanIp,omitempty"`
	Port           string `json:"Port,omitempty"`
	ResourceId     string `json:"ResourceId,omitempty"`
	ResourceName   string `json:"ResourceName,omitempty"`
	RsPortId       string `json:"RsPortId,omitempty"`
	Status         string `json:"Status,omitempty"`
	StatusZh       string `json:"StatusZh,omitempty"`
	VmType         string `json:"VmType,omitempty"`
	WanIp          string `json:"WanIp,omitempty"`
	Weight         string `json:"Weight,omitempty"`
}

type BandwidthBindResourceReq struct {
	BandwidthId string `json:"BandwidthId"`
	ResourceId  string `json:"ResourceId"`
	BindType    string `json:"BindType"`
}

type BandwidthBindResourceResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}

type BandwidthUnbindResourceReq struct {
	BandwidthId string `json:"BandwidthId"`
}

type BandwidthUnbindResourceResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId,omitempty"`
}
