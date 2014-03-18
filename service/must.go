// Author: sheppard(ysf1026@gmail.com) 2014-03-06

package service

import (
	"github.com/yangsf5/claw/center"
)


func Register() {
	services := map[string]center.Service{
		"Master": &Master{},
		"Harbor": &Harbor{},
		"Error": &Error{},
		"Test": &Test{},
		"Gate": &Gate{},
	}

	for name, cb := range services {
		center.Register(name, cb)
	}
}

func send(source, destination string, session int, msg interface{}) {
	center.Send(source, destination, session, msg)
}
