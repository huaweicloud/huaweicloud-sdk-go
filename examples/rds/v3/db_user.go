package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/db_user"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	fmt.Println("rds db_user test  start...")
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
	CreateTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01")
	ListTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01")
	DeleteDbUseruserTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01","rds_008")
}

func CreateTest(client *gophercloud.ServiceClient, InstanceId string) {
	opts := db_user.CreateDbUserOpts{
		Username: "rds_009",
		Password: "{your Password}",
	}
	resp, CreateDbUserErr := db_user.Create(client, opts, InstanceId).Extract()
	if CreateDbUserErr != nil {
		fmt.Println("CreateDbUserErr:", CreateDbUserErr)
		if ue, ok := CreateDbUserErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers CreateDbUserTest  is success!")
}

func ListTest(sc *gophercloud.ServiceClient, InstanceId string) {
	opts := db_user.ListDbUsersOpts{
		Limit: 6,
		Page:  1,
	}
	err := db_user.List(sc, opts, InstanceId).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := db_user.ExtractDbUsers(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}
		for _, v := range resp.UsersList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("DbUsers info is :", string(jsServer))
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


func DeleteDbUseruserTest(client *gophercloud.ServiceClient, instanceId string, dbName string) {

	resp,DeleteDbUserErr := db_user.Delete(client, instanceId,dbName).Extract()
	if DeleteDbUserErr != nil {
		fmt.Println("DeleteDbUserErr:", DeleteDbUserErr)
		if ue, ok := DeleteDbUserErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp =",resp)
	fmt.Println("Servers DeleteDbUseruserTest  Test is success!")
}