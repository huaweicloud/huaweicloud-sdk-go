package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/services"
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
	ListServices(sc)
	GetServices(sc)
	GetCatalog(sc)
	fmt.Println("main end...")
}

// 查询服务列表
// Query a service list
// GET /v3/services
func ListServices(client *gophercloud.ServiceClient) {
	opts := services.ListServiceOpts{
		Type: "",
	}

	// Retrieve a pager (i.e. a paginated collection)
	allPages, err := services.ListService(client, opts).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	// Define an anonymous function to be executed on each page's iteration
	allServices, err := services.ExtractListServices(allPages)

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, service := range allServices {
		// " service " will be a services.Service
		b, _ := json.MarshalIndent(service, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test TestListService success!")
}

// 查询服务详情
// Query the servic detail
// GET /v3/services/{service_id}
func GetServices(client *gophercloud.ServiceClient) {

	service, err := services.Get(client, "").ExtractService()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(service, "", " ")
	fmt.Printf(string(bytes))
	fmt.Println("Test Get Service success!")
}

// 查询服务目录
// Query the service catalog
// GET /v3/auth/catalog
func GetCatalog(client *gophercloud.ServiceClient) {

	catalog, err := services.GetCatalog(client).ExtractCatalog()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("catalog: %+v\r\n", catalog)
	for _, catalog := range catalog.Catalog {
		b, _ := json.MarshalIndent(catalog, "", " ")
		fmt.Println(string(b))
	}
	fmt.Println("Test Get Catalog success!")
}
