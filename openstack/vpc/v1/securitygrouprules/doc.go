/*
This interface is used to query network resource quotas for the VPC service of a tenant. The network resources include VPCs, subnets, security groups, security group rules, elastic IP addresses, and VPNs.

Sample Code, This interface is used to create a security group rule.


    result, err := securitygrouprules.Create(client, securitygrouprules.CreateOpts{
       SecurityGroupRule: securitygrouprules.CreateSecurityGroupRule {
           Description: "Test SecurityGroup",
           TenantId: tenantID,
           SecurityGroupId: "7af80d49-0a43-462d-aed8-a1e12ac91af6",
           Direction: "egress",
           Protocol: "tcp",
           RemoteIpPrefix: "10.10.0.0/24",
       },
    }).Extract()

    if err != nil {
       panic(err)
    }

Sample Code, This interface is used to query details about a security group rule.


    result, err := securitygrouprules.Get(client, "26243298-ae79-46a3-bad9-34395762e033").Extract()

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to query security group rules using search criteria and to display the security group rules in a list.


    allPages, err := securitygrouprules.List(client, securitygrouprules.ListOpts{
        Limit: 2,
        Protocol: "tcp",
    }).AllPages()

    result, err := securitygrouprules.ExtractList(allPages.(securitygrouprules.ListPage))

    if err != nil {
        panic(err)
    }

Sample Code, This interface is used to delete a security group rule.


    result = securitygrouprules.Delete(client, "26243298-ae79-46a3-bad9-34395762e033")
*/
package securitygrouprules
