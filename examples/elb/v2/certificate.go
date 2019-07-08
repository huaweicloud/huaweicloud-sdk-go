package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/certificates"

	"encoding/json"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
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

	id:=CertCreate(sc)
	CertList(sc)
	CertUpdate(sc, id)
	CertGet(sc, id)
	CertDelete(sc, id)

	fmt.Println("main end...")
}

func CertCreate(sc *gophercloud.ServiceClient) (id string) {

	//('HTTP', 'TCP', 'UDP_CONNECT')
	opts:=certificates.CreateOpts{
		Certificate:"*******************",
		PrivateKey:"******************",
		//not necessary
		Name:"mmmmm",
		Type:"client",
		Description:"test create cert",
		Domain:"www.test.com",
	}

	resp,err:=certificates.Create(sc,opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("certificates Create success!")
	p,_:=json.MarshalIndent(*resp,""," ")
	fmt.Println(string(p))
	id=(*resp).ID
	return id
}

func CertList(sc *gophercloud.ServiceClient)  {
	allPages, err := certificates.List(sc, certificates.ListOpts{}).AllPages()
		if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Cert List success!")

	allData,_:=certificates.ExtractCertificates(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _,v :=range allData{

		p,_:=json.MarshalIndent(v,""," ")
		fmt.Println(string(p))
	}

}

func CertGet(sc *gophercloud.ServiceClient, id string)  {

	resp,err:= certificates.Get(sc,id).Extract()

	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("Cert get success!",resp)

	p,_:=json.MarshalIndent(*resp,""," ")
	fmt.Println(string(p))


}

func CertUpdate(sc *gophercloud.ServiceClient, id string)  {

	updatOpts:=certificates.UpdateOpts{
		//not necessary
		Name:"mmmmm",
		Domain:"tt",
		Description:"test cert update",
		PrivateKey:"*********************",
		Certificate:"********************",
	}

	resp,err:=certificates.Update(sc,id,updatOpts).Extract()


	if err!=nil{
		fmt.Println(err)
		if ue,ok:=err.(*gophercloud.UnifiedError); ok{
			fmt.Println("ErrCode",ue.ErrCode)
			fmt.Println("ErrMessage",ue.ErrMessage)
		}
	}
	fmt.Println("Cert update success!")
	p,_:=json.MarshalIndent(*resp,""," ")
	fmt.Println(string(p))
}

func CertDelete(sc *gophercloud.ServiceClient, id string)  {
	err:=certificates.Delete(sc,id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete Cert success!")
}



