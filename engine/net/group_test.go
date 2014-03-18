// Author: sheppard(ysf1026@gmail.com) 2014-01-28

package net

import (
	"testing"
)

type PeerTmp struct {
	uid string
}

func (u *PeerTmp) Send(msg []byte) {
}


var (
	g *Group
)

func init() {
	g = NewGroup()
}


func TestGroup(t *testing.T) {
	g.AddPeer("user1", &PeerTmp{"user1"})
	g.AddPeer("user2", &PeerTmp{"user2"})
}
