package testing

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/rds/v3/instances"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)
	opts := instances.ListRdsInstanceOpts{
		Limit:  10,
		Offset: 0,
		Id:     "ed7cc6166ec24360a5ed5c5c9c2ed726in01",
	}
	count := 0
	err := instances.List(client.ServiceClient(), opts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := instances.ExtractRdsInstances(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, "ed7cc6166ec24360a5ed5c5c9c2ed726in01", actual.Instances[0].Id)
		th.CheckDeepEquals(t, "mysql-0820-022709-01", actual.Instances[0].Name)
		th.CheckDeepEquals(t, "postPaid", actual.Instances[0].ChargeInfo.ChargeMode)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
func TestListAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	allPages, err := instances.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allRds, err := instances.ExtractRdsInstances(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allRds.Instances))
}
func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateSuccessfully(t)
	InstancesStruct := instances.CreateRdsOpts{
		Name:           "trove-instance-rep2",
		Datastore:      instances.Datastore{Type: "MySQL", Version: "5.6"},
		BackupStrategy: &instances.BackupStrategy{StartTime: "08:15-09:15", KeepDays: 12},
		FlavorRef:        "rds.mysql.s1.large",
		Volume:           &instances.Volume{Type: "ULTRAHIGH", Size: 100},
		AvailabilityZone: "cn-north-4",
		VpcId:            "490a4a08-ef4b-44c5-94be-3051ef9e4fce",
		SubnetId:         "0e2eda62-1d42-4d64-a9d1-4e9aa9cd994f",
		SecurityGroupId:  "2a1f7fc8-3307-42a7-aa6f-42c8b9b8f8c5",
		Password:         "YpurPassword",
		Port:             "8635",
		Region:           "cn-north-4",
	}
	result, err := instances.Create(client.ServiceClient(), InstancesStruct).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "trove-instance-rep2", result.Instance.Name)
	th.CheckEquals(t, "dsfae23fsfdsae3435in01", result.Instance.Id)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteSuccessfully(t)
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	result, err := instances.Delete(client.ServiceClient(), instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "dff1d289-4d03-4942-8b9f-463ea07c000d", result.JobId)

}
func TestResizeFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleResizeFlavorSuccessfully(t)
	resizeFlavorBody := instances.ResizeFlavorOpts{
		ResizeFlavor: &instances.SpecCode{Speccode: "rds.mysql.s1.large.ha"},
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	result, err := instances.Resize(client.ServiceClient(),resizeFlavorBody, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "dff1d289-4d03-4942-8b9f-463ea07c000d", result.JobId)

}
func TestEnlargeVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleEnlargeVolumeSuccessfully(t)
	EnlargeVolumeBody := instances.EnlargeVolumeRdsOpts{
		EnlargeVolume: &instances.EnlargeVolumeSize{Size: 200},
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	result, err := instances.EnlargeVolume(client.ServiceClient(),EnlargeVolumeBody, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "dff1d289-4d03-4942-8b9f-463ea07c000d", result.JobId)

}
func TestRestarRdsInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleRestarRdsInstanceSuccessfully(t)
	Restart := instances.RestartRdsInstanceOpts{
		Restart: " ",
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	result, err := instances.Restart(client.ServiceClient(),Restart, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "dff1d289-4d03-4942-8b9f-463ea07c000d", result.JobId)

}
func TestSingleToHaRdsInstance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleSingleToHaRdsInstanceSuccessfully(t)
	singleToHaRdsBody := instances.SingleToHaRdsOpts{

		SingleToHa: &instances.SingleToHaRds{AzCodeNewNode: "cn-north-4b", Password: "YourPassword"},
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	result, err := instances.SingleToHa(client.ServiceClient(),singleToHaRdsBody, instanceId).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "dff1d289-4d03-4942-8b9f-463ea07c000d", result.JobId)

}
func TestListErrorLogAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListErrorLogSuccessfully(t)

	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	allPages, err := instances.ListErrorLog(client.ServiceClient(),nil, instanceId).AllPages()
	allRds, err := instances.ExtractErrorLog(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allRds.ErrorLogList))

}
func TestListErrorLogEach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListErrorLogSuccessfully(t)
	opts := instances.DbErrorlogOpts{
		StartDate: "",
		EndDate:   "",
		Level:     "",
		Limit:     "",
		Offset:    "",
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	count := 0
	err := instances.ListErrorLog(client.ServiceClient(), opts,instanceId).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := instances.ExtractErrorLog(page)
		th.AssertNoErr(t, err)
		fmt.Println(actual.ErrorLogList)
		th.CheckDeepEquals(t, "ERROR", actual.ErrorLogList[0].Level)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
func TestListSlowLogEach(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSlowLogSuccessfully(t)
	opts := instances.DbSlowLogOpts{
		StartDate: "2019-08-30T10:41:14+0800",
		EndDate: "",
		Type: "",
		Limit: "",
		Offset: "",
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	count := 0
	err := instances.ListSlowLog(client.ServiceClient(), opts,instanceId).EachPage(func(page pagination.Page) (bool, error) {
		count++
		allSlowLog, err := instances.ExtractSlowLog(page)
		th.AssertNoErr(t, err)
		fmt.Println(allSlowLog.TotalRecord)
		th.CheckDeepEquals(t, "INSERT", allSlowLog.Slowloglist[0].Type)
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}
func TestListSlowLogAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSlowLogSuccessfully(t)
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	allPages, err := instances.ListSlowLog(client.ServiceClient(),nil, instanceId).AllPages()
	th.AssertNoErr(t, err)
	allSlowLog, err := instances.ExtractSlowLog(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allSlowLog.Slowloglist))
}
func TestListSlowLogOptSuccessfully(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleListSlowLogOptSuccessfully(t)
	opts := instances.DbSlowLogOpts{
		StartDate: "2019-08-30T10:41:14+0800",
		EndDate: "2019-09-10T10:41:14+0800",
		Type: "INSERT",
		Limit: "10",
		Offset: "",
	}
	instanceId := "3e93d3eb20b34bfbbdcc81a79c1c3045"
	count := 0
	err := instances.ListSlowLog(client.ServiceClient(), opts,instanceId).EachPage(func(page pagination.Page) (bool, error) {
		count++
		allSlowLog, err := instances.ExtractSlowLog(page)
		th.AssertNoErr(t, err)
		fmt.Println(allSlowLog.TotalRecord)
		th.CheckDeepEquals(t, "INSERT", allSlowLog.Slowloglist[0].Type)
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}