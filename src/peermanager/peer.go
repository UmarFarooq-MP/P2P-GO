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
	status   bool
}

func (peer *Peer) Connect(ip string, port string) bool {
	peer.socket, _ = net.Dial("tcp", ip+":"+port)
	return true
}

func (peer *Peer) Disconnect() {
	peer.socket.Close()
}

func (peer *Peer) Read(newMessageChannel chan string, onDisconnectChannel chan string) {
	buffer := make([]byte, 512)
	peer.socket.SetReadDeadline(time.Now().Add(time.Millisecond * 5))
	n, err := peer.socket.Read(buffer)
	message := string(buffer[:n])

	if err == io.EOF {
		onDisconnectChannel <- "Peer Disconnected"
		peer.status = false
		fmt.Println("HAHAHAHAH")
	}
	if n > 0 {
		newMessageChannel <- message
	}

}
