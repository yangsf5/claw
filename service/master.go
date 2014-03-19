// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package service

import (
	"fmt"
	"net"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service/master"
)


type Master struct {
}

func (s *Master) ClawCallback(session int, source string, msg interface{}) {
	master.HandleClawCallback(msg)
}

func (s *Master) ClawStart() {
	if *center.IsMaster {
		go s.Listen()
	} else {
		fmt.Println("this server is not master, so service.master not start")
	}
}

func (s *Master) Listen() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", center.BaseConfig.Master.ListenAddr)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Service.Master listening", center.BaseConfig.Master.ListenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Service.Master err", err.Error())
			continue
		}
		fmt.Println("Service.Master new connection")

		master.HandleConnection(conn)
	}
}

