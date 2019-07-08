package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
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

	//v2
	sc, err := openstack.NewECSV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ecs v2 client failed")
		fmt.Println(err.Error())
		return
	}

	//TestResetPassword(sc)
	//TestChangeOS(sc)
	TestReinstallOS(sc)
	//TestResizeFlavor(sc)

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

func TestReinstallOS(sc *gophercloud.ServiceClient) {
	serverID := "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a"
	opts := cloudserversV2.ReinstallOpts{
		AdminPass: "asdf",
		UserID:    "aef",
	}

	job, err := cloudserversV2.ReinstallOS(sc, serverID, opts).ExtractJob()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("jobID:", job.ID)

	fmt.Println(" TestReinstallOS success!")
}

//func TestResizeFlavor(sc *gophercloud.ServiceClient) {
//
//	opts := cloudserversV2.ResizeOpts{
//		Limit:        4,
//		InstanceUUID: "2e8c5857-45d2-4f92-bd1c-14fd815f5a5a",
//	}
//
//	data, err := cloudserversV2.ResizeFlavor(sc, opts).Extract()
//	if err != nil {
//		fmt.Println("err:", err)
//		if ue, ok := err.(*gophercloud.UnifiedError); ok {
//			fmt.Println("ErrCode:", ue.ErrorCode())
//			fmt.Println("Message:", ue.Message())
//		}
//		return
//	}
//
//	b, err := json.MarshalIndent(data, "", " ")
//	fmt.Println(string(b))
//
//	fmt.Println(" TestResizeFlavor success!")
//}
