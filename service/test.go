// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package service

import (
	"fmt"
)

type Test struct {
}

func (s* Test) ClawCallback(session int, source string, msg interface{}) {
	fmt.Printf("Service.Test, session=%v source=%v, msg=%s\n", session, source, msg)
	send("Test", "Error", 1, "this from test");
}

func (s* Test) ClawStart() {
	fmt.Println("Service.Test, funcion Start is called, test passes.")
}
