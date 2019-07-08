package main

import (
	"fmt"
	"os"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/imagedata"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
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
	provider, err := openstack.AuthenticatedClient(tokenOpts)
	if err != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", err)
		return
	}
	//初始化service client
	sc, clientErr := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{})

	if clientErr != nil {
		fmt.Println("Failed to get the NewImageServiceV2 client: ", clientErr)
		return
	}
	imageDataUpload(sc)
	imageCreate(sc)
	imageGet(sc)
	imageUpdate(sc)
	imageList(sc)
	imageDelete(sc)

}

func imageDataUpload(sc *gophercloud.ServiceClient) {

	//上传镜像文件
	imageData, err := os.Open("/opt/cirros.img")
	if err != nil {
		fmt.Printf("open file err: %s", err)
		return
	}
	defer imageData.Close()
	imageID := ""
	err = imagedata.Upload(sc, imageID, imageData).ExtractErr()

	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

}

func imageCreate(sc *gophercloud.ServiceClient) {
	//创建镜像元数据
	visibility := images.ImageVisibilityPrivate
	createOpts := images.CreateOpts{
		Name:            "go-sdk-test",
		Visibility:      &visibility,
		ContainerFormat: "bare",
		DiskFormat:      "raw",
		MinDisk:         40,
	}

	imageResult, err := images.Create(sc, createOpts).Extract()
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("image ID is", imageResult.ID)
	fmt.Println("image Status is", imageResult.Status)
	fmt.Println("image Name is", imageResult.Name)
	fmt.Println("image SizeBytes is", imageResult.SizeBytes)
}

func imageUpdate(sc *gophercloud.ServiceClient) {
	//更新指定镜像信息
	updateOpts := images.UpdateOpts{
		images.ReplaceImageName{
			NewName: "go-sdk-test-2",
		},
	}
	imageID := ""
	imageResult, err := images.Update(sc, imageID, updateOpts).Extract()
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("image ID is", imageResult.ID)
	fmt.Println("image Status is", imageResult.Status)
	fmt.Println("image Name is", imageResult.Name)
	fmt.Println("image SizeBytes is", imageResult.SizeBytes)
}

func imageGet(sc *gophercloud.ServiceClient) {
	//查询指定的镜像
	imageID := ""
	imageResult, err := images.Get(sc, imageID).Extract()

	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("image ID is", imageResult.ID)
	fmt.Println("image Status is", imageResult.Status)
	fmt.Println("image Name is", imageResult.Name)
	fmt.Println("image SizeBytes is", imageResult.SizeBytes)
}

func imageList(sc *gophercloud.ServiceClient) {

	//查询镜像列表
	listOpts := images.ListOpts{
		Visibility: images.ImageVisibilityPrivate,
		Owner:      "owner_id",
		Status:     "active",
		Marker:     "marker_id",
		SortKey:    "name",
		SortDir:    "asc",
		Limit:      2,
	}

	allPages, err := images.List(sc, listOpts).AllPages()
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	for _, imageResult := range allImages {
		fmt.Println("image ID is", imageResult.ID)
		fmt.Println("image Status is", imageResult.Status)
		fmt.Println("image Name is", imageResult.Name)
		fmt.Println("image SizeBytes is", imageResult.SizeBytes)
	}
}

func imageDelete(sc *gophercloud.ServiceClient) {
	//删除指定镜像
	imageID := ""
	err := images.Delete(sc, imageID).ExtractErr()
	if err != nil {
		fmt.Printf("err: %s", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
	}

	return
}
