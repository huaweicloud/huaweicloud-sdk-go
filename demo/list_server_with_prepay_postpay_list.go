package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/serversext"
	"github.com/gophercloud/gophercloud/auth/aksk"
)

func main() {

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		AccessKey:        "XXXXXXXXXXXXXXXXXXXX",
		SecretKey:        "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
		DomainID:         "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	}

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Fail to get the provider: ", err_auth)
		return
	}

	client, err_client := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})

	if err_client != nil {
		fmt.Println("Fail to get the computer client: ", err_client)
		return
	}


	postpay_server_list, prepay_server_list, err_list := serversext.ListServers(client)

	if err_list != nil {
		if se, ok := err_list.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		} else{
			fmt.Println("Error:", err_list)
		}
		return
	}

	fmt.Println("GetServerList success!")

	fmt.Println("===================postpay_server_list===================")
	for _, s := range postpay_server_list {
		fmt.Println(s.ID)

	}

	fmt.Println("===================monthlySvrs===================")
	for _, s := range prepay_server_list {
		fmt.Println(s.ID)

	}


}


