## 查询节点列表

通过以下代码可以查询集群中的节点:

```go
params := eks.ListNodesReq{
  // eks集群的ID
  ClusterId:      "17463875-10d7-4e6b-b5e6-bfb29b79477f",
  // 机器状态:running(运行中)/error(错误)
  NodeStatus:     "Ready",
  // 需要查询的节点类型，分为:worker(工作节点)、master(控制平面节点)
  NodeType:       "worker",
  // 节点是否可调度:1(是)/0(否)
  SchedulableStr: "1",
  // k8s集群中节点的状态,可选:Ready(就绪)、NotReady(未就绪)
  Status:         "running",
  // 分页的大小，不传后端默认返回10条数据
  PageSize:       10,
  // 需要查询的页面号，不传默认返回第一页数据
  PageIndex:      1,
  // 查询的关键词(节点名或者节点ID，模糊搜索)
  Keyword:        "test-cluster",
}
response, err := eksClient.ListNodes(&params)
if err != nil {
  fmt.Println(err)

}
for _, node := range response.Data {
  // 节点的ID
  fmt.Println(node.NodeId)
  // 节点的名称
  fmt.Println(node.NodeName)
  // 机器的状态
  fmt.Println(node.NodeStatus)
  // 节点的类型，分为:worker(工作节点)、master(控制平面节点)
  fmt.Println(node.NodeType)
  // 节点所属集群的ID
  fmt.Println(node.ClusterId)
  // 节点所属集群的名称
  fmt.Println(node.ClusterName)
  // 所属主账号ID
  fmt.Println(node.CustomerId)
  // 所属子账号ID
  fmt.Println(node.UserId)
  // 节点所属的vpc的ID
  fmt.Println(node.VpcId)
  // 节点所属的子网的ID
  fmt.Println(node.SubnetId)
  // 节点的私网IP
  fmt.Println(node.PrivateIp)
  // 节点的cpu核数
  fmt.Println(node.Cpu)
  // 节点的内存,单位G
  fmt.Println(node.Ram)
  // 节点的gpu卡数量
  fmt.Println(node.Gpu)
  // 节点的gpu卡类型
  fmt.Println(node.GpuType)
  // 节点的规格配置名称
  fmt.Println(node.FamilyName)
  // 节点是否可调度，0(不可调度)/1(可调度)
  fmt.Println(node.Schedulable)
  // 节点在k8s集群中的状态，Ready(就绪)/NotReady(未就绪)
  fmt.Println(node.K8SStatus)
  // 节点所属地域名称
  fmt.Println(node.RegionName)
  // 节点所属地域的ID
  fmt.Println(node.RegionId)
  // 节点所属的可用区名称
  fmt.Println(node.AzName)
  // 节点所属的可用区的ID
  fmt.Println(node.AzId)
  // 节点来源，bms(裸金属)/ecs(GPU云主机)/other(第三方或外部节点)
  fmt.Println(node.SourceType)
  // 计费方式，0(按需付费)/1(包年包月)
  fmt.Println(node.BillingMethod)
  // 包年包月的预付费时长(单位月)
  fmt.Println(node.Duration)
  // 是否包月到月底:1(是)/0(否)
  fmt.Println(node.IsToMonth)
  // 是否自动续约:1(是)/0(否)
  fmt.Println(node.AutoRenew)
  // 节点的创建时间，格式 2006-01-02 15:04:05
  fmt.Println(node.CreateTime)
}
```



## 删除节点

通过以下代码可以删除集群中的指定节点

```go
params := eks.DeleteNodesReq{
  // eks集群的ID
  ClusterId: "5ff38515-e942-437d-8e16-60f12c382e08",
  // 需要删除的节点的列表
  NodeIds:   []string{"eks-4vdp8eu5nq3zppfg"},
}
response, err := eksClient.DeleteNodes(&params)
if err != nil {
  fmt.Println(err)

}
// 删除集群的任务ID
fmt.Println(response.Data.TaskId)

```