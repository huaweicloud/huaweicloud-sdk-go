package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/database"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	fmt.Println("rds datastore test  start...")

	//Set authentication parameters
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
	CreateDataBaseTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01")
	ListDataBaseTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01")
	DeleteDataBaseTest(client, "0477dd5f47c141809e4d0ea2ffde9655in01","rds_test09")
}

func CreateDataBaseTest(client *gophercloud.ServiceClient, InstanceId string) {
	opts := database.CreateOpts{
		Dbname:       "rds_test09",
		Characterset: "utf8",
	}
	resp, DatabaseErr := database.Create(client, opts, InstanceId).Extract()
	if DatabaseErr != nil {
		fmt.Println("CreateDatabaseTest:", DatabaseErr)
		if ue, ok := DatabaseErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("servers CreateDatabaseTest  is success!")
}

func ListDataBaseTest(sc *gophercloud.ServiceClient, InstanceId string) {
	opts := database.ListOpts{
		Limit: 6,
		Page:  1,
	}
	err := database.List(sc, opts, InstanceId).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := database.ExtractDataBase(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.DatabasesList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("DataBase info is :", string(jsServer))
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


func DeleteDataBaseTest(client *gophercloud.ServiceClient, instanceId string,dbName string) {

	resp,DeleteDataBaseErr := database.Delete(client, instanceId,dbName).Extract()
	if DeleteDataBaseErr != nil {
		fmt.Println("DeleteDataBaseErr:", DeleteDataBaseErr)
		if ue, ok := DeleteDataBaseErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp =",resp)
	fmt.Println("Servers DeleteDataBaseTest  Test is success!")
}