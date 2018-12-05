package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/certificates"
	"encoding/json"
)

func main() {

	fmt.Println("main start...")
	provider, err := common.AuthAKSK()
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
	TestCertificateList(sc)
	//TestCertificateCreate(sc)
	//TestCertificateUpdate(sc)/
	//TestCertificateGet(sc)
	//TestCertificateDelete(sc)

	fmt.Println("main end...")
}

func TestCertificateCreate(sc *gophercloud.ServiceClient) {
	//('HTTP', 'HTTPS', 'PING', 'TCP', 'UDP_CONNECT')
	opts := certificates.CreateOpts{
		Certificate: "******",
		Name:        "aaaaa",
		Description: "new cet",
		PrivateKey:  "******",
	}

	resp, err := certificates.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("certificates Create success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestCertificateList(sc *gophercloud.ServiceClient) {
	allPages, err := certificates.List(sc, certificates.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test Certificate List success!")

	allData, _ := certificates.ExtractCertificates(allPages)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range allData {

		p, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(p))
	}

}

func TestCertificateGet(sc *gophercloud.ServiceClient) {

	id := "042c646dd9214aa08a3f3cb8d4ae85ca"

	resp, err := certificates.Get(sc, id).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Certificate get success!", resp)

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
func TestCertificateUpdate(sc *gophercloud.ServiceClient) {

	id := "042c646dd9214aa08a3f3cb8d4ae85ca"

	updatOpts := certificates.UpdateOpts{
		Name: "KAKAK A Certificate",
	}

	resp, err := certificates.Update(sc, id, updatOpts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Certificate update success!")
	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}

func TestCertificateDelete(sc *gophercloud.ServiceClient) {

	id := "042c646dd9214aa08a3f3cb8d4ae85ca"
	err := certificates.Delete(sc, id).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("delete Certificate success!")
}
