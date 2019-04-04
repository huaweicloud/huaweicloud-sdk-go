package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/bandwidths"
)

func main() {
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//初始化provider client。
	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}
	//初始化服务 client
	client, err_client := openstack.NewVPCV2(provider, gophercloud.EndpointOpts{})
	if err_client != nil {
		fmt.Println("Failed to get the NewVPCV2 client: ", err_client)
		return
	}

	size := 10
	//modify bandwidth size
	updateopts := bandwidths.UpdateOpts{
		Bandwidth: bandwidths.Bandwidth{
			Name: "eeeeeeeeeeeeeeeee",
			Size: size,
		},
		ExtendParam: &bandwidths.ExtendParam{
			IsAutoPay: "true",
		},
	}

	data, err := bandwidths.Update(client, "2a2ebbe0-a9c3-475a-b1ac-089aa435a426", updateopts)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test Update  bandwidths success!")

	if order, ok := data.(bandwidths.PrePaid); ok {
		fmt.Println("its order id ")
		fmt.Println("order id is", order.OrderID)
	}

	if on, ok := data.(bandwidths.PostPaid); ok {
		fmt.Println("its bandwidth info")
		fmt.Println("bandwidth id is ", on.ID)
		fmt.Println("bandwidth Size is ", on.Size)
		fmt.Println("bandwidth Name is ", on.Name)
		fmt.Println("bandwidth ShareType is ", on.ShareType)
		fmt.Println("bandwidth ChargeMode is ", on.ChargeMode)
		fmt.Println("bandwidth PublicipInfo is ", on.PublicipInfo)
	}

}
