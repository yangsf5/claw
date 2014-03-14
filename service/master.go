// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package service

import (
	"net"

	"github.com/yangsf5/claw/service/master"
)


type Master struct {

}

func (s *Master) ClawCallback(session int, source string, msg []byte) {
}

func (s *Master) ClawStart() {
	go s.Listen()
}

func (s *Master) Listen() {
	addr := ":8889"
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

		master.HandleConnection(conn)
	}
}

