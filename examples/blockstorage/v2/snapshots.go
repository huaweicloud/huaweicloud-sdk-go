package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/snapshots"
)

func main() {

	fmt.Println("main start...")

	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{DomainID}",
	}

	provider, err_auth := openstack.AuthenticatedClient(opts)
	if err_auth != nil {
		fmt.Println("get provider client failed: ", err_auth)
		return
	}

	sc, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get BlockStorage v2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	SnapshotsListDetails(sc)
	SnapshotsCreate(sc)
	fmt.Println("main end...")
}

// list snapshot detail
func SnapshotsListDetails(sc *gophercloud.ServiceClient) {
	fmt.Println("start list snapshot detail...")
	allPages, err := snapshots.Detail(sc, snapshots.ListOpts{Limit: 5}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allData, err1 := snapshots.ExtractSnapshots(allPages)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	fmt.Println("list snapshot detail success!")
	for _, v := range allData.SnapshotsLinks {

		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}

	for _, v := range allData.Snapshots {
		b, _ := json.MarshalIndent(v, "", " ")
		fmt.Println(string(b))
	}
}

// create snapshot
func SnapshotsCreate(sc *gophercloud.ServiceClient) {
	fmt.Println("srart create snapshot...")
	opts := snapshots.CreateOpts{
		VolumeID: "xxx",
		Name:     "xxx",
	}

	resp, err := snapshots.Create(sc, opts).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("create snapshot success!")
	b, _ := json.MarshalIndent(resp, "", " ")
	fmt.Println(string(b))
}
