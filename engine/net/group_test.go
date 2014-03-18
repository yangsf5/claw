// Author: sheppard(ysf1026@gmail.com) 2014-01-28

package net

import (
	"testing"
)

type UserTmp struct {
	uid string
}

func (u *UserTmp) Send(msg string) {
}


var (
	g *Group
)

func init() {
	g = NewGroup()
}


func TestGroup(t *testing.T) {
	g.AddUser("user1", &UserTmp{"user1"})
	g.AddUser("user2", &UserTmp{"user2"})
}
