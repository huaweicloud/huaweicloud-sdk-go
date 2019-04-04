package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
	"github.com/gophercloud/gophercloud/auth/aksk"

	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
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

	lbid:=LBCreate(sc)
	LBList(sc)
	LBGetStatus(sc,lbid)
	LBGet(sc,lbid)
	LBUpdate(sc,lbid)
	LBDelete(sc,lbid)
	fmt.Println("main end...")
}



func LBCreate(sc *gophercloud.ServiceClient) (lbId string){
	opts:=loadbalancers.CreateOpts{
		Name:"newlb",
		Description:"a new lb",
		VipSubnetID:"5de13914-bd0c-4387-81a7-2d6618cd4824",
		VipAddress:"192.168.0.90",
		Provider:"vlb",
	}

	resp,err:=loadbalancers.Create(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("lb Create success!")
	lbId = (*resp).ID
	return lbId
}



func LBList(sc *gophercloud.ServiceClient)  {
	allPages, err := loadbalancers.List(sc, &loadbalancers.ListOpts{}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("lb list success!")
	fmt.Println(allPages)
}



func LBGet(sc *gophercloud.ServiceClient, id string)  {

	resp,err:= loadbalancers.Get(sc,id).Extract()

	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}

	fmt.Println("lb get success!")
	fmt.Println(resp)


}
func LBUpdate(sc *gophercloud.ServiceClient, id string)  {

	updatOpts:=loadbalancers.UpdateOpts{
		Name:"KAKAK",
		Description:"update test",
	}

	resp,err:=loadbalancers.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}

	fmt.Println("lb update success!")
	fmt.Println(resp)


}

func LBDelete(sc *gophercloud.ServiceClient, id string)  {

	err:=loadbalancers.Delete(sc,id).ExtractErr()

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

func LBGetStatus(sc *gophercloud.ServiceClient, id string)  {

	resp,err:=loadbalancers.GetStatuses(sc,id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("lb get status success!")
	p,_:=json.MarshalIndent(*resp.Loadbalancer,""," ")
	fmt.Println(string(p))
}