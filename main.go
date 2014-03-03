// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package main

import (
	"fmt"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service"
)

func main() {
	center.Register("test", service.Test)

	center.Send("haha", "test", 1, []byte("hello, test service"))

	fmt.Println("hell, claw!")
}

