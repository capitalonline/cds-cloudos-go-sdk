package main

import (
	"encoding/json"
	"fmt"
	"github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
)

func ListNodes() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)

	params := eks.ListNodesReq{
		// eks集群的ID
		ClusterId: "17463875-10d7-4e6b-b5e6-bfb29b79477f",
		// // 机器状态:running(运行中)/error(错误)/shutdown(关机)/deleting(删除中)
		NodeStatus: "running",
		// 需要查询的节点类型，分为:worker(工作节点)、master(控制平面节点)
		NodeType: "worker",
		// 节点是否可调度:1(是)/0(否)
		SchedulableStr: "1",
		// k8s集群中节点的状态,可选:Ready(就绪)、NotReady(未就绪)
		Status: "Ready",
		// 分页的大小，不传后端默认返回10条数据
		PageSize: 10,
		// 需要查询的页面号，不传默认返回第一页数据
		PageIndex: 1,
		// 查询的关键词(节点名或者节点ID，模糊搜索)
		Keyword: "test-cluster",
	}

	response, err := eksClient.ListNodes(&params)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.RequestId)
	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}

func DeleteNodes() {
	ak, sk := "Your Ak", "Your Sk"

	eksClient, _ := eks.NewClient(ak, sk)

	params := eks.DeleteNodesReq{
		// eks集群的ID
		ClusterId: "5ff38515-e942-437d-8e16-60f12c382e08",
		// 需要删除的节点的列表
		NodeIds: []string{"eks-4vdp8eu5nq3zppfg"},
	}
	response, err := eksClient.DeleteNodes(&params)
	if err != nil {
		fmt.Println(err)

	}

	bytes, _ := json.Marshal(response)
	fmt.Println(string(bytes))
}
