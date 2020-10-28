package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/iam/v3/agency"
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
	ListAgencies(sc)
	QueryAgencyDetails(sc)
	CreateAgency(sc)
	UpdateAgency(sc)
	DeleteAgency(sc)
	ListPermissionsForAgencyOnDomain(sc)
	ListPermissionsForAgencyOnProject(sc)
	GrantPermissionToAgencyOnDomain(sc)
	GrantPermissionToAgencyOnProject(sc)
	CheckPermissionForAgencyOnDomain(sc)
	CheckPermissionForAgencyOnProject(sc)
	RemovePermissionFromAgencyOnDomain(sc)
	RemovePermissionFromAgencyOnProject(sc)

	fmt.Println("main end...")
}

// 查询指定条件下的委托列表
// Query a agency list
// GET /v3.0/OS-AGENCY/agencies
func ListAgencies(client *gophercloud.ServiceClient) {
	opts := agency.AgenciesListOpts{
		DomainId:      "",
		TrustDomainId: "",
		Name:          "",
	}

	result, err := agency.List(client, opts).ExtractList()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Agencies is:", result.Agencies)
	for _, service := range result.Agencies {
		b, _ := json.MarshalIndent(service, "", " ")
		fmt.Println(string(b))
	}
	fmt.Println("Test list Agency success!")
}

// 查询委托详情
// Query the agency detail
// GET /v3.0/OS-AGENCY/agencies/{agency_id}
func QueryAgencyDetails(client *gophercloud.ServiceClient) {
	result, err := agency.Get(client, "").ExtractGet()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Agency is:", string(b))
	fmt.Println("Test query Agency success!")
}

// 创建委托
// Create a agency
// POST /v3.0/OS-AGENCY/agencies
func CreateAgency(client *gophercloud.ServiceClient) {
	opts := agency.CreateAgencyOpts{
		Name:            "",
		DomainID:        "",
		TrustDomainID:   "",
		Description:     "",
		Duration:        "",
		TrustDomainName: "",
	}
	result, err := agency.Create(client, opts).ExtractCreate()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Agency is:", string(b))
	fmt.Println("Test create Agency success!")
}

// 修改委托
// Update the agnecy
// PUT  /v3.0/OS-AGENCY/agencies/{agency_id}
func UpdateAgency(client *gophercloud.ServiceClient) {
	opts := agency.UpdateAgencyOpts{
		TrustDomainID:   "",
		Description:     "",
		Duration:        "",
		TrustDomainName: "",
	}
	result, err := agency.Update(client, "", opts).ExtractUpdate()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	b, _ := json.MarshalIndent(result, "", " ")
	fmt.Println("Agency is:", string(b))
	fmt.Println("Test update Agency success!")
}

// 删除委托
// Delete the agency
// DELETE /v3.0/OS-AGENCY/agencies/{agency_id}
func DeleteAgency(client *gophercloud.ServiceClient) {
	err := agency.Delete(client, "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test delete Agency success!")
}

// 查询全局服务中的委托权限
// Query the agency permissions in global services
// GET /v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles
func ListPermissionsForAgencyOnDomain(client *gophercloud.ServiceClient) {
	result, err := agency.ListPermissionsForAgencyOnDomain(client, "", "").ExtractListRolesDomain()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	for _, service := range result.Roles {
		b, _ := json.MarshalIndent(service, "", " ")
		fmt.Println(string(b))
	}
	fmt.Println("Test List Permission to Agency on Domain success!")
}

// 查询项目服务中的委托权限
// Query the agency permissions in the project
// GET /v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles
func ListPermissionsForAgencyOnProject(client *gophercloud.ServiceClient) {
	result, err := agency.ListPermissionsForAgencyOnProject(client, "", "").ExtractListRolesProject()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	for _, service := range result.Roles {
		b, _ := json.MarshalIndent(service, "", " ")
		fmt.Println(string(b))
	}
	fmt.Println("Test List Permission to Agency on Project success!")
}

// 为委托授予全局服务权限
// Grant the global service permissions to the agency
// PUT /v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}
func GrantPermissionToAgencyOnDomain(client *gophercloud.ServiceClient) {
	err := agency.GrantPermissionToAgencyOnDomain(client, "", "", "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test grant Permission to Agency on Domain success!")
}

// 为委托授予项目服务权限
// Grant the project service permissions to the agency
// PUT /v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}
func GrantPermissionToAgencyOnProject(client *gophercloud.ServiceClient) {
	err := agency.GrantPermissionToAgencyOnProject(client, "", "", "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test grant Permission to Agency on Project success!")
}

// 查询委托是否拥有全局服务权限
// Check whether the agency has the global service permission
// HEAD /v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}
func CheckPermissionForAgencyOnDomain(client *gophercloud.ServiceClient) {
	err := agency.CheckPermissionForAgencyOnDomain(client, "", "", "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test check Permission to Agency on Domain success!")
}

// 查询委托是否拥有项目服务权限
// Check whether the agnecy has the project service permission
// HEAD /v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}
func CheckPermissionForAgencyOnProject(client *gophercloud.ServiceClient) {
	err := agency.CheckPermissionForAgencyOnProject(client, "", "", "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test check Permission to Agency on Project success!")
}

// 移除委托的全局服务权限
// Remove the glabal service permission of the agency
// DELETE /v3.0/OS-AGENCY/domains/{domain_id}/agencies/{agency_id}/roles/{role_id}
func RemovePermissionFromAgencyOnDomain(client *gophercloud.ServiceClient) {
	err := agency.RemovePermissionFromAgencyOnDomain(client, "", "", "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test remove Permission to Agency on Domain success!")
}

// 移除委托的项目服务权限
// Remove the project service permissions of the agency
// PUT /v3.0/OS-AGENCY/projects/{project_id}/agencies/{agency_id}/roles/{role_id}
func RemovePermissionFromAgencyOnProject(client *gophercloud.ServiceClient) {
	err := agency.RemovePermissionFromAgencyOnProject(client, "", "", "").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test remove Permission to Agency on Project success!")
}
