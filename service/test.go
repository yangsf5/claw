package service

import (
	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
)

type Test struct {
}

func (s* Test) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Infof("Service.Test, session=%v source=%v, msg=%s", session, source, msg)
	send("Test", "Error", 1, center.MsgTypeText, "this from test");
}

func (s* Test) ClawStart() {
	glog.Info("Service.Test, funcion Start is called, test passes.")
}
