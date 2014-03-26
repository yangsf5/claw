// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package service

import (
	"bytes"
	"encoding/gob"
	"net"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service/master"
)

type RemoteMessage struct {
	Destination string
	MessageType int
	Message []byte
}

type Master struct {
}

func (s *Master) ClawCallback(session int, source string, msgType int, msg interface{}) {
	glog.Infof("Service.Master recv type=%v msg=%v", msgType, msg)
	switch msgType {
	case center.MsgTypeHarbor:
		if concrete, ok := msg.(*RemoteMessage); ok {
			var buffer bytes.Buffer
			if err := gob.NewEncoder(&buffer).Encode(concrete); err != nil {
				glog.Error("err" + err.Error())
				break
			}
			master.Broadcast(buffer.Bytes())
		}
	}
}

func (s *Master) ClawStart() {
	if center.BaseConfig.Master.IsMaster {
		go s.Listen()
	} else {
		glog.Info("this server is not master, so service.master not start")
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

	glog.Infof("Service.Master listening, addr=%s", center.BaseConfig.Master.ListenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			glog.Errorf("Service.Master accept error, err=%s", err.Error())
			continue
		}
		glog.Info("Service.Master new connection")

		master.HandleConnection(conn)
	}
}

