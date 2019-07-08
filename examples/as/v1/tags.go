package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/tags"
)

func main() {
	fmt.Println("main start...")
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}
	gophercloud.EnableDebug = true
	//初始化provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}
	//初始化服务 client
	sc, errClient := openstack.NewASV1(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewASV1 client: ", errClient)
		return
	}
	//开始测试
	ListResourceTag(sc)
	ListTag(sc)
	CreateTag(sc)
	DeleteTag(sc)
	ListTagInstance(sc)
	fmt.Println("main end...")
}

//--------------------------------------查询资源标签实例----------------------------------------
func ListTagInstance(client *gophercloud.ServiceClient) {
	matches := []tags.Tag{
		{
			Key:   "resource_name",
			Value: "as-group-",
		},
	}
	Tags := []tags.Tags{
		{
			Key:    "123",
			Values: []string{"456"},
		},
		{
			Key:    "22",
			Values: []string{"2256"},
		},
	}
	opts := tags.InstanceOpts{
		Action:     "filter",
		Offset:     "0",
		Limit:      "1000",
		Matches:    matches,
		NotTags:    nil,
		Tags:       Tags,
		NotTagsAny: nil,
		TagsAny:    nil,
	}
	result, err := tags.ListInstanceTags(client, "scaling_group_tag", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("List Tag Instance: %+v\r\n", result)
	fmt.Println("Test List Tag Instance success!")
}

//--------------------------------------删除标签----------------------------------------
func DeleteTag(client *gophercloud.ServiceClient) {
	ResourceType := "scaling_group_tag"
	tag1 := tags.Tag{
		Key:   "123",
		Value: "trtrtr",
	}
	tag2 := tags.Tag{
		Key:   "343",
		Value: "",
	}
	Tags := []tags.Tag{tag1, tag2}
	opts := tags.UpdateOpts{
		Tags:   Tags,
		Action: "delete",
	}
	err := tags.Update(client, ResourceType, "99df713f-3170-4c2d-9d47-efb861e18a42", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Delete Tag success!")
}

//--------------------------------------创建标签----------------------------------------
func CreateTag(client *gophercloud.ServiceClient) {
	ResourceType := "scaling_group_tag"
	tag1 := tags.Tag{
		Key:   "skjfe",
		Value: "fdksalee",
	}
	tag2 := tags.Tag{
		Key:   "123",
		Value: "erer567777",
	}
	Tags := []tags.Tag{tag1, tag2}
	opts := tags.UpdateOpts{
		Tags:   Tags,
		Action: "create",
	}
	err := tags.Update(client, ResourceType, "f9857b94-5c2d-41e1-9a29-c0077f965752", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Create Tag success!")
}

//--------------------------------------查询资源标签----------------------------------------
func ListResourceTag(client *gophercloud.ServiceClient) {
	ResourceType := "scaling_group_tag"
	result, err := tags.ListResourceTags(client, ResourceType, "99df713f-3170-4c2d-9d47-efb861e18a42").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("List Resource Tag: %+v\r\n", result)
	fmt.Println("Test List Resource Tag success!")
}

//--------------------------------------查询标签----------------------------------------
func ListTag(client *gophercloud.ServiceClient) {
	ResourceType := "scaling_group_tag"
	result, err := tags.ListTenantTags(client, ResourceType).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("List Tag: %+v\r\n", result)
	fmt.Println("Test List Tag success!")
}
