/*
This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.

Sample Code, This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.


    result, err := quotas.List(client.ServiceClient(), tenantID, quotas.ListOpts{
       Type: "vpc",
    }).Extract()
*/
package quotas
