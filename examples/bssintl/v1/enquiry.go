package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/enquiry"
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

	QueryRating(sc)
}

func QueryRating(client *gophercloud.ServiceClient) {
	var a = 0
	opts := enquiry.QueryRatingOpts{
		TenantId:                   "74610f3a5ad941998e91f076297ecf27",
		RegionId:                   "cn-north-1",
		AvaliableZoneId:            "cn-north-1",
		ChargingMode:               &a,
		PeriodType:                 1,
		PeriodNum:                  10,
		PeriodEndDate:              "",
		RelativeResourceId:         "546568dsdcsc",
		RelativeResourcePeriodType: 1,
		SubscriptionNum:            10,
		ProductInfo: 				[]enquiry.ProductInfo{
			{
				Id: "1",
				CloudServiceType: "hws.service.type.ec2",
				ResourceType: "hws.resource.type.vm",
				ResourceSpecCode: "s2.small.1.linux",
			},
			{
				Id: "1",
				CloudServiceType: "hws.service.type.ec2",
				ResourceType: "hws.resource.type.vm",
				ResourceSpecCode: "s2.small.1.linux",
			},
		},
	}

	enquiry.QueryRating(client, opts)
}
