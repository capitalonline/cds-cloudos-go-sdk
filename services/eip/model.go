/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package eip

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

type CreateEIPReq struct {
	RegionCode    string `json:"RegionCode"`
	BandwidthType string `json:"BandwidthType"`
	BillScheme    string `json:"BillScheme"`
	Qos           int    `json:"Qos"`
	Size          int    `json:"Size"`
	Description   string `json:"Description"`
	SubjectId     string `json:"SubjectId"`
	ProjectId     int    `json:"ProjectId"`
}

type CreateEipResData struct {
	EIPId string `json:"EIPId"`
	IP    string `json:"IP"`
}

type CreateEIPResult struct {
	Data []CreateEipResData `json:"data,omitempty"`
	OpenApiCommonResp
}

type GetEipReq struct {
	Keyword string `json:"Keyword"`
}

type BandwidthInfo struct {
	AvailableZoneCode string `json:"AvailableZoneCode"`
	BandwidthType     string `json:"BandwidthType"`
	BillScheme        string `json:"BillScheme"`
	CreateTime        string `json:"CreateTime"`
	Id                string `json:"Id"`
	Name              string `json:"Name"`
	Qos               int    `json:"Qos"`
	Status            string `json:"Status"`
}

type BindResourceInfo struct {
	SubnetIP string `json:"SubnetIP"`
	SubnetId string `json:"SubnetId"`
}

type Eip struct {
	AvailableZoneCode string           `json:"AvailableZoneCode"`
	CreateTime        string           `json:"CreateTime"`
	Description       string           `json:"Description"`
	IP                string           `json:"IP"`
	Id                string           `json:"Id"`
	IsBind            bool             `json:"IsBind"`
	RegionCode        string           `json:"RegionCode"`
	Status            string           `json:"Status"`
	BandwidthInfo     BandwidthInfo    `json:"BandwidthInfo"`
	BindResourceInfo  BindResourceInfo `json:"BindResourceInfo"`
}

type GetEipData struct {
	Total   int   `json:"Total"`
	EIPList []Eip `json:"EIPList"`
}

type GetEipResult struct {
	Data GetEipData `json:"data,omitempty"`
	OpenApiCommonResp
}

type ListEipsReq struct {
	Keyword           string `json:"Keyword,omitempty"`
	RegionCode        string `json:"RegionCode"`
	AvailableZoneCode string `json:"AvailableZoneCode,omitempty"`
}

type ListEipsData struct {
	Total   int   `json:"Total"`
	EIPList []Eip `json:"EIPList"`
}

type ListEipsResult struct {
	Data ListEipsData `json:"data,omitempty"`
	OpenApiCommonResp
}

type ReleaseEipReq struct {
	EIPId string `json:"EIPId"`
}

type ReleaseEipResult struct {
	Data map[string]interface{} `json:"Data"`
	OpenApiCommonResp
}

type UpdateEIPReq struct {
	EIPId       string `json:"EIPId"`
	Qos         int    `json:"Qos,omitempty"`
	Description string `json:"Description,omitempty"`
}

type UpdateEIPResult struct {
	Data map[string]interface{} `json:"Data"`
	OpenApiCommonResp
	CommonTask
}
