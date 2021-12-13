package peermanager

import "net"

type Listener struct {
	ip     string
	port   string
	status bool
}

func ConstructListener(ip string, port string, status bool) Listener {
	return Listener{ip, port, status}
}

func (listener *Listener) Listen() {
	net.Listen("tcp", listener.ip+":"+listener.port)
}
