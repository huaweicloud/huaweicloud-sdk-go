package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygroups"
)

func main() {

	fmt.Println("main start...")
	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get network client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	CreateSecurityGroup(sc)
	GetSecurityGroup(sc)
	ListSecurityGroup(sc)
	DeleteSecurityGroup(sc)
	fmt.Println("main end...")
}

func CreateSecurityGroup(client *gophercloud.ServiceClient) {
	result, err := securitygroups.Create(client, securitygroups.CreateOpts{
		Name: "EricSG",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securitygroup: %+v\r\n", result)
	fmt.Println("securityGroup ID is:", result.ID)
	fmt.Println("securityGroup Name is:", result.Name)
	fmt.Println("securityGroup Description is:", result.Description)
	fmt.Println("securityGroup EnterpriseProjectId is:", result.EnterpriseProjectId)
	fmt.Println("securityGroup SecurityGroupRules is:", result.SecurityGroupRules)
	fmt.Println("securityGroup VpcId is:", result.VpcId)
	//if use SecurityGroupRules.PortRangeMax and SecurityGroupRules.PortRangeMin,you need Judge nil like this
	//if PortRangeMax and PortRangeMin are nil,means null in the API documentation

	for index := 0; index < len(result.SecurityGroupRules); index++ {
		if result.SecurityGroupRules[index].PortRangeMin != nil {
			fmt.Println("securityGroupRule PortRangeMin is:", *result.SecurityGroupRules[index].PortRangeMin)
			fmt.Println("securityGroupRule PortRangeMax is:", *result.SecurityGroupRules[index].PortRangeMax)
		} else {
			fmt.Println("securityGroupRule PortRangeMin is:", result.SecurityGroupRules[index].PortRangeMin)
			fmt.Println("securityGroupRule PortRangeMax is:", result.SecurityGroupRules[index].PortRangeMax)
		}
	}

	fmt.Println("Create success!")
}

func GetSecurityGroup(client *gophercloud.ServiceClient) {
	result, err := securitygroups.Get(client, "05291b64-cdc0-43a4-8c55-75ec8571b06e").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securitygroup: %+v\r\n", result)
	fmt.Println("securityGroup ID is:", result.ID)
	fmt.Println("securityGroup Name is:", result.Name)
	fmt.Println("securityGroup Description is:", result.Description)
	fmt.Println("securityGroup EnterpriseProjectId is:", result.EnterpriseProjectId)
	fmt.Println("securityGroup SecurityGroupRules is:", result.SecurityGroupRules)
	fmt.Println("securityGroup VpcId is:", result.VpcId)
	//if use SecurityGroupRules.PortRangeMax and SecurityGroupRules.PortRangeMin,you need Judge nil like this
	//if PortRangeMax and PortRangeMin are nil,means null in the API documentation

	for index := 0; index < len(result.SecurityGroupRules); index++ {
		if result.SecurityGroupRules[index].PortRangeMin != nil {
			fmt.Println("securityGroupRule PortRangeMin is:", *result.SecurityGroupRules[index].PortRangeMin)
			fmt.Println("securityGroupRule PortRangeMax is:", *result.SecurityGroupRules[index].PortRangeMax)
		} else {
			fmt.Println("securityGroupRule PortRangeMin is:", result.SecurityGroupRules[index].PortRangeMin)
			fmt.Println("securityGroupRule PortRangeMax is:", result.SecurityGroupRules[index].PortRangeMax)
		}
	}
	fmt.Println("Get success!")
}

func ListSecurityGroup(client *gophercloud.ServiceClient) {
	allPages, err := securitygroups.List(client, securitygroups.ListOpts{
		Limit: 2,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result, err1 := securitygroups.ExtractSecurityGroups(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("securitygroup: %+v\r\n", result)
	for _, resp := range result {

		fmt.Println("securityGroup ID is:", resp.ID)
		fmt.Println("securityGroup Name is:", resp.Name)
		fmt.Println("securityGroup Description is:", resp.Description)
		fmt.Println("securityGroup EnterpriseProjectId is:", resp.EnterpriseProjectId)
		fmt.Println("securityGroup SecurityGroupRules is:", resp.SecurityGroupRules)
		fmt.Println("securityGroup VpcId is:", resp.VpcId)
		//if use SecurityGroupRules.PortRangeMax and SecurityGroupRules.PortRangeMin,you need Judge nil like this
		//if PortRangeMax and PortRangeMin are nil,means null in the API documentation

		for index := 0; index < len(resp.SecurityGroupRules); index++ {
			if resp.SecurityGroupRules[index].PortRangeMin != nil {
				fmt.Println("securityGroupRule PortRangeMin is:", *resp.SecurityGroupRules[index].PortRangeMin)
				fmt.Println("securityGroupRule PortRangeMax is:", *resp.SecurityGroupRules[index].PortRangeMax)
			} else {
				fmt.Println("securityGroupRule PortRangeMin is:", resp.SecurityGroupRules[index].PortRangeMin)
				fmt.Println("securityGroupRule PortRangeMax is:", resp.SecurityGroupRules[index].PortRangeMax)
			}
		}
	}
	fmt.Println("List success!")
}

func DeleteSecurityGroup(client *gophercloud.ServiceClient) {
	resp := securitygroups.Delete(client, "a830ab5f-9282-4a32-bfd4-710bfae864d1")
	if resp.Err != nil {
		fmt.Println(resp.Err)
		if ue, ok := resp.Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Delete success!")
}
