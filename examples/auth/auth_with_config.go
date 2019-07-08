package examples

import (
	"encoding/json"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"net/http"
	"crypto/tls"
)

func getServerList() {
	//设置认证参数
	akskOptions := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.cn-north-1.myhuaweicloud.com/v3",
		DomainID:         "replace-your-domainID",
		ProjectID:        "replace-your-ProjectID",
		Cloud:            "myhuaweicloud.com",
		Region:           "cn-north-1",
		AccessKey:        "replace-your-ak",
		SecretKey:        "replace-your-sk",
	}

	gophercloud.EnableDebug = true

	//初始化配置项，使用默认配置参数。
	conf := gophercloud.NewConfig()
	// 配置超时时间为10s,单位为纳秒。
	conf.WithTimeout(10000000000)
	// 配置http链接，跳过HTTPS 证书验证。
	conf.WithHttpTransport(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},})
	//初始化provider client
	provider, err := openstack.AuthenticatedClientWithOptions(akskOptions, conf)

	//初始化service client
	sc, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	//列出所有服务器
	err = servers.List(sc, servers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {

		//解析返回值
		allServers, errExt := servers.ExtractServers(page)
		if errExt != nil {
			fmt.Println(errExt)
			return false, errExt
		}
		//打印信息
		fmt.Println("List Servers:")
		for _, s := range allServers {
			b, e := json.MarshalIndent(s, "", "")
			if e != nil {
				return false, e
			}
			fmt.Println(string(b))
		}
		return true, nil
	})

	if err != nil {
		fmt.Println("err:", err.Error())
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

}
