// Author: sheppard(ysf1026@gmail.com) 2014-03-06

package service

import (
	"fmt"
	"net"
)

type Gate struct {
}


func (s* Gate) ClawCallback(session int, source string, msg []byte) {
}

func (s* Gate) ClawStart() {
//	go gateListen()
}

func gateListen() {
	//TODO config addr
	addr := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println(conn)
	}
}
