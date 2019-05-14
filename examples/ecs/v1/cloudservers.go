package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"encoding/json"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"time"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/job"
)

func main() {
	fmt.Println("main start...")
	gophercloud.EnableDebug = true
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
	//Init service client
	client, clientErr := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewECSV1 client: ", clientErr)
		return
	}
	serverId := "{serverId}"
	//ServerGetRecoveryStatus(client, serverId)
	ServerGetDetails(client, serverId)
	ServerBatchChangeOS(client)
	fmt.Println("main end...")

}

//Query whether the server is configured with automatic recovery action
//func ServerGetRecoveryStatus(client *gophercloud.ServiceClient, serverId string) {
//	serverRecoveryStatus, getServerRecoveryStatusErr := cloudservers.GetServerRecoveryStatus(client, serverId).Extract()
//	if getServerRecoveryStatusErr != nil {
//		fmt.Println("getServerRecoveryStatusErr:", getServerRecoveryStatusErr)
//		if ue, ok := getServerRecoveryStatusErr.(*gophercloud.UnifiedError); ok {
//			fmt.Println("ErrCode:", ue.ErrorCode())
//			fmt.Println("Message:", ue.Message())
//		}
//		return
//	}
//	serverRecoveryStatusJson, _ := json.MarshalIndent(serverRecoveryStatus, "", " ")
//	fmt.Println("server recovery status is : " + string(serverRecoveryStatusJson))
//}

//Query server details
func ServerGetDetails(client *gophercloud.ServiceClient, serverId string) {
	serverDetails, serverGetDetailsErr := cloudservers.Get(client, serverId).Extract()
	if serverGetDetailsErr != nil {
		fmt.Println("serverGetDetailsErr:", serverGetDetailsErr)
		if ue, ok := serverGetDetailsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	serverDetailsJson, _ := json.MarshalIndent(serverDetails, "", " ")
	fmt.Println("server get details is : " + string(serverDetailsJson))
}

//ServerBatchChangeOS is used for servers batch change OS
func ServerBatchChangeOS(client *gophercloud.ServiceClient) {
	batchChangeOsOpts := cloudservers.BatchChangeOpts{
		KeyName: "{KeyName}",
		UserID:  "{UserID}",
		ImageID: "{ImageID}",
		Servers: []cloudservers.Server{
			{ID: "{serverID}"},
			{ID: "{serverID}"},
		},
		MetaData: &cloudservers.MetaData{
			UserData: "{UserData}",
		},
	}
	resp, batchChangeOSErr := cloudservers.BatchChangeOS(client, batchChangeOsOpts).ExtractJob()
	if batchChangeOSErr != nil {
		fmt.Println("batchChangeOSErr:", batchChangeOSErr)
		if ue, ok := batchChangeOSErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)

	for {
		time.Sleep(time.Duration(20) * time.Second)
		jobRst,jobErr := job.GetJobResult(client, resp.ID)
		if jobErr != nil{
			fmt.Println(jobErr.Error())
			return
		}

		jsJob,_ := json.MarshalIndent(jobRst,"","   ")
		fmt.Println(string(jsJob))

		if jobRst.Status == "SUCCESS" {
			fmt.Println("servers batch change OS is success!")
			break
		} else if jobRst.Status == "FAIL" {
			fmt.Println("servers batch change OS is failed!")
			break
		}
	}
}
