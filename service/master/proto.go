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
	addNode(node)
}

type Command struct {
	Name string
}

var (
	handlers map[uint16]Handler
)

func init() {
	handlers = make(map[uint16]Handler)

	handlers[LOGIN] = &Login{}
}
