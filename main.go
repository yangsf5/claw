// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package main

import (
	"fmt"

	"github.com/yangsf5/claw/service"
)

func main() {
	var services map[string]func()
	services = make(map[string]func())
	services["test"] = service.Test

	services["test"]()

	fmt.Println("hell, claw!")
}

