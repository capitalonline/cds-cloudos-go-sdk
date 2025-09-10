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

