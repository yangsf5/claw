package master

import (
	"net"
)

func HandleConnection(conn net.Conn) {
	node := NewNode(conn)
	go node.Handle()
}

func Broadcast(content []byte) {
	nodes.Broadcast(content)
}
