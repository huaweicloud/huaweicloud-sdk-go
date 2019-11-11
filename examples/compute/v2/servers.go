package main

import (
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"encoding/json"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/startstop"
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/attachinterfaces"
)

func main() {
	fmt.Println("main start...")
	//Set authentication parameters
	akskOptions := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{DomainID}",
	}
	//Init provider client
	provider, authErr := openstack.AuthenticatedClient(akskOptions)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		return
	}
	//Init service client
	client, clientErr := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if clientErr != nil {
		fmt.Println("Failed to get the NewComputeV2 client: ", clientErr)
		return
	}
	serverID := "{serverID}"
	metaData := "{metaData}"
	length := "-1"
	reqID := "{reqID}"
	ServerCreate(client)
	ServersList(client)
	ServersListV247(client)
	ServerGet(client, serverID)
	ServerUpdate(client, serverID)
	ServerDelete(client, serverID)
	ServerRebuild(client, serverID)
	ServerResize(client, serverID)
	ServerConfirmResize(client, serverID)
	ServerRevertResize(client, serverID)
	ServerStart(client, serverID)
	ServerStop(client, serverID)
	ServerCreateFromVolume(client)
	ServerReboot(client, serverID)
	ServerResetMetadata(client, serverID)
	ServerMetadataList(client, serverID)
	ServerUpdateMetadata(client, serverID)
	ServerMetadataDetails(client, serverID, metaData)
	ServerDeleteMetadata(client, serverID, metaData)
	ServerAttachInterfaceList(client, serverID)
	ServerGetConsoleLog(client, serverID, length)
	ServerListInstanceActions(client, serverID)
	ServerGetInstanceActions(client, serverID, reqID)
	fmt.Println("main end...")

}

//Create server
func ServerCreate(client *gophercloud.ServiceClient) {
	createOpts := servers.CreateOpts{
		Name:      "createTestV2",
		FlavorRef: "c2.large",
		ImageRef:  "82372460-4f18-46a6-9315-25b8a606f847",
		Networks: []servers.Network{
			servers.Network{UUID: "cc7953b3-110f-4e87-b240-ff4915548875"},
		},
		AvailabilityZone: "kvmxen.dc1",
		KeyName:          "KeyPair-4037",
	}

	resp, createErr := servers.Create(client, keypairs.CreateOptsExt{
		CreateOptsBuilder: createOpts,
	}).Extract()
	if createErr != nil {
		fmt.Println("createErr:", createErr)
		if ue, ok := createErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(resp)
	fmt.Println("server create success!")

}

//List servers query server list
func ServersList(client *gophercloud.ServiceClient) {
	listOpt := servers.ListOpts{
		Name: "update",
	}
	// Query all servers list information
	allPages, allPagesErr := servers.List(client, listOpt).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform servers structure
	allServers, allServersErr := servers.ExtractServers(allPages)
	if allServersErr != nil {
		fmt.Println("allServersErr:", allServersErr)
		if ue, ok := allServersErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("servers list is :")
	for _, server := range allServers {
		serverJson, _ := json.MarshalIndent(server, "", " ")
		fmt.Println(string(serverJson))
	}
}

// ServersListV247 requests server details list with microversion 2.47
func ServersListV247(client *gophercloud.ServiceClient) {
	client.SetMicroversion("2.47")
	defer client.UnsetMicroversion()
	// Query all servers list information
	allPages, allPagesErr := servers.List(client, servers.ListOpts{}).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform servers structure
	allServers, allServersErr := servers.ExtractServers(allPages)
	if allServersErr != nil {
		fmt.Println("allServersErr:", allServersErr)
		return
	}
	fmt.Println("Servers list is :")
	for _, server := range allServers {
		serverJson, _ := json.MarshalIndent(server, "", " ")
		fmt.Println("Server info is :", string(serverJson))

		if vcpus, ok := server.Flavor["vcpus"].(float64); ok {
			fmt.Println("Flavor cpu is :", vcpus)
		}
		if ram, ok := server.Flavor["ram"].(float64); ok {
			fmt.Println("Flavor ram is :", ram)
		}
	}
}

//Get server details
func ServerGet(client *gophercloud.ServiceClient, serverId string) {
	server, serversErr := servers.Get(client, serverId).Extract()
	if serversErr != nil {
		fmt.Println("serversErr:", serversErr)
		if ue, ok := serversErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	serverJson, _ := json.MarshalIndent(server, "", " ")
	fmt.Println("server details is :" + string(serverJson))
}

//Update server
func ServerUpdate(client *gophercloud.ServiceClient, serverId string) {
	updateOpt := servers.UpdateOpts{
		Name: "updateServer",
	}
	resp, serversErr := servers.Update(client, serverId, updateOpt).Extract()
	if serversErr != nil {
		fmt.Println("serversErr:", serversErr)
		if ue, ok := serversErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(resp)
	fmt.Println("server update success!")
}

//Delete server
func ServerDelete(client *gophercloud.ServiceClient, serverId string) {
	deleteErr := servers.Delete(client, serverId).ExtractErr()
	if deleteErr != nil {
		fmt.Println("deleteErr:", deleteErr)
		if ue, ok := deleteErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("server delete success!")
}

//Rebuild server
func ServerRebuild(client *gophercloud.ServiceClient, serverId string) {
	rebuildOpt := servers.RebuildOpts{
		Name: "RebuildName",
		Metadata: map[string]string{
			"rebuild": "yes",
		},
		ImageID: "82372460-4f18-46a6-9315-25b8a606f847",
	}
	resp, rebuildErr := servers.Rebuild(client, serverId, rebuildOpt).Extract()
	if rebuildErr != nil {
		fmt.Println("rebuildErr", rebuildErr)
		if ue, ok := rebuildErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(resp)
	fmt.Println("server rebuild success!")
}

//Resize server
func ServerResize(client *gophercloud.ServiceClient, serverId string) {
	opts := &servers.ResizeOpts{
		FlavorRef: "c3.large.2",
	}
	resp := servers.Resize(client, serverId, opts)
	fmt.Println(resp)
	fmt.Println("resize server start ......")
}

//Confirmation reset operation
func ServerConfirmResize(client *gophercloud.ServiceClient, serverId string) {
	resp := servers.ConfirmResize(client, serverId)
	fmt.Println(resp)
	fmt.Println("server confirm resize success!")
}

//Revert resize server
func ServerRevertResize(client *gophercloud.ServiceClient, serverId string) {
	resp := servers.RevertResize(client, serverId)
	fmt.Println(resp)
	fmt.Println("server revert resize success!")
}

//Start the server
func ServerStart(client *gophercloud.ServiceClient, serverId string) {
	startErr := startstop.Start(client, serverId).ExtractErr()
	if startErr != nil {
		fmt.Println("startErr:", startErr)
		if ue, ok := startErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("server start success!")
}

//Close the server
func ServerStop(client *gophercloud.ServiceClient, serverId string) {
	stopErr := startstop.Stop(client, serverId).ExtractErr()
	if stopErr != nil {
		fmt.Println("stopErr:", stopErr)
		if ue, ok := stopErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("server stop success!")
}

//Create server from volumes
func ServerCreateFromVolume(client *gophercloud.ServiceClient) {
	blockDevices := []bootfromvolume.BlockDevice{
		bootfromvolume.BlockDevice{
			BootIndex:           0,
			DestinationType:     "volume",
			SourceType:          "image",
			VolumeSize:          40,
			UUID:                "ee5c7dc8-acb8-4d93-8d47-b27610b3477d",
			DeleteOnTermination: true,
		},
	}
	serverCreateOpts := servers.CreateOpts{
		Name:      "createTest",
		FlavorRef: "c2.large",
		ImageRef:  "82372460-4f18-46a6-9315-25b8a606f847",
		Networks: []servers.Network{
			{UUID: "cc7953b3-110f-4e87-b240-ff4915548875"},
		},
		AvailabilityZone: "kvmxen.dc1",
	}
	serverResult, serversErr := bootfromvolume.Create(client, bootfromvolume.CreateOptsExt{
		CreateOptsBuilder: serverCreateOpts,
		BlockDevice:       blockDevices,
	}).Extract()
	if serversErr != nil {
		fmt.Println("serversErr:", serversErr)
		if ue, ok := serversErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println(serverResult)
	fmt.Println("server create from volume success!")
}

//Reboot the server
func ServerReboot(client *gophercloud.ServiceClient, serverId string) {
	rebootOpts := &servers.RebootOpts{
		Type: servers.HardReboot,
	}
	resp := servers.Reboot(client, serverId, rebootOpts)
	fmt.Println(resp)
	fmt.Println("server reboot success!")
}

//Reset server metadata
func ServerResetMetadata(client *gophercloud.ServiceClient, serverId string) {
	metadata, resetMetadataErr := servers.ResetMetadata(client, serverId, servers.MetadataOpts{
		"Metadata": "testing",
	}).Extract()
	if resetMetadataErr != nil {
		fmt.Println("resetMetadataErr", resetMetadataErr)
		if ue, ok := resetMetadataErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	metadataJson, _ := json.MarshalIndent(metadata, "", " ")
	fmt.Println("metadata is :" + string(metadataJson))
	fmt.Println("server reset metadata success!")
}

//Query server metadata list
func ServerMetadataList(client *gophercloud.ServiceClient, serverId string) {
	metadata, metadataListErr := servers.Metadata(client, serverId).Extract()
	if metadataListErr != nil {
		fmt.Println("metadataListErr", metadataListErr)
		if ue, ok := metadataListErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	metadataJson, _ := json.MarshalIndent(metadata, "", " ")
	fmt.Println("metadata is :" + string(metadataJson))
}

//Update server metadata
func ServerUpdateMetadata(client *gophercloud.ServiceClient, serverId string) {
	metadata, updateMetadataErr := servers.UpdateMetadata(client, serverId, servers.MetadataOpts{
		"Metadata": "updateMetadatatesting",
	}).Extract()
	if updateMetadataErr != nil {
		fmt.Println("updateMetadataErr:", updateMetadataErr)
		if ue, ok := updateMetadataErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	metadataJson, _ := json.MarshalIndent(metadata, "", " ")
	fmt.Println(metadataJson)
	fmt.Println("server update metadata success!")
}

//Get server metadata details
func ServerMetadataDetails(client *gophercloud.ServiceClient, serverId string, metaData string) {
	metadata, metadataDetailsErr := servers.Metadatum(client, serverId, metaData).Extract()
	if metadataDetailsErr != nil {
		fmt.Println("metadataDetailsErr:", metadataDetailsErr)
		if ue, ok := metadataDetailsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	metadataJson, _ := json.MarshalIndent(metadata, "", " ")
	fmt.Println("metadata details is :" + string(metadataJson))
}

//Delete server metadata
func ServerDeleteMetadata(client *gophercloud.ServiceClient, serverId string, metaData string) {
	deleteErr := servers.DeleteMetadatum(client, serverId, metaData).ExtractErr()
	if deleteErr != nil {
		fmt.Println("deleteErr", deleteErr)
		if ue, ok := deleteErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("server delete metadata success!")
}

//Query server network card list
func ServerAttachInterfaceList(client *gophercloud.ServiceClient, serverId string) {
	// Query all attachinterfaces list information
	allPages, allPagesErr := attachinterfaces.List(client, serverId).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform attachinterfaces structure
	attachInterfaces, attachInterfacesErr := attachinterfaces.ExtractInterfaces(allPages)
	if attachInterfacesErr != nil {
		fmt.Println("attachInterfacesErr:", attachInterfacesErr)
		if ue, ok := attachInterfacesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("server attach interface list is :")
	for _, attachInterfaces := range attachInterfaces {
		attachInterfacesJson, _ := json.MarshalIndent(attachInterfaces, "", " ")
		fmt.Println(string(attachInterfacesJson))
	}
}

//ServerGetConsoleLog gets server console log
func ServerGetConsoleLog(sc *gophercloud.ServiceClient, serverID string, length string) {
	resp, err := servers.GetConsoleLog(sc, serverID, length).Extract()
	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server console log:", resp)
	fmt.Println("Server get console log success!")
}

//ServerListInstanceActions gets instance action list of a server
func ServerListInstanceActions(sc *gophercloud.ServiceClient, serverID string) {
	resp, err := servers.ListInstanceActions(sc, serverID).ExtractInstanceActionsListResult()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server list instance actions success!")

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))
}

//ServerGetInstanceActions gets instance action by request ID
func ServerGetInstanceActions(sc *gophercloud.ServiceClient, serverID string, reqID string) {
	resp, err := servers.GetInstanceActions(sc, serverID, reqID).ExtractInstanceActionsResult()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	fmt.Println("Server get instance actions success!")

	p, _ := json.MarshalIndent(*resp, "", " ")
	fmt.Println(string(p))

}
