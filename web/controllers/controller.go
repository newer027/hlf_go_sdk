package controllers
import (
	"github.com/chainHero/heroes-service/blockchain"
)
type Application struct {
	Fabric *blockchain.FabricSetup
}