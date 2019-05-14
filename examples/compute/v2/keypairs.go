package main

import (
	"github.com/gophercloud/gophercloud/auth/token"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud"
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"encoding/json"
	"crypto/rsa"
	"crypto/rand"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("main start...")
	//Set authentication parameters
	tokenOpts := token.TokenOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		Username:         "{Username}",
		Password:         "{Password}",
		DomainID:         "{DomainID}",
		ProjectID:        "{ProjectID}",
	}
	//Init provider client
	provider, authErr := openstack.AuthenticatedClient(tokenOpts)
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
	name := "{name}"
	KeyPairCreate(client)
	KeyPairGet(client, name)
	KeyPairDelete(client, name)
	KeyPairsList(client)
	fmt.Println("main end...")
}

//Create keypair
func KeyPairCreate(client *gophercloud.ServiceClient) {
	privateKey, keyPairsCreateKeyErr := rsa.GenerateKey(rand.Reader, 2048)
	if keyPairsCreateKeyErr != nil {
		fmt.Println("keyPairsCreateKeyErr:", keyPairsCreateKeyErr)
		if ue, ok := keyPairsCreateKeyErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	publicKey := privateKey.PublicKey
	pub, publicKeyErr := ssh.NewPublicKey(&publicKey)
	if publicKeyErr != nil {
		fmt.Println("publicKeyErr", publicKeyErr)
		if ue, ok := publicKeyErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	pubBytes := ssh.MarshalAuthorizedKey(pub)
	createOpts := keypairs.CreateOpts{
		Name:      "TestKey",
		PublicKey: string(pubBytes),
	}
	keyPair, keyPairsCreateErr := keypairs.Create(client, createOpts).Extract()
	fmt.Println(keyPair)
	if keyPairsCreateErr != nil {
		fmt.Println("keyPairsCreateErr:", keyPairsCreateErr)
		if ue, ok := keyPairsCreateErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("keypair create success!")
}

//Get keypair details
func KeyPairGet(client *gophercloud.ServiceClient, name string) {
	keyPair, keyPairGetErr := keypairs.Get(client, name).Extract()
	if keyPairGetErr != nil {
		fmt.Println("keyPairGetErr:", keyPairGetErr)
		if ue, ok := keyPairGetErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	keyPairJson, _ := json.MarshalIndent(keyPair, "", " ")
	fmt.Println("keypair detail is : " + string(keyPairJson))
}

//Delete keypair
func KeyPairDelete(client *gophercloud.ServiceClient, name string) {
	deleteErr := keypairs.Delete(client, name).ExtractErr()
	if deleteErr != nil {
		fmt.Println("deleteErr:", deleteErr)
		if ue, ok := deleteErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("keypair delete success!")
}

//Query keypairs list
func KeyPairsList(client *gophercloud.ServiceClient) {
	// Query all keypairs list information
	allPages, allPagesErr := keypairs.List(client).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform keypairs structure
	allKeys, allKeyPairsErr := keypairs.ExtractKeyPairs(allPages)
	if allKeyPairsErr != nil {
		fmt.Println("allKeyPairsErr:", allKeyPairsErr)
		if ue, ok := allKeyPairsErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	fmt.Println("keypairs list is : ")
	for _, keyPair := range allKeys {
		keyPairJson, _ := json.MarshalIndent(keyPair, "", " ")
		fmt.Println(string(keyPairJson))
	}
}
