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

package main

import (
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

// GetScalingGroupDetail 以下为示例代码，实际开发中请根据需要进行修改和补充
func GetScalingGroupDetail() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	scaleGroupId := "52427fe1-514f-4efe-8528-e4554342dea9"

	response, err := eksClient.GetScalingGroupDetail(scaleGroupId)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	fmt.Println(response.Data.ScalingGroupId)
	fmt.Println(response.Data.StatusStr)
}

func AddScalingGroupNode() {
	ak, sk := "E890861-ak", "E890861-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	addNodeGroupArgs := &eks.AddScalingGroupNodeReq{
		ScalingGroupId: "37b75503-4853-4783-8a73-6080a99c32be",
		AddNum:         1,
	}

	response, err := eksClient.AddScalingGroupNode(addNodeGroupArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	fmt.Println(response.Data.TaskId)
}

func main() {
	//GetScalingGroupDetail()
	AddScalingGroupNode()
}
