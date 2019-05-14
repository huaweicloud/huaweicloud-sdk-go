package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudserversext"
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

	//TestCreate(sc)
	//TestExtCreateServer(sc)
	TestGetJobResult(sc)
	fmt.Println("main end...")
}

func TestExtCreateServer(sc *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		cloudservers.Nic{
			SubnetId: "9a56640e-5503-4b8d-8231-963fc59ff91c",
		},
	}

	rv := cloudservers.RootVolume{
		VolumeType: "SATA",
	}

	dvs := []cloudservers.DataVolume{
		cloudservers.DataVolume{
			VolumeType: "SATA",
			Size:       60,
		},
		cloudservers.DataVolume{
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

func TestCreate(sc *gophercloud.ServiceClient) {
	nics := []cloudservers.Nic{
		cloudservers.Nic{
			SubnetId: "9a56640e-5503-4b8d-8231-963fc59ff91c",
		},
	}

	rv := cloudservers.RootVolume{
		VolumeType: "SATA",
	}

	dvs := []cloudservers.DataVolume{
		cloudservers.DataVolume{
			VolumeType: "SATA",
			Size:       60,
		},
		cloudservers.DataVolume{
			VolumeType: "SATA",
			Size:       70,
		},
	}

	securityGroups := []cloudservers.SecurityGroup{
		cloudservers.SecurityGroup{
			ID: "45678904567890456789HUdd",
		},
		cloudservers.SecurityGroup{
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
