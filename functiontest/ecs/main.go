package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudserversext"
	cloudserversV2 "github.com/gophercloud/gophercloud/openstack/ecs/v2/cloudservers"
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


	//v2
	//sc, err := openstack.NewECSV2(provider, gophercloud.EndpointOpts{})
	//if err != nil {
	//	fmt.Println("get ecs v2 client failed")
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	////TestResetPassword(sc)
	//TestChangeOS(sc)

	fmt.Println("main end...")
}

func TestChangeOS(sc *gophercloud.ServiceClient) {
	opts := cloudserversV2.ChangeOpts{
		AdminPass: "xxx",
		ImageID:   "7affa924-93bf-43c1-91e4-6234dbb822ae",

		//		KeyName: "kkkkkkkkkkkk",
		//		UserID:  "ertyuiopiiiiiiiiiiii",
		//		MetaData: cloudserversV2.MetaData{
		//			UserData: "uuuuuuuuuuu",
		//		},
	}

	serverID := "c00e66fa-e7e0-444f-aa8c-51e1559f40d2"
	job, err := cloudserversV2.ChangeOS(sc, serverID, opts).ExtractJob()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("jobID:", job.ID)
}

func TestResetPassword(sc *gophercloud.ServiceClient) {
	serverID := "1378c4be-8ca4-45bf-b1d9-8aaa3989c047"
	pwd := "xxx"
	err := cloudserversV2.ResetPassword(sc, serverID, pwd).ExtractErr()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("reset password success!")
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
