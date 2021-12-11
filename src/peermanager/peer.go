package peermanager

import "github.com/UmarFarooq-MP/P2P-GO/src/socket"

type Peer struct {
	socket   socket.Socket
	Messages []PeerMessage
}
