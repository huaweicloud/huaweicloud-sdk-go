package main

import (
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/ecs/v2/cloudservers"
)

func main() {
	fmt.Println("main start...")
	//Set authentication parameters
	akskOptions := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{DomainID}",
	}
	//Init provider client
	provider, authErr := openstack.AuthenticatedClient(akskOptions)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		return
	}
	//Init service client
	client, clientErr := openstack.NewECSV2(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewComputeV2 client: ", clientErr)
		return
	}
	serverId := "{serverId}"
	newPassword := "{newPassword}"
	ReinstallOS(client, serverId)
	ChangeOS(client, serverId)
	ResetPassword(client, serverId, newPassword)
	fmt.Println("main end...")

}

//Re-install server operating system (install Cloud-init)
func ReinstallOS(client *gophercloud.ServiceClient, serverId string) {
	reInstallOpts := cloudservers.ReinstallOpts{
		KeyName: "TestKey",
	}
	job, reinstallOSErr := cloudservers.ReinstallOS(client, serverId, reInstallOpts).ExtractJob()
	if reinstallOSErr != nil {
		fmt.Println("reinstallOSErr:", reinstallOSErr)
		if ue, ok := reinstallOSErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(job)
	fmt.Println("servers reinstall OS success!")
}

//Change server operating system (installing Cloud-init)
func ChangeOS(client *gophercloud.ServiceClient, serverId string) {
	changeOsOpts := cloudservers.ChangeOpts{
		ImageID: "2a50f694-b8e7-4a7a-8a51-0ff7f83d1345",
		KeyName: "TestKey",
	}
	resp, changeOSErr := cloudservers.ChangeOS(client, serverId, changeOsOpts).ExtractJob()
	if changeOSErr != nil {
		fmt.Println("changeOSErr:", changeOSErr)
		if ue, ok := changeOSErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers change OS is success!")
}

//One-click reset server password
func ResetPassword(client *gophercloud.ServiceClient, serverId string, newPassword string) {
	resetPasswordErr := cloudservers.ResetPassword(client, serverId, newPassword).ExtractErr()
	if resetPasswordErr != nil {
		fmt.Println("resetPasswordErr:", resetPasswordErr)
		if ue, ok := resetPasswordErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("server reset password success!")
}
