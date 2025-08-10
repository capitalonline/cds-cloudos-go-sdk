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
