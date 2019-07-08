/*

Package zones provides information and interaction with the zone API
resource for the OpenStack DNS service.

Example to List Zones

	listOpts := zones.ListOpts{
		Email: "jdoe@example.com",
	}

	allPages, err := zones.List(dnsClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allZones, err := zones.ExtractZones(allPages)
	if err != nil {
		panic(err)
	}

	for _, zone := range allZones {
		fmt.Printf("%+v\n", zone)
	}

Example to Create a Zone

	createOpts := zones.CreateOpts{
		Name:        "example.com.",
		Email:       "jdoe@example.com",
		Type:        "PRIMARY",
		TTL:         7200,
		Description: "This is a zone.",
	}

	zone, err := zones.Create(dnsClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Zone

	zoneID := "99d10f68-5623-4491-91a0-6daafa32b60e"
	err := zones.Delete(dnsClient, zoneID).ExtractErr()
	if err != nil {
		panic(err)
	}
package zones




 */

/*
When users need to access a server on the Internet, an IP address is required. However, it is hard to memorize an IP address. Therefore, a domain name easy to memorize is used to identify the IP address. After registering a domain name, you get the right to use it but users cannot access the website, send or receive emails using your domain name until the name is resolved. Domain name resolution is to map domain names to IP addresses. In this case, to access a website on the Internet, visitors only need to enter a domain name in the browser address bar, instead of an IP address.

Sample Code, Create a private zone.

    result, err := zones.Create(client, zones.CreateOpts{
        Name:        "www.ba1.com.",
        Description: "My Zone",
        ZoneType:    "private",
        Email:       "test@test.com",
        Ttl:         3600,
        Router: struct {
            RouterId     string `json:"router_id,omitempty"`
            RouterRegion string `json:"router_region,omitempty"`
        }{
            RouterId:     "773c3c42-d315-417b-9063-87091713148c",
            RouterRegion: "cn-north-1",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }

Sample Code, Delete a zone.


    result, err := zones.Delete(client, "ff80808262baef150162bbf1bc4a167a").Extract()
    if err != nil {
        panic(err)
    }

Sample Code, Query a zone.

    result, err := zones.Get(client, "ff80808262baef150162bbf1bc4a167a").Extract()
    if err != nil {
        panic(err)
    }
Sample Code, Query zones in list.

    allPages, err := zones.List(client, zones.ListOpts{
        Limit: "2",
        Type:  "private",
    }).AllPages()

    result, err := zones.ExtractList(allPages.(zones.ListPage))

    if err != nil {
        panic(err)
    }
Sample Code, Query name servers in a zone.

    result, err := zones.ListNameServers(client, "ff80808262baef150162bbf1bc4a167a").Extract()
    if err != nil {
        panic(err)
    }
Sample Code, Associate a zone with a VPC.

    result, err := zones.AssociateRouter(client, "ff80808262baef150162bbf1bc4a167a", zones.AssociateRouterOpts{
        Router: struct {
            RouterId     string `json:"router_id,omitempty"`
            RouterRegion string `json:"router_region,omitempty"`
        }{
            RouterId:     "773c3c42-d315-417b-9063-87091713148c",
            RouterRegion: "cn-north-1",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }
Sample Code, Disassociate a VPC from a zone.


    result, err := zones.DisassociateRouter(client, "ff80808262baef150162bbf1bc4a167a", zones.DisassociateRouterOpts{
        Router: struct {
            RouterId     string `json:"router_id,omitempty"`
            RouterRegion string `json:"router_region,omitempty"`
        }{
            RouterId:     "773c3c42-d315-417b-9063-87091713148c",
            RouterRegion: "cn-north-1",
        },
    }).Extract()

    if err != nil {
        panic(err)
    }
*/
package zones
