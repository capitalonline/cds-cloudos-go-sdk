package main

import (
	"fmt"

	"github.com/capitalonline/cds-cloudos-go-sdk/services/subnet"
)

//ListSubnet 查询子网
func ListSubnet() {
	ak, sk :=  "ak", "sk"

	subnetClient, _ := subnet.NewClient(ak, sk)
	ListSubnetArgs := &subnet.ListSubnetsReq{
		RegionCode: "CN_Qingyang",  // 区域code
	}
	response, err := subnetClient.ListSubnets(ListSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

// GetSubnet 获取子网数据
func GetSubnet() {
	ak, sk :=  "ak", "sk"
	subnetClient, _ := subnet.NewClient(ak, sk)
	GetSubnetArgs := &subnet.GetSubnetReq{
		Keyword: "919a5290-799e-11f0-adfa-6e18e986f14e",  // 子网id或者名称
	}
	response, err := subnetClient.GetSubnet(GetSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}

// CreateSubnet 创建子网
func CreateSubnet() {
	ak, sk :=  "ak", "sk"
	subnetClient, _ := subnet.NewClient(ak, sk)
	CreateSubnetArgs := &subnet.CreateSubnetReq{
        VPCId:        "9197340c-799e-11f0-adfa-6e18e986f14e",  // 私有网络VPC ID
        SubnetList: []subnet.CreateSubnetData{
            {
                AvailableZoneCode: "CN_Qingyang_A",  // 可用区code
                SubnetName:        "子网2",
                SubnetSegment:     "10.21.2.0/24", // 子网网段
            },
        },
    }
	response, err := subnetClient.CreateSubnet(CreateSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)

}

// DeleteSubnet 删除子网
func DeleteSubnet() {
	ak, sk :=  "ak", "sk"

	subnetClient, _ := subnet.NewClient(ak, sk)
	DeleteSubnetArgs := &subnet.DeleteSubnetReq{
		SubnetId: "919a5290-799e-11f0-adfa-6e18e986f14e",  
    }

	response, err := subnetClient.DeleteSubnet(DeleteSubnetArgs)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)
	fmt.Println(response.Data)
}


func main(){
	// ListSubnet()
	// GetSubnet()
//    CreateSubnet()
    DeleteSubnet()
}