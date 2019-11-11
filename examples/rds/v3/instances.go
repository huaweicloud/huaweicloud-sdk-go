package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/instances"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {
	fmt.Println("rds create instance test  start......")

	akskopts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}



	provider, authErr := openstack.AuthenticatedClient(akskopts)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		fmt.Println("Failed to get the provider: ", provider)
		return
	}

	client, clientErr := openstack.NewRDSV3(provider, gophercloud.EndpointOpts{Region:"xxx"})
	if clientErr != nil {
		fmt.Println("Failed to get the NewRDSV3 client: ", clientErr)
		return
	}

	RdsListDetailAllTest(client)
	ListErrorLogTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01")
	ListSlowLogTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01")
	CreateRdsInstanceTest(client)
	DeleteRdsInstanceTest(client)
	RestarRdsInstanceTest(client,"ab954e0bad034849abb7a2666041cc17in01")
	SingleToHaRdsInstanceTest(client,"9787515fe01746e192fa872e85ed61bein01")
	ResizeFlavorRdsInstanceTest(client,"9787515fe01746e192fa872e85ed61bein01")
	EnlargeVolumeRdsTest(client,"0477dd5f47c141809e4d0ea2ffde9655in01")

	fmt.Println("main end...")
}

func RdsListDetailAllTest(sc *gophercloud.ServiceClient) {
	opts := instances.ListRdsInstanceOpts{
		Limit:  10,
		Offset: 0,
		//Id:"2d452d4f2e344156b19bdbec4b32873ain01",
	}
	err := instances.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := instances.ExtractRdsInstances(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.Instances {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
			fmt.Println("Server id is :", v.Id)
			vpcID := v.VpcId
			fmt.Println("Server vpc id is :", vpcID)
		}
		// When returns false, current page of data will be returned.
		// Otherwise,when true,all pages of data will be returned.
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
}

func CreateRdsInstanceTest(client *gophercloud.ServiceClient) {

	InstancesStruct := instances.CreateRdsOpts{
		Name:           "GoT_Callback100_2_S-2-20190731-043903-fab4",
		Datastore:      instances.Datastore{Type: "MySQL", Version: "5.6"},
		BackupStrategy: &instances.BackupStrategy{StartTime: "06:15-07:15", KeepDays: 7},
		Ha:             &instances.Ha{Mode: "Ha", ReplicationMode: "semisync"},
		//FlavorRef : "rds.mysql.s1.medium",
		FlavorRef:        "rds.mysql.s1.medium.ha",
		Volume:           &instances.Volume{Type: "ULTRAHIGH", Size: 100},
		AvailabilityZone: "cn-north-4b,cn-north-4b",
		VpcId:            "3138ce3d-8837-49a6-b68a-4cdbc5b30a45",
		SubnetId:         "0f48e1d1-c244-422a-baa0-acfb1133c148",
		SecurityGroupId:  "702e9e18-34a2-4eda-a847-59546c3f5fa5",
		Password:         "{Password}",
		Port:             "{Port}",
		Region:           "{Region}",
	}
	jsonstr, err := json.Marshal(InstancesStruct)
	if err != nil {
		fmt.Println("error= ", err)
	}
	fmt.Println("jsonstr = ", string(jsonstr))

	job, createRdsInstanceErr := instances.Create(client, InstancesStruct).Extract()
	if createRdsInstanceErr != nil {
		fmt.Println("CreateRdsInstanceTest error:", createRdsInstanceErr)
		if ue, ok := createRdsInstanceErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(job)
	fmt.Println("servers CreateRdsInstanceTest success!")
}

func DeleteRdsInstanceTest(client *gophercloud.ServiceClient) {
	instancesId := instances.DeleteInstance{
		InstanceId: "7def68f0453b4ad586f4a529caf54ebano01",
	}
	resp, deleteRdsErr := instances.Delete(client, instancesId.InstanceId).Extract()
	if deleteRdsErr != nil {
		fmt.Println("deleteRdsErr:", deleteRdsErr)
		if ue, ok := deleteRdsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers deleteRds is success!")
}

func RestarRdsInstanceTest(client *gophercloud.ServiceClient, InstanceId string) {
	instancesIdRst := instances.RestartRdsInstanceOpts{

		Restart: " ",
	}
	resp, RestartRdsErr := instances.Restart(client, instancesIdRst, InstanceId).Extract()
	if RestartRdsErr != nil {
		fmt.Println("RestartRdsErr:", RestartRdsErr)
		if ue, ok := RestartRdsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers Restart Rds instance is success!")
}

func SingleToHaRdsInstanceTest(client *gophercloud.ServiceClient, InstanceId string) {
	singleToHaRdsBody := instances.SingleToHaRdsOpts{

		SingleToHa: &instances.SingleToHaRds{AzCodeNewNode: "cn-north-4b", Password: ""},
	}
	resp, SingleToHaRdsErr := instances.SingleToHa(client, singleToHaRdsBody, InstanceId).Extract()
	if SingleToHaRdsErr != nil {
		fmt.Println("SingleToHaRdsErr:", SingleToHaRdsErr)
		if ue, ok := SingleToHaRdsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers Single To Ha Rds instance is success!")
}

func ResizeFlavorRdsInstanceTest(client *gophercloud.ServiceClient, InstanceId string) {
	resizeFlavorBody := instances.ResizeFlavorOpts{
		ResizeFlavor: &instances.SpecCode{Speccode: "rds.mysql.s1.large"},
	}
	resp, ResizeFlavorRdsErr := instances.Resize(client, resizeFlavorBody, InstanceId).Extract()
	if ResizeFlavorRdsErr != nil {
		fmt.Println("ToResizeFlavorRds:", ResizeFlavorRdsErr)
		if ue, ok := ResizeFlavorRdsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers Resize Flavor Rds Instance Test  is success!")
}

func EnlargeVolumeRdsTest(client *gophercloud.ServiceClient, InstanceId string) {
	EnlargeVolumeBody := instances.EnlargeVolumeRdsOpts{
		EnlargeVolume: &instances.EnlargeVolumeSize{Size: 160},
	}
	resp, EnlargeVolumeRdsErr := instances.EnlargeVolume(client, EnlargeVolumeBody, InstanceId).Extract()
	if EnlargeVolumeRdsErr != nil {
		fmt.Println("EnlargeVolumeRdsErr:", EnlargeVolumeRdsErr)
		if ue, ok := EnlargeVolumeRdsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("Servers Enlarge Volume Test is success!")
}

func ListErrorLogTest(sc *gophercloud.ServiceClient, instanceID string) {
	opts := instances.DbErrorlogOpts{
		StartDate: "2019-08-06T10:41:14+0800",
		EndDate:   "2019-08-30T10:41:14+0800",
		Level:     "ALL",
		Limit:     "25", // Limit is set to 25 as default in Go SDK.If servers are too many, 200 is recommended.
		Offset:    "1",
	}
	err := instances.ListErrorLog(sc, opts, instanceID).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := instances.ExtractErrorLog(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.ErrorLogList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Database error log info is :", string(jsServer))

		}

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
}

func ListSlowLogTest(sc *gophercloud.ServiceClient, instanceID string) {
	opts := instances.DbSlowLogOpts{
		StartDate: "2019-09-30T10:41:14+0800",
		EndDate:   "2019-10-12T10:41:14+0800",
		Type:     "INSERT",
		Limit:     "10",
		Offset:    "1",
	}
	err := instances.ListSlowLog(sc, opts, instanceID).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := instances.ExtractSlowLog(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}
		for _, v := range resp.Slowloglist {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Database slow log is :", string(jsServer))

		}
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
}
