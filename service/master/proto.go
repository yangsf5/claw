// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master

import (
	"fmt"
)

type Handler interface {
	Handle(node *Node)
}

const (
	LOGIN = iota
	BROADCAST
)

type Login struct {
	Name string
}

func (m *Login) Handle(node *Node) {
	if ret := nodes.AddPeer(node.Name, node); !ret {
		fmt.Println("master.proto add peer fail, repeated")
	}
}

type Broadcast struct {
	Content []byte
}

func (m *Broadcast) Handle(node *Node) {
	nodes.Broadcast(m.Content)
}

var (
	handlers map[uint16]Handler
)

func init() {
	handlers = make(map[uint16]Handler)

	handlers[LOGIN] = &Login{}
	handlers[BROADCAST] = &Broadcast{}
}
