package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/ptrs"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	opts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "your username",
		Password:         "your password",
		DomainID:         "your domainId",
		ProjectID:        "your projectID",
	}
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}
	dnsClient, err2 := openstack.NewDNSV2(provider, gophercloud.EndpointOpts{})
	if err2 != nil {
		panic(err2)
	}

	//setup Ptr record
	setupPtr(dnsClient)
	//get ptr record
	getPtr(dnsClient)
	//update ptrs
	updatePtr(dnsClient)
	//list ptrs
	listPtr(dnsClient)
	//restore ptr
	restorePtr(dnsClient)
}

func setupPtr(sc *gophercloud.ServiceClient) {
	opts := ptrs.SetupOpts{
		Ptrdname: "www.xxxx.com", //your ptr name
	}
	region := "region name" //region name
	fip := "your fip id"
	ptr, err := ptrs.Setup(sc, region, fip, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(ptr.ID, ptr.Ptrdname)
	ptrjson, _ := json.MarshalIndent(ptr, "", " ")
	fmt.Println(string(ptrjson))
}

func getPtr(sc *gophercloud.ServiceClient) {
	regionID := "region name"
	fip := "your fip id"

	ptr, err := ptrs.Get(sc, regionID, fip).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(ptr.ID, ptr.Ptrdname)
	ptrjson, _ := json.MarshalIndent(ptr, "", " ")
	fmt.Println(string(ptrjson))

}

func updatePtr(sc *gophercloud.ServiceClient) {
	opts := ptrs.UpdateOpts{
		Ptrdname: "www.yyy.com",
	}
	regionID := "region name"
	fip := "your fip id"
	resp, err := ptrs.Update(sc, regionID, fip, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

}

func listPtr(sc *gophercloud.ServiceClient) {

	opts := ptrs.ListOpts{
		Limit: 2,
	}

	ptrs.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		listptrs, err := ptrs.ExtractPtrs(page)
		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, err
		} else {
			for _, floatingIp := range listptrs.Floatingips {
				fmt.Println(floatingIp.ID, floatingIp.Ptrdname)
				floatingIpjson, _ := json.MarshalIndent(floatingIp, "", " ")
				fmt.Println(string(floatingIpjson))
			}
		}
		return false, err
	})
}

func restorePtr(sc *gophercloud.ServiceClient) {
	regionID := "region name" //region name
	fip := "your fip id"
	err := ptrs.Restore(sc, regionID, fip).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

}
