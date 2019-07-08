package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/configures"
)

func main() {
	fmt.Println("main start...")
	//AKSK 认证，初始化认证参数。
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}
	//初始化provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}
	//初始化服务 client
	sc, errClient := openstack.NewASV1(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewASV1 client: ", errClient)
		return
	}
	//开始测试
	TestCreateConfigWithAllPara(sc)
	TestCreateConfigWinWithAdmin(sc)
	TestCreateConfigWithInstance(sc)
	TestGetConfig(sc)
	TestListConfig(sc)
	TestDeleteConfig(sc)
	TestBatchDeleteConfig(sc)
	fmt.Println("main end...")
}

//---------------------------------------------创建伸缩配置--使用新模板--带所有参数--linux--秘钥------------------------------------------------
func TestCreateConfigWithAllPara(client *gophercloud.ServiceClient) {
	disk1 := configures.Disk{
		Size:       40,
		VolumeType: "SATA",
		DiskType:   "SYS",
	}
	disk2 := configures.Disk{
		Size:       40,
		VolumeType: "SATA",
		DiskType:   "DATA",
	}
	disk := []configures.Disk{disk1, disk2}
	personality1 := configures.Personality{
		Path:    "/etc/foo.txt",
		Content: "MTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMQ==",
	}
	personality2 := configures.Personality{
		Path:    "/etc/foo1.txt",
		Content: "MTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMQ==",
	}
	personality := []configures.Personality{personality1, personality2}
	var userData = "MjMyMzIzMjMyMzIzMjMyMzJmZGZnZGdnZmc="
	//安全组
	SecurityGroupID1 := configures.SecurityGroup{
		ID: "bb8fcedf-9f9e-4815-8372-e9031f685b79",
	}
	SecurityGroupID := []configures.SecurityGroup{SecurityGroupID1}
	//eip
	PublicIP := &configures.PublicIP{
		EIP: configures.EIP{
			IpType: "5_sbgp",
			Bandwidth: configures.BandwidthInfo{
				ID:           "",
				Size:         5,
				ShareType:    "PER",
				ChargingMode: "bandwidth",
			},
		},
	}
	opts := configures.CreateOpts{
		ScalingConfigurationName: "as-sdk-CreateConfigWithAllPara",
		InstanceConfig: configures.CreateInstanceConfig{
			InstanceId:      "",
			FlavorRef:       "s3.small.1",
			ImageRef:        "79bee4ee-0025-4645-b004-23d2a66f6eec",
			Disk:            disk,
			KeyName:         "KeyPair-f958",
			Personality:     personality,
			UserData:        userData,
			PublicIP:        PublicIP,
			SecurityGroups:  SecurityGroupID,
			ServerGroupID:   "",
			Tenancy:         "",
			DedicatedHostID: "",
			Metadata:        nil,
		},
	}
	result, err := configures.Create(client, opts).Extract()
	if err != nil{
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingConfigurationId is:", result.ScalingConfigurationId)
}

//---------------------------------------------创建伸缩配置--使用新模板--带所有参数---Windows--密码-----------------------------------------------
func TestCreateConfigWinWithAdmin(client *gophercloud.ServiceClient) {
	disk1 := configures.Disk{
		Size:       40,
		VolumeType: "SATA",
		DiskType:   "SYS",
	}
	disk2 := configures.Disk{
		Size:       100,
		VolumeType: "SATA",
		DiskType:   "DATA",
	}
	disk := []configures.Disk{disk1, disk2}
	personality1 := configures.Personality{
		Path:    "/etc/foo.txt",
		Content: "MTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMQ==",
	}
	personality2 := configures.Personality{
		Path:    "/etc/foo1.txt",
		Content: "MTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMQ==",
	}
	personality := []configures.Personality{personality1, personality2}
	var userData = "MjMyMzIzMjMyMzIzMjMyMzJmZGZnZGdnZmc="
	var value = "******"
	var metadata map[string]interface{}
	metadata = make(map[string]interface{})
	metadata["admin_pass"] = value
	opts := configures.CreateOpts{
		ScalingConfigurationName: "as-sdk-CreateConfigWithAllPara",
		InstanceConfig: configures.CreateInstanceConfig{
			FlavorRef:   "s3.small.1",
			ImageRef:    "79bee4ee-0025-4645-b004-23d2a66f6eec",
			Disk:        disk,
			Personality: personality,
			UserData:    userData,
			Metadata:    metadata,
		},
	}
	result, err := configures.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingConfigurationId is:", result.ScalingConfigurationId)
}

//---------------------------------------------创建伸缩配置--使用已有云服务器----windows----------------------------------------------
func TestCreateConfigWithInstance(client *gophercloud.ServiceClient) {
	disk1 := configures.Disk{
		Size:       40,
		VolumeType: "SATA",
		DiskType:   "SYS",
	}
	disk2 := configures.Disk{
		Size:       40,
		VolumeType: "SATA",
		DiskType:   "DATA",
	}
	disk := []configures.Disk{disk1, disk2}
	personality1 := configures.Personality{
		Path:    "foo",
		Content: "MTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMQ==",
	}
	personality2 := configures.Personality{
		Path:    "foo1",
		Content: "MTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMTExMQ==",
	}
	personality := []configures.Personality{personality1, personality2}
	var userData = "MjMyMzIzMjMyMzIzMjMyMzJmZGZnZGdnZmc="
	var value = "****"
	var metadata map[string]interface{}
	metadata = make(map[string]interface{})
	metadata["admin_pass"] = value
	PublicIP := &configures.PublicIP{
		EIP: configures.EIP{
			IpType: "5_sbgp",
			Bandwidth: configures.BandwidthInfo{
				ID:           "",
				Size:         5,
				ShareType:    "PER",
				ChargingMode: "bandwidth",
			},
		},
	}
	opts := configures.CreateOpts{
		ScalingConfigurationName: "as-sdk-CreateConfigWithAllPara",
		InstanceConfig: configures.CreateInstanceConfig{
			InstanceId:      "fed6b677-fc7d-4fdb-936e-a2ae43e654d2",
			FlavorRef:       "s3.small.1",
			ImageRef:        "79bee4ee-0025-4645-b004-23d2a66f6eec",
			Disk:            disk,
			KeyName:         "KeyPair-f958",
			Personality:     personality,
			UserData:        userData,
			PublicIP:        PublicIP,
			SecurityGroups:  nil,
			ServerGroupID:   "",
			Tenancy:         "",
			DedicatedHostID: "",
			Metadata:        metadata,
		},
	}
	result, err := configures.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingConfigurationId is:", result.ScalingConfigurationId)
}

//---------------------------------------------查询伸缩配置详情--------------------------------------------------
func TestGetConfig(client *gophercloud.ServiceClient) {

	result, err := configures.Get(client, "54c4d98b-3eeb-4d6c-b8a1-87bdb404d7ec").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Get Config detail:")
	fmt.Println("subnet ScalingConfigurationId is:", result.ScalingConfigurationId)
	fmt.Println("FlavorRef is:", result.InstanceConfig.FlavorRef)
	fmt.Println("ImageRef is:", result.InstanceConfig.ImageRef)
	fmt.Println("Disk is:", result.InstanceConfig.Disk)
	fmt.Println("ScalingConfigurationName is:", result.ScalingConfigurationName)
}

//---------------------------------------------查询伸缩配置列表--------------------------------------------------
func TestListConfig(client *gophercloud.ServiceClient) {
	opts := configures.ListOpts{
		ScalingConfigurationName: "as-sdk-",
		ImageId:                  "79bee4ee-0025-4645-b004-23d2a66f6eec",
		StartNumber:              0,
		Limit:                    3,
	}
	result, err := configures.List(client, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	resultList,err:= configures.ExtractConfigs(result)
	for _, resp := range resultList.ScalingConfigurations{
		fmt.Println("Config list:")
		fmt.Println("ScalingConfigurationId is:", resp.ScalingConfigurationId)
		fmt.Println("ScalingConfigurationName is:", resp.ScalingConfigurationName)
		fmt.Println("ImageRef is:", resp.InstanceConfig.ImageRef)
		fmt.Println("Disk is:", resp.InstanceConfig.Disk)
		fmt.Println("FlavorRef is:", resp.InstanceConfig.FlavorRef)
		fmt.Println("KeyName is:", resp.InstanceConfig.KeyName)
		fmt.Println("Personality is:", resp.InstanceConfig.Personality)
		fmt.Println("------------------------")
	}
}

//---------------------------------------------删除伸缩配置表--------------------------------------------------
func TestDeleteConfig(client *gophercloud.ServiceClient) {
	err := configures.Delete(client, "d724958f-1e3b-4ff2-9d9e-438120d36e6b").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test delete AS Config success!")
}

//---------------------------------------------批量删除伸缩配置表--------------------------------------------------
func TestBatchDeleteConfig(client *gophercloud.ServiceClient) {
	var Ids = []string{"26682044-c7ae-4cd7-b0be-9c85248fcb24","bedf5bd4-ed0e-4b38-80b8-3d364735e558"}
	opts := configures.DeleteWithBatchOpts{
		ScalingConfigurationId: Ids,
	}
	err := configures.DeleteWithBatch(client, opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test delete AS Configs success!")
}
