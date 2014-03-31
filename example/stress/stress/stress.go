// Author: sheppard(ysf1026@gmail.com) 2014-03-07

package main

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
	myService "github.com/yangsf5/claw/example/stress/stress/service"
)

func main() {
	fmt.Println("Stress start!")

	service.Register()
	myService.Register()
	center.Use([]string{"Error", "Master", "Harbor", "StressAdd"})

	if center.BaseConfig.Master.IsMaster {
		for {
			prompt()
			var i int
			fmt.Scanf("%d", &i)
			fmt.Println()
			switch(i) {
			case 1:
				fmt.Println("show")
			case 2:
				center.Send("main", "Master", 0, center.MsgTypeHarbor, &service.RemoteMessage{"StressAdd", center.MsgTypeHarbor, []byte("start")})
			case 3:
				center.Send("main", "Master", 0, center.MsgTypeHarbor, &service.RemoteMessage{"StressAdd", center.MsgTypeHarbor, []byte("stop")})
			default:
				fmt.Println("Unkown operation.")
			}
			fmt.Println()
		}
	} else {
		for {
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("Stress exit!")
	glog.Flush()
}

func prompt() {
	fmt.Println("Prompt:")
	fmt.Println("1 - show")
	fmt.Println("2 - start")
	fmt.Println("3 - stop")
	fmt.Print("\n>")
}
