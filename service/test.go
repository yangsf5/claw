// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package service

import (
	"fmt"

	"github.com/yangsf5/claw/center"
)

func Test(session int, source string, msg []byte) {
	fmt.Printf("Service.Test, session=%v source=%v, msg=%s\n", session, source, msg)
	center.Send("Test", "Error", 1, []byte("this from test"));
}

