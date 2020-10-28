package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/iam/v3/custompolicy"
)

func main() {
	fmt.Println("main start...")
	// AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		DomainID:         "{domainID}",
	}

	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}
	// 初始化服务 client
	sc, errClient := openstack.NewIAMV3(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewIdentityV3 client: ", errClient)
		return
	}

	// 开始测试
	ListCustomPolicies(sc)
	QueryCustomPolicyDetails(sc)
	CreateCloudServiceCustomPolicy(sc)
	CreateAgencyCustomPolicy(sc)
	UpdateCloudServiceCustomPolicy(sc)
	UpdateAgencyCustomPolicy(sc)
	DeleteCustomPolicy(sc)
	fmt.Println("main end...")
}

// 查询自定义策略列表
// Query a custom policy list
// GET /v3.0/OS-ROLE/roles
func ListCustomPolicies(client *gophercloud.ServiceClient) {
	result, err := custompolicy.ListCustomPolicies(client).ExtractList()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Agencies is:", result.Roles)
	for _, service := range result.Roles {
		b, _ := json.MarshalIndent(service, "", " ")
		fmt.Println(string(b))
	}
	c, _ := json.MarshalIndent(result.Links, "", " ")
	fmt.Println(string(c))
	fmt.Println("Test List CustomPolicies success!")
}

// 查询自定义策略详情
// Query the custom policy detail
// GET /v3.0/OS-ROLE/roles/{role_id}
func QueryCustomPolicyDetails(client *gophercloud.ServiceClient) {
	result, err := custompolicy.QueryCustomPolicyDetails(client, "").ExtractQuery()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Role is:", result.Role)
	b, _ := json.MarshalIndent(result.Role, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test Query CustomPolicyDetails success!")
}

// 创建云服务自定义策略
// Create a cloud custom policy
// POST /v3.0/OS-ROLE/roles
func CreateCloudServiceCustomPolicy(client *gophercloud.ServiceClient) {
	condition := map[string]map[string][]string{
		"StringStartWith": map[string][]string{
			"g:ProjectName": []string{""},
		},
		"StringNotEquals": map[string][]string{
			"g:UserName": []string{""},
		},
	}

	statement := custompolicy.CloudServiceCustomPolicyStatement{
		Effect:    "",
		Action:    []string{},
		Condition: condition,
		Resource:  []string{},
	}
	opts := custompolicy.CreateCloudServiceCustomPolicyOpts{
		DisplayName:   "",
		Type:          "",
		Description:   "",
		DescriptionCn: "",
		Policy: custompolicy.CloudServiceCustomPolicy{
			Version:   "",
			Statement: []custompolicy.CloudServiceCustomPolicyStatement{statement},
		},
	}
	result, err := custompolicy.CreateCloudServiceCustomPolicy(client, opts).ExtractCreate()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Role is:", result.Role)
	b, _ := json.MarshalIndent(result.Role, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test Create CloudServiceCustomPolicy success!")
}

// 创建委托自定义策略
// Create a custom policy
// POST /v3.0/OS-ROLE/roles
func CreateAgencyCustomPolicy(client *gophercloud.ServiceClient) {
	statement := custompolicy.AgencyCustomPolicyStatement{
		Effect: "",
		Action: []string{},
		Resource: custompolicy.AgencyCustomPolicyStatementResource{
			Uri: []string{""},
		},
	}
	opts := custompolicy.CreateAgencyCustomPolicyOpts{
		DisplayName:   "",
		Type:          "",
		Description:   "",
		DescriptionCn: "",
		Policy: custompolicy.AgencyCustomPolicy{
			Version:   "",
			Statement: []custompolicy.AgencyCustomPolicyStatement{statement},
		},
	}
	result, err := custompolicy.CreateAgencyCustomPolicy(client, opts).ExtractCreate()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Role is:", result.Role)
	b, _ := json.MarshalIndent(result.Role, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test Create AgencyCustomPolicy success!")
}

// 修改云服务自定义策略
// Update the cloud custom policy
// PATCH /v3.0/OS-ROLE/roles/{role_id}
func UpdateCloudServiceCustomPolicy(client *gophercloud.ServiceClient) {
	condition := map[string]map[string][]string{
		"StringStartWith": map[string][]string{
			"g:ProjectName": []string{},
		},
		"StringNotEquals": map[string][]string{
			"g:UserName": []string{""},
		},
	}

	statement := custompolicy.CloudServiceCustomPolicyStatement{
		Effect:    "",
		Action:    []string{},
		Condition: condition,
		Resource:  []string{},
	}
	opts := custompolicy.UpdateCloudServiceCustomPolicyOpts{
		DisplayName:   "",
		Type:          "",
		Description:   "",
		DescriptionCn: "",
		Policy: custompolicy.CloudServiceCustomPolicy{
			Version:   "",
			Statement: []custompolicy.CloudServiceCustomPolicyStatement{statement},
		},
	}
	result, err := custompolicy.UpdateCloudServiceCustomPolicy(client, "", opts).ExtractPatch()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Role is:", result.Role)
	b, _ := json.MarshalIndent(result.Role, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test Update CloudServiceCustomPolicy success!")
}

// 修改委托自定义策略
// Update the custom policy
// PATCH /v3.0/OS-ROLE/roles/{role_id}
func UpdateAgencyCustomPolicy(client *gophercloud.ServiceClient) {
	statement := custompolicy.AgencyCustomPolicyStatement{
		Effect: "",
		Action: []string{},
		Resource: custompolicy.AgencyCustomPolicyStatementResource{
			Uri: []string{},
		},
	}
	opts := custompolicy.UpdateAgencyCustomPolicyOpts{
		DisplayName:   "",
		Type:          "",
		Description:   "",
		DescriptionCn: "",
		Policy: custompolicy.AgencyCustomPolicy{
			Version:   "",
			Statement: []custompolicy.AgencyCustomPolicyStatement{statement},
		},
	}
	result, err := custompolicy.UpdateAgencyCustomPolicy(client, "", opts).ExtractPatch()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Role is:", result.Role)
	b, _ := json.MarshalIndent(result.Role, "", " ")
	fmt.Println(string(b))
	fmt.Println("Test Update AgencyCustomPolicy success!")
}

// 删除自定义角色
// Delete the custom policy
// DELETE /v3.0/OS-ROLE/roles/{role_id}
func DeleteCustomPolicy(client *gophercloud.ServiceClient) {
	err := custompolicy.DeleteCustomPolicy(client, "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test delete CustomPolicy success!")
}
