package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/policies"
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

	id:=PolicyCreate(sc)
	PolicyList(sc)
	PolicyGet(sc, id)
	PolicyUpdate(sc, id)
	PolicyDelete(sc, id)

	fmt.Println("main end...")
}



func PolicyCreate(sc *gophercloud.ServiceClient) (id string) {

	opts:=policies.CreateOpts{
		RedirectPoolID:"13a887d0-cce3-4d2a-8961-7ad855d054c9",
		ListenerID:"bf392d78-3783-4d8d-9ec1-621e606e6074",
		Action:"REDIRECT_TO_POOL",
		Name:"asd",
		TenantID:"601240b9c5c94059b63d484c92cfe308",
		Description:"create test",
		AdminStateUp:true,
		Position:50,
	}

	resp,err:=policies.Create(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("policy Create success!")
	id=(*resp).ID
	return id
}


func PolicyList(sc *gophercloud.ServiceClient)  {
	allPages, err := policies.List(sc,policies.ListOpts{}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test policy List success!")
	fmt.Println(allPages)

}



func PolicyGet(sc *gophercloud.ServiceClient, id string)  {

	resp,err:= policies.Get(sc,id).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("policy get success!")

	fmt.Println(resp)
}

func PolicyUpdate(sc *gophercloud.ServiceClient, id string)  {

	updatOpts:=policies.UpdateOpts{
		Name:"up test",
		Description:"asdddddddddddddd",
		RedirectPoolID:"2b5a4280-bf51-459e-99ea-31ee24424579",
	}

	resp,err:=policies.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("policy update success!")
	fmt.Println(resp)


}

func PolicyDelete(sc *gophercloud.ServiceClient, id string)  {
	err:=policies.Delete(sc,id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete policy success!")
}



