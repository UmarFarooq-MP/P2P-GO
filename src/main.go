package main

import (
	"github.com/UmarFarooq-MP/P2P-GO/src/peermanager"
)

func main() {

	peerManager := peermanager.PeersManager{}
	peerManager.LoadConfigs("")
	go peerManager.Listen()

}
