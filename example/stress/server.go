// Author: sheppard(ysf1026@gmail.com) 2014-03-07

package main

import (
	"fmt"
	"time"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
)

func main() {
	fmt.Println("Mock server of stress start!")

	service.Register()
	center.Use([]string{"Error", "Test", "Gate"})

	center.Send("haha", "Test", 1, []byte("hello, test service"))
	center.Send("haha", "Error", 1, []byte("sth. is wrong"))

	for {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Claw exit!")
}

