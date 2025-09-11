package eks

import (
	"encoding/json"
	"github.com/capitalonline/cds-cloudos-go-sdk/cds"
	"github.com/capitalonline/cds-cloudos-go-sdk/http"
)

const (
	CreateClusterAction    = "CreateCluster"
	ListClustersAction     = "ListClusters"
	GetClusterEventsAction = "GetClusterEvents"
	DeleteClusterAction    = "DeleteCluster"
	GetClusterAction       = "GetCluster"
	GetTaskStatusAction    = "TaskStatus"
	AddClusterSubnetAction = "AddClusterSubnet"
)

func (c *Client) CreateCluster(args *CreateClusterReq) (*CreateClusterResult, error) {
	result := &CreateClusterResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", CreateClusterAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListClusters(args *ListClustersReq) (*ListClustersResult, error) {
	result := &ListClustersResult{}

	bytes, _ := json.Marshal(args)
	params := make(map[string]string)
	_ = json.Unmarshal(bytes, &params)
	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", ListClustersAction).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetClusterEvents(clusterId string) (*GetClusterEventsResult, error) {
	result := &GetClusterEventsResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", GetClusterEventsAction).
		WithQueryParam("ClusterId", clusterId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteCluster(clusterId string) (*DeleteClusterResult, error) {
	result := &DeleteClusterResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", DeleteClusterAction).
		WithBody(map[string]string{
			"ClusterId": clusterId,
		}).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetCluster(clusterId string) (*GetClusterResult, error) {
	result := &GetClusterResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", GetClusterAction).
		WithQueryParam("ClusterId", clusterId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) GetTaskStatus(taskId string) (*TaskStatusResult, error) {
	result := &TaskStatusResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.GET).
		WithQueryParam("Action", GetTaskStatusAction).
		WithQueryParam("TaskId", taskId).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) AddClusterSubnet(args *AddClusterSubnetReq) (*AddClusterSubnetResult, error) {
	result := &AddClusterSubnetResult{}

	err := cds.NewRequestBuilder(c).
		WithURI(eksURI).
		WithMethod(http.POST).
		WithQueryParam("Action", AddClusterSubnetAction).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}
