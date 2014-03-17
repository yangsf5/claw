// Author: sheppard(ysf1026@gmail.com) 2014-03-07

package main

import (
	"fmt"
	"time"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
	myService "github.com/yangsf5/claw/example/stress/service"
)

func main() {
	fmt.Println("Stress start!")

	service.Register()
	myService.Register()
	center.Use([]string{"Error", "Master", "Harbor", "Gate", "StressAdd"})

	for {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Stress exit!")
}

