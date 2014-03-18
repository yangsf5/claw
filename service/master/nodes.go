// Author: sheppard(ysf1026@gmail.com) 2014-03-17

package master

import (
	"github.com/yangsf5/claw/engine/net"
)

var (
	nodes *net.Group
)

func init() {
	nodes = net.NewGroup()
}


