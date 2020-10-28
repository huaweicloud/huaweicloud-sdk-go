package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		DomainID:         "{domainID}",
	}

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	sc, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get IAM v3 failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	CreateProject(sc)
	ListProject(sc)
	ListProjectsByConditions(sc)
	ListProjectsForUser(sc)
	GetProjectDetail(sc)
	UpdateProjectsDetails(sc)

	fmt.Println("main end...")
}

// 创建项目
// Create a project
// POST /v3/projects
func CreateProject(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestCreateProject")
	opts := projects.CreateOpts{
		DomainID:    "",
		Name:        "",
		ParentID:    "",
		Description: "",
	}
	resp := projects.Create(sc, opts)
	result, err := resp.Extract()
	if err != nil {
		fmt.Println("create project failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestCreateProject")
}

// 查询用户可以访问的项目列表
// Listing the project accessible to user
// GET /v3/auth/projects
func ListProject(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestListProject")

	resp := projects.ListProjects(sc)
	result, err := resp.ExtractProjectList()
	if err != nil {
		fmt.Println("get project failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestListProject")
}

// 查询指定条件下的项目列表
// Query the project list on specified conditions
// GET /v3/projects
func ListProjectsByConditions(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestListProjectsByConditions")
	opts := projects.ListOpts{
		DomainID: "",
		Enabled:  nil,
		IsDomain: nil,
		Name:     "",
		ParentID: "",
		Page:     1,
		PerPage:  10,
	}
	resp := projects.ListProjectsByConditions(sc, opts)
	result, err := resp.ExtractProjectList()
	if err != nil {
		fmt.Println("list projects failed")
		return
	} else {
		fmt.Println("projects", result)
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestListProjectsByConditions")
}

// 查询指定用户的项目列表
// Query the project list by the user id
// GET /v3/users/{user_id}/projects
func ListProjectsForUser(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestListProjectsForUser")
	userID := ""
	resp := projects.ListProjectsForUser(sc, userID)
	result, err := resp.ExtractProjectList()
	if err != nil {
		fmt.Println("list projects failed")
		return
	} else {
		fmt.Println("projects", result)
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestListProjectsForUser")
}

// 查询项目详情
// Query the project detail
// GET /v3/projects/{project_id}
func GetProjectDetail(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetProjectDetail")
	projectID := ""
	resp := projects.GetProjectDetails(sc, projectID)
	result, err := resp.Extract()
	if err != nil {
		fmt.Println("get project details failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestGetProjectDetail")
}

// 修改项目
// Update thee project
// PATCH /v3/projects/{project_id}
func UpdateProjectsDetails(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestListProjectsForUser")
	projectID := ""
	opts := projects.UpdateOpts{
		Description: "",
		Name:        "",
	}
	resp := projects.Update(sc, projectID, opts)
	result, err := resp.ExtractProjectUpdated()
	if err != nil {
		fmt.Println("patch project details failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestListProjectsForUser")
}
