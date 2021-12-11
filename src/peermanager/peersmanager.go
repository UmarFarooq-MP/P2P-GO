package peermanager

import "github.com/UmarFarooq-MP/P2P-GO/src/socket"

type PeersManager struct {
	listener  Peer
	peersList []Peer
}

func (peers *PeersManager) LoadConfigs(path string) {
	//TODO::once YML is ready we will up it
	//for now lets use dummy values
	peers.listener.socket = socket.ConstructSocket("127.0.0.1", "9999")
}

func (peers *PeersManager) addNewPeerToList(socket socket.Socket) {
	peers.peersList = append(peers.peersList, Peer{socket: socket})
}

func (peers *PeersManager) Listen() {
	peers.listener.socket.Listen()
}
