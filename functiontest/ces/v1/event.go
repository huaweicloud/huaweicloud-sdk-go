package main
/*
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/events"

)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	sc, err := openstack.NewCESV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get ces client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	EventCreate(sc)
	fmt.Println("main end...")
}

func EventCreate(sc *gophercloud.ServiceClient) {
	opts := events.CreateOpts{
		{
			EventName:   "systemInvaded",
			EventSource: "financial.Sytem",
			Time:        time.Now().Unix() * 1000,
			Detail: events.EventItemDetail{
				Content: "The financial system was invaded",
				//GroupId:"rg15221211517051YWWkEnVd",
				//ResourceId:"1234567890sjgggad",
				ResourceName: "ecs001",
				EventLevel:   "Major",
				EventState:   "normal",
			},
		},
	}

	events, err := events.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(events, "", " ")
	fmt.Println(string(bytes))
}
*/