## 创建集群
通过以下代码创建集群
```go
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
    // 集群主节点的数量,当前支持3,5,7
    MasterNumber: eks.MasterNumber3,
    // 集群主节点的配置信息
    MasterConfig: eks.NodePoolNodeConfig{
        // 主节点规格
        InstanceTypeIds: []string{eks.EcsCpuC11Compute2XLarge},
        // 主节点付费方式
        BillingSpec: eks.NodePoolBillingSpec{
            // 包年包月时，是否自动续费，0:否，1:是。按需时无需传入
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
        // 节点初始化完成后向master节点设置标签
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
            // 节点初始化完成后向worker节点标签
            Labels: map[string]string{},
        },
    },
}
response, err := eksClient.CreateCluster(&req)
if err != nil {
    fmt.Println(err)
}
bytes, _ := json.Marshal(response)
fmt.Println(string(bytes))
```
## 获取集群详细信息
通过以下示例代码，获取指定集群的详细信息。
```go
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
```

## 查询集群列表



使用以下代码可以根据指定的过滤条件查询相关集群

```go
req := eks.ListClustersReq{
  // 集群所使用的vpc的ID
  VpcId:   "d4ceb516-6def-11f0-a509-5672334e2706",
  // 集群的名称(模糊搜素)或者ID
  Keyword: "test-cluster",
}

response, err := eksClient.ListClusters(&req)
if err != nil {
  fmt.Println(err)

}
for _, cluster := range response.Data {
  // eks集群ID
  fmt.Println("cluster id:", cluster.ClusterId)
  // 集群的入网公网IP
  fmt.Println("cluster ip:", cluster.ClusterIp)
  // 集群的名称
  fmt.Println("cluster name:", cluster.ClusterName)
  // 集群的状态，running(运行中)/error(错误)
  fmt.Println("cluster status:", cluster.ClusterStatus)
  // 集群的创建时间，格式：2006-01-02 15:04:05
  fmt.Println("cluster create time:",cluster.CreateTime)
  // 集群的k8s版本,v1.26.15/v1.30.8/v1.30.14
  fmt.Println("cluster version:",cluster.K8SVersion)
  // 集群中的节点数量
  fmt.Println("cluster node sum:",cluster.NodeSum)
  // 集群所在的区域的ID
  fmt.Println("region id:",cluster.RegionId)
  // 集群所在的区域的名称
  fmt.Println("region name:",cluster.RegionName)
  // 搭建集群所使用的slb的ID(通过slb搭建的集群才有)
  fmt.Println("cluster slb id:",cluster.SlbId)
  // 集群的ssh公网入口端口
  fmt.Println("cluster ssh port:",cluster.SshPort)
  // 集群的域名
  fmt.Println("cluster sub domain:",cluster.SubDomain)
  // 集群的内网虚拟IP(通过nat网关搭建的集群才有)
  fmt.Println("cluster vip:",cluster.Vip)
  // 集群所使用的vpc的ID
  fmt.Println("cluster vpc id:",cluster.VpcId)
  // 集群所使用的vpc的名称
  fmt.Println("cluster vpc name:",cluster.VpcName)
}
```



## 删除集群

使用以下代码可以删除指定的集群

```go
// 集群ID
response, err := eksClient.DeleteCluster("4a5da7b9-0c3d-4dc3-96db-4c041c273d05")
if err != nil {
  fmt.Println(err)

}
// 删除集群的任务的ID
fmt.Println(response.Data.TaskId)
// 删除的集群的ID
fmt.Println(response.Data.ClusterId)
```



## 获取集群事件

使用以下代码可以查询指定集群的事件(集群的创建进度)

```go
response, err := eksClient.GetClusterEvents("466603a8-32c4-4bd7-98c5-bb7b6b152e9e")
if err != nil {
  fmt.Println(err)

}
for _, event := range response.Data {
  // 事件的中文名称
  fmt.Println("sub task name:", event.SubtaskName)
  // 事件的类型
  fmt.Println("sub task type:", event.SubtaskType)
  // 事件的状态
  fmt.Println("sub task status:", event.Status)
  // 事件的开始时间
  fmt.Println("sub task start at:", event.SubtaskStartTime)
  // 事件的耗时
  fmt.Println("sub task cost:", event.SubtaskElapsedTime)
}
```



## 为cni增加一个网段

使用以下代码可以给集群cni增加一个网段配置

```go
params := eks.AddClusterSubnetReq{
  // 集群ID
  ClusterId: "861deabe-832d-4266-ade5-4aa411104a68",
  SubnetList: []eks.ClusterSubnet{{
    // 需要使用的子网的ID
    SubnetId: "cc2ff5e6-74f2-11f0-bd15-0a6c401afcb2",
    // 需要使用的网段
    Segment: "192.168.0.0/22",
    // 需要使用的ip数量
    UsedIpNum: 256,
  }},
}
response, err := eksClient.AddClusterSubnet(&params)
if err != nil {
  fmt.Println(err)
}

```

