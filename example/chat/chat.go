// Author: sheppard(ysf1026@gmail.com) 2014-03-25

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
	glog.Info("Chat start!")

	service.Register()
	center.Use([]string{"Error", "Master", "Harbor", "Gate"})

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Chat exit!")
	glog.Flush()
}

