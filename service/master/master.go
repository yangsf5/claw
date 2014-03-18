// Author: sheppard(ysf1026@gmail.com) 2014-03-17

package master

import (
	"net"
)

func HandleConnection(conn net.Conn) {
	node := NewNode(conn)
	go node.Handle()
}

func HandleClawCallback(msg interface{}) {
	if handler, ok := msg.(Handler); ok {
		//TODO tidy this type handle
		handler.Handle(nil)
	}
}
