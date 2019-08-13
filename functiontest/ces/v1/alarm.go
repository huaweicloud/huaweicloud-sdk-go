package main

import (
	"encoding/json"
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/ces/v1/alarms"
	"github.com/gophercloud/gophercloud/pagination"
)

var alarmId string

func main() {

	fmt.Println("main start...")
	provider, err := common.AuthAKSK()
	//provider, err := common.AuthToken()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
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

func AlarmList(sc *gophercloud.ServiceClient) {
	opts := alarms.ListOpts{
		Limit: 10,
	}

	var alarmsOnePage alarms.Alarms
	err := alarms.List(sc, opts).EachPage(func(page pagination.Page) (bool, error) {
		alarmList, err := alarms.ExtractAlarms(page)
		if err != nil {
			fmt.Println(err)
			return false, err
		}
		alarmsOnePage = alarmList
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

	bytes, _ := json.MarshalIndent(alarmsOnePage, "", " ")
	fmt.Println(string(bytes))

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
	bytes, _ = json.MarshalIndent(alarmsAllPage, "", " ")
	fmt.Println(string(bytes))
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
	fmt.Println("Test AlarmCreate success！")
	bytes, _ := json.MarshalIndent(createInfo, "", " ")
	fmt.Println(string(bytes))
}

func AlarmGet(sc *gophercloud.ServiceClient) {
	createInfo, err := alarms.Get(sc, alarmId).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test AlarmGet success！")
	bytes, _ := json.MarshalIndent(createInfo, "", " ")
	fmt.Println(string(bytes))
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
