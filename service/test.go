// Author: sheppard(ysf1026@gmail.com) 2014-03-02

package service

import (
	"fmt"
)

func Test(session int, source string, msg []byte) {
	fmt.Printf("Service.test, session=%v source=%v, msg=%s\n", session, source, msg)
}

