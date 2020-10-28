package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/endpoints"
)

func main() {
	fmt.Println("main start...")
	// AKSK 认证
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		DomainID:         "{domainID}",
	}

	// init provider client
	provider, errAuth := openstack.AuthenticatedClient(opts)

	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}
	// 初始化服务 client
	sc, errClient := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewIdentityV3 client: ", errClient)
		return
	}

	// 开始测试
	ListEndpoints(sc)
	GetEndpoints(sc)
	fmt.Println("main end...")
}

// 查询终端节点列表
// Query a endpoint list
// GET /v3/endpoints
func ListEndpoints(client *gophercloud.ServiceClient) {

	// List enumerates the services available to a specific user.
	listOpts := endpoints.ListEndPointsOpts{
		ServiceID: "",
		Interface: "",
	}

	// Retrieve a pager (i.e. a paginated collection)
	allPages, err := endpoints.ListEndPoint(client, listOpts).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	// Define an anonymous function to be executed on each page's iteration
	allEndpoints, err := endpoints.ExtractEndpointsList(allPages)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, endpoint := range allEndpoints {
		// " endpoint " will be a endpoints.Endpoint
		b, _ := json.MarshalIndent(endpoint, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test TestListEndpoints success!")
}

// 查询终端节点详情
// Query the endpoint detail
// GET /v3/endpoints/{endpoint_id}
func GetEndpoints(client *gophercloud.ServiceClient) {
	endpoint, err := endpoints.Get(client, "").ExtractDetail()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, err := json.MarshalIndent(endpoint, "", " ")
	fmt.Printf(string(bytes))
	fmt.Println("Test Get Endpoints success!")
}
