package main

import (
	"fmt"
	"encoding/json"

	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ecs/v1/job"
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
	TestGetJobStatus(sc)

	fmt.Println("main end...")
}

func TestGetJobStatus(sc *gophercloud.ServiceClient)  {
	jobrs,err := job.GetJobResult(sc,"ff808082665cebfe016661ff5c8a47f1")
	if err !=nil{
		fmt.Println(err.Error())
		return
	}

	jsjobrs,_ := json.MarshalIndent(jobrs,"","   ")
	fmt.Println(string(jsjobrs))
}