package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/backups"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	fmt.Println("rds backups test  start...")
	gophercloud.EnableDebug = true
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
	//Init service client
	client, clientErr := openstack.NewRDSV3(provider, gophercloud.EndpointOpts{Region:"xxx"})
	if clientErr != nil {
		fmt.Println("Failed to get the NewRDSV3 client: ", clientErr)
		return
	}

	GetAutoBackupsPolicyTest(client,"0477dd5f47c141809e4d0ea2ffde9655in01")
	ListBackupsFilesTest(client)
	CreateBackupsTest(client)

	AutoBackupsPolicyTest(client,"490fdcd8353441c5b463534b39290507in03")
	DeleteBackupsTest(client,"0477dd5f47c141809e4d0ea2ffde9655in01")
	ListBackupsTest(client,"0477dd5f47c141809e4d0ea2ffde9655in01")
	ListRestoreTimeTest(client,"9787515fe01746e192fa872e85ed61bein01")
	RestoreNewRdsTest(client)
	RecoveryTest(client)

}

func CreateBackupsTest(client *gophercloud.ServiceClient) {
	CreateBackupsBody := backups.CreateBackupsOpts{
		InstanceId:  "3705d1a019684147a8e2c7f87147d043in01",
		Name:        "gobackup",
		Description: "mannual backup",
	}
	resp, ToCreateBackupsErr := backups.Create(client, CreateBackupsBody).Extract()
	if ToCreateBackupsErr != nil {
		fmt.Println("ToCreateBackupsErr:", ToCreateBackupsErr)
		if ue, ok := ToCreateBackupsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("Servers CreateBackups  Test is success!")
}

func AutoBackupsPolicyTest(client *gophercloud.ServiceClient, instanceID string) {
	var a = 7
	AutoBackupsPolicyBody := backups.AutoBackupsPolicyOpts{
		BackupPolicy: &backups.BackupsPolicy{KeepDays: &a, StartTime: "08:15-09:15", Period: "2"},
	}

	AutoBackupsPolicyErr := backups.UpdatePolicy(client, AutoBackupsPolicyBody, instanceID).ExtractErr()
	if AutoBackupsPolicyErr != nil {
		fmt.Println("AutoBackupsPolicyErr:", AutoBackupsPolicyErr)
		if ue, ok := AutoBackupsPolicyErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Servers AutoBackupsPolicy  Test is success!")
}

func GetAutoBackupsPolicyTest(client *gophercloud.ServiceClient, instanceID string) {

	resp, GetAutoBackupsPolicyErr := backups.GetPolicy(client, instanceID).Extract()
	if GetAutoBackupsPolicyErr != nil {
		fmt.Println("GetAutoBackupsPolicyErr:", GetAutoBackupsPolicyErr)
		if ue, ok := GetAutoBackupsPolicyErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("Servers GetAutoBackupsPolicyTest  Test is success!")
}

func DeleteBackupsTest(client *gophercloud.ServiceClient, backupid string) {

	DeleteBackupsErr := backups.Delete(client, backupid).ExtractErr()
	if DeleteBackupsErr != nil {
		fmt.Println("DeleteBackupsErr:", DeleteBackupsErr)
		if ue, ok := DeleteBackupsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Servers DeleteBackupsTest  Test is success!")
}

func ListBackupsTest(sc *gophercloud.ServiceClient, InstanceId string) {
	opts := backups.ListBackupsOpts{
		InstanceId: InstanceId,
		Limit:      0,
		Offset:     0,
	}
	err := backups.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := backups.ExtractBackups(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.Backups {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
			fmt.Println("Server id is :", v.Id)
			Datastore := v.Datastore
			fmt.Println("Server backups  Datastore is :", Datastore)
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

func RestoreNewRdsTest(client *gophercloud.ServiceClient) {

	instancesTestStruct := backups.RestoreNewRdsOpts{
		Name:             "TestRestore_132_S-2-20190731-043903-fab4",
		BackupStrategy:   &backups.BackupStrategy{StartTime: "06:15-07:15", KeepDays: 7},
		FlavorRef:        "rds.mysql.s1.medium",
		Volume:           &backups.Volume{Type: "ULTRAHIGH", Size: 110},
		AvailabilityZone: "cn-north-4b",
		VpcId:            "3138ce3d-8837-49a6-b68a-4cdbc5b30a45",
		SubnetId:         "0f48e1d1-c244-422a-baa0-acfb1133c148",
		SecurityGroupId:  "702e9e18-34a2-4eda-a847-59546c3f5fa5",
		Password:         "{Your Password}",
		Port:             "{Your Port}",
		RestorePoint:     &backups.RestorePoint{InstanceId: "3705d1a019684147a8e2c7f87147d043in01", Type: "backup", BackupId: "126624c8d178426aa3916aea36471dc5br01"},
	}

	job, RestoreNewRdsErr := backups.Restore(client, instancesTestStruct).Extract()
	if RestoreNewRdsErr != nil {
		fmt.Println("RestoreNewRdsTest error:", RestoreNewRdsErr)
		if ue, ok := RestoreNewRdsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(job)
	fmt.Println("servers RestoreNewRdsTest success!")
}

func ListBackupsFilesTest(sc *gophercloud.ServiceClient) {
	opts := backups.ListBackupFilesOpts{
		BackupId: "983baba2112147e184ffadc3e509f162br01",
	}
	err := backups.ListFiles(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := backups.ExtractBackupsFiles(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}
		for _, v := range resp.FilesList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
			fmt.Println("Backupsfiles DownloadLink is :", v.DownloadLink)
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

func ListRestoreTimeTest(sc *gophercloud.ServiceClient, instanceId string) {
	opts := backups.ListRestoreTimeOpts{
		Date: "",
	}
	err := backups.ListRestoreTime(sc, opts, instanceId).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := backups.ExtractRestoreTime(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.RestoreTimeList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
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

func RecoveryTest(client *gophercloud.ServiceClient) {
	var source backups.Source
	source = backups.Source{InstanceId: "3705d1a019684147a8e2c7f87147d043in01", Type: "", BackupId: "126624c8d178426aa3916aea36471dc5br01", RestoreTime: 0, DatabaseName: nil}
	Body := backups.RecoveryOpts{
		Source: source,
		Target: backups.Target{InstanceId: "9787515fe01746e192fa872e85ed61bein01"},
	}
	resp, Err := backups.Recovery(client, Body).Extract()
	if Err != nil {
		fmt.Println("ToRecoveryTest Err:", Err)
		if ue, ok := Err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("Servers ToRecoveryTest  is success!")
}
