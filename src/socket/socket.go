package socket

import (
	"fmt"
	"net"
)

type Socket struct {
	IP     string
	Port   string
	Status bool
}

func ConstructSocket(ip string, port string) Socket {
	return Socket{ip, port, false}
}

func (s *Socket) Listen() {
	fmt.Printf("Started Listening on %v : %v", s.IP, s.Port)
	net.Listen("tcp", s.IP+":"+s.Port)
}

func (s *Socket) Connect(ip string, port string) {}

func (s *Socket) DisConnect(ip string, port string) {}
