package slb

type OpenapiCommonPage struct {
	Total int `json:"Total"`
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
	BandwidthId string   `json:"BandwidthId"`
	EipList     []string `json:"EipList"`
	LanVipList  []string `json:"LanVipList"`
}

type ListVpcSlbResult struct {
	// Data []Slb  `json:Data`
	Data []map[string]interface{} `json:"Data"`
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
	VipList       []map[string]interface{} `json:"VipList"`
}

type GetVpcSlbResult struct {
	Data SlbResultData `json:"data"`
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
	TaskId  string           `json:"TaskId"`
	Data    CreateVpcSlbData `json:"Data"`
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
	TaskId  string `json:"TaskId"`
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
	TaskId  string `json:"TaskId"`
}

type ListVpcSlbDetailReq struct {
	VpcId   string `json:"VpcId"`
	Keyword string `json:"Keyword"`
}

type DescribeVpcSlbListResult struct {
	Code    string                 `json:"Code"`
	Message string                 `json:"Message"`
	Data    DescribeVpcSlbListData `json:"Data"`
}

type DescribeVpcSlbListData struct {
	SlbList []DescribeVpcSlbEntry `json:"SlbList"`
	//Total   int                   `json:"Total"`
}

type DescribeVpcSlbEntry struct {
	AzId          string `json:"AzId"`
	AzName        string `json:"AzName"`
	BandwidthId   string `json:"BandwidthId"`
	BandwidthName string `json:"BandwidthName"`
	BillingMethod string `json:"BillingMethod"`
	ConfName      string `json:"ConfName"`
	CreateTime    string `json:"CreateTime"`
	Id            string `json:"Id"`
	Name          string `json:"Name"`
	NetType       string `json:"NetType"`
	//NetTypeZh     string     `json:"NetTypeZh"`
	RegionId   string     `json:"RegionId"`
	RegionName string     `json:"RegionName"`
	Status     string     `json:"Status"`
	StatusZh   string     `json:"StatusZh"`
	VipList    []VipEntry `json:"VipList"`
	VpcId      string     `json:"VpcId"`
	VpcName    string     `json:"VpcName"`
	//VpcSegments   string     `json:"VpcSegments"`
}

type VipEntry struct {
	Address string `json:"Address"`
	Type    string `json:"Type"`
	Vip     string `json:"Vip"`
	VipId   string `json:"VipId"`
	VipType string `json:"VipType"`
}

type GetVpcSlbReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbDetailInfoResult struct {
	Code    string                       `json:"Code"`
	Message string                       `json:"Message"`
	Data    DescribeVpcSlbDetailInfoData `json:"Data"`
}

type DescribeVpcSlbDetailInfoData struct {
	AzId              string        `json:"AzId"`
	AzName            string        `json:"AzName"`
	BandwidthInfo     BandwidthInfo `json:"BandwidthInfo"`
	BillingMethod     string        `json:"BillingMethod"`
	BillingSchemeId   string        `json:"BillingSchemeId"`
	BillingSchemeName string        `json:"BillingSchemeName"`
	ConfName          string        `json:"ConfName"`
	CreateTime        string        `json:"CreateTime"`
	Id                string        `json:"Id"`
	Name              string        `json:"Name"`
	NetType           string        `json:"NetType"`
	NetTypeZh         string        `json:"NetTypeZh"`
	RegionId          string        `json:"RegionId"`
	RegionName        string        `json:"RegionName"`
	Status            string        `json:"Status"`
	StatusZh          string        `json:"StatusZh"`
	VipList           []VipEntry    `json:"VipList"`
	VpcId             string        `json:"VpcId"`
	VpcName           string        `json:"VpcName"`
	VpcSegments       string        `json:"VpcSegments"`
}

type BandwidthInfo struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
}

type GetVpcSlbListenCreateInfoReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbListenCreateInfoResult struct {
	Code    string                             `json:"Code"`
	Message string                             `json:"Message"`
	Data    DescribeVpcSlbListenCreateInfoData `json:"Data"`
}

type DescribeVpcSlbListenCreateInfoData struct {
	SlbAclList []AclEntry `json:"SlbAclList"`
	SlbVipList []VipEntry `json:"SlbVipList"`
}

type AclEntry struct {
	AclId   string `json:"AclId"`
	AclName string `json:"AclName"`
}

type CreateVpcSlbListenReq struct {
	SlbId           string          `json:"SlbId"`
	ListenName      string          `json:"ListenName"`
	Vip             string          `json:"Vip"`
	VipId           string          `json:"VipId"`
	VipType         string          `json:"VipType"`
	ListenProtocol  string          `json:"ListenProtocol"`
	ListenPort      int             `json:"ListenPort"`
	ListenTimeout   int             `json:"ListenTimeout"`
	AclId           string          `json:"AclId"`
	Scheduler       string          `json:"Scheduler"`
	HealthCheckInfo HealthCheckInfo `json:"HealthCheckInfo"`
}

type HealthCheckInfo struct {
	Protocol         string `json:"Protocol"`
	Port             int    `json:"Port"`
	ConnectTimeout   int    `json:"ConnectTimeout,omitempty"`
	Virtualhost      string `json:"Virtualhost"`
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
	Data    CreateVpcSlbListenData `json:"Data"`
}

type CreateVpcSlbListenData struct {
	ListenId string `json:"ListenId"`
}

type DeleteVpcSlbListenReq struct {
	ListenIds []string `json:"ListenIds"`
}

type DeleteVpcSlbListenResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}

type UpdateVpcSlbListenReq struct {
	ListenId        string          `json:"ListenId"`
	ListenName      string          `json:"ListenName"`
	AclId           string          `json:"AclId"`
	ListenTimeout   int             `json:"ListenTimeout"`
	Scheduler       string          `json:"Scheduler"`
	Virtualhost     string          `json:"Virtualhost"`
	HealthCheckInfo HealthCheckInfo `json:"HealthCheckInfo"`
}

type UpdateVpcSlbListenResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}

type DescribeVpcSlbListenListReq struct {
	SlbId string `json:"SlbId"`
}

type DescribeVpcSlbListenListResult struct {
	Code    string                       `json:"Code"`
	Message string                       `json:"Message"`
	Data    DescribeVpcSlbListenListData `json:"Data"`
}

type DescribeVpcSlbListenListData struct {
	ListenList []DescribeVpcListenEntry `json:"ListenList"`
	Total      int                      `json:"Total"`
}

type DescribeVpcListenEntry struct {
	ListenId       string        `json:"ListenId"`
	ListenName     string        `json:"ListenName"`
	ListenPort     string        `json:"ListenPort"`
	ListenProtocol string        `json:"ListenProtocol"`
	ListenVip      string        `json:"ListenVip"`
	Scheduler      string        `json:"Scheduler"`
	SchedulerCn    string        `json:"SchedulerCn"`
	Status         string        `json:"Status"`
	StatusCn       string        `json:"StatusCn"`
	CheckInfo      CheckInfoData `json:"CheckInfo"`
}

type CheckInfoData struct {
	Error int `json:"Error"`
	Ok    int `json:"Ok"`
}

type GetVpcSlbListenReq struct {
	SlbId    string `json:"SlbId"`
	ListenId string `json:"ListenId"`
}

type QueryVpcSlbListenResult struct {
	Code    string           `json:"Code"`
	Message string           `json:"Message"`
	Data    QueryListenEntry `json:"Data"`
}

type QueryListenEntry struct {
	ListenId        string          `json:"ListenId"`
	ListenName      string          `json:"ListenName"`
	ListenProtocol  string          `json:"ListenProtocol"`
	ListenVip       string          `json:"ListenVip"`
	Vip             string          `json:"Vip"`
	VipId           string          `json:"VipId"`
	VipType         string          `json:"VipType"`
	ListenPort      int             `json:"ListenPort"`
	ListenTimeout   int             `json:"ListenTimeout"`
	AclId           string          `json:"AclId"`
	Scheduler       string          `json:"Scheduler"`
	SchedulerCn     string          `json:"SchedulerCn"`
	Status          string          `json:"Status"`
	StatusCn        string          `json:"StatusCn"`
	HealthCheckInfo HealthCheckInfo `json:"HealthCheckInfo"`
}

type GetVpcSlbListenRsInfoReq struct {
	VmType string `json:"VmType"`
	VpcId  string `json:"VpcId"`
}

type DescribeVpcSlbListenRsInfoResult struct {
	Code    string                           `json:"Code"`
	Message string                           `json:"Message"`
	Data    []DescribeVpcSlbListenRsInfoData `json:"Data"`
}

type DescribeVpcSlbListenRsInfoData struct {
	VmId        string `json:"vm_id"`
	VmName      string `json:"vm_name"`
	VmPrivateIp string `json:"vm_private_ip"`
	VmPublicIp  string `json:"vm_public_ip"`
	VmType      string `json:"vm_type"`
}

type BackendRs struct {
	RsId    string `json:"RsId"`
	Address string `json:"Address"`
	Port    int    `json:"Port"`
	Weight  int    `json:"Weight"`
	Status  string `json:"Status"`
}

type CreateVpcSlbRsPortReq struct {
	ListenId string       `json:"ListenId"`
	RsList   []RsPortItem `json:"RsList"`
}

type RsPortItem struct {
	RsPortId    string `json:"RsPortId"`
	VmId        string `json:"VmId"`
	VmName      string `json:"VmName"`
	VmPublicIp  string `json:"VmPublicIp"`
	VmType      string `json:"VmType"`
	VmPrivateIp string `json:"VmPrivateIp"`
	Port        string `json:"Port"`
	Weight      string `json:"Weight"`
}

type RsPort struct {
	RsId   string `json:"RsId"`
	Port   int    `json:"Port"`
	Weight int    `json:"Weight"`
}

type CreateVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}

type DeleteVpcSlbRsPortReq struct {
	RsPortIds []string `json:"RsPortIds"`
}

type DeleteVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}

type UpdateVpcSlbRsPortReq struct {
	RsPortList []RsPortItem `json:"RsPortList"`
}

type UpdateVpcSlbRsPortResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}

type GetVpcSlbRsPortReq struct {
	Keyword  string `json:"Keyword"`
	ListenId string `json:"ListenId"`
}

type QueryVpcSlbRsPortResult struct {
	Code    string                       `json:"Code"`
	Message string                       `json:"Message"`
	Data    QueryVpcSlbRsPortResultEntry `json:"Data"`
}

type QueryVpcSlbRsPortResultEntry struct {
	RsPortList []RsPortEntry `json:"RsPortList"`
	Total      int           `json:"Total"`
}

type RsPortEntry struct {
	CheckStatus    string `json:"CheckStatus"`
	CheckStatusStr string `json:"CheckStatusStr"`
	LanIp          string `json:"LanIp"`
	Port           string `json:"Port"`
	ResourceId     string `json:"ResourceId"`
	ResourceName   string `json:"ResourceName"`
	RsPortId       string `json:"RsPortId"`
	Status         string `json:"Status"`
	StatusZh       string `json:"StatusZh"`
	VmType         string `json:"VmType"`
	WanIp          string `json:"WanIp"`
	Weight         string `json:"Weight"`
}

type BandwidthBindResourceReq struct {
	BandwidthId string `json:"BandwidthId"`
	ResourceId  string `json:"ResourceId"`
	BindType    string `json:"BindType"`
}

type BandwidthBindResourceResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}

type BandwidthUnbindResourceReq struct {
	BandwidthId string `json:"BandwidthId"`
}

type BandwidthUnbindResourceResult struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
	TaskId  string `json:"TaskId"`
}
