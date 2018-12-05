package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/monitors"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/pagination"
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

	id:=HealthMonitorCreate(sc)
	HealthMonitorList(sc)
	HealthMonitorUpdate(sc, id)
	HealthMonitorGet(sc, id)
	HealthMonitorDelete(sc, id)
	fmt.Println("hmid:", id)

	fmt.Println("main end...")
}

func HealthMonitorCreate(sc *gophercloud.ServiceClient) (id string) {

	trueValue:=true
	//('HTTP', 'TCP', 'UDP_CONNECT')
	opts:=monitors.CreateOpts{
		PoolID:"13a887d0-cce3-4d2a-8961-7ad855d054c9",
		Type:"HTTP",
		Delay:10,
		Timeout:10,
		MaxRetries:3,
		Name:"mmmmm",
		AdminStateUp:&trueValue,
		MonitorPort: 520,
		URLPath: "/test",
		HTTPMethod: "GET",
		ExpectedCodes: "200",
		DomainName:"www.test.com",
	}

	resp,err:=monitors.Create(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("monitors Create success!")
	fmt.Println(resp)
	id=(*resp).ID
	return id
}


func HealthMonitorList(sc *gophercloud.ServiceClient) (allPages pagination.Page) {
	allPages, err := monitors.List(sc,monitors.ListOpts{}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test monitor List success!")
	fmt.Println(allPages)
	return  allPages

}



func HealthMonitorGet(sc *gophercloud.ServiceClient, id string)  {

	resp,err:= monitors.Get(sc,id).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("monitor get success!\n",resp)

}

func HealthMonitorUpdate(sc *gophercloud.ServiceClient, id string)  {
	trueValue:=true
	updatOpts:=monitors.UpdateOpts{
		Delay:10,
		Timeout:10,
		MaxRetries:3,
		Name:"mmmmm",
		AdminStateUp:&trueValue,
		MonitorPort: 520,
		URLPath: "/test",
		HTTPMethod: "GET",
		ExpectedCodes: "200",
	}

	resp,err:=monitors.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("monitor update success!")
	fmt.Println(*resp)
}

func HealthMonitorDelete(sc *gophercloud.ServiceClient, id string)  {
	err:=monitors.Delete(sc,id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete monitor success!")
}



