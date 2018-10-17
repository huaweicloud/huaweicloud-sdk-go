package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"

	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
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

	//TestLBCreate(sc)
	//TestLBList(sc)
	//TestLBGetStatus(sc)
	//TestLBGet(sc)
	//TestLBUpdate(sc)
	TestLBCreate(sc)
	//TestLBGetStats(sc)
	//TestLBDelete(sc)
	fmt.Println("main end...")
}

func TestLBCreate(sc *gophercloud.ServiceClient) {

	trueVlaue := true
	opts := loadbalancers.CreateOpts{
		Name:        "newlb",
		Description: "a new lb",
		//VipSubnetID:"e9181c01-697c-47be-b9a9-502b940c3297",
		VipSubnetID: "20b8a44b-e724-4103-8233-f70c7aa1bbc2",
		//VipAddress:"192.168.10.7",
		AdminStateUp: &trueVlaue,
		Provider:     "vlb",
	}
for true{
		resp, err := loadbalancers.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("lb Create success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}
	//resp, err := loadbalancers.Create(sc, opts).Extract()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	if ue, ok := err.(*gophercloud.UnifiedError); ok {
	//		fmt.Println("ErrCode:", ue.ErrorCode())
	//		fmt.Println("Message:", ue.Message())
	//	}
	//	return
	//}
	//fmt.Println("lb Create success!")
	//p, _ := json.MarshalIndent(*resp, "", " ")
	//fmt.Println(string(p))

}

func TestLBList(sc *gophercloud.ServiceClient) {
	l := 1
	allPages, err := loadbalancers.List(sc, &loadbalancers.ListOpts{Limit: &l}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("TestLBListsuccess!")

	allData, _ := loadbalancers.ExtractLoadBalancers(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestLBGet(sc *gophercloud.ServiceClient) {

	id := "9eb9ef27-25d1-45d3-b860-84791d97f328"

	resp, err := loadbalancers.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}

	fmt.Println("lb get success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestLBUpdate(sc *gophercloud.ServiceClient) {

	id := "9eb9ef27-25d1-45d3-b860-84791d97f328"

	updatOpts := loadbalancers.UpdateOpts{
		Name: "KAKAK",
	}

	resp, err := loadbalancers.Update(sc, id, updatOpts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("lb update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestLBDelete(sc *gophercloud.ServiceClient) {

	id := "111ea5fc-69e1-4687-8ff0-f9ad454558c0"
	err := loadbalancers.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("lb delete success!")
}

func TestLBGetStatus(sc *gophercloud.ServiceClient) {

	id := "fd18b88e-e75f-46d6-984e-753eb56d7b17"
	resp, err := loadbalancers.GetStatuses(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	p, _ := json.MarshalIndent(*resp.Loadbalancer, "", " ")
	fmt.Println(string(p))
	fmt.Println(*resp.Loadbalancer)

}

//func TestLBGetStats(sc *gophercloud.ServiceClient)  {
//
//		id:="111ea5fc-69e1-4687-8ff0-f9ad454558c0"
//		resp,err:=loadbalancers.GetStats(sc,id).Extract()
//		if err!=nil{
//
//			fmt.Println(err)
//			if ue,ok:=err.(*gophercloud.UnifiedError);ok{
//			fmt.Println("ErrCode:", ue.ErrorCode())
//			fmt.Println("Message:", ue.Message())
//			}
//			return
//		}
//		p,_:=json.MarshalIndent(*resp,""," ")
//		fmt.Println("result is :",string(p))
//
//}
