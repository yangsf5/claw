// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package main

import (
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
)

var (
)

func main() {
	glog.Info("Claw start!")

	service.Register()
	center.Use([]string{"Error", "Master", "Harbor", "Test"})

	center.Send("main", "Test", 1, center.MsgTypeText, "hello, test service")
	center.Send("main", "Error", 1, center.MsgTypeText, "sth. is wrong")

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Claw exit!")
	glog.Flush()
}

