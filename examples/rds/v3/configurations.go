package main

import (
	"encoding/json"
	//"crypto/tls"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/configurations"
	"github.com/gophercloud/gophercloud/pagination"
	//"net/http"
)

func main() {

	fmt.Println("rds config test  start...")
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
	//Init service client
	client, clientErr := openstack.NewRDSV3(provider, gophercloud.EndpointOpts{Region:"xxx"})
	if clientErr != nil {
		fmt.Println("Failed to get the NewRDSV3 client: ", clientErr)
		return
	}
	CreateConfigurationsTest(client)
	ListConfigurationsTest(client)

}

func ListConfigurationsTest(sc *gophercloud.ServiceClient) {

	err := configurations.List(sc).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := configurations.ExtractGetConfigurations(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}

		for _, v := range resp.ConfigurationsList {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Server info is :", string(jsServer))
			fmt.Println("Configurations  Name is :", v.Name)
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

func CreateConfigurationsTest(client *gophercloud.ServiceClient) {
	value:=map[string]string{
		"auto_increment_increment": "3",
	}
	CreateConfigurations := configurations.CreateConfigurationsOpts{
		Name: "config-191",
		Description: "config_test191",
		Values:value,
		Datastore: &configurations.Datastore{Type: "MySQL", Version: "5.6"},
	}

	resp, CreateConfigurationsErr := configurations.Create(client, CreateConfigurations).Extract()
	if CreateConfigurationsErr != nil {
		fmt.Println("CreateConfigurationsErr:", CreateConfigurationsErr)
		if ue, ok := CreateConfigurationsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("resp:", resp)
	fmt.Println("resp.Configurations.Description:", resp.CreateConfigurationsRes.Description)
	fmt.Println("Servers CreateConfigurationsTest  Test is success!")
}
