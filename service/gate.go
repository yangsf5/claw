// Author: sheppard(ysf1026@gmail.com) 2014-03-06

package service

import (
	"net"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
)

var (
	gateConnHandler func(conn net.Conn)
)

func GateRegisterConnHandler(handler func(conn net.Conn)) {
	gateConnHandler = handler
}

type Gate struct {
}

func (s *Gate) ClawCallback(session int, source string, msgType int, msg interface{}) {
}

func (s *Gate) ClawStart() {
	go s.Listen()
}

func (s *Gate) Listen() {
	if gateConnHandler == nil {
		panic("Service.Gate lack gateConnHandler")
	}

	addr := center.BaseConfig.Gate.ListenAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	glog.Infof("Service.Gate listening, addr=%s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			glog.Errorf("Service.Gate accept error, err=%s", err.Error())
			continue
		}
		glog.Info("Service.Gate new connection")

		gateConnHandler(conn)
	}
}

