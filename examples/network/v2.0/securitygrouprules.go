package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
)

func main() {
	fmt.Println("main start...")
	//AKSK authentication, initialization authentication parameters
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	//Initialization provider client
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//Initialization service client
	sc, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get Network client failed")
		fmt.Println(err)
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

// List SecurityGroupRule
func ListSecurityGroupRule(sc *gophercloud.ServiceClient) {
	allPages, err := rules.List(sc, rules.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := rules.ExtractRules(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, resp := range result {

		fmt.Printf("securityGroupRule: %+v\r\n", resp)
		fmt.Println("securityGroupRule ID is:", resp.ID)
		fmt.Println("securityGroupRule Direction is:", resp.Direction)
		fmt.Println("securityGroupRule EtherType is:", resp.EtherType)
		fmt.Println("securityGroupRule SecGroupID is:", resp.SecGroupID)
		fmt.Println("securityGroupRule PortRangeMin is:", resp.PortRangeMin)
		fmt.Println("securityGroupRule PortRangeMax is:", resp.PortRangeMax)
		fmt.Println("securityGroupRule Protocol is:", resp.Protocol)
		fmt.Println("securityGroupRule ProjectId is:", resp.ProjectId)
		fmt.Println("securityGroupRule UpdatedAt is:", resp.UpdatedAt)
	}

	fmt.Println("List success!")

}

// Get a SecurityGroupRule
func GetSecurityGroupRule(sc *gophercloud.ServiceClient) {
	result, err := rules.Get(sc, "xxxxxx").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securityGroupRule: %+v\r\n", result)
	fmt.Println("securityGroupRule ID is:", result.ID)
	fmt.Println("securityGroupRule Direction is:", result.Direction)
	fmt.Println("securityGroupRule EtherType is:", result.EtherType)
	fmt.Println("securityGroupRule SecGroupID is:", result.SecGroupID)
	fmt.Println("securityGroupRule PortRangeMin is:", result.PortRangeMin)
	fmt.Println("securityGroupRule PortRangeMax is:", result.PortRangeMax)
	fmt.Println("securityGroupRule Protocol is:", result.Protocol)
	fmt.Println("securityGroupRule ProjectId is:", result.ProjectId)
	fmt.Println("securityGroupRule UpdatedAt is:", result.UpdatedAt)

	fmt.Println("Get success!")

}

// Create a SecurityGroupRule
func CreateSecurityGroupRule(sc *gophercloud.ServiceClient) {
	opts := rules.CreateOpts{
		Direction:  "ingress",
		EtherType:  "IPv4",
		SecGroupID: "xxxxxx",
	}
	result, err := rules.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securityGroupRule: %+v\r\n", result)
	fmt.Println("securityGroupRule ID is:", result.ID)
	fmt.Println("securityGroupRule Direction is:", result.Direction)
	fmt.Println("securityGroupRule EtherType is:", result.EtherType)
	fmt.Println("securityGroupRule SecGroupID is:", result.SecGroupID)
	fmt.Println("securityGroupRule PortRangeMin is:", result.PortRangeMin)
	fmt.Println("securityGroupRule PortRangeMax is:", result.PortRangeMax)
	fmt.Println("securityGroupRule Protocol is:", result.Protocol)
	fmt.Println("securityGroupRule ProjectId is:", result.ProjectId)
	fmt.Println("securityGroupRule UpdatedAt is:", result.UpdatedAt)

	fmt.Println("Create success!")

}

// Delete a SecurityGroupRule
func DeleteSecurityGroupRule(sc *gophercloud.ServiceClient) {
	err := rules.Delete(sc, "xxxxxx").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Delete success!")
}
