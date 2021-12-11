package socket

import "net"

type Socket struct {
	IP     string
	port   string
	status bool
}

func (s *Socket) Listen(ip string, port string) {
	s.port = port
	s.IP = ip
	s.status = false
	net.Listen("tcp", ip+":"+port)
}

func (s *Socket) Connect(ip string, port string) {}

func (s *Socket) DisConnect(ip string, port string) {}
