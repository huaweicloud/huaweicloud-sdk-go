package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/realnameauth"
)

func main() {
	//AKSK auth，initial parameter.
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		DomainID:         "{domainID}",
	}

	//initial provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("get provider client failed")
		fmt.Println(errAuth.Error())
		return
	}

	// initial client
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}
	IndividualRealNameAuth(sc)
	EnterpriseRealNameAuth(sc)
	QueryRealNameAuth(sc)
	ChangeEnterpriseRealNameAuth(sc)
	fmt.Println("realnameauth end...")
}

func IndividualRealNameAuth(client *gophercloud.ServiceClient) {
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
	fmt.Println("IndividualAuth success")
}

func EnterpriseRealNameAuth(client *gophercloud.ServiceClient) {
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
	fmt.Println("EnterpriseAuth success")
}



func QueryRealNameAuth(client *gophercloud.ServiceClient) {
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
	fmt.Println("SearchAuth success")

}

func ChangeEnterpriseRealNameAuth(client *gophercloud.ServiceClient) {
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
	fmt.Println("ChangeAuth success")
}

