// Author: sheppard(ysf1026@gmail.com) 2014-03-17

package master

import (
)

var (
	nodes map[string]*Node
)

func init() {
	nodes = make(map[string]*Node)
}

func addNode(node *Node) {
	//TODO check repeat
	nodes[node.Name] = node
}

func delNode(node *Node) {
	delete(nodes, node.Name)
}
