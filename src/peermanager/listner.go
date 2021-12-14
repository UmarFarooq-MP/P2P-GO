package peermanager

import (
	"net"
)

type Listener struct {
	ip            string
	port          string
	status        bool
	socket        net.Listener
	newConnection []net.Conn
}

func ConstructListener(ip string, port string, status bool) Listener {
	return Listener{ip, port, status, nil, nil}
}

func (listener *Listener) Listen() {
	listener.socket, _ = net.Listen("tcp", listener.ip+":"+listener.port)
}

func (listener *Listener) Accept() {
	newConnection, error := listener.socket.Accept()
	if error != nil {
		listener.newConnection = append(listener.newConnection, newConnection)
	}
}
