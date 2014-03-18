// Author: sheppard(ysf1026@gmail.com) 2014-03-13

package service

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"net"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service/master"
)

type Harbor struct {
	masterConn net.Conn
}

func (s *Harbor) ClawCallback(session int, source string, msg interface{}) {
}

func (s *Harbor) ClawStart() {
	if !*center.IsMaster {
		go s.connect()
	}
}

func (s *Harbor) connect() {
	var err error
	s.masterConn, err = net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	login := &master.Login{"harbor1"}
	err = gob.NewEncoder(&buffer).Encode(login)
	if err != nil {
		panic(err)
	}

	var headBuffer bytes.Buffer
	binary.Write(&headBuffer, binary.BigEndian, uint16(2 + buffer.Len()))
	binary.Write(&headBuffer, binary.BigEndian, uint16(master.LOGIN))
	s.send(headBuffer.Bytes())
	s.send(buffer.Bytes())
}

func (s *Harbor) send(buf []byte) {
	_, err := s.masterConn.Write(buf)
	if err != nil {
		panic(err)
	}
}
