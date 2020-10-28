package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/tms/v1/predefinetags"
)

func main() {
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.huaweicloud.com/v3",
		Username:         "xxxxxxxxxxx",
		Password:         "********",
		DomainID:         "yyyyyyyyyyyyyy",
	}
	provider, err := openstack.AuthenticatedClient(tokenOpts)
	if err != nil {
		fmt.Println("Failed to authenticate:", err)
	}
	//创建tms服务的client
	tmsClient, err := openstack.NewTMSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		// 异常处理
		panic(err)
	}
	predefineTagsList(tmsClient)
	createOrDeletePredefineTags(tmsClient)
	updatePredefineTags(tmsClient)
}

func createOrDeletePredefineTags(sc *gophercloud.ServiceClient) {
	opts := predefinetags.CreateOrDeleteOpts{
		Action: "create",//删除时为delete
		Tags:   []predefinetags.Tag {
			{
				Key: "test1",
				Value: "value1",
			},
		},
	}
	err := predefinetags.CreateOrDelete(sc, opts).ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Create success!")
}

func updatePredefineTags(sc *gophercloud.ServiceClient) {
	opts := predefinetags.UpdateOpts{
		OldTag: predefinetags.Tag {
				Key: "test1",
				Value: "value3",
			},
		NewTag: predefinetags.Tag {
			Key: "test1",
			Value: "value1",
		},
	}
	err := predefinetags.Update(sc, opts).ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Update success!")
}


func predefineTagsList(sc *gophercloud.ServiceClient) {
	opts := predefinetags.ListOpts{
		Limit: 0,
	}
	allPages, err := predefinetags.List(sc, opts).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	tagList, extractErr := predefinetags.ExtractTags(allPages)

	if extractErr != nil {
		fmt.Println("err1:", extractErr.Error())
		return
	}
	fmt.Printf("tags: %+v\r\n", tagList)
	for _, resp := range tagList {
		fmt.Println("Key is:", resp.Key)
		fmt.Println("Value is:", resp.Value)
		fmt.Println("Update_Time is:", resp.Update_Time)
	}
	fmt.Println("List success!")
}
