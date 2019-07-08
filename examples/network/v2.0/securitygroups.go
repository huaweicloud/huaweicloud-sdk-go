package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
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
		fmt.Println("get network client failed")
		fmt.Println(err)
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
	UpdateSecurityGroup(sc)

	fmt.Println("main end...")
}

// List SecurityGroup
func ListSecurityGroup(sc *gophercloud.ServiceClient) {
	allPages, err := groups.List(sc, groups.ListOpts{}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	result, err := groups.ExtractGroups(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, resp := range result {
		fmt.Printf("securityGroup: %+v\r\n", resp)
		fmt.Println("securityGroup ID is:", resp.ID)
		fmt.Println("securityGroup Name is:", resp.Name)
		fmt.Println("securityGroup Description is:", resp.Description)
		fmt.Println("securityGroup TenantID is:", resp.TenantID)
		fmt.Println("securityGroup SecurityGroupRules is:", resp.Rules)
	}

	fmt.Println("List success!")

}

// Get a SecurityGroup
func GetSecurityGroup(sc *gophercloud.ServiceClient) {
	result, err := groups.Get(sc, "xxxxxx").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securityGroup: %+v\r\n", result)
	fmt.Println("securityGroup ID is:", result.ID)
	fmt.Println("securityGroup Name is:", result.Name)
	fmt.Println("securityGroup Description is:", result.Description)
	fmt.Println("securityGroup TenantID is:", result.TenantID)
	fmt.Println("securityGroup SecurityGroupRules is:", result.Rules)

	fmt.Println("Get success!")

}

// Update a SecurityGroup
func UpdateSecurityGroup(sc *gophercloud.ServiceClient) {
	opts := groups.UpdateOpts{
		Name:        "xxxxxx",
		Description: "xxxxxx",
	}

	result, err := groups.Update(sc, "xxxxxx", opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securityGroup: %+v\r\n", result)
	fmt.Println("securityGroup ID is:", result.ID)
	fmt.Println("securityGroup Name is:", result.Name)
	fmt.Println("securityGroup Description is:", result.Description)
	fmt.Println("securityGroup TenantID is:", result.TenantID)
	fmt.Println("securityGroup SecurityGroupRules is:", result.Rules)

	fmt.Println("Update success!")

}

// Create a SecurityGroup
func CreateSecurityGroup(sc *gophercloud.ServiceClient) {
	opts := groups.CreateOpts{
		Name:        "xxxxxx",
		Description: "xxxxxx",
	}

	result, err := groups.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("securityGroup: %+v\r\n", result)
	fmt.Println("securityGroup ID is:", result.ID)
	fmt.Println("securityGroup Name is:", result.Name)
	fmt.Println("securityGroup Description is:", result.Description)
	fmt.Println("securityGroup TenantID is:", result.TenantID)
	fmt.Println("securityGroup SecurityGroupRules is:", result.Rules)

	fmt.Println("Create success!")

}

// Delete a SecurityGroup
func DeleteSecurityGroup(sc *gophercloud.ServiceClient) {
	err := groups.Delete(sc, "xxxxxx").ExtractErr()
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
