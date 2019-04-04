package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"	
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
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

	lsId:=ListenerCreate(sc)
	ListenerGet(sc,lsId)
	ListenerList(sc)
	ListenerUpdate(sc,lsId)
	ListenerUpdateCloseCAAndSni(sc,lsId)
	ListenerDelete(sc,lsId)

	fmt.Println("main end...")
}

func ListenerCreate(sc *gophercloud.ServiceClient) (lsID string) {
	fmt.Printf("start listener Create\n")
	var	sni = []string{
		"e15d1b5000474adca383c3cd9ddc06d4",
		"5882325fd6dd4b95a88d33238d293a0f"}
	opts:=listeners.CreateOpts{
		Name:						"new_listener",
		Description:				"test_listener",
		Protocol:					listeners.ProtocolTerminatedHTTPS,
		ProtocolPort:				20,
		LoadbalancerID:				"6bb85e33-4953-457a-85a9-336d76125b7b",
		DefaultTlsContainerRef:		"b912ae6dd0b24cd0984b96c1cbd928a1",
		ClientCaTlsContainerRef:	"417a0976969f497db8cbb083bff343ba",
		SniContainerRefs:           sni,
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
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
	} else {
		fmt.Println("listener get success!")
	}
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}

func ListenerUpdate(sc *gophercloud.ServiceClient, id string)  {
	fmt.Printf("start listener Update\n")
	//close Mutual Authentication
	ClientCaTlsContainerRef := "0bb3c0edd72d4ff4b15addfdba0bfb2d"
	DefaultTlsContainerRef := "02dcd56799e045bf8b131533cc911dd6"
	SniContainerRefs := []string{"7b53376c55a442a494f44952999ff2de"}
	updatOpts:=listeners.UpdateOpts{
		Name:						"listener_update_test",
		Description:				"ls update test",
		DefaultTlsContainerRef:     &DefaultTlsContainerRef,
		ClientCaTlsContainerRef:	&ClientCaTlsContainerRef,
		SniContainerRefs:           &SniContainerRefs,
	}

	resp,err:=listeners.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	} else {
		fmt.Println("listener update success!")
	}

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}

func ListenerUpdateCloseCAAndSni(sc *gophercloud.ServiceClient, id string)  {
	fmt.Printf("start listener Update\n")
	//close Mutual Authentication
	updatOpts:=listeners.UpdateOpts{
		Name:						"listener_update_test",
		Description:				"ls update test",
		ClientCaTlsContainerRef:	new(string),
		SniContainerRefs:           &[]string{},
	}

	resp,err:=listeners.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	} else {
		fmt.Println("listener update success!")
	}

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
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



