// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package main

import (
	"fmt"
	"time"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
)

var (
)

func main() {
	fmt.Println("Claw start!")

	service.Register()
	center.Use([]string{"Error", "Master", "Harbor", "Test", "Gate"})

	center.Send("main", "Test", 1, "hello, test service")
	center.Send("main", "Error", 1, "sth. is wrong")

	for {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Claw exit!")
}

