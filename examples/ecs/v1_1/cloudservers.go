package main

import (
	"github.com/gophercloud/gophercloud/openstack"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	"time"
	"encoding/json"
	"strings"
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
			SubnetId: "cc7953b3-110f-4e87-b240-ff4915548875",
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
		VpcId:            "b7ff7a9b-cc95-4dd0-b76a-f586c88e6556",
		Nics:             nics,
		RootVolume:       rv,
		DataVolumes:      dvs,
		AvailabilityZone: "az1.dc1",
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
	fmt.Println("jobId is ", jobId)
	fmt.Println("orderId is ", orderId)
	fmt.Println("server create success!")
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
	entity,createErr:=cloudservers.CreateServer(client, opts)
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