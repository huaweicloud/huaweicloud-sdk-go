package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/events"
)

func main() {
	fmt.Println("main start...")
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
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
			EventSource: "financial.System",
			Time:        time.Now().Unix() * 1000,
			Detail: events.EventItemDetail{
				Content: "The financial system was invaded",
				ResourceName: "ecs001",
				EventLevel:   "Major",
				EventState:   "normal",
			},
		},
	}

	createInfo, err := events.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	res, marshalErr := json.MarshalIndent(createInfo, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal createInfo error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}