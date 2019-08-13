package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
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

	//TestCreateZone(sc)
	//TestGetZone(sc)
	TestListZone(sc)
	//TestListNameServerZone(sc)
	//TestDeleteZone(sc)
	//TestAssociateZone(sc)
	//TestDisAssociateZone(sc)
	//TestUpdateZone(sc)
	fmt.Println("main end...")
}

func TestCreateZone(sc *gophercloud.ServiceClient) {
	opts := zones.CreateOpts{
		Name:     "kaka",
		ZoneType: "private",
		Router: zones.RouterCreateOpts{
			RouterId: "90de16ce-3bd5-42e8-a6b3-d275d26ceb33",
		},
	}

	resp, err := zones.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestCreateZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestGetZone(sc *gophercloud.ServiceClient) {
	zoneId := "4011a19d695bc22f01695c0d084e0690"
	resp, err := zones.Get(sc, zoneId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test TestGetZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

}

func TestListZone(sc *gophercloud.ServiceClient) {

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

	fmt.Println("Test TestListZone success!")
}

func TestDeleteZone(sc *gophercloud.ServiceClient) {
	zoneId := "ff808082695bc23301695c3853d60bab"

	response, err := zones.Delete(sc, zoneId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test TestDeleteZone success!")

}

func TestListNameServerZone(sc *gophercloud.ServiceClient) {
	zoneId := "4011a19d695bc22f01695c0d084e0690"

	resp, err := zones.ListNameServers(sc, zoneId).Extract()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))

	fmt.Println("Test TestListNameServerZone success!")
}

func TestAssociateZone(sc *gophercloud.ServiceClient) {
	opts := zones.AssociateRouterOpts{
		Router: zones.Router{
			RouterId: "90de16ce-3bd5-42e8-a6b3-d275d26ceb33",
		},
	}
	zoneId := "4011a19d695bc22f01695c0d084e0690"
	resp, err := zones.AssociateRouter(sc, zoneId, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestAssociateZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestDisAssociateZone(sc *gophercloud.ServiceClient) {
	opts := zones.DisassociateRouterOpts{
		Router: zones.Router{
			RouterId: "90de16ce-3bd5-42e8-a6b3-d275d26ceb33",
		},
	}
	zoneId := "4011a19d695bc22f01695c0d084e0690"
	resp, err := zones.DisassociateRouter(sc, zoneId, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestDisAssociateZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}

func TestUpdateZone(sc *gophercloud.ServiceClient) {
	opts := zones.UpdateOpts{
		TTL: 900,
	}
	zoneId := "4011a19d695bc22f01695c0d084e0690"
	resp, err := zones.Update(sc, opts, zoneId).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test TestUpdateZone success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}
