package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bssintl/v1/customermanagement"
)

func main() {
	fmt.Println("customermanagerment start...")

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

	TestCheckUserName(sc)

	TestCheckEmail(sc)

	TestCheckMobile(sc)

	TestCreateCustomer(sc)

	TestQueryCustomer(sc)

	TestFrozenCustomer(sc)

	TestUnFrozenCustomer(sc)

	fmt.Println("customermanagerment end...")
}

func TestCheckUserName(client *gophercloud.ServiceClient) {
	opts := customermanagement.CheckCustomerRegisterInfoOpts{
		SearchType: "name",
		SearchKey: "bsstest02",
	}

	checkUserRsp,err := customermanagement.CheckCustomerRegisterInfo(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if checkUserRsp.ErrorCode == "CBC.0000" {
		fmt.Println("TestCheckUserName success, status=", checkUserRsp.Status)
	}else {
		fmt.Println("TestCheckUserName failed, ErrorCode=", checkUserRsp.ErrorCode, checkUserRsp.ErrorMsg)
	}
}

func TestCheckEmail(client *gophercloud.ServiceClient) {
	opts := customermanagement.CheckCustomerRegisterInfoOpts{
		SearchType: "email",
		SearchKey: "bss.test01@huawei.com",
	}

	checkUserRsp,err := customermanagement.CheckCustomerRegisterInfo(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if checkUserRsp.ErrorCode == "CBC.0000" {
		fmt.Println("TestCheckEmail success, status=", checkUserRsp.Status)
	}else {
		fmt.Println("TestCheckEmail failed, ErrorCode=", checkUserRsp.ErrorCode, checkUserRsp.ErrorMsg)
	}
}

func TestCheckMobile(client *gophercloud.ServiceClient) {
	opts := customermanagement.CheckCustomerRegisterInfoOpts{
		SearchType: "mobile",
		SearchKey: "0086-13951700000",
	}

	checkUserRsp,err := customermanagement.CheckCustomerRegisterInfo(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}


	if checkUserRsp.ErrorCode == "CBC.0000" {
		fmt.Println("TestCheckMobile success, status=", checkUserRsp.Status)
	}else {
		fmt.Println("TestCheckMobile failed, ErrorCode=", checkUserRsp.ErrorCode, checkUserRsp.ErrorMsg)
	}
}

func TestCreateCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.CreateCustomerOpts{
		DomainName:       "cwx521847",
		XAccountId:       "mwx642314",
		XAccountType:     "cwx521847_IDP",
		Password:         "msb121212",
	}

	createCustomerRsp,err := customermanagement.CreateCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}


	bytes, _ := json.MarshalIndent(createCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestCreateCustomer success")
}

func TestQueryCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.QueryCustomerOpts{
		DomainName:           "1",
	}

	queryCustomerRsp,err := customermanagement.QueryCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(queryCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestQueryCustomer success")
}

func TestFrozenCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.FrozenCustomerOpts{
		CustomerIds: []string{"1"},
		Reason: "abc",
	}

	queryCustomerRsp,err := customermanagement.FrozenCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(queryCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestQueryCustomer success")
}

func TestUnFrozenCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.UnFrozenCustomerOpts{
		CustomerIds: []string{"1"},
		Reason: "abc",
	}

	queryCustomerRsp,err := customermanagement.UnFrozenCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(queryCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestQueryCustomer success")
}