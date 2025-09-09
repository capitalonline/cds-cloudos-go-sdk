package main

import (
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/eip"
)

// ListEips 查询eip列表
func ListEips() {
	ak, sk := "ak", "sk"
	EipClient, _ := eip.NewClient(ak, sk)
	ListEipArgs := &eip.ListEipsReq{
		RegionCode:"CN_Qingyang",
	}
	response, err := EipClient.ListEips(ListEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
// GetEip 获取某个eip信息
func GetEip() {
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	GetEipArgs := &eip.GetEipReq{
		Keyword: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",  // eip地址或者id
	}
	response, err := EipClient.GetEip(GetEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
	fmt.Println(response.Data.Total)
}
// CreateEip 创建eip
func CreateEip(){
	ak, sk := "ak", "sk"
	EipClient, _ := eip.NewClient(ak, sk)
	CreateEipArgs := &eip.CreateEIPReq{
		RegionCode: "CN_Qingyang",  // 区域code
        BandwidthType: "Bandwidth_Multi_ISP_BGP",  // 带宽类型
        BillScheme: "BandwIdth",  // 计费方案
        Qos: 5,  // 带宽大小
        Size: 1,  // 数量
        Description: "go create",
	}
	response, err := EipClient.CreateEip(CreateEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}
// UpdateEip 更新eip带宽或者描述
func UpdateEip(){
	ak, sk := "ak", "sk"

	EipClient, _ := eip.NewClient(ak, sk)
	UpdateEipArgs := &eip.UpdateEIPReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
        Qos: 10,
		Description: "go create",
	}

	response, err := EipClient.UpdateEip(UpdateEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

// ReleaseEip 删除弹性eip
func ReleaseEip(){
	ak, sk := "ak", "sk"


	EipClient, _ := eip.NewClient(ak, sk)
	ReleaseEipArgs := &eip.ReleaseEipReq{
		EIPId: "70cf50e2-79a3-11f0-9be8-6e18e986f14e",
	}

	response, err := EipClient.ReleaseEip(ReleaseEipArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

func main() {
	// ListEips()
	// GetEip()
	// CreateEip()
	// UpdateEip()
	ReleaseEip()
}