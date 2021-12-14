package peermanager

import (
	"net"
)

var peersManagerInstance *PeersManager

type PeersManager struct {
	listener           Peer
	connectedPeersList []Peer
}

func getInstance() *PeersManager {
	if peersManagerInstance == nil {
		peersManagerInstance = &PeersManager{Peer{"", "", nil, nil}, nil}
		return peersManagerInstance
	}
	return peersManagerInstance
}

func (peers *PeersManager) LoadConfigs(path string) {
	//TODO::once YML is ready we will up it
	//for now lets use dummy values
	//peers.listener.socket = socket.ConstructSocket("127.0.0.1", "9999")
}

func (peers *PeersManager) AddNewPeerToList(connection net.Conn) {
	//peers.connectedPeersList = append(peers.connectedPeersList, Peer{socket: socket})
}

func (peers *PeersManager) Listen() {
	//peers.listener.socket.Listen()
}

func Connect(ip string, port string) {

}
