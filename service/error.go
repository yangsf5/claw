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
