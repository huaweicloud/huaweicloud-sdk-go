/*
A public IP address is an IP address that can be directly accessed over the Internet. Private IP addresses are all IP addresses on the local area network (LAN) of the public cloud and cannot exist on the Internet. An EIP is a static, public IP address. You can bind an EIP to and unbind an EIP from an ECS in your subnet. An EIP enables an ECS in your VPC to communicate with the Internet through a fixed public IP address.

Sample Code, This interface is used to apply for a private IP address.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := privateips.Create(client, tenantID, privateips.CreateOpts{
       Privateips: [] privateips.CreatePrivateIp{
           {
               SubnetId:"5ae24488-454f-499c-86c4-c0355704005d",
               IpAddress: "192.168.0.12",
           },
       },
    }).Extract()

    if err != nil {
       panic(err)
    }

Sample Code, This interface is used to query details about a private IP address using the specified ID.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result, err := privateips.Get(client, tenantID, "ea274524-f1cc-4078-8e67-c002be25c413").Extract()

    if err != nil {
      panic(err)
    }

Sample Code, This interface is used to query private IP addresses using search criteria and to display the private IP addresses in a list.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    subnetID := "5ae24488-454f-499c-86c4-c0355704005d"
    result, err := privateips.List(client, tenantID, subnetID, privateips.ListOpts{
      Limit: 2,
    }).Extract()

    if err != nil {
      panic(err)
    }

Sample Code, This interface is used to delete a private IP address.


    tenantID := "57e98940a77f4bb988a21a7d0603a626"
    result := privateips.Delete(client, tenantID, "ea274524-f1cc-4078-8e67-c002be25c413")

    if err != nil {
      panic(err)
    }

*/
package privateips
