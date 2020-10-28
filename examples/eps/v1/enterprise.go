package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/eps/v1/enterpriseprojects"
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
	//创建eps服务的client
	epsClient, err := openstack.NewEPSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		// 异常处理
		panic(err)
	}
	createEP(epsClient)
	updateEP(epsClient,"4b5d9cb4-183b-4adb-bd36-cc5eb0213860")
	epList(epsClient)
	getEP(epsClient, "bd668a9e-059a-402d-8b83-b1647b5d7c04")
	EnableOrDisableEP(epsClient,"0e00f529-7d13-4d47-8c62-496a27ecb3c6")
	epQuotas(epsClient)
	filterEPResources(epsClient, "0e00f529-7d13-4d47-8c62-496a27ecb3c6")
	migrateEPResources(epsClient, "4b5d9cb4-183b-4adb-bd36-cc5eb0213860")
}
func epQuotas(sc *gophercloud.ServiceClient) {
	resp, err := enterpriseprojects.GetQuotas(sc).ExtractQuotas()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	for _, respData := range resp.Resources {
		fmt.Println("Type is:", respData.Type)
		fmt.Println("Used is:", respData.Used)
		fmt.Println("Quota is:", respData.Quota)
	}
	fmt.Println("GetQuotas success!")
}

func getEP(sc *gophercloud.ServiceClient, epID string) {

	resp, err := enterpriseprojects.Get(sc, epID).ExtractEP()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("id is:", resp.Id)
	fmt.Println("Name is:", resp.Name)
	fmt.Println("Description is:", resp.Description)
	fmt.Println("Status is:", resp.Status)
	fmt.Println("Create success!")
}

func EnableOrDisableEP(sc *gophercloud.ServiceClient, epID string) {

	opts := enterpriseprojects.ActionOpts{
		Action: "disable",
	}
	err := enterpriseprojects.EnableOrDisableEP(sc, opts, epID).ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Action success!")
}

func updateEP(sc *gophercloud.ServiceClient, epID string) {
	opts := enterpriseprojects.UpdateOpts{
		Name:        "create10",
		Description: "zzzzzzzzzzzz",
	}
	resp, err := enterpriseprojects.Update(sc, opts, epID).ExtractEP()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("id is:", resp.Id)
	fmt.Println("Name is:", resp.Name)
	fmt.Println("Description is:", resp.Description)
	fmt.Println("Status is:", resp.Status)
	fmt.Println("Update success!")
}

func createEP(sc *gophercloud.ServiceClient) {
	opts := enterpriseprojects.CreateOpts{
		Name:        "create10",
		Description: "ttttttttt",
	}
	resp, err := enterpriseprojects.Create(sc, opts).ExtractEP()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("id is:", resp.Id)
	fmt.Println("Name is:", resp.Name)
	fmt.Println("Description is:", resp.Description)
	fmt.Println("Status is:", resp.Status)
	fmt.Println("Create success!")
}
func epList(sc *gophercloud.ServiceClient) {
	opts := enterpriseprojects.ListOpts{
		//Id: "0",
		Name:   "test",
		Limit:  5,
		Offset: 0,
	}
	result, err := enterpriseprojects.List(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("########TotalCount##########: %+v\r\n", result.Total_count)
	for _, respData := range result.Enterprise_projects {
		fmt.Println("id is:", respData.Id)
		fmt.Println("Name is:", respData.Name)
		fmt.Println("Description is:", respData.Description)
		fmt.Println("Status is:", respData.Status)
	}
	fmt.Println("List success!")
}

func filterEPResources(sc *gophercloud.ServiceClient, epID string) {
	opts := enterpriseprojects.FilterResourcesOpts{
		Projects: [] string{
			"0605767f6f00d5762ff9c001c70e7359",
			"0596cf891180d3e72fa0c001fa4e20aa",
		},
		ResourceTypes: [] string{
			"ecs",
			"disk",
			"DNS_private_zone",
		},
		Offset: 0,
		Limit:  10,
	}
	resp, err := enterpriseprojects.FilterResources(sc, opts, epID).ExtractFilterResources()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("##########TotalCount is########### :", resp.Total_count)
	fmt.Println("##########Resource Start########### ")
	for _, respData := range resp.Resources {
		fmt.Println("ProjectId is:", respData.Project_id)
		fmt.Println("ProjectName is:", respData.Project_name)
		fmt.Println("ResourceType is:", respData.Resource_type)
		fmt.Println("ResourceName is:", respData.Resource_name)
	}
	fmt.Println("##########Resource End########### ")
	fmt.Println("##########Error Start########### ")
	for _, errResp := range resp.Errors {
		fmt.Println("ProjectId is:", errResp.Project_id)
		fmt.Println("ResourceType is:", errResp.Resource_type)
		fmt.Println("ErrorCode is:", errResp.Error_code)
		fmt.Println("ErrorMsg is:", errResp.Error_msg)
	}
	fmt.Println("##########Error End########### ")
	fmt.Println("FilterResources success!")
}

func migrateEPResources(sc *gophercloud.ServiceClient, epID string) {
	opts := enterpriseprojects.MigrateResourceOpts{
		Action:       "bind",
		ProjectId:    "0605767f6f00d5762ff9c001c70e7359", //当ResourceType不为全局资源时,必传
		ResourceType: "disk",
		ResourceId:   "379cfe53-e14b-4e39-8d76-efacc35306b1",
		RegionId:     "br-iaas-odin1", //OBS资源迁移时必传
		Associated:   false,
	}
	err := enterpriseprojects.MigrateResources(sc, opts, epID).ExtractErr()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("MigrateResource success!")
}