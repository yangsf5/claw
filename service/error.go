// Author: sheppard(ysf1026@gmail.com) 2014-03-05

package service

import (
	"fmt"
)

func errorReportCallback(session int, source string, msg []byte) {
	fmt.Printf("Error, session=%v source=%v msg=[%s]\n", session, source, msg)
}
