package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/quotasets"
	"github.com/gophercloud/gophercloud/auth/token"
)

func main() {
	fmt.Println("main start...")
	//Set authentication parameters
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "{Username}",
		Password:         "{Password}",
		DomainID:         "{DomainID}",
		ProjectID:        "{ProjectID}",
	}
	//Init provider client
	provider, authErr := openstack.AuthenticatedClient(tokenOpts)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		return
	}

	sc, computeV2Err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if computeV2Err != nil {
		fmt.Println("Get compute v2 client failed!", computeV2Err)
		if ue, ok := computeV2Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	projectID := "{projectID}"
	GetQuotaSetLimit(sc)
	GetQuotaSetDefault(sc, projectID)
	GetQuotaSet(sc, projectID)

	fmt.Println("main end...")
}

// GetQuotaSetLimit gets tenant quota limits
func GetQuotaSetLimit(sc *gophercloud.ServiceClient) {
	resp, err := quotasets.GetLimits(sc).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server get QuotaSetLimit success!")
	fmt.Println("maxServerMeta is ",resp.Absolute.MaxImageMeta)
	fmt.Println("MaxPersonality is ",resp.Absolute.MaxPersonality)
	fmt.Println("MaxPersonalitySize is ",resp.Absolute.MaxPersonalitySize)
	fmt.Println("MaxSecurityGroupRules is ",resp.Absolute.MaxSecurityGroupRules)
	fmt.Println("MaxSecurityGroups is ",resp.Absolute.MaxSecurityGroups)
	fmt.Println("MaxServerGroupMembers is ",resp.Absolute.MaxServerGroupMembers)
	fmt.Println("MaxServerGroups is ",resp.Absolute.MaxServerGroups)
	fmt.Println("MaxServerMeta is ",resp.Absolute.MaxServerMeta)
	fmt.Println("MaxTotalCores is ",resp.Absolute.MaxTotalCores)
	fmt.Println("MaxTotalFloatingIps is ",resp.Absolute.MaxTotalFloatingIps)
	fmt.Println("MaxTotalInstances is ",resp.Absolute.MaxTotalInstances)
	fmt.Println("MaxTotalKeypairs is ",resp.Absolute.MaxTotalKeypairs)
	fmt.Println("MaxTotalRAMSize is ",resp.Absolute.MaxTotalRAMSize)
	fmt.Println("TotalCoresUsed is ",resp.Absolute.TotalCoresUsed)
	fmt.Println("TotalFloatingIpsUsed is ",resp.Absolute.TotalFloatingIpsUsed)
	fmt.Println("TotalInstancesUsed is ",resp.Absolute.TotalInstancesUsed)
	fmt.Println("TotalRAMUsed is ",resp.Absolute.TotalRAMUsed)
	fmt.Println("TotalSecurityGroupsUsed is ",resp.Absolute.TotalSecurityGroupsUsed)
	fmt.Println("TotalServerGroupsUsed is ",resp.Absolute.TotalServerGroupsUsed)
}

// GetQuotaSet gets tenant quota
func GetQuotaSet(sc *gophercloud.ServiceClient, projectID string) {
	resp, err := quotasets.Get(sc, projectID).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server get QuotaSet success!")
	fmt.Println("ID is ",resp.ID)
	fmt.Println("Cores is ",resp.Cores)
	fmt.Println("FixedIPs is ",resp.FixedIPs)
	fmt.Println("FloatingIPs is ",resp.FloatingIPs)
	fmt.Println("InjectedFileContentBytes is ",resp.InjectedFileContentBytes)
	fmt.Println("InjectedFiles is ",resp.InjectedFiles)
	fmt.Println("Instances is ",resp.Instances)
	fmt.Println("KeyPairs is ",resp.KeyPairs)
	fmt.Println("MetadataItems is ",resp.MetadataItems)
	fmt.Println("RAM is ",resp.RAM)
	fmt.Println("SecurityGroupRules is ",resp.SecurityGroupRules)
	fmt.Println("SecurityGroups is ",resp.SecurityGroups)
	fmt.Println("InjectedFilePathBytes is ",resp.InjectedFilePathBytes)
	fmt.Println("ServerGroupMembers is ",resp.ServerGroupMembers)
	fmt.Println("ServerGroups is ",resp.ServerGroups)

}

// GetQuotaSetDefault gets default quota
func GetQuotaSetDefault(sc *gophercloud.ServiceClient, projectID string) {
	resp, err := quotasets.GetDefault(sc, projectID).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server get QuotaSet success!")
	fmt.Println("ID is ",resp.ID)
	fmt.Println("Cores is ",resp.Cores)
	fmt.Println("FixedIPs is ",resp.FixedIPs)
	fmt.Println("FloatingIPs is ",resp.FloatingIPs)
	fmt.Println("InjectedFileContentBytes is ",resp.InjectedFileContentBytes)
	fmt.Println("InjectedFiles is ",resp.InjectedFiles)
	fmt.Println("Instances is ",resp.Instances)
	fmt.Println("KeyPairs is ",resp.KeyPairs)
	fmt.Println("MetadataItems is ",resp.MetadataItems)
	fmt.Println("RAM is ",resp.RAM)
	fmt.Println("SecurityGroupRules is ",resp.SecurityGroupRules)
	fmt.Println("SecurityGroups is ",resp.SecurityGroups)
	fmt.Println("InjectedFilePathBytes is ",resp.InjectedFilePathBytes)
	fmt.Println("ServerGroupMembers is ",resp.ServerGroupMembers)
	fmt.Println("ServerGroups is ",resp.ServerGroups)

}
