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

package eks

type OpenApiCommonResp struct {
	Code      interface{} `json:"Code"`
	Message   interface{} `json:"Msg"`
	RequestId interface{} `json:"RequestId"`
	CommonOpenApiPage
}

type CommonOpenApiPage struct {
	TotalCount int `json:"TotalCount,omitempty"`
	PageIndex  int `json:"PageIndex,omitempty"`
	PageSize   int `json:"PageSize,omitempty"`
}

type CommonScalingGroupTask struct {
	TaskId string `json:"TaskId"`
}

type GetScalingGroupDetailResult struct {
	Data ScalingGroupDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ScalingGroupDetail struct {
	Cpu              int    `json:"Cpu"`
	Ram              int    `json:"Ram"`
	Gpu              int    `json:"Gpu"`
	GpuShowType      string `json:"GpuShowType"`
	MaxSize          int    `json:"MaxSize"`
	MinSize          int    `json:"MinSize"`
	ActivitySize     int    `json:"ActivitySize"`
	PendingSize      int    `json:"PendingSize"`
	RemovingSize     int    `json:"RovingSize"`
	TotalScalingSize int    `json:"TotalScalingSize"`
	ScalingGroupId   string `json:"ScalingGroupId"`
	Status           string `json:"Status"`
	StatusStr        string `json:"StatusStr"`
}

type AddScalingGroupNodeReq struct {
	ScalingGroupId string `json:"ScalingGroupId"`
	AddNum         int    `json:"AddNum"`
}

type AddScalingGroupNodeResult struct {
	Data CommonScalingGroupTask `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ListClustersReq struct {
	Keyword string `json:"Keyword,omitempty"`
	VpcId   string `json:"VpcId,omitempty"`
	Status  string `json:"Status,omitempty"`
	Version string `json:"Version,omitempty"`
}

type ListClustersResult struct {
	Data []ListClustersDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type ListClustersDetail struct {
	ClusterId     string `json:"ClusterId"`
	ClusterIp     string `json:"ClusterIp"`
	ClusterName   string `json:"ClusterName"`
	ClusterStatus string `json:"ClusterStatus"`
	CreateTime    string `json:"CreateTime"`
	K8SVersion    string `json:"K8sVersion"`
	NodeSum       int    `json:"NodeSum"`
	RegionId      string `json:"RegionId"`
	RegionName    string `json:"RegionName"`
	SlbId         string `json:"SlbId"`
	SshPort       int    `json:"SshPort"`
	StatusStr     string `json:"StatusStr"`
	SubDomain     string `json:"SubDomain"`
	UpdateTime    string `json:"UpdateTime"`
	Vip           string `json:"Vip"`
	VpcId         string `json:"VpcId"`
	VpcName       string `json:"VpcName"`
}

type GetClusterEventsResult struct {
	Data []GetClusterEventsDetail `json:"Data,omitempty"`
	OpenApiCommonResp
}

type GetClusterEventsDetail struct {
	ErrorInfo          string `json:"ErrorInfo"`
	Frontend           string `json:"Frontend"`
	Status             string `json:"Status"`
	SubtaskElapsedTime string `json:"SubtaskElapsedTime"`
	SubtaskName        string `json:"SubtaskName"`
	SubtaskStartTime   string `json:"SubtaskStartTime"`
	SubtaskType        string `json:"SubtaskType"`
}

type DeleteClusterResult struct {
	Data DeleteClusterData `json:"Data,omitempty"`
	OpenApiCommonResp
}

type DeleteClusterData struct {
	ClusterId string `json:"ClusterId"`
	TaskId    string `json:"TaskId"`
}
