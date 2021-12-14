package main

import (
	"fmt"
	"github.com/UmarFarooq-MP/P2P-GO/src/peermanager"
	"time"
)

func main() {

	var peerManager *peermanager.PeersManager = peermanager.GetInstance()
	peerManager.ConstructListener("127.0.0.1", "9999", false)
	peerManager.Listen()
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Accepting Connection .............")
		peerManager.Accept()
	}()

	for {
		time.Sleep(2 * time.Second)
		fmt.Println("Reading Messages .............")
		//peerManager.ReadMessages()
	}

}
