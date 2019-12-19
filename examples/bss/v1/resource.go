package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/bss/v1/resource"
)

func main() {
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

	// 初始化服务的client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	customerId := ""
	pageNo := 1
	pageSize := 100
	ListResourceDetailPage(sc, customerId, pageNo, pageSize)
}


//查询客户包周期资源列表
func ListResourceDetailPage(client *gophercloud.ServiceClient, customerId string, pageNo int, pageSize int) {
	resourceListTmp := resource.ListOpts{
		CustomerId: customerId,
		//ResourceIds: "2251f59c-b1ef-4398-bfa8-321782f670a5",
		//OrderId: "xxxx",
		//OnlyMainResource: 1,
		//StatusList:"1,2",
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	resResourceListTmp, err := resource.ListDetail(client, resourceListTmp).Extract()
	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if resResourceListTmp.ErrorCode == "CBC.0000" {
		fmt.Println("ListResourceDetailPage success!")
		fmt.Println("ListResourceDetailPage.TotalCount:", resResourceListTmp.TotalCount)
		fmt.Println("ListResourceDetailPage.ErrorCode:", resResourceListTmp.ErrorCode)
		fmt.Println("ListResourceDetailPage.ErrorMsg:", resResourceListTmp.ErrorMsg)
		fmt.Println("ListResourceDetailPage.Data:", len(resResourceListTmp.Data))

		b, _ := json.MarshalIndent(resResourceListTmp, "", " ")
		fmt.Println(string(b))

		return
	} else {
		fmt.Println("ListResourceDetailPage failed!")
	}
}
