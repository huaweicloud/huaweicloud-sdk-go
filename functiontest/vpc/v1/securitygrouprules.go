package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygrouprules"
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
	CreateSecurityGroupRule(sc)
	GetSecurityGroupRule(sc)
	ListSecurityGroupRule(sc)
	DeleteSecurityGroupRule(sc)
	fmt.Println("main end...")
}

func CreateSecurityGroupRule(client *gophercloud.ServiceClient) {
	result, err := securitygrouprules.Create(client, securitygrouprules.CreateOpts{
		Description:     " SecurityGroup",
		SecurityGroupId: "a830ab5f-9282-4a32-bfd4-710bfae864d1",
		Direction:       "egress",
		Protocol:        "tcp",
		RemoteIpPrefix:  "10.10.0.0/24",
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securitygrouprule: %+v\r\n", result)
	fmt.Println("securityGroupRule Description is:", result.Description)
	fmt.Println("securityGroupRule Direction is:", result.Direction)
	fmt.Println("securityGroupRule EtherType is:", result.Ethertype)
	//if PortRangeMax and PortRangeMin are nil,means null in the API documentation
	if result.PortRangeMax != nil {
		fmt.Println("securityGroupRule PortRangeMax is:", *result.PortRangeMax)
		fmt.Println("securityGroupRule PortRangeMin is:", *result.PortRangeMin)
	} else {
		fmt.Println("securityGroupRule PortRangeMax is:", result.PortRangeMax)
		fmt.Println("securityGroupRule PortRangeMin is:", result.PortRangeMin)
	}
	fmt.Println("securityGroupRule ID is:", result.ID)
	fmt.Println("Create success!")
}

func GetSecurityGroupRule(client *gophercloud.ServiceClient) {
	result, err := securitygrouprules.Get(client, "0159cf15-7f02-40c7-90b0-dae9a2d321ef").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securitygrouprule: %+v\r\n", result)
	fmt.Println("securityGroupRule Description is:", result.Description)
	fmt.Println("securityGroupRule Direction is:", result.Direction)
	fmt.Println("securityGroupRule EtherType is:", result.Ethertype)
	//if PortRangeMax and PortRangeMin are nil,means null in the API documentation
	if result.PortRangeMax != nil {
		fmt.Println("securityGroupRule PortRangeMax is:", *result.PortRangeMax)
		fmt.Println("securityGroupRule PortRangeMin is:", *result.PortRangeMin)
	} else {
		fmt.Println("securityGroupRule PortRangeMax is:", result.PortRangeMax)
		fmt.Println("securityGroupRule PortRangeMin is:", result.PortRangeMin)
	}
	fmt.Println("securityGroupRule ID is:", result.ID)
	fmt.Println("Create success!")
}

func ListSecurityGroupRule(client *gophercloud.ServiceClient) {

	allPages, err := securitygrouprules.List(client, securitygrouprules.ListOpts{
		Limit: 20,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err1 := securitygrouprules.ExtractSecurityGroupRules(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return
	}

	fmt.Printf("securitygrouprule: %+v\r\n", result)
	for _, resp := range result {
		fmt.Println("securityGroupRule Description is:", resp.Description)
		fmt.Println("securityGroupRule Direction is:", resp.Direction)
		fmt.Println("securityGroupRule EtherType is:", resp.Ethertype)
		//if PortRangeMax and PortRangeMin are nil,means null in the API documentation
		if resp.PortRangeMax != nil {
			fmt.Println("securityGroupRule PortRangeMax is:", *resp.PortRangeMax)
			fmt.Println("securityGroupRule PortRangeMin is:", *resp.PortRangeMin)
		} else {
			fmt.Println("securityGroupRule PortRangeMax is:", resp.PortRangeMax)
			fmt.Println("securityGroupRule PortRangeMin is:", resp.PortRangeMin)
		}
		fmt.Println("securityGroupRule ID is:", resp.ID)
		fmt.Println(" Create success!")
	}
	fmt.Println("List success!")
}

func DeleteSecurityGroupRule(client *gophercloud.ServiceClient) {
	resp := securitygrouprules.Delete(client, "3ddbcf95-7fe8-4161-8903-fe46ff09b1ae")
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
