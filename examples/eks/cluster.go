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
		// 集群名称，长度为1-26个字符，只能包含数字、字母和"-"，且首尾只能是字母或数字
		ClusterName: "cluster_name",
		// 集群所使用的vpc的ID
		VpcId: "VpcId",
		// 集群所使用的cni相关参数
		Cni: eks.CniInfo{
			// 集群所使用的cni类型, 目前支持vpc-cni和flannel
			//以下为vpc-cni示例
			CniType: eks.CniTypeVpcCni,
			// 集群所使用的cni相关配置
			CniConfig: eks.CniConfig{
				// 每个node节点所能创建的最大pod数，默认64。
				NodePodsNum: 64,
				// 集群所使用的子网ID列表，使用vpc-cni时必填。使用flannel非必填。
				SubnetIds: []string{"subnet-01", "subnet-02"},
			},
			//以下为flannel示例
			// CniType: eks.CniTypeFlannel,
			// 集群所使用的cni相关配置
			//CniConfig: eks.CniConfig{
			//	// 集群所使用的pod网段,使用vpc-cni时，非必填。使用flannel是必填
			//	PodCidr: "172.16.0.0/16",
			//	// 每个node节点所能创建的最大pod数，默认64。
			//	NodePodsNum: 64,
			//},
			// 代理模式,目前支持iptables和ipvs
			ProxyConfig: eks.ProxyConfigIptables,
			// 集群service网段
			ServiceCidr: eks.ServiceCidr192_16,
		},
		// k8s版本,目前支持1.30.14和1.26.5
		K8sVersion: eks.K8sVersion1_30_14,
		// 集群所使用的出网关ID
		NatId: "NatId",
		// 集群出网EIP的IP地址
		SourceEip: "SourceEip",
		// 集群入网EIP的IP地址
		DestinationEip: "DestinationEip",
		// 集群入网所使用的负载均衡ID
		SlbId: "SlbId",
		// 集群所使用的ak和sk
		Ak: "Ak",
		Sk: "Sk",
		// 创建集群使用的密码
		Password: "password",
		// 集群主节点的数量
		MasterNumber: eks.MasterNumber3,
		// 集群主节点的配置信息
		MasterConfig: eks.NodePoolNodeConfig{
			// 主节点规格
			InstanceTypeIds: []string{eks.EcsCpuC11Compute2XLarge},
			// 主节点付费方式
			BillingSpec: eks.NodePoolBillingSpec{
				// 包年包月时，是否自动续费，0:否，1:是。按需是无需传入
				AutoRenew: 0,
				// 计费类型，按需和包年包月
				BillingMethod: eks.NodePoolBillingMethodPostPaid,
				// 包年包月计费时长，当计费类型为包年包月时必填。单位为月，取值范围1-12,24,36。
				Duration: 0,
				// 包年包月时，是否购买至月底，0:否，1:是
				IsToMonth: 0,
			},
			// 主节点数据盘信息
			DataDisk: []eks.NodePoolDiskInfo{
				{
					// 数据盘类型,当前只支持ssd
					DiskType: eks.NodePoolDiskTypeSSD,
					// 数据盘大小,初始值80，步长80
					DiskSize: 80,
				},
			},
			// 集群所使用的镜像名称
			OsImageName: eks.EcsUbuntu2204K8s13014Cpu,
			// 集群所使用的子网ID,支持多选，必须同一vpc下的子网
			SubnetIds: []string{"subnet-01", "subnet-02"},
			// 节点初始化完成后执行脚本命令
			Shell: "",
			// 节点初始化完成后的标签
			Labels: map[string]string{},
		},
		// 节点池配置信息
		NodePoolConfig: eks.NodePoolConfiguration{
			// 节点池名称,长度为1-26个字符，只能包含数字、字母和"-"，且首尾只能是字母或数字
			PoolName: "node-pool-1",
			// 节点池类型,目前支持ecs和bms
			NodeType: eks.NodePoolNodeTypeECS,
			// 节点池使用的测试金ID
			SubjectId: 0,
			// 节点池所使用的节点配置信息
			NodeConfig: eks.NodePoolNodeConfig{
				// 节点规格
				InstanceTypeIds: []string{eks.EcsCpuC11Compute2XLarge},
				// 付费方式
				BillingSpec: eks.NodePoolBillingSpec{
					// 包年包月时，是否自动续费，0:否，1:是。按需是无需传入
					AutoRenew: 0,
					// 计费类型，按需和包年包月
					BillingMethod: eks.NodePoolBillingMethodPostPaid,
					// 包年包月计费时长，当计费类型为包年包月时必填。单位为月，取值范围1-12,24,36。
					Duration: 0,
					// 包年包月时，是否购买至月底，0:否，1:是
					IsToMonth: 0,
				},
				// 节点数据盘信息
				DataDisk: []eks.NodePoolDiskInfo{
					{
						// 数据盘类型,当前只支持ssd
						DiskType: eks.NodePoolDiskTypeSSD,
						// 数据盘大小,初始值80，步长80
						DiskSize: 80,
					},
				},
				// 节点所使用的镜像名称
				OsImageName: eks.EcsUbuntu2204K8s13014Cpu,
				// 节点所使用的子网ID,支持多选，必须同一vpc下的子网
				SubnetIds: []string{"subnet-01", "subnet-02"},
				// 节点初始化完成后执行脚本命令
				Shell: "",
				// 节点初始化完成后的标签
				Labels: map[string]string{},
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

func GetTaskStatus() {
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
