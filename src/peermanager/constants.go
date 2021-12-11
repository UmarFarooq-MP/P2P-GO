package peermanager

type MSGTYPE int

const (
	BROADCAST  MSGTYPE = 0x007B
	DBENTRY            = 0x007C
	DISCONNECT         = 0x007D
)
