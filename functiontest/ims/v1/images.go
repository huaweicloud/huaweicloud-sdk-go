package main
/*
import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"encoding/json"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ims/v1/cloudimages"
)

func main() {

	fmt.Println("main start...")

	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}

	//获取ims服务客户端
	sc, err := openstack.NewIMSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ims v1 client failed")
		fmt.Println(err.Error())
		return
	}
	//OK
	//TestListImageTags(sc)
	//TestSetImageTag(sc)
	//TestGetImageGetQuota(sc)
	//TestAddImageMembers(sc)
	//TestUpdateImageMember(sc)
	//TestDeleteImageMembers(sc)
	//TestCopyImage(sc)

	//extarct job
	TestImportImage(sc)
	TestExportImage(sc)

	fmt.Println("main end...")
}

func TestListImageTags(sc *gophercloud.ServiceClient) {
	cloudimagesTag, err := cloudimages.ListImageTags(sc, cloudimages.ListImageTagsOpts{
		Limit: 2,
	}).Extract()
	if err != nil {
		fmt.Println("err:", err)
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	b, e := json.MarshalIndent(cloudimagesTag, "", "")

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(b))

	fmt.Println("Test TestListImageTags sucessfull")

}

func TestCopyImage(sc *gophercloud.ServiceClient) {
	imageid := "75261fea-e8bf-42b3-81ab-33954ea12d2f"
	copyOpts := cloudimages.CopyImageOpts{
		Name: "asdf",
	}
	tasks, err := cloudimages.CopyImage(sc, imageid, copyOpts).ExtractJob()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	p, _ := json.MarshalIndent(tasks, "", " ")
	fmt.Println(string(p))

	fmt.Println("Test TestCopyImage success!")
}

func TestAddImageMembers(sc *gophercloud.ServiceClient) {
	tasks, err := cloudimages.AddImageMembers(sc, cloudimages.AddImageMembersOpts{

		Images:   []string{"80597a42-1daf-476e-aac0-a2ef82a3a0b5"},
		Projects: []string{"054efa2069a64785a196efe56c05ee74"},
	}).ExtractJob()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	p, _ := json.MarshalIndent(tasks, "", " ")
	fmt.Println(string(p))

	fmt.Println("Test TestAddImageMembers success!")
}

func TestUpdateImageMember(sc *gophercloud.ServiceClient) {
	opts := cloudimages.UpdateImageMemberOpts{

		Images:    []string{"80597a42-1daf-476e-aac0-a2ef82a3a0b5"},
		ProjectID: "054efa2069a64785a196efe56c05ee74",
		Status:    "accepted",
	}

	image, err := cloudimages.UpdateImageMember(sc, opts).ExtractJob()
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}
	p, _ := json.MarshalIndent(image, "", " ")
	fmt.Println(string(p))

	fmt.Println("Test update TestUpdateImageMember success!")
}

func TestSetImageTag(sc *gophercloud.ServiceClient) {
	opts := cloudimages.SetImageTagOpts{
		ImageID: "80597a42-1daf-476e-aac0-a2ef82a3a0b5",
		Tag:     "sdfasdf.121212",
	}

	err := cloudimages.SetImageTag(sc, opts).Err
	if err != nil {
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	fmt.Println("Test TestSetImageTag success!")
}

func TestDeleteImageMembers(sc *gophercloud.ServiceClient) {
	opts := cloudimages.DeleteImageMembersOpts{

		Images:   []string{"80597a42-1daf-476e-aac0-a2ef82a3a0b5"},
		Projects: []string{"054efa2069a64785a196efe56c05ee74"},
	}
	imagetask, err := cloudimages.DeleteImageMembers(sc, opts).ExtractJob()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	p, _ := json.MarshalIndent(imagetask, "", " ")
	fmt.Println(string(p))
	fmt.Println("Test TestDeleteImageMembers success!")
}

func TestImportImage(sc *gophercloud.ServiceClient) {
	imageid := "4ee0ad06-346b-45d4-a4d4-754c4e5d593b"
	imageurl := "ims-image:centos7_5.qcow2"
	imagetask, err := cloudimages.ImportImage(sc, imageid, imageurl).ExtractJob()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	p, _ := json.MarshalIndent(imagetask, "", " ")
	fmt.Println(string(p))

	fmt.Println("Test TestImportImage success!")
}

func TestExportImage(sc *gophercloud.ServiceClient) {

	imageid := "8e4947ee-0507-45ee-bbb9-e7e3d5cdd869"
	bucketUrl := "obs-6662:centos7_5.qcow2"
	fileformat := "qcow2"
	imagetask, err := cloudimages.ExportImage(sc, imageid, bucketUrl, fileformat).ExtractJob()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	p, _ := json.MarshalIndent(imagetask, "", " ")
	fmt.Println(string(p))

	fmt.Println("Test TestExportImage success!")
}

func TestGetImageGetQuota(sc *gophercloud.ServiceClient) {
	cloudimageQuota, err := cloudimages.GetQuota(sc).Extract()
	if err != nil {
		fmt.Println("err:", err)
		if se, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", se.ErrorCode())
			fmt.Println("Message:", se.Message())
		}
		return
	}

	b, e := json.MarshalIndent(cloudimageQuota, "", "")

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(b))

	fmt.Println("Test TestGetImageGetQuota sucessfull")

}
*/