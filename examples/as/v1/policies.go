package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v1/policies"
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
	CreateSchedulePoliceV1(sc)
	CreateRecurrenceDayPoliceV1(sc)
	CreateRecurrenceWeekPoliceV1(sc)
	CreateRecurrenceMonthPoliceV1(sc)
	CreateAlarmPoliceV1(sc)
	ModifyPoliceV1(sc)
	ListPoliceV1(sc)
	GetPoliceV1(sc)
	ActionPoliceV1(sc)
	DeletePoliceV1(sc)
	fmt.Println("main end...")
}

//--------------------------------------启用、停止、立即执行伸缩策略---------------------------------------------------
func ActionPoliceV1(client *gophercloud.ServiceClient) {
	opts := policies.ActionOpts{
		Action: "execute",
	}
	err := policies.Action(client, "f56fd124-3530-42c1-a11a-6572e2d58870", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Action Police success!")
}

//--------------------------------------删除伸缩策略---------------------------------------------------
func DeletePoliceV1(client *gophercloud.ServiceClient) {
	err := policies.Delete(client, "f56fd124-3530-42c1-a11a-6572e2d58870").ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Delete Police success!")
}

//--------------------------------------查看策略详情---------------------------------------------------
func GetPoliceV1(client *gophercloud.ServiceClient) {
	result, err := policies.Get(client, "9f18e8e6-f5da-4e62-ad84-ac53f263b342").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingGroupId is:", result.ScalingGroupId)
	fmt.Println("CoolDownTime is:", result.CoolDownTime)
	fmt.Println("AlarmId is:", result.AlarmId)
	fmt.Println("CreateTime is:", result.CreateTime)
	fmt.Println("PolicyStatus is:", result.PolicyStatus)
	fmt.Println("InstanceNumber is:", *result.ScalingPolicyAction.InstanceNumber)
	fmt.Println("Operation is:", result.ScalingPolicyAction.Operation)
	fmt.Println("ScalingPolicyName is:", result.ScalingPolicyName)
	fmt.Println("LaunchTime is:", result.ScheduledPolicy.LaunchTime)
	fmt.Println("RecurrenceType is:", result.ScheduledPolicy.RecurrenceType)
	fmt.Println("RecurrenceValue is:", result.ScheduledPolicy.RecurrenceValue)
	fmt.Println("EndTime is:", result.ScheduledPolicy.EndTime)
	fmt.Println("StartTime is:", result.ScheduledPolicy.StartTime)
}

//--------------------------------------查看策略列表---------------------------------------------------
func ListPoliceV1(client *gophercloud.ServiceClient) {
	opts := policies.ListOpts{
		ScalingPolicyID:   "ff166faf-06f3-42f8-90e9-985610bc7cbd",
		ScalingPolicyName: "as-policy-sdk-RecurrenceDayPolicy",
		ScalingPolicyType: "RECURRENCE",
		StartNumber:       0,
		Limit:             20,
	}
	result, err := policies.List(client, "86f3b2dc-de0f-4e63-84ac-a4fddb713555", opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List,err:= policies.ExtractPolicies(result)
	for _, resp := range result1List.ScalingPolicies{
		fmt.Println("Policy detail:")
		fmt.Println("CreateTime is:", resp.CreateTime)
		fmt.Println("ScalingGroupId is:", resp.ScalingGroupId)
		fmt.Println("CoolDownTime is:", resp.CoolDownTime)
		fmt.Println("ScalingPolicyId is:", resp.ScalingPolicyId)
		fmt.Println("ScalingPolicyType is:", resp.ScalingPolicyType)
		fmt.Println("PolicyStatus is:", resp.PolicyStatus)
		fmt.Println("Operation is:", resp.ScalingPolicyAction.Operation)
		fmt.Println("InstanceNumber is:", *resp.ScalingPolicyAction.InstanceNumber)
		fmt.Println("ScalingPolicyName is:", resp.ScalingPolicyName)
		fmt.Println("LaunchTime is:", resp.ScheduledPolicy.LaunchTime)
		fmt.Println("StartTime is:", resp.ScheduledPolicy.StartTime)
		fmt.Println("EndTime is:", resp.ScheduledPolicy.EndTime)
		fmt.Println("RecurrenceValue is:", resp.ScheduledPolicy.RecurrenceValue)
		fmt.Println("RecurrenceType is:", resp.ScheduledPolicy.RecurrenceType)
		fmt.Println("AlarmId is:", resp.AlarmId)
		fmt.Println("------------------------")
	}
}

//--------------------------------------修改策略---------------------------------------------------
func ModifyPoliceV1(client *gophercloud.ServiceClient) {
	InstanceNumber := 300
	CoolDownTime := 200
	opts := policies.UpdateOpts{
		ScalingPolicyName: "111234567",
		ScalingPolicyType: "RECURRENCE",
		AlarmId:           "al1556177913202WRqdpNnVx",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "12:03",
			RecurrenceType:  "Weekly",
			RecurrenceValue: "1,2",
			StartTime:       "2019-04-30T12:03Z",
			EndTime:         "2029-04-30T12:03Z",
		},
		ScalingPolicyAction: policies.ScalingPolicyAction{
			Operation:      "",
			InstanceNumber: &InstanceNumber,
		},
		CoolDownTime: &CoolDownTime,
	}
	result, err := policies.Update(client, "ce651ef2-7369-43c6-aec2-c7d2c9382bc0", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
}

//--------------------------------------创建策略--告警----------------------------------------
func CreateAlarmPoliceV1(client *gophercloud.ServiceClient) {
	InstanceNumber := 1
	CoolDownTime := 200
	opts := policies.CreateOpts{
		ScalingPolicyName: "as-policy-sdk-alarmPolicy",
		ScalingGroupId:    "04ffce3e-2424-49ef-859d-885e85ee1fde",
		ScalingPolicyType: "ALARM",
		AlarmId:           "al1556177913202WRqdpNnVx",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "2019-04-30T12:03Z",
			RecurrenceType:  "Weekly",
			RecurrenceValue: "1,3,5",
			StartTime:       "2019-04-30T12:03Z",
			EndTime:         "2020-04-30T12:03Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation:      "ADD",
			InstanceNumber: &InstanceNumber,
		},
		CoolDownTime: &CoolDownTime,
	}
	result, err := policies.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
}

//--------------------------------------创建策略--定时----------------------------------------
func CreateSchedulePoliceV1(client *gophercloud.ServiceClient) {
	InstanceNumber := 1
	CoolDownTime := 200
	opts := policies.CreateOpts{
		ScalingPolicyName: "as-policy-sdk-SchedulePolicy",
		ScalingGroupId:    "04ffce3e-2424-49ef-859d-885e85ee1fde",
		ScalingPolicyType: "SCHEDULED",
		AlarmId:           "123",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "2019-04-30T12:03Z",
			RecurrenceType:  "xxx",
			RecurrenceValue: "xxx",
			StartTime:       "xxx",
			EndTime:         "xxx",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation:      "ADD",
			InstanceNumber: &InstanceNumber,
		},
		CoolDownTime: &CoolDownTime,
	}
	result, err := policies.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
}

//--------------------------------------创建策略--周期-按天----------------------------------------
func CreateRecurrenceDayPoliceV1(client *gophercloud.ServiceClient) {
	InstanceNumber := 1
	CoolDownTime := 200
	opts := policies.CreateOpts{
		ScalingPolicyName: "as-policy-sdk-RecurrenceDayPolicy",
		ScalingGroupId:    "04ffce3e-2424-49ef-859d-885e85ee1fde",
		ScalingPolicyType: "RECURRENCE",
		AlarmId:           "1234",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "12:03",
			RecurrenceType:  "Daily",
			RecurrenceValue: "1,3,5,xxx",
			StartTime:       "2019-04-30T12:03Z",
			EndTime:         "2020-04-30T12:03Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation:      "ADD",
			InstanceNumber: &InstanceNumber,
		},
		CoolDownTime: &CoolDownTime,
	}
	result, err := policies.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
}

//--------------------------------------创建策略--周期-按周----------------------------------------
func CreateRecurrenceWeekPoliceV1(client *gophercloud.ServiceClient) {
	InstancePercentage := 20000
	CoolDownTime := 86400
	opts := policies.CreateOpts{
		ScalingPolicyName: "as-policy-sdk-RecurrenceWeekPolicy",
		ScalingGroupId:    "04ffce3e-2424-49ef-859d-885e85ee1fde",
		ScalingPolicyType: "RECURRENCE",
		AlarmId:           "",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "12:03",
			RecurrenceType:  "Weekly",
			RecurrenceValue: "1,3,5,6,7",
			StartTime:       "2019-04-30T12:03Z",
			EndTime:         "2020-04-30T12:03Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation:          "ADD",
			InstancePercentage: &InstancePercentage,
		},
		CoolDownTime: &CoolDownTime,
	}
	result, err := policies.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
}

//--------------------------------------创建策略--周期-按月----------------------------------------
func CreateRecurrenceMonthPoliceV1(client *gophercloud.ServiceClient) {
	InstanceNumber := 1
	CoolDownTime := 200
	opts := policies.CreateOpts{
		ScalingPolicyName: "as-policy-sdk-RecurrenceMonthPolicy",
		ScalingGroupId:    "04ffce3e-2424-49ef-859d-885e85ee1fde",
		ScalingPolicyType: "RECURRENCE",
		AlarmId:           "",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "12:03",
			RecurrenceType:  "Monthly",
			RecurrenceValue: "1,3,5,20,31,31,31,31,31,1,3,5,20,31,31,31,31,31",
			StartTime:       "2019-04-30T12:03Z",
			EndTime:         "2020-04-30T12:03Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation:      "ADD",
			InstanceNumber: &InstanceNumber,
		},
		CoolDownTime: &CoolDownTime,
	}
	result, err := policies.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
}
