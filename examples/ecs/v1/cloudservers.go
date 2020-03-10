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
	"strings"
	"time"
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
	BatchStartServers(client)
	BatchRebootServers(client)
	BatchStopServers(client)
	BatchCreateServerTags(client)
	BatchDeleteServerTags(client)
	ListProjectTags(client)
	ListServerTags(client)

	// Batch Update Servers Name example
	BatchUpdateServersNameExample(tokenOpts)

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

//BatchStartServers requests to batch start servers.
func BatchStartServers(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchStartOpts{
		Servers: []cloudservers.Server{
			{ID: "ca5b7bdb-4f3b-494e-8563-018ca0b18c3d"},
			{ID: "4c8c776d-050c-4216-950e-c2807666d86c"},
		},
	}

	resp, err := cloudservers.BatchStart(sc, opts).ExtractJob()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("jobID:", resp.ID)
	var jobObj job.JobResult
	for {
		time.Sleep(10 * time.Second)
		jobRst, jobErr := job.GetJobResult(sc, resp.ID)
		if jobErr != nil {
			fmt.Println("getJobResultErr:", jobErr)
			if ue, ok := jobErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return
		}
		jsJob, _ := json.MarshalIndent(jobRst, "", "   ")
		fmt.Println(string(jsJob))

		if strings.Compare("SUCCESS", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers batch start is success!")
			break
		} else if strings.Compare("FAIL", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers batch start is failed!")
			break
		}
	}
	subJobs := jobObj.Entities.SubJobs
	var successServers []string
	var failServers []string
	for _, value := range subJobs {
		if strings.Compare("SUCCESS", value.Status) == 0 {
			successServers = append(successServers, value.Entities.ServerId)
		} else {
			failServers = append(failServers, value.Entities.ServerId)
		}
	}
	fmt.Println("successServers is ", successServers)
	fmt.Println("failServers is ", failServers)
}

//BatchRebootServers requests to batch reboot servers.
func BatchRebootServers(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchRebootOpts{
		Type: cloudservers.Type(cloudservers.Soft),
		Servers: []cloudservers.Server{
			{ID: "ca5b7bdb-4f3b-494e-8563-018ca0b18c3d"},
			{ID: "4c8c776d-050c-4216-950e-c2807666d86c"},
		},
	}

	resp, err := cloudservers.BatchReboot(sc, opts).ExtractJob()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("jobID:", resp.ID)
	var jobObj job.JobResult
	for {
		time.Sleep(10 * time.Second)
		jobRst, jobErr := job.GetJobResult(sc, resp.ID)
		if jobErr != nil {
			fmt.Println("getJobResultErr:", jobErr)
			if ue, ok := jobErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return
		}
		jsJob, _ := json.MarshalIndent(jobRst, "", "   ")
		fmt.Println(string(jsJob))

		if strings.Compare("SUCCESS", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers batch reboot is success!")
			break
		} else if strings.Compare("FAIL", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers batch reboot is failed!")
			break
		}
	}
	subJobs := jobObj.Entities.SubJobs
	var successServers []string
	var failServers []string
	for _, value := range subJobs {
		if strings.Compare("SUCCESS", value.Status) == 0 {
			successServers = append(successServers, value.Entities.ServerId)
		} else {
			failServers = append(failServers, value.Entities.ServerId)
		}
	}
	fmt.Println("successServers is ", successServers)
	fmt.Println("failServers is ", failServers)
}

//BatchStopServers requests to batch stop servers.
func BatchStopServers(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchStopOpts{
		Type: cloudservers.Type(cloudservers.Hard),
		Servers: []cloudservers.Server{
			{ID: "ca5b7bdb-4f3b-494e-8563-018ca0b18c3d"},
			{ID: "4c8c776d-050c-4216-950e-c2807666d86c"},
		},
	}

	resp, err := cloudservers.BatchStop(sc, opts).ExtractJob()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("jobID:", resp.ID)
	var jobObj job.JobResult
	for {
		time.Sleep(10 * time.Second)
		jobRst, jobErr := job.GetJobResult(sc, resp.ID)
		if jobErr != nil {
			fmt.Println("getJobResultErr:", jobErr)
			if ue, ok := jobErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return
		}
		jsJob, _ := json.MarshalIndent(jobRst, "", "   ")
		fmt.Println(string(jsJob))

		if strings.Compare("SUCCESS", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers batch stop is success!")
			break
		} else if strings.Compare("FAIL", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers batch stop is failed!")
			break
		}
	}
	subJobs := jobObj.Entities.SubJobs
	var successServers []string
	var failServers []string
	for _, value := range subJobs {
		if strings.Compare("SUCCESS", value.Status) == 0 {
			successServers = append(successServers, value.Entities.ServerId)
		} else {
			failServers = append(failServers, value.Entities.ServerId)
		}
	}
	fmt.Println("successServers is ", successServers)
	fmt.Println("failServers is ", failServers)
}

func BatchUpdateServersNameExample(tokenOpts token.TokenOptions) {

	// BatchUpdate supports a maximum of 1000 servers, which will take longer and requires timeout adjustments.
	conf := gophercloud.NewConfig().WithTimeout(time.Second * 30)

	//Init provider client
	provider, authErr := openstack.AuthenticatedClientWithOptions(tokenOpts, conf)
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

	opts := cloudservers.BatchUpdateOpts{
		Name: "test-name",
		Servers: []cloudservers.Server{
			{ID: "fafdaaf7-0b75-4208-b05c-269ed3a8f45a"},
			{ID: "cb92348b-e8d3-424a-8dc9-2a982177fb23"},
		},
	}
	resp, err := cloudservers.BatchUpdate(client, opts).ExtractBatchUpdate()
	if err != nil {
		if err1, ok := err.(*cloudservers.BatchOperateError); ok {
			fmt.Println("ErrorCode:", err1.ErrorCode())
			fmt.Println("Message:", err1.Message())
			fmt.Println("ErrorInfo:", err1.Error())
		}
		if err2, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrorCode:", err2.ErrorCode())
			fmt.Println("Message:", err2.Message())
			fmt.Println("ErrorInfo:", err2.Error())
		}
		return
	}
	servers := resp.Response
	for _, server := range servers {
		fmt.Println("the server update name success: ", server.ID)
	}
}

//BatchCreateServerTags requests to batch create server tags.
func BatchCreateServerTags(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchTagCreateOpts{
		Tags: []cloudservers.TagCreate{
			{Key: "key1", Value: "value1"},
		},
	}
	err := cloudservers.BatchCreateServerTags(sc, "f935d800-b801-4f05-829a-4688e2caaf06", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" BatchCreateServerTags success!")
}

//BatchDeleteServerTags requests to batch delete server tags.
func BatchDeleteServerTags(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchTagDeleteOpts{
		Tags: []cloudservers.TagDelete{
			{Key: "key1", Value: ""},
		},
	}
	err := cloudservers.BatchDeleteServerTags(sc, "f935d800-b801-4f05-829a-4688e2caaf06", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" BatchDeleteServerTags success!")
}

//ListProjectTags requests to query project tags.
func ListProjectTags(sc *gophercloud.ServiceClient) {
	resp, err := cloudservers.ListProjectTags(sc).Extract()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, jsErr := json.MarshalIndent(resp, "", " ")
	if jsErr != nil {
		fmt.Println(jsErr)
		return
	}
	fmt.Println("List project tags success!")
	fmt.Println(string(b))
}

//ListServerTags requests to query server tags.
func ListServerTags(sc *gophercloud.ServiceClient) {
	resp, err := cloudservers.ListServerTags(sc, "f935d800-b801-4f05-829a-4688e2caaf06").Extract()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	b, jsErr := json.MarshalIndent(resp, "", " ")
	if jsErr != nil {
		fmt.Println(jsErr)
		return
	}
	fmt.Println("List server tags success!")
	fmt.Println(string(b))
}
