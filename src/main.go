package main

import (
	"fmt"
	"github.com/UmarFarooq-MP/P2P-GO/src/peermanager"
	"time"
)

func main() {

	peerManager := peermanager.PeersManager{}
	peerManager.LoadConfigs("")
	peerManager.Listen()
	time.Sleep(time.Duration(10) * time.Second)
	fmt.Println(peerManager)
}
