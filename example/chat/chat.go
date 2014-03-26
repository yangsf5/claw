// Author: sheppard(ysf1026@gmail.com) 2014-03-25

package main

import (
	"net"
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
)

var (
)

func main() {
	service.Register()
	service.GateRegisterConnHandler(connHandle)
	center.Use([]string{"Error", "Master", "Harbor", "Gate"})

	glog.Info("Chat start!")

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Chat exit!")
	glog.Flush()
}

func connHandle(conn net.Conn) {
	glog.Info(conn)
}
