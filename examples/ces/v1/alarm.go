package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/alarms"
	"github.com/gophercloud/gophercloud/pagination"
)

var alarmId string

func main() {
	fmt.Println("main start...")
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Domain:           "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}

	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("Failed to get the provider: ", errAuth)
		return
	}

	sc, err := openstack.NewCESV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get ces client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	AlarmCreate(sc)
	AlarmList(sc)
	AlarmUpdate(sc)
	AlarmGet(sc)
	AlarmDelete(sc)
	fmt.Println("main end...")
}

func AlarmCreate(sc *gophercloud.ServiceClient) {
	metric := alarms.MetricInfo{
		MetricName: "down_stream",
		Namespace:  "SYS.VPC",
		Dimensions: []alarms.MetricsDimension{
			{
				Name:  "publicip_id",
				Value: "94e1983d-22f6-4e41-8c83-ac0d1c82622d",
			},
		},
	}
	alarmEnabled := true
	alarmActionEnabled := false
	opts := alarms.CreateOpts{
		AlarmName:          "Test",
		AlarmEnabled:       &alarmEnabled,
		AlarmLevel:         2,
		AlarmActionEnabled: &alarmActionEnabled,
		Metric:             metric,
		Condition: alarms.Condition{
			Period:             1,
			Filter:             "average",
			ComparisonOperator: ">=",
			Value:              100,
			Unit:               "Byte",
			Count:              3,
		},
	}

	createInfo, err := alarms.Create(sc, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	alarmId = createInfo.AlarmId
	fmt.Printf("Test AlarmCreate success, alarmID: %s\n", alarmId)

	res, marshalErr := json.MarshalIndent(createInfo, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal createInfo error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}

func AlarmDelete(sc *gophercloud.ServiceClient) {
	err := alarms.Delete(sc, alarmId).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test AlarmDelete success！")
}

func AlarmGet(sc *gophercloud.ServiceClient) {
	getInfo, err := alarms.Get(sc, alarmId).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test AlarmGet success！")

	res, marshalErr := json.MarshalIndent(getInfo, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal getInfo error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}

func AlarmList(sc *gophercloud.ServiceClient) {
	opts := alarms.ListOpts{
		Limit: 10,
	}

	var alarmsOnePage alarms.Alarms
	err := alarms.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		alarmTemp, err := alarms.ExtractAlarms(page)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		alarmsOnePage = alarmTemp
		return false, err
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	res, marshalErr := json.MarshalIndent(alarmsOnePage, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal alarmsOnePage error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))

	allPages, err := alarms.List(sc, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	alarmsAllPage, _ := alarms.ExtractAlarms(allPages)
	fmt.Println("Test AlarmList success！")

	res, marshalErr = json.MarshalIndent(alarmsAllPage, "", " ")
	if marshalErr != nil {
		fmt.Printf("Marshal alarmsAllPage error: %s\n", marshalErr.Error())
	}
	fmt.Println(string(res))
}

func AlarmUpdate(sc *gophercloud.ServiceClient) {
	alarmEnabled := false
	opts := alarms.UpdateOpts{
		AlarmEnabled: &alarmEnabled,
	}

	err := alarms.Update(sc, alarmId, opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test AlarmUpdate success！")
}
