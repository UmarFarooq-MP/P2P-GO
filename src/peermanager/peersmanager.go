package peermanager

import (
	"net"
)

var PeersManagerInstance *PeersManager

type PeersManager struct {
	listener           Listener
	connectedPeersList []Peer
}

func (peers *PeersManager) ReadMessages() {
	for _, peer := range peers.connectedPeersList {
		peer.Read()
	}
}

func GetInstance() *PeersManager {
	if PeersManagerInstance == nil {
		PeersManagerInstance = &PeersManager{Listener{"", "", false, nil, nil}, nil}
		return PeersManagerInstance
	}
	return PeersManagerInstance
}

func (peers *PeersManager) ConstructListener(ip string, port string, status bool) {
	peers.listener.ConstructListener(ip, port, status)
}

func (peers *PeersManager) LoadConfigs(path string) {
	//TODO::once YML is ready we will up it
	//for now lets use dummy values
	//peers.listener.socket = socket.ConstructSocket("127.0.0.1", "9999")
}

func (peers *PeersManager) AddNewPeerToList(connection *net.Conn) {
	peers.connectedPeersList = append(peers.connectedPeersList, Peer{})
}

func (peers *PeersManager) Listen() {
	peers.listener.Listen()
}

func (peers *PeersManager) Accept() {
	peers.listener.Accept()
}

func Connect(ip string, port string) {

}
