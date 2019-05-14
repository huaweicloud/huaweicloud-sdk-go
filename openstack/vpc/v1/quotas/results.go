package quotas

import (
	"github.com/gophercloud/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

type Quota struct {

	// Specifies the resource list objects.
	Resources []Resource `json:"resources"`
}

type Resource struct {

	// Specifies the resource type. The value can be vpc, subnet,
	// securityGroup, securityGroupRule, publicIp, vpn, physicalConnect, virtualInterface,
	// vpcPeer, loadbalancer, listener, firewall, or shareBandwidthIP.
	Type string `json:"type"`

	// Specifies the number of created network resources. The value
	// ranges from 0 to the value of quota.
	Used int `json:"used"`

	// Specifies the maximum quota values for the resources. The
	// quotas can be changed only in the FusionSphere OpenStack system. If it is left blank,
	// -1 is displayed and the resources cannot be created. The default quotas for different
	// resources are as follows:  VPC: 2 Subnet: 100 Security group: 100 Security group
	// rule: 5000 Elastic IP address: 10  VPN: 5 Physical connection: 10 Virtual interface:
	// 50 Load balancer: 10 Listener: 10 VPC peering connection: 50 Firewall: 200 IP address
	// with shared bandwidth: 20 The value ranges from the default quota value to the
	// maximum quota value.
	Quota int `json:"quota"`

	// Specifies the minimum quota value allowed.
	Min int `json:"min"`
}

type ListResult struct {
	commonResult
}

func (r ListResult) Extract() (*Quota, error) {
	var entity Quota
	err := r.ExtractIntoStructPtr(&entity, "quotas")
	return &entity, err
}
