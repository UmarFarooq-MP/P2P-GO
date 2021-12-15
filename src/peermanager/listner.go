package peermanager

import (
	"fmt"
	"net"
)

type Listener struct {
	ip            string
	port          string
	status        bool
	socket        net.Listener
	newConnection []net.Conn
}

func (listener *Listener) ConstructListener(ip string, port string, status bool) {
	listener.ip = ip
	listener.port = port
	listener.status = false
}

func (listener *Listener) Listen() {
	listener.socket, _ = net.Listen("tcp", listener.ip+":"+listener.port)
}

func (listener *Listener) Accept(newConnectionChannel chan net.Conn) {
	newConnection, error := listener.socket.Accept()
	if error != nil {
		fmt.Println("NO :(")
	} else {
		newConnectionChannel <- newConnection
	}
}
