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

import (
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

// GetScalingGroupDetail 以下为示例代码，实际开发中请根据需要进行修改和补充
func GetScalingGroupDetail() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	eksClient, _ := eks.NewClient(ak, sk, endpoint)

	scaleGroupId := "52427fe1-514f-4efe-8528-e4554342dea9"

	response, err := eksClient.GetScalingGroupDetail(scaleGroupId)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(response)

}
