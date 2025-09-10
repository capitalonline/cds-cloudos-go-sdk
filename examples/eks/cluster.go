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
				NodePodsNum:        64,
				SubnetIds:          []string{"subnet-01", "subnet-02"},
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
		MasterConfig: eks.NodePoolNodeConfig{
			InstanceTypeIds: []string{eks.EcsCpuC11Compute2XLarge},
			BillingSpec: eks.NodePoolBillingSpec{
				AutoRenew:     0,
				BillingMethod: eks.NodePoolBillingMethodPostPaid,
				Duration:      0,
				IsToMonth:     0,
			},
			SystemDisk: eks.NodePoolDiskInfo{
				DiskType: eks.NodePoolDiskTypeSSD,
				DiskSize: 40,
			},
			DataDisk: []eks.NodePoolDiskInfo{
				{
					DiskType: eks.NodePoolDiskTypeSSD,
					DiskSize: 80,
				},
			},
			OsImageName: eks.EcsUbuntu2204K8s13014Cpu,
			SubnetIds:   []string{"subnet-01", "subnet-02"},
			Shell:       "",
			Password:    "password",
			Labels:      map[string]string{},
		},
		NodePoolConfig: eks.NodePoolConfiguration{

			PoolName:  "node-pool-1",
			NodeType:  eks.NodePoolNodeTypeECS,
			SubjectId: 0,
			NodeConfig: eks.NodePoolNodeConfig{
				InstanceTypeIds: []string{eks.EcsCpuC11Compute2XLarge},
				BillingSpec: eks.NodePoolBillingSpec{
					AutoRenew:     0,
					BillingMethod: eks.NodePoolBillingMethodPostPaid,
					Duration:      0,
					IsToMonth:     0,
				},
				SystemDisk: eks.NodePoolDiskInfo{
					DiskType: eks.NodePoolDiskTypeSSD,
					DiskSize: 40,
				},
				DataDisk: []eks.NodePoolDiskInfo{
					{
						DiskType: eks.NodePoolDiskTypeSSD,
						DiskSize: 80,
					},
				},
				OsImageName: eks.EcsUbuntu2204K8s13014Cpu,
				SubnetIds:   []string{"subnet-01", "subnet-02"},

				Shell:    "",
				Password: "password",
				Labels:   map[string]string{},
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

func TaskStatus() {
	ak, sk := "E101277-ak", "E101277-sk"

	eksClient, _ := eks.NewClient(ak, sk)

	response, err := eksClient.GetTaskStatus("task_id")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func ListClusters() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)

	req := eks.ListClustersReq{
		// 集群所使用的vpc的ID
		VpcId: "d4ceb516-6def-11f0-a509-5672334e2706",
		// 集群的名称(模糊搜素)或者ID
		Keyword: "test-cluster",
	}

	response, err := eksClient.ListClusters(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func GetClusterEvents() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)
	// 传集群的ID
	response, err := eksClient.GetClusterEvents("b9ee61c2-87fa-4035-9ca5-a580980e3dde")
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteCluster() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)
	// 传集群的ID
	response, err := eksClient.DeleteCluster("4a5da7b9-0c3d-4dc3-96db-4c041c273d05")
	if err != nil {
		fmt.Println(err)
		return
	}

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

func AddClusterSubnet() {
	ak, sk := "your-ak", "your-sk"

	eksClient, _ := eks.NewClient(ak, sk)
	params := eks.AddClusterSubnetReq{
		// 集群ID
		ClusterId: "861deabe-832d-4266-ade5-4aa411104a68",
		SubnetList: []eks.ClusterSubnet{{
			// 需要使用的子网的ID
			SubnetId: "cc2ff5e6-74f2-11f0-bd15-0a6c401afcb2",
		}},
	}
	response, err := eksClient.AddClusterSubnet(&params)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}
