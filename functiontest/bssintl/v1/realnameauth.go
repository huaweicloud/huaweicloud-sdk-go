package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/realnameauth"
)

func main() {
	fmt.Println("realnameauth start...")

	//打开debug日志
    gophercloud.EnableDebug = true

	provider, err := common.AuthToken()
	//provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("auth success!")

	// 初始化服务的client
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}
	TestIndividualRealNameAuth(sc)
	TestEnterpriseRealNameAuth(sc)
	TestQueryRealNameAuth(sc)
	TestChangeEnterpriseRealNameAuth(sc)
	fmt.Println("realnameauth end...")
}

func TestIndividualRealNameAuth(client *gophercloud.ServiceClient) {
	var a = 4

	opts := realnameauth.IndividualRealNameAuthOpts{
		CustomerId:      "name",
		IdentifyType:    &a,
		VerifiedType:    1,
		VerifiedFileURL: []string{"123","312"},
		Name:            "123",
		VerifiedNumber:  "123",
		ChangeType:      0,
		XaccountType:    "1",
	}
	realNameAuthAuthRsp,err := realnameauth.IndividualRealNameAuth(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(realNameAuthAuthRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestIndividualAuth success")
}

func TestEnterpriseRealNameAuth(client *gophercloud.ServiceClient) {
	var a = 1
	opts := realnameauth. EnterpriseRealNameAuthOpts{
		CustomerId:       "name",
		IdentifyType:     &a,
		CertificateType:     1,
		VerifiedFileURL:  []string{"123", "312"},
		CorpName:         "aaa",
		VerifiedNumber:   "123",
		RegCountry:       "",
		RegAddress:       "",
		XaccountType:     "1",
		EnterprisePerson: nil,
	}
	realNameAuthAuthRsp,err := realnameauth.EnterpriseRealNameAuth(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(realNameAuthAuthRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestEnterpriseAuth success")
}



func TestQueryRealNameAuth(client *gophercloud.ServiceClient) {
	opts := realnameauth.QueryRealNameAuthOpts{
		CustomerId:        "name",
	}

	realNameAuthAuthRsp,err := realnameauth.QueryRealNameAuth(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(realNameAuthAuthRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestSearchAuth success")

}

func TestChangeEnterpriseRealNameAuth(client *gophercloud.ServiceClient) {
	var a = 1
	opts := realnameauth.ChangeEnterpriseRealNameAuthOpts{
		CustomerId:       "name",
		IdentifyType:     &a,
		CertificateType:     1,
		VerifiedFileURL:  []string{"123", "312"},
		CorpName:         "aaa",
		VerifiedNumber:   "123",
		RegCountry:       "",
		RegAddress:       "",
		XaccountType:     "1",
		ChangeType:          &a,
		EnterprisePerson: nil,
	}

	realNameAuthAuthRsp,err := realnameauth.ChangeEnterpriseRealNameAuth(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(realNameAuthAuthRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestChangeAuth success")
}

