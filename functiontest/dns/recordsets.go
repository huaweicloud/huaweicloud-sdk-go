package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	sc, err := openstack.NewDNSV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get DNS v2 client failed")
		fmt.Println(err.Error())
		return
	}

	//TestCreateRecordSetsByZone(sc)
	//TestGetRecordSetsByZone(sc)
	TestUpdateRecordSetsByZone(sc)
	//TestDeleteRecordSetsByZone(sc)
	//TestListRecordSetsByZone(sc)
	//TestListRecordSets

	fmt.Println("main end...")
}

func TestCreateRecordSetsByZone(sc *gophercloud.ServiceClient) {
	opts := recordsets.CreateOpts{
		Type: "A",
		Name: "www.example.kaka.",
		Records: []string{"192.168.10.1",
			"192.168.10.2"},
	}
	zoneId := "4011afa2695c457701695c49a0f7007c"
	resp, err := recordsets.Create(sc, zoneId, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestCreateRecordSetsByZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestGetRecordSetsByZone(sc *gophercloud.ServiceClient) {
	zoneId := "4011afa2695c457701695c49a0f7007c"
	rrd := "4011afa2695c457701695c49a0f7007d"
	resp, err := recordsets.Get(sc, zoneId, rrd).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestGetRecordSetsByZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

}

func TestUpdateRecordSetsByZone(sc *gophercloud.ServiceClient) {
	zoneId := "4011afa2695c457701695c49a0f7007c"
	rrd := "4011afa2695c457701695c54dd780177"
	updateOpts := recordsets.UpdateOpts{
		TTL:         1098,
		Description: "imaok",
		Records:     []string{"192.168.10.200"},
	}

	resp, err := recordsets.Update(sc, zoneId, rrd, updateOpts).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Test TestUpdateRecordSetsByZone update success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestDeleteRecordSetsByZone(sc *gophercloud.ServiceClient) {
	zoneId := "4011afa2695c457701695c49a0f7007c"
	rrd := "asdfasdfas"

	response, err := recordsets.Delete(sc, zoneId, rrd).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test TestDeleteRecordSetsByZone success!")

}

func TestListRecordSetsByZone(sc *gophercloud.ServiceClient) {
	zoneId := "4011afa2695c457701695c49a0f7007c"
	opts := recordsets.ListByZoneOpts{}

	resp, err := recordsets.ListByZone(sc, zoneId, opts).AllPages()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	rs, err := recordsets.ExtractRecordSets(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, d := range rs.Recordsets {

		b, _ := json.MarshalIndent(d, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test TestListRecordSetsByZone success!")
}

func TestListRecordSets(sc *gophercloud.ServiceClient) {
	opts := recordsets.ListOpts{
		Type: "A",
	}

	resp, err := recordsets.List(sc, opts).AllPages()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	rs, err := recordsets.ExtractRecordSets(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, d := range rs.Recordsets {

		b, _ := json.MarshalIndent(d, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("Test TestListRecordSetsByZone success!")
}
