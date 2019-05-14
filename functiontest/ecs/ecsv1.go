package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudserversext"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/job"
	"encoding/json"
	"time"
)

func main() {
	fmt.Println("main start...")

	//provider, err := common.AuthToken()
	provider, err := common.AuthAKSK()
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
	TestBatchChangeOS(sc)
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

//func TestGetEcsAutoRecovery(sc *gophercloud.ServiceClient) {
//	//2c2cd6a9-c501-42a9-a679-53518e6757cc
//	resp, err := cloudservers.GetServerRecoveryStatus(sc, "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a").Extract()
//	if err != nil {
//		fmt.Println(err)
//	}
//	b, err := json.MarshalIndent(*resp, "", " ")
//
//	if err != nil {
//
//		fmt.Println(err)
//	}
//	fmt.Println(string(b))
//
//}
//
//func TestConfigEcsAutoRecovery(sc *gophercloud.ServiceClient) {
//	//2c2cd6a9-c501-42a9-a679-53518e6757cc
//	err := cloudservers.ConfigServerRecovery(sc, "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a", "true").ExtractErr()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(" TestConfigEcsAutoRecovery success!")
//}
//
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
		jobRst,jobErr := job.GetJobResult(sc, jobObj.ID)
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
