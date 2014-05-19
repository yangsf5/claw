// Author: sheppard(ysf1026@gmail.com) 2014-05-14

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
	web.Start()
}

