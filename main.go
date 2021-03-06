package main

import (
	"fmt"
	"github.com/chainHero/servntire/blockchain"
	"os"
)

func main() {
	
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters 
		OrdererID: "orderer.servntire.io",

		// Channel parameters
		ChannelID:     "servntire",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/serventire/fixtures/artifacts/servntire.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "servntire",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainHero/servntire/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()	

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}
}
