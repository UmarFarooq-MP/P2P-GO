package peermanager

import (
	"net"
)

type Peer struct {
	ip       string
	port     string
	socket   net.Conn
	Messages []string
}

func (peer *Peer) Connect(ip string, port string) bool {
	peer.socket, _ = net.Dial("tcp", ip+":"+port)
	return true
}

func (peer *Peer) Disconnect() {
	peer.socket.Close()
}

func ConstructPeer(ip string, port string) Peer {
	return Peer{ip, port, nil, nil}
}
