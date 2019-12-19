package main

import (
	"fmt"
	"encoding/json"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
)

//AuthWithTokenID using token ID auth method ,list zones.
func AuthWithTokenID() {

	fmt.Println("main start...")
	gophercloud.EnableDebug = true
	// init token auth options
	tokenOpts := token.TokenIdOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		AuthToken:        "{TokenID}",
		}
	// get provider client
	provider, err := openstack.AuthenticatedClient(tokenOpts)

	if err != nil {
		fmt.Println(err)
	}

	sc, err := openstack.NewDNSV2(provider, gophercloud.EndpointOpts{Region: "{your region}"})
	if err != nil {
		fmt.Println(err)
	}
	opts := zones.ListOpts{
		Type: "private",
	}

	resp, err := zones.List(sc, opts).AllPages()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	rs, err := zones.ExtractZones(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, d := range rs.Zones {
		b, _ := json.MarshalIndent(d, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("AuthWithTokenID success!")
	fmt.Println("main end...")
}

func main() {

	AuthWithTokenID()
}
