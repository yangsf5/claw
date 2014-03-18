// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package service

import (
	"encoding/xml"
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

type MasterConfigPack struct {
	XMLName xml.Name `xml:"clawconfig"`
	Master MasterConfig `xml:"master"`
}

type MasterConfig struct {
	ListenAddr string `xml:"listenAddr,attr"`
}

func (s *Master) Listen() {
	var config MasterConfigPack
	center.GetConfig(&config)

	tcpAddr, err := net.ResolveTCPAddr("tcp", config.Master.ListenAddr)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Service.Master listening", config.Master.ListenAddr)

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

