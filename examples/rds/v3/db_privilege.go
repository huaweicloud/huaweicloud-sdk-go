package main

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/db_privilege"
)

func main() {

	fmt.Println("rds Dbprivilege test  start...")

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

	CreateDbprivilegeTest(client, "9787515fe01746e192fa872e85ed61bein01")
	DeleteDbprivilegeTest(client, "9787515fe01746e192fa872e85ed61bein01")

}

func CreateDbprivilegeTest(client *gophercloud.ServiceClient, InstanceId string) {

	var user = []db_privilege.User{
		{Name: "rdsusr_009", Readonly: true},
		{Name: "rdsusr_1109", Readonly: true},
	}
	opts := db_privilege.DbprivilegeOpts{
		Dbname: "rds_test06",
		Users:  user,
	}
	resp, DbprivilegeErr := db_privilege.Create(client, opts, InstanceId).Extract()
	if DbprivilegeErr != nil {
		fmt.Println("CreateDbprivilegeTest:", DbprivilegeErr)
		if ue, ok := DbprivilegeErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers CreateDbprivilegeTest  is success!")
}

func DeleteDbprivilegeTest(client *gophercloud.ServiceClient, InstanceId string) {
	var user = []db_privilege.DeleteUsers{
		{Name: "rdsusr_009"},{Name: "rdsusr_1109"},
	}
	opts := db_privilege.DeleteDbprivilegeOpts{
		Dbname: "rds_008",
		Users:  user,
	}
	resp, DbprivilegeErr := db_privilege.Delete(client, opts, InstanceId).Extract()
	if DbprivilegeErr != nil {
		fmt.Println("DeleteDbprivilegeTest:", DbprivilegeErr)
		if ue, ok := DbprivilegeErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers DeleteDbprivilegeTest  is success!")
}

