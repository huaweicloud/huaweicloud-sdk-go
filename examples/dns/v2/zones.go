package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
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

	//Create public zone
	createPublicZone(dnsClient)
	//Get zone
	getZone(dnsClient)
	//List Zones
	listZone(dnsClient)
	//update Zone
	updateZone(dnsClient)
	//Delete zone
	deleteZone(dnsClient)
	//Create private Zone
	createPrivateZone(dnsClient)
	//List nameserver
	listNameServerZone(dnsClient)
	//Associate private Zone with router
	associateZone(dnsClient)
	//Disassociate private Zone with router
	disAssociateZone(dnsClient)

}
func listZone(dnsClient *gophercloud.ServiceClient) {
	listOpts := zones.ListOpts{
		Type:  "public",
		Limit: 2,
	}

	zones.List(dnsClient, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		allZones, err := zones.ExtractZones(page)
		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, err
		} else {
			for _, zone := range allZones.Zones {
				zonejson, _ := json.MarshalIndent(zone, "", " ")
				fmt.Println(string(zonejson))
				fmt.Println(zone.Links)
			}
		}
		return false, err
	})
}

func createPublicZone(dnsClient *gophercloud.ServiceClient) {
	createOpts := zones.CreateOpts{
		Name:        "your zone name",
		Email:       "your email",
		ZoneType:    "public",
		TTL:         7200,
		Description: "your description.",
	}

	zone, err := zones.Create(dnsClient, createOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(string(zonejson))
}

func createPrivateZone(dnsClient *gophercloud.ServiceClient) {
	createOpts := zones.CreateOpts{
		Name:        "your private zone name",
		Email:       "your email",
		ZoneType:    "private",
		TTL:         7200,
		Description: "your description.",
		Router: zones.RouterCreateOpts{
			RouterRegion: "regionname",
			RouterId:     "vpc id",
		},
	}

	zone, err := zones.Create(dnsClient, createOpts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(zone.Router)
	fmt.Println(string(zonejson))
}

func getZone(sc *gophercloud.ServiceClient) {
	zoneId := "your zone id"
	zone, err := zones.Get(sc, zoneId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(string(zonejson))

}

func listNameServerZone(sc *gophercloud.ServiceClient) {
	zoneId := "your zone id"

	nameserver, err := zones.ListNameServers(sc, zoneId).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	nameserverjson, _ := json.MarshalIndent(nameserver, "", " ")
	fmt.Println(string(nameserverjson))

}

func updateZone(dnsClient *gophercloud.ServiceClient) {
	updateOpts := zones.UpdateOpts{
		Email:       "your email",
		TTL:         900,
		Description: "your description",
	}
	zoneID := "your zone id"
	zone, err := zones.Update(dnsClient, updateOpts, zoneID).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(string(zonejson))
}

func deleteZone(dnsClient *gophercloud.ServiceClient) {
	zoneID := "your zone id"
	zone, err := zones.Delete(dnsClient, zoneID).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(string(zonejson))
}

func associateZone(sc *gophercloud.ServiceClient) {
	opts := zones.AssociateRouterOpts{
		Router: zones.Router{
			RouterId: "vpc id",
		},
	}
	zoneId := "your zone id"
	zone, err := zones.AssociateRouter(sc, zoneId, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(string(zonejson))
}

func disAssociateZone(sc *gophercloud.ServiceClient) {
	opts := zones.DisassociateRouterOpts{
		Router: zones.Router{
			RouterId: "vpc id",
		},
	}
	zoneId := "your zone id"
	zone, err := zones.DisassociateRouter(sc, zoneId, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	zonejson, _ := json.MarshalIndent(zone, "", " ")
	fmt.Println(string(zonejson))
}
