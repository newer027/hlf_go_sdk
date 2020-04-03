package main

import (
	"fmt"
	"github.com/chainHero/heroes-service/blockchain"
	"github.com/chainHero/heroes-service/web"
	"github.com/chainHero/heroes-service/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters 
		OrdererID: "orderer.example.com",

		// Channel parameters
		ChannelID:     "mychannel",
		ChannelConfig: "/Users/Jacob/Desktop/fabric-samples-master/first-network/channel-artifacts/channel.tx",

		// Chaincode parameters
		ChainCodeID:     "marbles",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "/Users/Jacob/Desktop/fabric-samples-master/chaincode/marbles02/go",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "/Users/Jacob/go/src/github.com/chainHero/heroes-service/config.yaml",

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
	// err = fSetup.InstallAndInstantiateCC()
	// if err != nil {
	// 	fmt.Printf("Unable to")
	// 	return
	// }
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}