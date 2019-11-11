package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/instances"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/securitygroups"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/vpcs"
	"github.com/gophercloud/gophercloud/pagination"
	"time"
)

func main() {
	fmt.Println("rds create instance ext test  start......")

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

	//Initialization service client
	clientvpc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get vpc v1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	GetRdsInstanceInfo(client,"")
	CreateAndGetRdsInstanceInfoTest(client)
    CreateAndGetRdsReplicaInfoTest(client)

	for _, resp := range ListRdsVPCTest(clientvpc) {
		fmt.Println("vpc name is:", resp.Name)
		fmt.Println("vpc Id is:", resp.ID)
	}

	for _, resp := range ListRdsSubnetTest(clientvpc, "3138ce3d-8837-49a6-b68a-4cdbc5b30a45") {
		fmt.Println("subnet Id is:", resp.ID)
		fmt.Println("subnet Status is:", resp.Status)
		fmt.Println("subnet Name is:", resp.Name)
		fmt.Println("subnet VpcId is:", resp.VpcID)

	}
	for _, resp := range ListRdsSecurityGroupTest(clientvpc) {
		fmt.Println("securityGroup ID is:", resp.ID)
		fmt.Println("securityGroup Name is:", resp.Name)
	}

	fmt.Println("main end...")
}

func CreateAndGetRdsInstanceInfoTest(client *gophercloud.ServiceClient) {

	InstancesTestStruct := instances.CreateRdsOpts{
		Name:           "GoT_Callback100_2_S-2-20190731-043903-fab4",
		Datastore:      instances.Datastore{Type: "MySQL", Version: "5.6"},
		BackupStrategy: &instances.BackupStrategy{StartTime: "06:15-07:15", KeepDays: 7},
		Ha:             &instances.Ha{Mode: "Ha", ReplicationMode: "semisync"},
		FlavorRef:        "rds.mysql.s1.medium.ha",
		Volume:           &instances.Volume{Type: "ULTRAHIGH", Size: 100},
		AvailabilityZone: "cn-north-4b,cn-north-4b",
		VpcId:            "3138ce3d-8837-49a6-b68a-4cdbc5b30a45",
		SubnetId:         "0f48e1d1-c244-422a-baa0-acfb1133c148",
		SecurityGroupId:  "702e9e18-34a2-4eda-a847-59546c3f5fa5",
		Password:         "{your Password}",
		Port:             "{your Port}",
		Region:           "cn-north-4",
	}

	var instanceInfor *instances.RdsInstanceResponse
	instanceInfor = nil
	instanceInfor, err := CreateAndGetRdsInstanceInfo(client, InstancesTestStruct, 25)
	if instanceInfor == nil {
		if err == nil {
			fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest is timeout !!  ")
		} else {
			fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest instanceInfor is fail !!  ", err)
		}
	} else {
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest instanceInfor  =  ", instanceInfor)
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest instanceInfor.Name =  ", instanceInfor.Name)
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest instanceInfor.Status =  ", instanceInfor.Status)
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest instanceInfor.Id =  ", instanceInfor.Id)
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfoTest instanceInfor.VpcId =  ", instanceInfor.VpcId)
	}

}

func CreateAndGetRdsReplicaInfoTest(client *gophercloud.ServiceClient) {

	instancesTestStruct := instances.CreateReplicaOpts{
		Name:             "Test_reply_fab4",
		FlavorRef:        "rds.mysql.s1.medium.rr",
		Volume:           &instances.Volume{Type: "ULTRAHIGH", Size: 100},
		AvailabilityZone: "cn-north-4a",
		Region:           "cn-north-4",
		ReplicaOfId:      "aadd7e737bb24deda16f1d15d89fc87ein01",
	}
	var instanceInfor *instances.RdsInstanceResponse
	instanceInfor = nil
	instanceInfor = CreateAndGetRdsReplica(client, instancesTestStruct, 25)
	if instanceInfor == nil {
		fmt.Println("[LogInfor]CreateAndGetRdsReplicaInfoTest instanceInfor is fail !!  ", instanceInfor)
	} else {
		fmt.Println("[LogInfor]CreateAndGetRdsReplicaInfoTest instanceInfor  =  ", instanceInfor)
		fmt.Println("[LogInfor]CreateAndGetRdsReplicaInfoTest instanceInfor.Name =  ", instanceInfor.Name)
		fmt.Println("[LogInfor]CreateAndGetRdsReplicaInfoTest instanceInfor.Id =  ", instanceInfor.Id)
		fmt.Println("[LogInfor]CreateAndGetRdsReplicaInfoTest instanceInfor.PrivateIps =  ", instanceInfor.PrivateIps)
	}
}

func CreateAndGetRdsReplica(client *gophercloud.ServiceClient, instancesCreate instances.CreateReplicaOpts, trytimes int) *instances.RdsInstanceResponse {

	job, createRdsInstanceErr := instances.CreateReplica(client, instancesCreate).Extract()
	if createRdsInstanceErr != nil {
		fmt.Println("[LogInfor]CreateAndGetRdsReplica error:", createRdsInstanceErr)
		if ue, ok := createRdsInstanceErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return nil
	}
	instanceIDS := job.Instance.Id
	if instanceIDS == "" {
		fmt.Println("[LogInfor]CreateAndGetRdsReplica Instance.Id is null !")
		return nil
	}

	var checktimes = 0
	var instanceResTmp *instances.RdsInstanceResponse

	for checktimes < trytimes {
		time.Sleep(60 * time.Second)
		checktimes = checktimes + 1
		instanceResTmp := GetRdsInstanceInfo(client, instanceIDS)
		fmt.Println("[LogInfor]GetRdsInstanceInfo instanceResTmp is ", instanceResTmp)
		if instanceResTmp != nil {
			if instanceResTmp.Status == "ACTIVE" {
				return instanceResTmp
			} else {
				continue
			}
		} else {

			continue
		}
	}
	return instanceResTmp
}

func CreateAndGetRdsInstanceInfo(client *gophercloud.ServiceClient, instancesCreate instances.CreateRdsOpts, trytimes int) (*instances.RdsInstanceResponse, error) {

	job, createRdsInstanceErr := instances.Create(client, instancesCreate).Extract()
	if createRdsInstanceErr != nil {
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfo error:", createRdsInstanceErr)
		if ue, ok := createRdsInstanceErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return nil, createRdsInstanceErr
	}
	instanceIDS := job.Instance.Id
	if instanceIDS == "" {
		fmt.Println("[LogInfor]CreateAndGetRdsInstanceInfo Instance.Id is null !")
		return nil, createRdsInstanceErr
	}

	var checktimes = 0
	var instanceResTmp *instances.RdsInstanceResponse
	for checktimes < trytimes {
		time.Sleep(60 * time.Second)

		checktimes = checktimes + 1
		instanceResTmp := GetRdsInstanceInfo(client, instanceIDS)
		fmt.Println("[LogInfor]GetRdsInstanceInfo instanceResTmp is ", instanceResTmp)
		if instanceResTmp != nil {
			if instanceResTmp.Status == "ACTIVE" {
				return instanceResTmp, nil
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return instanceResTmp, nil
}

func GetRdsInstanceInfo(sc *gophercloud.ServiceClient, instanceID string) *instances.RdsInstanceResponse {
	opts := instances.ListRdsInstanceOpts{
		Limit:  0,
		Offset: 0,
		Id:     instanceID,
	}
	var instanceResTmp *instances.RdsInstanceResponse
	instanceResTmp = nil
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
			if v.Id == instanceID && v.Status == "ACTIVE" {
				instanceResTmp = &v
				fmt.Println("[LogInfor]GetRdsInstanceInfo create instance status is  ACTIVE infor: ", instanceResTmp)
				return true, nil
			}
			if v.Id == instanceID && v.Status != "ACTIVE" {
				fmt.Println("[LogInfor]GetRdsInstanceInfo create instance status is: ", v.Status)
				return true, nil
			}
		}
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return nil
	}
	fmt.Println("[LogInfor] GetRdsInstanceInfo instanceTmp is ", instanceResTmp)
	return instanceResTmp
}

func ListRdsVPCTest(sc *gophercloud.ServiceClient) []vpcs.VPC {

	allpages, err := vpcs.List(sc, vpcs.ListOpts{
		//Limit: 2,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return nil
	}

	result, err1 := vpcs.ExtractVpcs(allpages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return nil
	}
	return  result

}

func ListRdsSecurityGroupTest(client *gophercloud.ServiceClient) []securitygroups.SecurityGroup {
	allPages, err := securitygroups.List(client, securitygroups.ListOpts{
		Limit: 1,
	}).AllPages()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return nil
	}
	result, err1 := securitygroups.ExtractSecurityGroups(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return nil
	}
	return  result
}

func ListRdsSubnetTest(sc *gophercloud.ServiceClient, vpcid string) []subnets.Subnet {

	allPages, err := subnets.List(sc, subnets.ListOpts{
		VpcID: vpcid,
		//Limit: 1,
	}).AllPages()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return nil
	}

	result, err1 := subnets.ExtractSubnets(allPages)

	if err1 != nil {
		fmt.Println("err1:", err1.Error())
		return nil
	}
	return result

}