package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/as/v1/groups"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/auth/aksk"
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
	TestCreateGroupWithAllPara(sc)
	TestQueryGroupDetail(sc)
	TestQueryGroupList(sc)
	UpdateGroupInStopAndNoneInstance(sc)
	EnableGroup(sc)
	DisableGroup(sc)
	DeleteGroup(sc)
	fmt.Println("main end...")
}

//---------------------------------------------创建伸缩组--------------------------------------------------
func TestCreateGroupWithAllPara(client *gophercloud.ServiceClient) {
	//Parameters
	ConfigurationId := "ef745d88-fdf9-48d6-b807-9e2cc8e1355b"
	DesireInstance := 0
	MinInstance := 0
	MaxInstance := 5
	CoolDownTime := 200
	//AZ
	var zone = []string{"az1.dc1", "kvmxen.dc1"}
	//子网
	networkID1 := groups.Network{
		ID: "ea5080b9-1317-403e-bc69-4d33660ca0fb",
	}
	networkID2 := groups.Network{
		ID: "2c610576-df25-456e-9468-a83a41fae137",
	}
	networkID := []groups.Network{networkID1, networkID2,}
	//安全组
	SecurityGroupID1 := groups.SecurityGroup{
		ID: "bb8fcedf-9f9e-4815-8372-e9031f685b79",
	}
	SecurityGroupID := []groups.SecurityGroup{SecurityGroupID1}
	//VPC
	var vpcID = "8f4ad2fb-8a4b-4329-be7c-5c28ef3f32da"
	//healthPeriodicAuditTime
	healthPeriodicAuditTime := 15
	//通知
	var notification = []string{"EMAIL"}
	//deletePublicip
	var deletePublicip = true
	//lb
	var ListenerId = "7a281f9eed37448b80ad4b526273b04e,b512335c71f24d9d9950301800824cd0"
	//创建伸缩组带所有参数
	HealthPeriodicAuditTimeGracePeriod := 50
	opts := groups.CreateOpts{
		ScalingGroupName:                   "as-test-CreateGroup-WithAllPara",
		ScalingConfigurationId:             ConfigurationId,
		DesireInstanceNumber:               &DesireInstance,
		MinInstanceNumber:                  &MinInstance,
		MaxInstanceNumber:                  &MaxInstance,
		CoolDownTime:                       &CoolDownTime,
		LbListenerId:                       ListenerId,
		AvailableZones:                     zone,
		Networks:                           networkID,
		SecurityGroups:                     SecurityGroupID,
		VpcId:                              vpcID,
		HealthPeriodicAuditMethod:          "ELB_AUDIT",
		HealthPeriodicAuditTime:            &healthPeriodicAuditTime,
		HealthPeriodicAuditTimeGracePeriod: &HealthPeriodicAuditTimeGracePeriod,
		InstanceTerminatePolicy:            "OLD_CONFIG_OLD_INSTANCE",
		Notifications:                      notification,
		DeletePublicip:                     &deletePublicip,
		LBaasListeners:                     nil,
		EnterpriseProjectID:                "0",
	}
	result, err := groups.Create(client, opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("ScalingGroupId is:", result.ScalingGroupId)
}

//-------------------------------------------------查询伸缩组详情-------------------------------------------------------
func TestQueryGroupDetail(client *gophercloud.ServiceClient) {

	result, err := groups.Get(client, "b34fd165-e156-489c-bd92-8cf25d64a029").Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Get Group detail:")
	fmt.Println("ScalingGroupId is:", result.ScalingGroupId)
	fmt.Println("ScalingConfigurationName is:", result.ScalingConfigurationName)
	fmt.Println("ScalingConfigurationId is:", result.ScalingConfigurationId)
	fmt.Println("ScalingGroupStatus is:", result.ScalingGroupStatus)
	fmt.Println("AvailableZones is:", result.AvailableZones)
	fmt.Println("ScalingGroupId is:", result.ScalingGroupId)
	fmt.Println("CoolDownTime is:", result.CoolDownTime)
	fmt.Println("CurrentInstanceNumber is:", result.CurrentInstanceNumber)
	fmt.Println("DesireInstanceNumber is:", result.DesireInstanceNumber)
	fmt.Println("IsScaling is:", result.IsScaling)
	fmt.Println("Networks is:", result.Networks)
	fmt.Println("ScalingGroupId is:", result.ScalingGroupId)
	fmt.Println("MaxInstanceNumber is:", result.MaxInstanceNumber)
	fmt.Println("MinInstanceNumber is:", result.MinInstanceNumber)
	fmt.Println("ScalingGroupName is:", result.ScalingGroupName)
}

//--------------------------------------------------查询伸缩组列表------------------------------------------------------
func TestQueryGroupList(client *gophercloud.ServiceClient) {
	//查询伸缩组列表,不带参数
	querry := groups.ListOpts{
	}
	result, err := groups.List(client, querry).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	resultList,err:= groups.ExtractGroups(result)
	for _, resp := range resultList.ScalingGroups{
		fmt.Println("Group list:")
		fmt.Println("ScalingConfigurationName is:", resp.ScalingConfigurationName)
		fmt.Println("ScalingConfigurationId is:", resp.ScalingConfigurationId)
		fmt.Println("ScalingGroupId is:", resp.ScalingGroupId)
		fmt.Println("CoolDownTime is:", resp.CoolDownTime)
		fmt.Println("ScalingGroupName is:", resp.ScalingGroupName)
		fmt.Println("MinInstanceNumber is:", resp.MinInstanceNumber)
		fmt.Println("MaxInstanceNumber is:", resp.MaxInstanceNumber)
		fmt.Println("Networks is:", resp.Networks)
		fmt.Println("IsScaling is:", resp.IsScaling)
		fmt.Println("DesireInstanceNumber is:", resp.DesireInstanceNumber)
		fmt.Println("CurrentInstanceNumber is:", resp.CurrentInstanceNumber)
		fmt.Println("AvailableZones is:", resp.AvailableZones)
		fmt.Println("ScalingGroupStatus is:", resp.ScalingGroupStatus)
		fmt.Println("LbaasListeners is:", resp.LbaasListeners)
		fmt.Println("VpcId is:", resp.VpcId)
		fmt.Println("------------------------")
	}
	//查询伸缩组列表,带参数
	scalingGroupName := "as-"
	scalingConfigurationId := "774e80e1-987f-400b-a501-f242d2dc2349"
	querryWithParameters := groups.ListOpts{
		ScalingGroupName:       scalingGroupName,
		ScalingConfigurationId: scalingConfigurationId,
		ScalingGroupStatus:     "INSERVICE",
		StartNumber:            0,
		Limit:                  0,
	}
	result1, err := groups.List(client, querryWithParameters).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result1List,err:= groups.ExtractGroups(result1)
	for _, resp := range result1List.ScalingGroups{
		fmt.Println("Group list:")
		fmt.Println("ScalingConfigurationName is:", resp.ScalingConfigurationName)
		fmt.Println("ScalingConfigurationId is:", resp.ScalingConfigurationId)
		fmt.Println("ScalingGroupId is:", resp.ScalingGroupId)
		fmt.Println("CoolDownTime is:", resp.CoolDownTime)
		fmt.Println("ScalingGroupName is:", resp.ScalingGroupName)
		fmt.Println("MinInstanceNumber is:", resp.MinInstanceNumber)
		fmt.Println("MaxInstanceNumber is:", resp.MaxInstanceNumber)
		fmt.Println("IsScaling is:", resp.IsScaling)
		fmt.Println("DesireInstanceNumber is:", resp.DesireInstanceNumber)
		fmt.Println("CurrentInstanceNumber is:", resp.CurrentInstanceNumber)
		fmt.Println("AvailableZones is:", resp.AvailableZones)
		fmt.Println("ScalingGroupStatus is:", resp.ScalingGroupStatus)
		fmt.Println("LbaasListeners is:", resp.LbaasListeners)
		fmt.Println("VpcId is:", resp.VpcId)
		fmt.Println("------------------------")
	}
}

//--------------------------------------修改伸缩组--伸缩组无实例，且为停用状态------------------------------------------
func UpdateGroupInStopAndNoneInstance(client *gophercloud.ServiceClient) {
	ConfigurationId := "ef745d88-fdf9-48d6-b807-9e2cc8e1355b"
	DesireInstance := 0
	MinInstance := 0
	MaxInstance := 6
	CoolDownTime := 300
	//AZ
	var zone  = []string{"az1.dc1", "kvmxen.dc1"}
	//子网
	networkID1 := groups.Network{
		ID: "3dd9cd75-c03a-4d8e-95bf-757ab3996017",
	}
	networkID := []groups.Network{networkID1}
	//安全组
	SecurityGroupID1 := groups.SecurityGroup{
		ID: "bb8fcedf-9f9e-4815-8372-e9031f685b79",
	}
	SecurityGroupID := []groups.SecurityGroup{SecurityGroupID1}
	//healthPeriodicAuditTime
	healthPeriodicAuditTime := 15
	//通知
	var notification  = []string{"EMAIL"}
	//deletePublicip
	var deletePublicip = true
	ProtocolPort1 := 80
	Weight1 := 2
	ProtocolPort2 := 80
	Weight2 := 2
	LBaasListener1 := groups.LBaasListener{
		PoolID:       "ed495077-0703-4f9e-aa80-f07de3618613",
		ProtocolPort: &ProtocolPort1,
		Weight:       &Weight1,
	}
	LBaasListener2 := groups.LBaasListener{
		PoolID:       "159dd3ec-50a0-494f-8b20-e38a6a63dd9b",
		ProtocolPort: &ProtocolPort2,
		Weight:       &Weight2,
	}
	LBaasListeners := []groups.LBaasListener{LBaasListener1, LBaasListener2}
	opts := groups.UpdateOpts{
		ScalingGroupName:          "as-test-ModifyGroup",
		DesireInstanceNumber:      &DesireInstance,
		MinInstanceNumber:         &MinInstance,
		MaxInstanceNumber:         &MaxInstance,
		CoolDownTime:              &CoolDownTime,
		AvailableZones:            zone,
		Networks:                  networkID,
		SecurityGroups:            SecurityGroupID,
		LbListenerId:              "",
		HealthPeriodicAuditMethod: "NOVA_AUDIT",
		HealthPeriodicAuditTime:   &healthPeriodicAuditTime,
		InstanceTerminatePolicy:   "OLD_INSTANCE",
		ScalingConfigurationId:    ConfigurationId,
		Notifications:             notification,
		DeletePublicip:            &deletePublicip,
		LBaasListeners:            LBaasListeners,
		EnterpriseProjectID:       "",
	}
	_, err := groups.Update(client, "b34fd165-e156-489c-bd92-8cf25d64a029", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
}

//-------------------------------------------------启用伸缩组-----------------------------------------------------------
func EnableGroup(client *gophercloud.ServiceClient) {
	opts := groups.EnableOpts{
		Action: "resume",
	}
	_, err := groups.Enable(client, "b34fd165-e156-489c-bd92-8cf25d64a029", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
}

//------------------------------------------------停用伸缩组------------------------------------------------------------
func DisableGroup(client *gophercloud.ServiceClient) {
	opts := groups.EnableOpts{
		Action: "pause",
	}
	_, err := groups.Enable(client, "b34fd165-e156-489c-bd92-8cf25d64a029", opts).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
}

//-------------------------------------------------------删除伸缩组-----------------------------------------------------
func DeleteGroup(client *gophercloud.ServiceClient) {
	opts := groups.DeleteOpts{
		ForceDelete: "no",
	}
	err := groups.Delete(client, "0ea13935-75c8-4bcd-94db-a9eaff51eb70", opts).ExtractErr()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("Test delete AS Group success!")
}
