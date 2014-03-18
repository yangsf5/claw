// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master


type Handler interface {
	Handle(node *Node)
}

const (
	LOGIN = iota
)

type Login struct {
	Name string
}

func (m *Login) Handle(node *Node) {
	nodes.AddPeer(node.Name, node)
}

type Command struct {
	Name string
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
}
