package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ims/v2/cloudimages"
)

func main() {

	fmt.Println("main start...")

	//provider, err := common.AuthAKSK()
	provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	//获取ims服务客户端
	sc, err := openstack.NewIMSV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ims client failed")
		fmt.Println(err.Error())
		return
	}

	TestList(sc)
	//TestCreateByFile(sc)
	//TestCreateByServer(sc)
	//TestGetJobResult(sc)

	fmt.Println("main end...")
}

func TestList(sc *gophercloud.ServiceClient) {
	allPages, err := cloudimages.List(sc, cloudimages.ListOpts{Status: "queued"}).AllPages()
	if err != nil {
		fmt.Println("err:", err)
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	allImages, err := cloudimages.ExtractImages(allPages)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("len:", len(allImages))

	fmt.Println("File:", allImages[0].File)
	fmt.Println("Owner:", allImages[0].Owner)
	fmt.Println("ID:", allImages[0].ID)
	fmt.Println("Size:", allImages[0].Size)
	fmt.Println("Self:", allImages[0].Self)
	fmt.Println("Schema:", allImages[0].Schema)
	fmt.Println("Status:", allImages[0].Status)
	fmt.Println("Tags:", allImages[0].Tags)
	fmt.Println("Visibility:", allImages[0].Visibility)
	fmt.Println("Name:", allImages[0].Name)
	fmt.Println("Checksum:", allImages[0].Checksum)
	fmt.Println("Deleted:", allImages[0].Deleted)
	fmt.Println("Protected:", allImages[0].Protected)
	fmt.Println("ContainerFormat:", allImages[0].ContainerFormat)
	fmt.Println("MinRam:", allImages[0].MinRam)
	fmt.Println("UpdatedAt:", allImages[0].UpdatedAt)
	fmt.Println("OsBit:", allImages[0].OsBit)
	fmt.Println("OsVersion:", allImages[0].OsVersion)
	fmt.Println("Description:", allImages[0].Description)
	fmt.Println("DiskFormat:", allImages[0].DiskFormat)
	fmt.Println("Isregistered:", allImages[0].Isregistered)
	fmt.Println("Platform:", allImages[0].Platform)
	fmt.Println("OsType:", allImages[0].OsType)
	fmt.Println("MinDisk:", allImages[0].MinDisk)
	fmt.Println("VirtualEnvType:", allImages[0].VirtualEnvType)
	fmt.Println("ImageSourceType:", allImages[0].ImageSourceType)
	fmt.Println("Imagetype:", allImages[0].Imagetype)
	fmt.Println("CreatedAt:", allImages[0].CreatedAt)
	fmt.Println("VirtualSize:", allImages[0].VirtualSize)
	fmt.Println("DeletedAt:", allImages[0].DeletedAt)
	fmt.Println("Originalimagename:", allImages[0].Originalimagename)
	fmt.Println("BackupID:", allImages[0].BackupID)
	fmt.Println("Productcode:", allImages[0].Productcode)
	fmt.Println("ImageSize:", allImages[0].ImageSize)
	fmt.Println("DataOrigin:", allImages[0].DataOrigin)
	fmt.Println("SupportKvm:", allImages[0].SupportKvm)
	fmt.Println("SupportXen:", allImages[0].SupportXen)
	fmt.Println("SupportDiskintensive:", allImages[0].SupportDiskintensive)
	fmt.Println("SupportHighperformance:", allImages[0].SupportHighperformance)
	fmt.Println("SupportXenGpuType:", allImages[0].SupportXenGpuType)
	fmt.Println("IsConfigInit:", allImages[0].IsConfigInit)
	fmt.Println("SystemSupportMarket:", allImages[0].SystemSupportMarket)
}

func TestCreateByFile(sc *gophercloud.ServiceClient) {
	createOpts := cloudimages.CreateByFileOpts{
		Name:         "ims_obs_xx4",
		Description:  "OBS文件制作镜像",
		ImageUrl:     "obs-xx2:Ubuntu14.04_Server_64bit_WebService_Picture.qcow2",
		OsVersion:    "Ubuntu 14.04 server 64bit",
		IsConfigInit: true,
		MinDisk:      40,
		IsConfig:     true,
		Tags: []string{
			"aaa.111",
			"bbb.333",
			"ccc.444",
		},
	}

	job, err := cloudimages.CreateImageByFile(sc, createOpts).ExtractJob()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	fmt.Println("create image by file success!")
	fmt.Println("jobId:", job.Id)
}

func TestCreateByServer(sc *gophercloud.ServiceClient) {
	createOpts := cloudimages.CreateByServerOpts{
		Name:        "image_from_ecs_xx1",
		Description: "云服务器制作镜像",
		InstanceId:  "2251f59c-b1ef-4398-bfa8-321782f670a5",
		Tags: []string{
			"aaa.111",
			"bbb.333",
			"ccc.444",
		},
	}

	job, err := cloudimages.CreateImageByServer(sc, createOpts).ExtractJob()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	fmt.Println("create image by server success!")
	fmt.Println("jobId:", job.Id)
}

func TestGetJobResult(sc *gophercloud.ServiceClient) {
	jr, err := cloudimages.GetJobResult(sc, "ff8080826446419a0164495a436941df").ExtractJobResult()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	fmt.Println("Id:", jr.Id)
	fmt.Println("Type:", jr.Type)
	fmt.Println("Status:", jr.Status)
	fmt.Println("BeginTime:", jr.BeginTime)
	fmt.Println("EndTime:", jr.EndTime)
	fmt.Println("ErrorCode:", jr.ErrorCode)
	fmt.Println("FailReason:", jr.FailReason)
	fmt.Println("Entities:", jr.Entities)
	fmt.Println("Entities.ImageId:", jr.Entities.ImageId)
}
