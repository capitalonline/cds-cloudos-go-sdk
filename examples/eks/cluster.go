package main

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func CreateCluster() {
	ak, sk := "ak", "sk"

	eksClient, _ := eks.NewClient(ak, sk)
	req := eks.CreateClusterReq{
		ClusterName: "cluster_name",
		VpcId:       "VpcId",
		Cni: eks.CniInfo{
			CniType: eks.CniTypeVpcCni,
			CniConfig: eks.CniConfig{
				PodCidr:            "10.244.0.0/16",
				FlannelBackendType: eks.FlannelBackendTypeVxlan,
				NodePodsNum:        eks.NodePodsNum64,
				SubnetList: []eks.SubnetInfo{
					{
						SubnetId: "SubnetId",
					},
				},
			},
			ProxyConfig: eks.ProxyConfigIptables,
			ServiceCidr: eks.ServiceCidr192_16,
		},
		K8sVersion:       eks.K8sVersion1_30_14,
		NatId:            "NatId",
		SourceEipId:      "SourceEipId",
		DestinationEipId: "DestinationEipId",
		SlbId:            "SlbId",
		Ak:               "Ak",
		Sk:               "Sk",
		MasterNumber:     eks.MasterNumber3,
		MasterConfig: eks.NodeConfig{
			BillingSpec: eks.BillingSpec{
				AutoRenew:     eks.AutoRenew0,
				BillingMethod: eks.BillingMethod0,
				Duration:      eks.Duration1,
				IsToMonth:     eks.IsToMonth0,
			},
			SystemDisk: eks.DiskInfo{
				DiskFeature: eks.DiskFeatureSsd,
				DiskSize:    40,
			},
			DataDisk: []eks.DiskInfo{
				{
					DiskFeature: eks.DiskFeatureSsd,
					DiskSize:    40,
					Number:      1,
				},
			},
			OsImageName: eks.OsImageNameUbuntu2204_1_30_14,
			SubnetList: []eks.SubnetInfo{
				{
					SubnetId: "SubnetId",
				},
			},
			Shell: "",
			Specifics: []eks.Specifics{
				{
					ProductName: eks.ProductName1,
				},
			},
			Password:    "password",
			Labels:      map[string]string{},
			Annotations: map[string]string{},
		},
		NodePoolConfig: eks.NodePoolConfig{
			PoolName:  "node-pool-1",
			NodeType:  eks.NodeTypeEcs,
			SubjectId: 0,
			NodeConfig: eks.NodeConfig{
				BillingSpec: eks.BillingSpec{
					AutoRenew:     eks.AutoRenew0,
					BillingMethod: eks.BillingMethod0,
					Duration:      1,
					IsToMonth:     eks.IsToMonth1,
				},
				SystemDisk: eks.DiskInfo{
					DiskFeature: eks.DiskFeatureSsd,
					DiskSize:    40,
				},
				DataDisk: []eks.DiskInfo{
					{
						DiskFeature: eks.DiskFeatureSsd,
						DiskSize:    80,
						Number:      1,
					},
				},
				OsImageName: eks.OsImageNameUbuntu2204_1_30_14,
				SubnetList: []eks.SubnetInfo{
					{
						SubnetId: "SubnetId",
					},
				},
				Specifics: []eks.Specifics{
					{
						ProductName: eks.ProductName1,
					},
				},
				Shell:       "",
				Password:    "password",
				Labels:      map[string]string{},
				Annotations: map[string]string{},
			},
		},
	}
	response, err := eksClient.CreateCluster(&req)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func ListClusters() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := eks.ListClustersReq{
		VpcId: "d4ceb516-6def-11f0-a509-5672334e2706",
	}

	response, err := eksClient.ListClusters(&req)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetClusterEvents() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.GetClusterEvents("466603a8-32c4-4bd7-98c5-bb7b6b152e9e")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteCluster() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.DeleteCluster("4a5da7b9-0c3d-4dc3-96db-4c041c273d05")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetCluster() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.GetCluster("your-cluster-id")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}
