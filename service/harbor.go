// Author: sheppard(ysf1026@gmail.com) 2014-03-13

package service

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"net"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/service/master"
)

type Harbor struct {
	masterConn net.Conn
}

func (s *Harbor) ClawCallback(session int, source string, msgType int, msg interface{}) {
}

func (s *Harbor) ClawStart() {
	if !*center.IsMaster {
		go s.connect()
	}
}

func (s *Harbor) connect() {
	var err error
	s.masterConn, err = net.Dial("tcp", center.BaseConfig.Master.ListenAddr)
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

	s.recv()
}

func (s *Harbor) send(buf []byte) {
	_, err := s.masterConn.Write(buf)
	if err != nil {
		panic(err)
	}
}

func (s *Harbor) recv() {
	for {
		var sizeBuf [2]byte
		_, err := s.masterConn.Read(sizeBuf[:])
		if err != nil {
			break
		}
		var size uint16
		binary.Read(bytes.NewBuffer(sizeBuf[:]), binary.BigEndian, &size)

		msg := make([]byte, size)
		_, err = s.masterConn.Read(msg[:])
		if err != nil && err != io.EOF {
			break
		}

		remoteMsg := new(RemoteMessage)
		if err = gob.NewDecoder(bytes.NewBuffer(msg)).Decode(remoteMsg); err != nil {
			fmt.Println("Remote message decode error")
		}

		handleBroadcast(remoteMsg)
	}
}

func handleBroadcast(msg *RemoteMessage) {
	send("Harbor", msg.Destination, 0, msg.MessageType, msg.Message)
}

