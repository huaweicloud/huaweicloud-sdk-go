package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
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
	//Create RecordSets
	createRecordset(dnsClient)
	//Get Recordset by ID
	getRecordSetsByID(dnsClient)
	//List RecordSets
	listRecordsets(dnsClient)
	//List RecordSets by Zone
	listRecordsetsByZone(dnsClient)
	//update Recordset
	updateRecordset(dnsClient)
	//Delete Recordsets
	deleteRecordset(dnsClient)
}

func listRecordsets(dnsClient *gophercloud.ServiceClient) {

	listOpts := recordsets.ListOpts{
		Type:  "A",
		Limit: 20,
	}

	recordsets.List(dnsClient, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		allRRs, err := recordsets.ExtractRecordSets(page)
		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, err
		} else {
			for _, rr := range allRRs.Recordsets {
				rrjson, _ := json.MarshalIndent(rr, "", " ")
				fmt.Println(string(rrjson))
			}
		}
		return false, err
	})

}

func listRecordsetsByZone(dnsClient *gophercloud.ServiceClient) {

	listOpts := recordsets.ListByZoneOpts{
		Limit: 10,
	}

	zoneID := "your zone id"

	recordsets.ListByZone(dnsClient, zoneID, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		allRRs, err := recordsets.ExtractRecordSets(page)
		if err != nil {
			if ue, ok := err.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, err
		} else {
			for _, rr := range allRRs.Recordsets {
				rrjson, _ := json.MarshalIndent(rr, "", " ")
				fmt.Println(string(rrjson))
			}
		}
		return false, err
	})
}
func createRecordset(dnsClient *gophercloud.ServiceClient) {
	createOpts := recordsets.CreateOpts{
		Name:        "your recordset name",
		Type:        "A",
		TTL:         3600,
		Description: "your description",
		Records:     []string{"10.1.0.2"},
	}

	zoneID := "your zone id"

	rr, err := recordsets.Create(dnsClient, zoneID, createOpts).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(rr.ID, rr.Name, rr.Records)
	rrjson, _ := json.MarshalIndent(rr, "", " ")
	fmt.Println(string(rrjson))
}

func getRecordSetsByID(sc *gophercloud.ServiceClient) {
	zoneId := "your zone id"
	rrId := "your recordset id"
	rset, err := recordsets.Get(sc, zoneId, rrId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(rset.ID, rset.Name, rset.Records)
	fmt.Println(rset.Links)
	rsetjson, _ := json.MarshalIndent(rset, "", " ")
	fmt.Println(string(rsetjson))

}

func deleteRecordset(dnsClient *gophercloud.ServiceClient) {
	zoneID := "your zone id"
	recordsetID := "your recordset id"

	rr, err := recordsets.Delete(dnsClient, zoneID, recordsetID).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	rrjson, _ := json.MarshalIndent(rr, "", " ")
	fmt.Println(string(rrjson))
	fmt.Printf("%+v", rr.Links)
	fmt.Println(rr.Links, rr.Status)
}
func updateRecordset(dnsClient *gophercloud.ServiceClient) {
	updataOpts := recordsets.UpdateOpts{
		Description: "your description",
		TTL:         600,
		Records:     []string{"1.3.1.4", "3.3.4.4"}, //your ip address
	}
	zoneID := "your zone id"
	recordsetID := "your recordser id"

	rr, err := recordsets.Update(dnsClient, zoneID, recordsetID, updataOpts).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println(rr.ID, rr.Name, rr.Records)
	rrjson, _ := json.MarshalIndent(rr, "", " ")
	fmt.Println(string(rrjson))
}
