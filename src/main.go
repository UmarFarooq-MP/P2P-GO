package main

import (
	"github.com/UmarFarooq-MP/P2P-GO/src/peermanager"
	"time"
)

func main() {

	peerManager := peermanager.GetInstance()
	peerManager.ConstructListener("127.0.0.1", "9999", false)
	peerManager.Listen()
	go peerManager.Processor()

	for {
		time.Sleep(3 * time.Second)
	}

}
