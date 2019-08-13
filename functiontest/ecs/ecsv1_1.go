package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudserversext"
	"strings"
	"time"
	"encoding/json"
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

	//v1.1
	// 设置计算服务的client
	sc, err := openstack.NewECSV1_1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ecs v1.1 client failed")
		fmt.Println(err.Error())
		return
	}

	TestCreate(sc)
	TestExtCreateServer(sc)
	TestPostPaidServerCreate(sc)
	TestPrePaidServerCreate(sc)
	TestGetJobResult(sc)
	fmt.Println("main end...")
}

func TestExtCreateServer(sc *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		{
			SubnetId: "9a56640e-5503-4b8d-8231-963fc59ff91c",
		},
	}

	rv := cloudservers.RootVolume{
		VolumeType: "SATA",
	}

	dvs := []cloudservers.DataVolume{
		{
			VolumeType: "SATA",
			Size:       60,
		},
		{
			VolumeType: "SATA",
			Size:       70,
		},
	}

	opts := cloudservers.CreateOpts{
		Name:             "ecs_cloud_xx2",
		FlavorRef:        "c1.xlarge",
		ImageRef:         "2a50f694-b8e7-4a7a-8a51-0ff7f83d1345",
		VpcId:            "1d79d5ce-bc4c-48c6-88cd-4a8619f6ad2c",
		Nics:             nics,
		RootVolume:       rv,
		DataVolumes:      dvs,
		AvailabilityZone: "az1.dc1",
	}

	jobId, orderId, err := cloudserversext.CreateServer(sc, opts)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("create cloudservers success!")
	fmt.Println("jobId:", jobId)
	fmt.Println("orderId:", orderId)
}

//TestPostPaidServerCreate tests creating post paid server
func TestPostPaidServerCreate(client *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		{
			SubnetId: "a050d1ef-ff89-467a-a10c-f679d300c844",
		},
	}
	rv := cloudservers.RootVolume{
		VolumeType: "SAS",
		Size:       40,
	}
	opts := cloudservers.CreateOpts{
		Name:             "PostPaidServer",
		FlavorRef:        "s6.small.1",
		ImageRef:         "67f433d8-ed0e-4321-a8a2-a71838539e09",
		VpcId:            "fc7f54f1-703e-4f71-ab5c-e755f96b3c0c",
		Nics:             nics,
		RootVolume:       rv,
		AvailabilityZone: "cn-north-1a",
		Count: 2,
	}
	resp,createErr := cloudservers.CreateServer(client, opts)
	if createErr != nil {
		fmt.Println("createErr:", createErr)
		if ue, ok := createErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)

	var jobRst cloudservers.JobResult
	for {
		time.Sleep(time.Duration(10)*time.Second)
		job, getJobErr := cloudservers.GetJobResult(client, resp.Job.Id)
		if getJobErr != nil {
			fmt.Println("getJobResultErr:", getJobErr)
			if ue, ok := getJobErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return
		}
		jsJob, _ := json.MarshalIndent(job, "", "   ")
		fmt.Println(string(jsJob))

		if strings.Compare("SUCCESS", job.Status) == 0 {
			jobRst = job
			fmt.Println("Create server success!")
			break
		} else if strings.Compare("FAIL", job.Status) == 0 {
			jobRst = job
			fmt.Println("Create server failed!")
			break
		}
	}
	subJobs := jobRst.Entities.SubJobs
	var successServers []string
	var failServers []string
	for _, value := range subJobs {
		if strings.Compare("SUCCESS", value.Status) == 0 {
			successServers = append(successServers, value.Entities.ServerId)
		} else {
			failServers = append(failServers, value.Entities.ServerId)
		}
	}
	fmt.Println("jobId is ", resp.Job.Id)
	fmt.Println("successServers is ", successServers)
	fmt.Println("failServers is ", failServers)
}

//TestPrePaidServerCreate tests creating prepaid server
func TestPrePaidServerCreate(client *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		{
			SubnetId: "47822b98-c5bd-45b3-9454-d9773380f249",
		},
	}
	rv := cloudservers.RootVolume{
		VolumeType: "SAS",
		Size:       40,
	}
	extendParam := &cloudservers.ServerExtendParam{
		ChargingMode:     "prePaid",
		PeriodType:       "month",
		PeriodNum:        1,
		IsAutoPay:        "true",
	}
	opts := cloudservers.CreateOpts{
		Name:             "PrePaidServer",
		FlavorRef:        "c1.medium",
		ImageRef:         "4c4f6f5d-b198-44bf-96a5-f5e8a44bda37",
		VpcId:            "00a2e4fe-5295-4c0e-8bfd-ad3b8426cb93",
		Nics:             nics,
		RootVolume:       rv,
		AvailabilityZone: "AZ1",
		ExtendParam:      extendParam,
	}
	entity, createErr := cloudservers.CreateServer(client, opts)
	if createErr != nil {
		fmt.Println("createErr:", createErr)
		if ue, ok := createErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("orderId is ", entity.Order.Id)
	fmt.Println("serverIds is ", entity.Server.IDs)
}

func TestCreate(sc *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		{
			SubnetId: "9a56640e-5503-4b8d-8231-963fc59ff91c",
		},
	}

	rv := cloudservers.RootVolume{
		VolumeType: "SATA",
	}

	dvs := []cloudservers.DataVolume{
		{
			VolumeType: "SATA",
			Size:       60,
		},
		{
			VolumeType: "SATA",
			Size:       70,
		},
	}

	securityGroups := []cloudservers.SecurityGroup{
		{
			ID: "45678904567890456789HUdd",
		},
		{
			ID: "11111111111110456789HUdd",
		},
	}

	//isAR := false
	opts := cloudservers.CreateOpts{
		Name:             "ecs_cloud_xx1",
		FlavorRef:        "c1.xlarge",
		ImageRef:         "2a50f694-b8e7-4a7a-8a51-0ff7f83d1345",
		VpcId:            "1d79d5ce-bc4c-48c6-88cd-4a8619f6ad2c",
		Nics:             nics,
		RootVolume:       rv,
		DataVolumes:      dvs,
		AvailabilityZone: "az1.dc1",
		SecurityGroups:   securityGroups,
		//IsAutoRename:     &isAR,
	}

	jobId, orderId, err := cloudservers.Create(sc, opts)
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("create cloudservers success!")
	fmt.Println("jobId:", jobId)
	fmt.Println("orderId:", orderId)
}

func TestGetJobResult(sc *gophercloud.ServiceClient) {
	jr, err := cloudservers.GetJobResult(sc, "ff808082643fe15f01644548abd60cb8")
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Get Job Result success!")
	fmt.Println("jr:", jr)
}
