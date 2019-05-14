package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/configures"
	"encoding/json"
)

func main() {
	provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("Error Code:", ue.ErrorCode())
			fmt.Println("Error Message:", ue.Message())
		}
	}
	sc, err := openstack.NewASV1(provider, gophercloud.EndpointOpts{})
	TestAsConfigCreate(sc)
	TestAsConfigList(sc)
	TestAsConfigGet(sc)
	TestDelAsConfig(sc)
	TestBatchDelAsConfig(sc)
}

func TestAsConfigCreate(client *gophercloud.ServiceClient) {

	disk := configures.Disk{Size: 40, VolumeType: "SATA", DiskType: "SYS"}
	instance := configures.CreateInstanceConfig{
		ImageRef:  "ce886742-d292-4e02-8c6b-66dd426ca248",
		FlavorRef: "s3.small.1",
		KeyName:   "KeyPair-53da",
		Disk:      []configures.Disk{disk},
		PublicIP: &configures.PublicIP{
			EIP: configures.EIP{
				IpType: "ADFA",
				Bandwidth: configures.BandwidthInfo{
					ShareType: "asdf",
				},
			},
		},
	}
	opts := configures.CreateOpts{ScalingConfigurationName: "TESTCONFIG", InstanceConfig: instance}

	result, err := configures.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("config: %+v\r\n", result)
	fmt.Println("Test Create success!")

}

func TestAsConfigList(client *gophercloud.ServiceClient) {
	opts := configures.ListOpts{
		Limit:       1,
		StartNumber: 1,
	}

	err := configures.List(client, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, err := configures.ExtractConfigs(page)
		if err != nil {
			return false, err
		}
		b, _ := json.MarshalIndent(resp, "", " ")
		fmt.Println(string(b))
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Println("Test LIST success!")

}

func TestAsConfigGet(client *gophercloud.ServiceClient) {
	result, err := configures.Get(client, "6ba04cba-8c6b-49df-86bc-ad126dbbf6fe").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Printf("config: %+v\r\n", result)
	fmt.Println("Test GET success!")

}

func TestDelAsConfig(client *gophercloud.ServiceClient) {
	err := configures.Delete(client, "12e42c36-51ba-488a-adfb-8aeab7b91530").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test DEL success!")
}

func TestBatchDelAsConfig(client *gophercloud.ServiceClient) {
	opts := configures.DeleteWithBatchOpts{
		ScalingConfigurationId: []string{"604aaf63-85b5-4059-877b-ebaf1b42eaf0"}}

	err := configures.DeleteWithBatch(client, opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test DEL BATCH success!")

}
