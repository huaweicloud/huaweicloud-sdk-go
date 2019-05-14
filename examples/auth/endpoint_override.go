package examples

import (
	"fmt"
	"github.com/gophercloud/gophercloud/functiontest/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"os"
	"github.com/gophercloud/gophercloud/openstack/vpc/v1/bandwidths"
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
	os.Setenv("SDK_ANTIDDOS_ENDPOINT_OVERRIDE", "https://antiddos.%(region)s.%(domain)s/")
	os.Setenv("SDK_BSS_ENDPOINT_OVERRIDE", "https://bss.%(region)s.%(domain)s/")
	os.Setenv("SDK_BSS_ENDPOINT_OVERRIDE", "https://bss.cn-north-1.%(domain)s/")
	os.Setenv("SDK_VPC_ENDPOINT_OVERRIDE", "https://vpc.%(region)s.%(domain)s/v1/%(projectID)s/")
	os.Setenv("SDK_CESV1_ENDPOINT_OVERRIDE", "https://ces.%(region)s.%(domain)s/V1.0/%(projectID)s/")
	os.Setenv("SDK_VPCV2.0_ENDPOINT_OVERRIDE", "https://vpc.%(region)s.%(domain)s/v2.0/%(projectID)s/")
	os.Setenv("SDK_ASV1_ENDPOINT_OVERRIDE", "https://as.%(region)s.%(domain)s/autoscaling-api/v1/%(projectID)s/")
	os.Setenv("SDK_ASV2_ENDPOINT_OVERRIDE", "https://as.%(region)s.%(domain)s/autoscaling-api/v2/%(projectID)s/")
	os.Setenv("SDK_DNS_ENDPOINT_OVERRIDE", "https://dns.%(region)s.%(domain)s/")
}

func main() {
	fmt.Println("main start...")
	provider, err := common.AuthToken()
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
		"SDK_VPC_ENDPOINT_OVERRIDE", "https://vpc.%(region)s.%(domain)s/v1/%(projectID)s/")

	sc, err := openstack.NewVPCV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get VPC V1 client failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	result, err := bandwidths.List(sc, bandwidths.ListOpts{
		Limit: 100,
	}).Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	fmt.Printf("bandwidths: %+v\r\n", result)
	fmt.Println("Test List success!")
	fmt.Println("main end...")
}
