package peermanager

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var PeersManagerInstance *PeersManager

type PeersManager struct {
	ListenerSocket     Listener
	ConnectedPeersList []Peer
}

func (peers *PeersManager) ReadMessages(newMessageChannel chan string, onDisconnectChannel chan string) {
	for {
		time.Sleep(4 * time.Second)
		fmt.Println("Reading Messages.................")
		for index, _ := range peers.ConnectedPeersList {
			if peers.ConnectedPeersList[index].status == true {
				peers.ConnectedPeersList[index].Read(newMessageChannel, onDisconnectChannel)
			}
		}
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
	peers.ListenerSocket.ConstructListener(ip, port, status)
}

func (peers *PeersManager) LoadConfigs(path string) {
	//TODO::once YML is ready we will up it
	//for now lets use dummy values
	//peers.listener.socket = socket.ConstructSocket("127.0.0.1", "9999")
}

func (peers *PeersManager) AddNewPeerToList(connection net.Conn) {
	connectionString := connection.RemoteAddr().String()
	colunIndex := strings.IndexByte(connectionString, ':')
	peers.ConnectedPeersList = append(peers.ConnectedPeersList, Peer{connectionString[:colunIndex], connectionString[colunIndex:], connection, nil, true})
}

func (peers *PeersManager) Listen() {
	peers.ListenerSocket.Listen()
}

func (peers *PeersManager) Accept(newConnectionChannel chan net.Conn) {
	for {
		fmt.Println("Accepting Connection .............")
		peers.ListenerSocket.Accept(newConnectionChannel)
	}
}

func (peers *PeersManager) Connect(ip string, port string) {

}

func (peers *PeersManager) Processor() {
	newConnectionChannel := make(chan net.Conn)
	newMessageChannel := make(chan string)
	onDisconnectChannel := make(chan string)
	go peers.Accept(newConnectionChannel)
	go peers.ReadMessages(newMessageChannel, onDisconnectChannel)
	for {
		time.Sleep(1 * time.Second)
		select {
		case newConnection := <-newConnectionChannel:
			fmt.Printf("New Connection Arrived with address %v \n", newConnection.RemoteAddr())
			peers.AddNewPeerToList(newConnection)
		case newMessage := <-newMessageChannel:
			fmt.Printf("New Message: %v", newMessage)
		case peerDisconnect := <-onDisconnectChannel:
			fmt.Printf("Peer Disconnected with with Name %v", peerDisconnect)
		}
	}
}
