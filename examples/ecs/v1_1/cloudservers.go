package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	"go-sdk/openstack/ecs/v1/job"
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
	//Init  service client
	client, clientErr := openstack.NewECSV1_1(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewComputeV2 client: ", clientErr)
		return
	}
	ServerCreate(client)
	PostPaidServerCreate(client)
	PrePaidServerCreate(client)
	fmt.Println("main end...")

}

//Create a server (v1.1 version)
func ServerCreate(client *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		{
			SubnetId: "ddcc9be3-fcd3-433c-ac35-100c3208e97c",
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
		FlavorRef:        "s3.small.1",
		ImageRef:         "84663868-483c-4067-af7e-9d801e4a42f3",
		VpcId:            "4b95e7aa-b810-43cc-bc9f-8a92cf102100",
		Nics:             nics,
		RootVolume:       rv,
		DataVolumes:      dvs,
		AvailabilityZone: "br-iaas-odin1a",
	}
	jobId, orderId, createErr := cloudservers.Create(client, opts)
	if createErr != nil {
		fmt.Println("createErr:", createErr)
		if ue, ok := createErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	var jobObj job.JobResult
	for {
		time.Sleep(10 * time.Second)
		jobRst, jobErr := job.GetJobResult(client, jobId)
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
			fmt.Println("Servers create is success!")
			break
		} else if strings.Compare("FAIL", jobRst.Status) == 0 {
			jobObj = jobRst
			fmt.Println("Servers create is failed!")
			break
		}
	}
	fmt.Println("jobObj ", jobObj)
	fmt.Println("orderId is ", orderId)
}

//PostPaidServerCreate creates a postPaid server (v1.1 version)
func PostPaidServerCreate(client *gophercloud.ServiceClient) {
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
	}
	resp, createErr := cloudservers.CreateServer(client, opts)
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
		time.Sleep(time.Duration(10) * time.Second)
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

//PrePaidServerCreate creates a prepaid server (v1.1 version)
func PrePaidServerCreate(client *gophercloud.ServiceClient) {
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
		ChargingMode: "prePaid",
		PeriodType:   "month",
		PeriodNum:    1,
		IsAutoPay:    "true",
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
	fmt.Println("jobId is ", entity.Job.Id)
	fmt.Println("orderId is ", entity.Order.Id)
	fmt.Println("serverIds is ", entity.Server.IDs)
}
