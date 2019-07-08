package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/as/v2/policies"
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
	sc, errClient := openstack.NewASV2(provider, gophercloud.EndpointOpts{})
	if errClient != nil {
		fmt.Println("Failed to get the NewASV2 client: ", errClient)
		return
	}
	//开始测试
	CreateGroupPoliceV2(sc)
	ListPoliceV2(sc)
	GetPoliceV2(sc)
	UpdatePoliceV2(sc)
	ListAllPoliceV2(sc)
	fmt.Println("main end...")
}

//--------------------------------------修改策略----------------------------------------
func UpdatePoliceV2(client *gophercloud.ServiceClient) {
	opts := policies.UpdateOpts{
		ScalingPolicyName:   "12345",
		ScalingPolicyType:   "SCHEDULED",
		ScalingResourceId:   "7b9a766b-5d1f-41a4-9015-6726f6d0b99e",
		ScalingResourceType: "SCALING_GROUP",
		AlarmId:             "",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "2019-04-30T12:03Z",
			RecurrenceType:  "",
			RecurrenceValue: "",
			StartTime:       "",
			EndTime:         "",
		},
		ScalingPolicyAction: policies.ScalingPolicyAction{
			Operation:  "",
			Size:       0,
			Percentage: 0,
			Limits:     0,
		},
		CoolDownTime: nil,
	}
	result, err := policies.Update(client, "82331ff7-e0e2-46f0-91b5-7c2b27a2e460", opts).Extract()
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

//--------------------------------------查询详情----------------------------------------
func GetPoliceV2(client *gophercloud.ServiceClient) {
	result, err := policies.Get(client, "5cba0719-d424-46a7-9053-82ee643c3e69").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingPolicyId is:", result.ScalingPolicyId)
	fmt.Println("LaunchTime is:", result.ScheduledPolicy.LaunchTime)
	fmt.Println("ScalingPolicyType is:", result.ScalingPolicyType)
	fmt.Println("ScalingPolicyName is:", result.ScalingPolicyName)
	fmt.Println("Operation is:", result.ScalingPolicyAction.Operation)
	fmt.Println("Limits is:", result.ScalingPolicyAction.Limits)
	fmt.Println("Size is:", result.ScalingPolicyAction.Size)
	fmt.Println("PolicyStatus is:", result.PolicyStatus)
	fmt.Println("ScalingResourceId is:", result.ScalingResourceId)
	fmt.Println("ScalingPolicyType is:", result.ScalingPolicyType)
}

//--------------------------------------查询列表----------------------------------------
func ListPoliceV2(client *gophercloud.ServiceClient) {
	opts := policies.ResourceListOpts{
		ScalingPolicyName: "as-policy-8f5k",
		ScalingPolicyType: "RECURRENCE",
		StartNumber:       0,
		Limit:             20,
		ScalingPolicyID:   "afc4042b-4392-41ec-b23a-dfa73774c6b8",
	}
	result, err := policies.GetPolicyListByResourceID(client, "24a2211e-651e-42e6-bf94-348bc052e12c", opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List, err := policies.ExtractPolicies(result)
	for _, resp := range result1List.ScalingPolicies {
		fmt.Println("Policy detail:")
		fmt.Println("ScalingResourceType is:", resp.ScalingResourceType)
		fmt.Println("ScalingPolicyId is:", resp.ScalingPolicyId)
		fmt.Println("ScalingResourceId is:", resp.ScalingResourceId)
		fmt.Println("AlarmId is:", resp.AlarmId)
		fmt.Println("LaunchTime is:", resp.ScheduledPolicy.LaunchTime)
		fmt.Println("RecurrenceType is:", resp.ScheduledPolicy.RecurrenceType)
		fmt.Println("RecurrenceValue is:", resp.ScheduledPolicy.RecurrenceValue)
		fmt.Println("EndTime is:", resp.ScheduledPolicy.EndTime)
		fmt.Println("StartTime is:", resp.ScheduledPolicy.StartTime)
		fmt.Println("ScalingPolicyName is:", resp.ScalingPolicyName)
		fmt.Println("Operation is:", resp.ScalingPolicyAction.Operation)
		fmt.Println("Size is:", resp.ScalingPolicyAction.Size)
		fmt.Println("Limits is:", resp.ScalingPolicyAction.Limits)
		fmt.Println("Percentage is:", resp.ScalingPolicyAction.Percentage)
		fmt.Println("PolicyStatus is:", resp.PolicyStatus)
		fmt.Println("ScalingPolicyType is:", resp.ScalingPolicyType)
		fmt.Println("CoolDownTime is:", resp.CoolDownTime)
		fmt.Println("CreateTime is:", resp.CreateTime)
		fmt.Println("MataData is:", resp.MataData)
		fmt.Println("------------------------")
	}
}

//--------------------------------------查询全量列表----------------------------------------
func ListAllPoliceV2(client *gophercloud.ServiceClient) {
	opts := policies.ListOpts{
		ScalingPolicyName:   "as-",
		ScalingPolicyType:   "SCHEDULED",
		StartNumber:         0,
		Limit:               20,
		ScalingPolicyID:     "0c4507c5-1479-44f6-add4-adc4dfe80516",
		ScalingResourceID:   "85840689-2103-423c-a091-bebe157b3f37",
		ScalingResourceType: "BANDWIDTH",
		SortBy:              "CREATE_TIME",
		Order:               "DESC",
		EnterpriseProjectID: "00253d40-7ff5-4069-9ec8-1b3cbbf05d30",
	}
	result, err := policies.List(client, opts).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List, err := policies.ExtractPolicies(result)
	for _, resp := range result1List.ScalingPolicies {
		fmt.Println("Policy detail:")
		fmt.Println("ScalingResourceType is:", resp.ScalingResourceType)
		fmt.Println("ScalingPolicyId is:", resp.ScalingPolicyId)
		fmt.Println("ScalingResourceId is:", resp.ScalingResourceId)
		fmt.Println("AlarmId is:", resp.AlarmId)
		fmt.Println("ScheduledPolicy is:", resp.ScheduledPolicy)
		fmt.Println("RecurrenceType is:", resp.ScheduledPolicy.RecurrenceType)
		fmt.Println("RecurrenceValue is:", resp.ScheduledPolicy.RecurrenceValue)
		fmt.Println("StartTime is:", resp.ScheduledPolicy.StartTime)
		fmt.Println("EndTime is:", resp.ScheduledPolicy.EndTime)
		fmt.Println("LaunchTime is:", resp.ScheduledPolicy.LaunchTime)
		fmt.Println("ScalingPolicyName is:", resp.ScalingPolicyName)
		fmt.Println("Operation is:", resp.ScalingPolicyAction.Operation)
		fmt.Println("Size is:", resp.ScalingPolicyAction.Size)
		fmt.Println("Limits is:", resp.ScalingPolicyAction.Limits)
		fmt.Println("Percentage is:", resp.ScalingPolicyAction.Percentage)
		fmt.Println("PolicyStatus is:", resp.PolicyStatus)
		fmt.Println("ScalingPolicyType is:", resp.ScalingPolicyType)
		fmt.Println("CoolDownTime is:", resp.CoolDownTime)
		fmt.Println("CreateTime is:", resp.CreateTime)
		fmt.Println("MataData is:", resp.MataData)
		fmt.Println("------------------------")
	}
}

//--------------------------------------创建策略--带宽策略----------------------------------------
func CreateGroupPoliceV2(client *gophercloud.ServiceClient) {
	Size := 1
	CoolDownTime := 200
	opts := policies.CreateOpts{
		ScalingPolicyName:   "as-sdk-policyV2-Group",
		ScalingResourceId:   "83ba170e-2286-4166-aae4-a60668d78ad4",
		ScalingResourceType: "BANDWIDTH",
		ScalingPolicyType:   "ALARM",
		AlarmId:             "al155626897516238E4OldAx",
		ScheduledPolicy: policies.ScheduledPolicy{
			LaunchTime:      "12:03",
			RecurrenceType:  "Weekly",
			RecurrenceValue: "1,3,5,6,7",
			StartTime:       "2019-04-30T12:03Z",
			EndTime:         "2020-04-30T12:03Z",
		},
		ScalingPolicyAction: policies.CreateScalingPolicyAction{
			Operation: "ADD",
			Size:      &Size,
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
