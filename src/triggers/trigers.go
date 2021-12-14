package triggers

import (
	"fmt"
	//"github.com/UmarFarooq-MP/P2P-GO/src/peermanager"
	"net"
)

func OnNewConnection(conn net.Conn) {
	fmt.Printf("New Connection with address :%v \n", conn.RemoteAddr().String())
	//peermanager.GetInstance().AddNewPeerToList(&conn)
}

func OnDisConnect(conn net.Conn) {
	fmt.Printf("Connection Disconnected with address :%v \n", conn.RemoteAddr().String())
}
