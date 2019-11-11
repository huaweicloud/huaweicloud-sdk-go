package main

import (
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"encoding/json"
)

func main() {
	fmt.Println("main start...")
	//Set authentication parameters
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "{Username}",
		Password:         "{Password}",
		DomainID:         "{DomainID}",
		ProjectID:        "{ProjectID}",
	}
	//Init provider client
	provider, authErr := openstack.AuthenticatedClient(tokenOpts)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		return
	}
	//Init service client
	client, clientErr := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewComputeV2 client: ", clientErr)
		return
	}
	serverID := "{serverID}"
	attachmentID := "{attachmentID}"
	deleteFlag := 1
	AttachvolumesList(client, serverID)
	AttachvolumeCreate(client, serverID)
	AttachvolumeGet(client, serverID, attachmentID)
	AttachvolumeDelete(client, serverID, attachmentID)
	DetachVolumeWithDeleteFlag(client, serverID, attachmentID, deleteFlag)
	fmt.Println("main end...")
}

//Query attachvolumes list
func AttachvolumesList(client *gophercloud.ServiceClient, serverID string) {
	// Query all volumeattach list information
	allPages, allPagesErr := volumeattach.List(client, serverID).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform volumeattach structure
	allVolumes, allVolumesErr := volumeattach.ExtractVolumeAttachments(allPages)
	if allVolumesErr != nil {
		fmt.Println("allVolumesErr:", allVolumesErr)
		if ue, ok := allVolumesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("attachvolumes list is :")
	for _, volume := range allVolumes {
		volumeJson, _ := json.MarshalIndent(volume, "", " ")
		fmt.Println(string(volumeJson))
	}
}

//Create attachvolume
func AttachvolumeCreate(client *gophercloud.ServiceClient, serverID string) {
	volumeAttachOptions := volumeattach.CreateOpts{
		Device: "/dev/sdb",
		VolumeID: "640c1f2d-69ad-4d8a-9da8-c4b9abf83469",
	}
	resp, volumeAttachmentErr := volumeattach.Create(client, serverID, volumeAttachOptions).Extract()
	if volumeAttachmentErr != nil {
		fmt.Println("volumeAttachmentErr:", volumeAttachmentErr)
		if ue, ok := volumeAttachmentErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("attachvolume create success!")
}

//Get detail of the specified attachvolume
func AttachvolumeGet(client *gophercloud.ServiceClient, serverID string, attachmentId string) {
	volume, attachvolumesGetErr := volumeattach.Get(client, serverID,
		attachmentId).Extract()
	if attachvolumesGetErr != nil {
		fmt.Println("attachvolumesGetErr:", attachvolumesGetErr)
		if ue, ok := attachvolumesGetErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	volumeJson, _ := json.MarshalIndent(volume, "", " ")
	fmt.Println("attachvolume detail is " + string(volumeJson))
}

//AttachvolumeDelete requests detachment of volume on server
func AttachvolumeDelete(client *gophercloud.ServiceClient, serverID string, attachmentID string) {
	attachVolumesDetachErr := volumeattach.Delete(client, serverID,
		attachmentID).ExtractErr()
	if attachVolumesDetachErr != nil {
		fmt.Println("attachVolumesDetachErr:", attachVolumesDetachErr)
		if ue, ok := attachVolumesDetachErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Server detach volume success!")
}

//DetachVolumeWithDeleteFlag requests enforce detachment of volume with delete flag 1
func DetachVolumeWithDeleteFlag(sc *gophercloud.ServiceClient, serverID string, volumeID string, deleteFlag int) {
	err := volumeattach.DeleteWithFlag(sc, serverID, volumeID, deleteFlag).ExtractErr()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server detach volume with delete flag success!")

}
