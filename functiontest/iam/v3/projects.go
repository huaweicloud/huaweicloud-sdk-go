package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/iam/v3/projects"
)

func main() {

	fmt.Println("main start...")

	provider, err_auth := common.AuthAKSK()
	//provider, err_auth := common.AuthToken()

	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}
	sc, err := openstack.NewIAMV3Ext(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get IAM v3-ext failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	UpdateProjectDetailAndStatus(sc)
	GetProjectDetailAndStatus(sc)

	fmt.Println("main end...")
}

// 查询项目详情与状态
// Query the project detail and status
// GET /v3-ext/projects/{project_id}
func GetProjectDetailAndStatus(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetProjectDetailAndStatus")
	projectID := ""
	resp := projects.GetProjectDetailsAndStatus(sc, projectID)
	result, err := resp.Extract()
	if err != nil {
		fmt.Println("get project details failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestGetProjectDetailAndStatus")
}

// 设置项目状态
// Update the project status
// PUT /v3-ext/projects/{project_id}
func UpdateProjectDetailAndStatus(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestUpdateProjectDetailAndStatus")
	projectID := ""
	opts := projects.UpdateOpts{
		Status: "normal",
	}
	resp := projects.Update(sc, projectID, opts)
	if resp.Err.Error() == "EOF" {
		fmt.Println("update project detail and status successfully")
	} else {
		fmt.Println("update project detail and status failed")
	}
	fmt.Println("finish TestUpdateProjectDetailAndStatus")
}
