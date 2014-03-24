// Author: sheppard(ysf1026@gmail.com) 2014-03-05

package service

import (
	"github.com/golang/glog"
)

type Error struct {
}

func (s *Error) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Errorf("Error, session=%v source=%v msg=[%s]\n", session, source, msg)
}

func (s *Error) ClawStart() {
}
