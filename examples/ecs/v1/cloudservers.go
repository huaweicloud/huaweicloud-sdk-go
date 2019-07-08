package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/job"
	"github.com/gophercloud/gophercloud/pagination"
	"time"
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
	serverID := "{serverID}"
	ServerGetRecoveryStatus(client, serverID)
	ServerGetDetails(client, serverID)
	ServerBatchChangeOS(client)
	GetJobStatus(client)
	ConfigEcsAutoRecovery(client, serverID)
	ServersListDetailOnePage(client)
	ServersListDetailAllPages(client)
	fmt.Println("main end...")

}

//Query whether the server is configured with automatic recovery action
func ServerGetRecoveryStatus(client *gophercloud.ServiceClient, serverID string) {
	serverRecoveryStatus, getServerRecoveryStatusErr := cloudservers.GetServerRecoveryStatus(client, serverID).Extract()
	if getServerRecoveryStatusErr != nil {
		fmt.Println("getServerRecoveryStatusErr:", getServerRecoveryStatusErr)
		if ue, ok := getServerRecoveryStatusErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	serverRecoveryStatusJson, _ := json.MarshalIndent(serverRecoveryStatus, "", " ")
	fmt.Println("server recovery status is : " + string(serverRecoveryStatusJson))
}

//Query server details
func ServerGetDetails(client *gophercloud.ServiceClient, serverID string) {
	serverDetails, serverGetDetailsErr := cloudservers.Get(client, serverID).Extract()
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
		jobRst, jobErr := job.GetJobResult(client, resp.ID)
		if jobErr != nil {
			fmt.Println(jobErr.Error())
			return
		}

		jsJob, _ := json.MarshalIndent(jobRst, "", "   ")
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

//GetJobStatus is used for get job details via jobID
func GetJobStatus(sc *gophercloud.ServiceClient) {
	jobrs, err := job.GetJobResult(sc, "ff808082665cebfe016661ff5c8a47f1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsjobrs, _ := json.MarshalIndent(jobrs, "", "   ")
	fmt.Println(string(jsjobrs))
}

// ConfigEcsAutoRecovery manages automatic recovery of a server.
func ConfigEcsAutoRecovery(sc *gophercloud.ServiceClient, serverID string) {
	err := cloudservers.ConfigServerRecovery(sc, serverID, "false").ExtractErr()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("ConfigEcsAutoRecovery success!")

}

// ServersListDetailOnePage requests one page data of server list details by pagination.
func ServersListDetailOnePage(sc *gophercloud.ServiceClient) {
	opts := cloudservers.ListOpts{
		Limit:               1, // Limit is set to 25 as default in Go SDK.
		Offset:              1,
		Name:                "test",
		Flavor:              "s3.small.1",
		Status:              "SHUTOFF",
		Tags:                "testkey=testvalue",
		NotTags:             "now",
		EnterpriseProjectID: "0",
	}
	err := cloudservers.ListDetail(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := cloudservers.ExtractCloudServers(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		fmt.Println("Resp Count is :", resp.Count)
		for _, v := range resp.Servers {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
			fmt.Println("Server id is :", v.ID)
			vpcID, ok := v.Metadata["vpc_id"]
			if ok {
				fmt.Println("Server vpc id is :", vpcID)
			}
		}
		// When returns false, current page of data will be returned.
		// Otherwise,when true,all pages of data will be returned.
		return false, nil
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
}

// ServersListDetailAllPages requests all pages data of server list details by pagination.
func ServersListDetailAllPages(sc *gophercloud.ServiceClient) {
	opts := cloudservers.ListOpts{
		Limit:               1, // Limit is set to 25 as default in Go SDK.If servers are too many, 200 is recommended.
		Offset:              1,
		Name:                "test",
		Flavor:              "s3.small.1",
		Status:              "SHUTOFF",
		Tags:                "testkey=testvalue",
		NotTags:             "now",
		EnterpriseProjectID: "0",
	}
	err := cloudservers.ListDetail(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := cloudservers.ExtractCloudServers(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		fmt.Println("Resp Count is :", resp.Count)
		for _, v := range resp.Servers {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
			fmt.Println("Server id is :", v.ID)
			vpcID, ok := v.Metadata["vpc_id"]
			if ok {
				fmt.Println("Server vpc id is :", vpcID)
			}
		}
		// When returns false, current page of data will be returned.
		// Otherwise,when true,all pages of data will be returned.
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
}
