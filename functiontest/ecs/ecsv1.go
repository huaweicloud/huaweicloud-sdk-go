package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"encoding/json"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudserversext"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/job"
	"github.com/gophercloud/gophercloud/pagination"
	"strings"
	"time"
)

func main() {
	fmt.Println("main start...")

	provider, err := common.AuthToken()
	//provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	sc, err := openstack.NewECSV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get ecs v1 client failed")
		fmt.Println(err.Error())
		return
	}
	//TestGetEcs(sc)
	//TestGetEcsExtbyServerId(sc)
	//TestGetEcsExtbyOrderId(sc)
	//TestGetEcsAutoRecovery(sc)
	//TestConfigEcsAutoRecovery(sc)
	//TestAddServerOnMonitorList(sc)
	//TestBatchChangeOS(sc)
	//TestListDetailOnePage(sc)
	//TestListDetailAllPages(sc)
	TestBatchStartServers(sc)
	TestBatchRebootServers(sc)
	TestBatchStopServers(sc)
	TestBatchUpdateServersName(sc)
	TestBatchCreateServerTags(sc)
	TestBatchDeleteServerTags(sc)
	TestListProjectTags(sc)
	TestListServerTags(sc)
	fmt.Println("main end...")

}

func TestGetEcs(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	resp, err := cloudservers.Get(sc, "d26b697b-3a74-4ec2-bd9d-5c3829f5d8a5").Extract()
	if err != nil {
		fmt.Println(err)
	}
	b, errr := json.MarshalIndent(*resp, "", " ")

	if errr != nil {

		fmt.Println(errr)
	}
	fmt.Println(string(b))

}

func TestGetEcsExtbyServerId(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	//95b23c71-0016-4f80-b160-7c1e0341d205
	resp, err := cloudserversext.GetServerExt(sc, "2544b973-ba5b-4cbd-a060-771ba4ec73e2")
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("CloudServer id is:", resp.CloudServer.ID)
	fmt.Println("CloudServer charging mode is:", resp.Charging.ChargingMode)
	volumeAttached, _ := json.MarshalIndent(resp.VolumeAttached, "", "    ")
	fmt.Println("CloudServer volume attached is:", string(volumeAttached))
}

func TestGetEcsExtbyOrderId(sc *gophercloud.ServiceClient) {
	resp, err := cloudserversext.GetPrepaidServerDetailByOrderId(sc, "CS1811091456QYTEX")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range resp {
		fmt.Println("CloudServer id is:", v.CloudServer.ID)
		fmt.Println("CloudServer charging mode is:", v.Charging.ChargingMode)
		volumeAttached, _ := json.MarshalIndent(v.VolumeAttached, "", "    ")
		fmt.Println("CloudServer volume attached is:", string(volumeAttached))
	}
}

func TestGetEcsAutoRecovery(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	resp, err := cloudservers.GetServerRecoveryStatus(sc, "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a").Extract()
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.MarshalIndent(*resp, "", " ")

	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(string(b))

}

func TestConfigEcsAutoRecovery(sc *gophercloud.ServiceClient) {
	//2c2cd6a9-c501-42a9-a679-53518e6757cc
	err := cloudservers.ConfigServerRecovery(sc, "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a", "true").ExtractErr()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" TestConfigEcsAutoRecovery success!")
}

//func TestAddServerOnMonitorList(sc *gophercloud.ServiceClient) {
//	//2c2cd6a9-c501-42a9-a679-53518e6757cc
//	err := cloudservers.AddServerOnMonitorList(sc, "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a").ExtractErr()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(" TestAddServerOnMonitorList success!")
//
//}

//TestBatchChangeOS tests the batch change OS function.
func TestBatchChangeOS(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchChangeOpts{
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

	jobObj, err := cloudservers.BatchChangeOS(sc, opts).ExtractJob()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("jobID:", jobObj.ID)

	for {
		time.Sleep(time.Duration(20) * time.Second)
		jobRst, jobErr := job.GetJobResult(sc, jobObj.ID)
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

// TestListDetailOnePage requests one page data of server list details by pagination.
func TestListDetailOnePage(sc *gophercloud.ServiceClient) {
	opts := cloudservers.ListOpts{
		Limit:               1,
		Offset:              1,
		Name:                "test",
		Flavor:              "s3.small.1",
		Status:              "SHUTOFF",
		Tags:                "onePage",
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

// TestListDetailAllPages requests all pages data of server list details by pagination.
func TestListDetailAllPages(sc *gophercloud.ServiceClient) {
	opts := cloudservers.ListOpts{
		Limit:   1,
		Offset:  1,
		Name:    "test",
		Flavor:  "s3.small.1",
		Status:  "SHUTOFF",
		Tags:    "testkey=testvalue",
		NotTags: "now",
		//ReservationID: "123",
		EnterpriseProjectID: "0",
	}
	page, err := cloudservers.ListDetail(sc, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, pageErr := cloudservers.ExtractCloudServers(page)
	if pageErr != nil {
		fmt.Println(pageErr)
		if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Resp Count is :", len(resp.Servers))
	for _, v := range resp.Servers {
		jsServer, _ := json.MarshalIndent(v, "", "   ")
		fmt.Println("Server info is :", string(jsServer))
	}

}

//TestBatchStartServers tests the batch start servers function.
func TestBatchStartServers(sc *gophercloud.ServiceClient) {
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

//TestBatchRebootServers tests the batch reboot servers function.
func TestBatchRebootServers(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchRebootOpts{
		Type: cloudservers.Type(cloudservers.Soft),
		Servers: []cloudservers.Server{
			{ID: "f51ba5c4-4ac4-4725-9965-4106773e0499"},
			{ID: "f935d800-b801-4f05-829a-4688e2caaf06"},
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

//TestBatchStopServers tests the batch stop servers function.
func TestBatchStopServers(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchStopOpts{
		Type: cloudservers.Type(cloudservers.Hard),
		Servers: []cloudservers.Server{
			{ID: "f51ba5c4-4ac4-4725-9965-4106773e0499"},
			{ID: "f935d800-b801-4f05-829a-4688e2caaf06"},
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

//BatchUpdateServersName requests to batch update servers name.
func TestBatchUpdateServersName(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchUpdateOpts{
		Name: "test-name",
		Servers: []cloudservers.Server{
			{ID: "5a2b0b54-f45f-4144-8ad8-6ccb131b9c57"},
			{ID: "b0a9d2b4-2cae-4b66-a6ba-6af70f3bd7f8"},
		},
	}

	resp, err := cloudservers.BatchUpdate(sc, opts).ExtractBatchUpdate()
	if err != nil {
		if err1, ok := err.(*cloudservers.BatchOperateError); ok {
			fmt.Println("ErrorCode:", err1.ErrorCode())
			fmt.Println("Message:", err1.Message())
			fmt.Println("ErrorInfo:", err1.Error())
		}
		return
	}
	servers := resp.Response
	for _, server := range servers {
		fmt.Println("the server update name success: ", server.ID)
	}
}

//TestBatchCreateServerTags tests the batch create server tags function.
func TestBatchCreateServerTags(sc *gophercloud.ServiceClient) {
	opts := cloudservers.BatchTagCreateOpts{
		Tags: []cloudservers.TagCreate{
			{Key: "key1", Value: "value1"},
		},
	}
	err := cloudservers.BatchCreateServerTags(sc, "dc69a241-f192-47d7-be31-cd43b1106a45", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" TestBatchCreateServerTags success!")
}

//TestBatchDeleteServerTags tests the batch delete server tags function.
func TestBatchDeleteServerTags(sc *gophercloud.ServiceClient) {
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
	fmt.Println(" TestBatchDeleteServerTags success!")
}

//TestListProjectTags tests the query project tags function.
func TestListProjectTags(sc *gophercloud.ServiceClient) {
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
	fmt.Println("TestListProjectTags success!")
	fmt.Println(string(b))
}

//TestListServerTags tests the query project tags function.
func TestListServerTags(sc *gophercloud.ServiceClient) {
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
	fmt.Println("TestListServerTags success!")
	fmt.Println(string(b))
}
