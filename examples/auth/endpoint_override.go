package examples

import (
	"fmt"
	"os"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	//"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

// function setenv ,list override examples as below
func setenv() {
	os.Setenv("SDK_COMPUTE_ENDPOINT_OVERRIDE", "https://ecs.%(region)s.%(domain)s/v2/%(projectID)s/")
	os.Setenv("SDK_ECSV1.1_ENDPOINT_OVERRIDE", "https://ecs.%(region)s.%(domain)s/v1.1/%(projectID)s/")
	os.Setenv("SDK_ECSV2_ENDPOINT_OVERRIDE", "https://ecs.%(region)s.%(domain)s/v2/%(projectID)s/")
	os.Setenv("SDK_ECS_ENDPOINT_OVERRIDE", "https://ecs.%(region)s.%(domain)s/v1/%(projectID)s/")
	os.Setenv("SDK_IMAGE_ENDPOINT_OVERRIDE", "https://ims.%(region)s.%(domain)s/")
	os.Setenv("SDK_NETWORK_ENDPOINT_OVERRIDE", "https://vpc.%(region)s.%(domain)s/")
	os.Setenv("SDK_VOLUMEV2_ENDPOINT_OVERRIDE", "https://evs.%(region)s.%(domain)s/v2/%(projectID)s/")
	os.Setenv("SDK_BSSV1_ENDPOINT_OVERRIDE", "https://bss.%(domain)s/v1.0/")
	os.Setenv("SDK_BSSINTLV1_ENDPOINT_OVERRIDE", "https://cbc.ap-southeast-1.%(domain)s/v1.0/")
	os.Setenv("SDK_IDENTITY_ENDPOINT_OVERRIDE", "https://iam.%(region)s.%(domain)s/")
	os.Setenv("SDK_VPC_ENDPOINT_OVERRIDE", "https://vpc.%(region)s.%(domain)s/v1/%(projectID)s/")
	os.Setenv("SDK_CESV1_ENDPOINT_OVERRIDE", "https://ces.%(region)s.%(domain)s/V1.0/%(projectID)s/")
	os.Setenv("SDK_VPCV2.0_ENDPOINT_OVERRIDE", "https://vpc.%(region)s.%(domain)s/v2.0/%(projectID)s/")
	os.Setenv("SDK_ASV1_ENDPOINT_OVERRIDE", "https://as.%(region)s.%(domain)s/autoscaling-api/v1/%(projectID)s/")
	os.Setenv("SDK_ASV2_ENDPOINT_OVERRIDE", "https://as.%(region)s.%(domain)s/autoscaling-api/v2/%(projectID)s/")
	//os.Setenv("SDK_DNS_ENDPOINT_OVERRIDE", "https://dns.%(region)s.%(domain)s/")
	//os.Setenv("SDK_ANTIDDOS_ENDPOINT_OVERRIDE", "https://antiddos.%(region)s.%(domain)s/v1/%(projectID)s/")
	//os.Setenv("SDK_ANTIDDOSV2_ENDPOINT_OVERRIDE", "https://antiddos.%(region)s.%(domain)s/v2/%(projectID)s/")
	//os.Setenv("SDK_KMSV1_ENDPOINT_OVERRIDE","https://kms.%(region)s.%(domain)s/v1.0/%(projectID)s/ ")
}

/*
func authAKSK() (*gophercloud.ProviderClient, error) {
	akskOptions := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "your AK string",
		SecretKey:        "your SK string",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}
	provider, err := openstack.AuthenticatedClient(akskOptions)
	if err != nil {
		panic(err)
	}
	return provider, nil
}
*/

func authToken() (*gophercloud.ProviderClient, error) {

	tokenOpts := token.TokenOptions{
		Username:         "your username ",
		Password:         "your password",
		ProjectID:        "{ProjectID}",
		DomainID:         "{domainID}",
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
	}
	gophercloud.EnableDebug = true
	provider, err := openstack.AuthenticatedClient(tokenOpts)

	if err != nil {
		fmt.Println("Failed to authenticate:", err)
		return nil, err
	}
	return provider, nil
}

// EndpointOverrideExample shows how to use endpoint override mechanism to customize your service endpoint.
func EndpointOverrideExample() {
	fmt.Println("main start...")
	provider, err := authToken()
	//provider, err := authAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	//set OS environment, mark override format as：SDK_{service_type}_ENDPOINT_OVERRIDE
	os.Setenv(
		"SDK_COMPUTE_ENDPOINT_OVERRIDE", "https://ecs.xxx.yyy.com/v2/%(projectID)s/")

	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get compute V2 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	allPages, err := servers.List(sc, servers.ListOpts{Limit: 5}).AllPages()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("List Servers:")
	for _, s := range allServers {
		fmt.Println(s)
	}
}
