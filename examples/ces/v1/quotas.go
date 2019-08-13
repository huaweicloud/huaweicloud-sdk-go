package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/quotas"
)

func main() {
	fmt.Println("main start...")
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}

	sc, err := openstack.NewCESV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ces client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	QuotaList(sc)
	fmt.Println("main end...")
}

func QuotaList(sc *gophercloud.ServiceClient) {
	getInfo, err := quotas.Get(sc).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	res, marshalErr := json.MarshalIndent(getInfo, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal getInfo error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}
