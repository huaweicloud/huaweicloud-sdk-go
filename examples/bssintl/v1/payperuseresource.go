package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/payperuseresource"
)

func main() {

	//AKSK auth，initial parameter.
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		DomainID:         "{domainID}",
	}

	//initial provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("get provider client failed")
		fmt.Println(errAuth.Error())
		return
	}

	// initial client
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	QueryCustomerResource(sc)

}

func QueryCustomerResource(client *gophercloud.ServiceClient) {
	opts := payperuseresource.QueryCustomerResourceOpts{
		CustomerId: "3d2c6b3ab1fd4e26846b0f2c46e67bda",
		RegionCode: "cn-north-1",
		CloudServiceTypeCode: "hws.service.type.ebs",
		ResourceTypeCode: "hws.resource.type.volume",
		ResourceIds: []string{
			"71e3eeb5-4b77-44ae-9c42-119ee7976cf7",
			"39d90d01-4774-4af6-8136-83ba5732bccf"},
		StartTimeBegin: "2018-06-01T11:00:00Z",
		StartTimeEnd: "2018-06-30T11:00:00Z",
	}
	detailRsp,err := payperuseresource.QueryCustomerResource(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("QueryResources success")
}