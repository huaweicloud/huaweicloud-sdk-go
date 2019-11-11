package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/storagetype"
	"github.com/gophercloud/gophercloud/pagination"
)

func main() {

	fmt.Println("rds storagetype test  start...")

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
	ListStorageTypeTest(client, "MySQL")
	ListStorageTypeTest(client, "PostgreSQL")
	ListStorageTypeTest(client, "SQLServer")
}

func ListStorageTypeTest(sc *gophercloud.ServiceClient, databasename string) {
	opts := storagetype.ListOpts{
		VersionName: "5.7",
	}
	err := storagetype.List(sc, opts,databasename).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := storagetype.ExtractStorageType(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.StorageTypeList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Database info is :", string(jsServer))
			fmt.Println("Database  id is :", v.Name)

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
