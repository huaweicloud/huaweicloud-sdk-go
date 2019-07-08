package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/ims/v2/cloudimages"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
)

func main() {
	//设置认证参数
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "{Username}",
		Password:         "{Password}",
		DomainID:         "{DomainID}",
		ProjectID:        "{ProjectID}",
	}
	//初始化provider client
	provider, authErr := openstack.AuthenticatedClient(tokenOpts)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		return
	}
	//初始化service client
	sc, clientErr := openstack.NewIMSV2(provider, gophercloud.EndpointOpts{})

	if clientErr != nil {
		fmt.Println("Failed to get the NewIMSV2 client: ", clientErr)
		return
	}

	createImageByServer(sc)
	createImageByFile(sc)
	getJobStatus(sc, "{JobID}")
}

func createImageByServer(sc *gophercloud.ServiceClient) {
	//通过虚拟机创建镜像
	serverOpts := cloudimages.CreateByServerOpts{
		Name:        "go-sdk-test-1",
		Description: "test go",
		InstanceId:  "{InstanceId}",
	}

	jobInfo, err := cloudimages.CreateImageByServer(sc, serverOpts).ExtractJob()

	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("job ID is :", jobInfo.Id)
}

func createImageByFile(sc *gophercloud.ServiceClient) {
	//通过文件创建镜像
	fileOpts := cloudimages.CreateByFileOpts{
		Name:        "go-sdk-file-test",
		Description: "test go",
		ImageUrl:    "ims-huanan-image:Fedora.zvhd",
		IsConfig:    true,
		MinDisk:     40,
		CmkId:       "{CmkId}",
	}
	jobInfo, err := cloudimages.CreateImageByFile(sc, fileOpts).ExtractJob()
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("job ID is :", jobInfo.Id)
}

func getJobStatus(sc *gophercloud.ServiceClient, JobID string) {
	//查询JOB信息
	jobResult, err := cloudimages.GetJobResult(sc, JobID).ExtractJobResult()
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(jobResult)
}
