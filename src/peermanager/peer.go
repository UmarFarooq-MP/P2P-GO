package peermanager

import (
	"fmt"
	"io"
	"net"
	"time"
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
func ConstructPeer(ip string, port string) *Peer {
	return &Peer{ip, port, nil, nil}
}

func (peer *Peer) Read() {
	one := make([]byte, 1)
	peer.socket.SetReadDeadline(time.Now())
	if _, err := peer.socket.Read(one); err == io.EOF {
		fmt.Printf("Diconnected with address %v\n", peer.socket.RemoteAddr().String())
	}
}
