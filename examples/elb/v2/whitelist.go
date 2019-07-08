package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/whitelist"
	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	whitelistId := WhiteListCreate(sc)
	WhiteListList(sc)
	WhiteListGet(sc, whitelistId)
	WhiteListUpdate(sc, whitelistId)
	WhiteListDelete(sc, whitelistId)

	fmt.Println("main end...")
}

func WhiteListList(sc *gophercloud.ServiceClient) {
	allPages, err := whitelist.List(sc, whitelist.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test get whitelist list success!")
	allwhitelist, err := whitelist.ExtractWhiteLists(allPages)
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(allwhitelist, "", "   ")
	fmt.Println(string(b))
}

func WhiteListGet(sc *gophercloud.ServiceClient, whitelistid string) {
	whiteList, err := whitelist.Get(sc, whitelistid).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test get whitelist detail success!")

	p, _ := json.MarshalIndent(whiteList, "", " ")
	fmt.Println(string(p))

}

func WhiteListCreate(sc *gophercloud.ServiceClient) (whitelistid string) {
	fmt.Println("start wl create...")
	boolvalue := true
	opts := whitelist.CreateOpts{
		ListenerId:      "3374c58c-b376-41f3-b1c3-06f7f9c39a74",
		EnableWhitelist: &boolvalue,
		Whitelist:       "192.168.11.1,192.168.0.1/24",
	}
	whiteList, err := whitelist.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test create whitelist success!")
	whitelistid = whiteList.WhiteList.ID
	p, err := json.MarshalIndent(whiteList, "", "   ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(p))
	return whitelistid
}

func WhiteListUpdate(sc *gophercloud.ServiceClient, whitelistid string) {
	ip := "192.168.11.1"
	truevalue := false
	updatOpts := whitelist.UpdateOpts{
		EnableWhitelist: &truevalue,
		Whitelist:       &ip,
	}

	whiteList, err := whitelist.Update(sc, whitelistid, updatOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("Test update whitelist success!")
	p, _ := json.MarshalIndent(whiteList, "", " ")
	fmt.Println(string(p))
}

func WhiteListDelete(sc *gophercloud.ServiceClient, whitelistid string) {
	err := whitelist.Delete(sc, whitelistid).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test delete whitelist success!")
}
