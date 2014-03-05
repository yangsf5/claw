// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package main

import (
	"fmt"
	"time"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
)

func main() {
	center.Register("Test", service.Test)
	center.Register("Error", service.Error)

	center.Send("haha", "Test", 1, []byte("hello, test service"))
	center.Send("haha", "Error", 1, []byte("sth. is wrong"))

	fmt.Println("hello, claw!")

	for {
		time.Sleep(100 * time.Millisecond)
	}
}

