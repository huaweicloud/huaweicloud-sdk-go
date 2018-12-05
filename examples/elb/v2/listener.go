package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"	
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "myhuaweicloud.com",
		Region:           "cn-north-1",
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

	lsId:=ListenerCreate(sc)
	ListenerGet(sc,lsId)
	ListenerList(sc)
	ListenerUpdate(sc,lsId)
	ListenerDelete(sc,lsId)

	fmt.Println("main end...")
}

func ListenerCreate(sc *gophercloud.ServiceClient) (lsID string) {
	opts:=listeners.CreateOpts{
		Name:"new listener",
		Description:"AAAAAAA",
		Protocol:listeners.ProtocolTCP,
		ProtocolPort:20,
		LoadbalancerID:"812b636f-6287-420a-af23-307f2e490615",

	}

	resp,err:=listeners.Create(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("listener Create success!")
	fmt.Println(resp)
	lsId:=(*resp).ID
	return lsId

}


func ListenerList(sc *gophercloud.ServiceClient)  {
	allPages, err := listeners.List(sc,listeners.ListOpts{}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test ListenerList success!")
	fmt.Println(allPages)

}



func ListenerGet(sc *gophercloud.ServiceClient, id string)  {

	resp,err:= listeners.Get(sc,id).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("listener get success!")
	fmt.Println(resp)


}
func ListenerUpdate(sc *gophercloud.ServiceClient, id string)  {
	updatOpts:=listeners.UpdateOpts{
		Name:"KAKAK A listener",
		Description:"ls update test",
		DefaultPoolID:"2b5a4280-bf51-459e-99ea-31ee24424579",
		DefaultTlsContainerRef:"23b58a961a4d4c95be585e98046e657a",
		ClientCaTlsContainerRef:"23b58a961a4d4c95be585e98046e657a",
	}

	resp,err:=listeners.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("listener update success!")
	fmt.Println(resp)
}

func ListenerDelete(sc *gophercloud.ServiceClient, id string)  {

	err:=listeners.Delete(sc,id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete listener success!")
}



