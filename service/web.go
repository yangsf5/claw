package service

import (
	"github.com/golang/glog"

	"github.com/yangsf5/claw/service/web"
)

type Web struct {
}

func (s *Web) ClawCallback(session int, source string, msgType int, msg interface{}) {
}

func (s *Web) ClawStart() {
	glog.Infof("Claw.Web service start")
	go web.Start()
}

