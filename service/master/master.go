// Author: sheppard(ysf1026@gmail.com) 2014-03-17

package master

import (
	"net"
)

func HandleConnection(conn net.Conn) {
	node := NewNode(conn)
	go node.Handle()
}

func HandleHarborMsg(msg interface{}) {
	if handler, ok := msg.(Handler); ok {
		//TODO tidy this type handle
		handler.Handle(nil)
	}
}

func Broadcast(content []byte) {
	nodes.Broadcast(content)
}
