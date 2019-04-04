package main


import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/monitors"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/pagination"
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

	id:=HealthMonitorCreate(sc)
	HealthMonitorList(sc)
	HealthMonitorUpdate(sc, id)
	HealthMonitorUpdateDefault(sc, id)
	HealthMonitorGet(sc, id)
	HealthMonitorDelete(sc, id)
	fmt.Println("hmid:", id)

	fmt.Println("main end...")
}

func HealthMonitorCreate(sc *gophercloud.ServiceClient) (id string) {

	trueValue:=true
	//('HTTP', 'TCP', 'UDP_CONNECT')
	opts:=monitors.CreateOpts{
		PoolID:				"7292e873-775c-4d51-8a83-62005f4f92e5",
		Type:				"HTTP",
		Delay:				10,
		Timeout:			10,
		MaxRetries:			3,
		Name:				"mmmmm",
		AdminStateUp:		&trueValue,
		MonitorPort:		520,
		URLPath: 			"/test",
		HTTPMethod: 		"GET",
		ExpectedCodes: 		"200",
		DomainName:			"www.test.com",
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

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

func HealthMonitorUpdateDefault(sc *gophercloud.ServiceClient, id string)  {
	trueValue:=true
	// set monitor port and domain name to default value
	updatOpts:=monitors.UpdateOpts{
		Delay:				10,
		Timeout:			10,
		MaxRetries:			3,
		Name:				"mmmmm",
		AdminStateUp:		&trueValue,
		MonitorPort: 		new(int),
		DomainName: 		new(string),
		URLPath: 			"/test",
		HTTPMethod: 		"GET",
		ExpectedCodes: 		"200",
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}

func HealthMonitorUpdate(sc *gophercloud.ServiceClient, id string)  {
	trueValue:=true
	newPort := 8000
	newDomainName := "www.newtest.com"
	updatOpts:=monitors.UpdateOpts{
		Delay:				10,
		Timeout:			10,
		MaxRetries:			3,
		Name:				"mmmmm",
		AdminStateUp:		&trueValue,
		MonitorPort: 		&newPort,
		DomainName: 		&newDomainName,
		URLPath: 			"/test",
		HTTPMethod: 		"GET",
		ExpectedCodes: 		"200",
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

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
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



