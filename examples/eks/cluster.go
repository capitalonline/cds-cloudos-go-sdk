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

func AddClusterSubnet() {
	ak, sk := "a37e2b425ee411ef97433293ad7453f8", "57324ec8f5eb3f44750a66da2baed989"

	eksClient, _ := eks.NewClient(ak, sk)
	params := eks.AddClusterSubnetReq{
		ClusterId: "132e7235-5574-4a87-8b45-da73634a6617",
		SubnetList: []eks.ClusterSubnet{{
			SubnetId: "9ab4444a-8df7-11f0-b443-2eb410f96a95",
		}},
	}
	response, err := eksClient.AddClusterSubnet(&params)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf(">>> response: %+v", response)

	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func main() {
	AddClusterSubnet()
}
