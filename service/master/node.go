// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master

import (
	"bytes"
	"bufio"
	"encoding/binary"
	"encoding/gob"
	"net"

	"github.com/golang/glog"
	myNet "github.com/yangsf5/claw/engine/net"
)

type Node struct {
	Name string
	conn net.Conn
}

func NewNode(conn net.Conn) *Node {
	return &Node{
		conn: conn,
	}
}

func (n *Node) Handle() {
	defer n.conn.Close()

	cb := func(reader *bufio.Reader, err error) {
		if err != nil {
			nodes.DelPeer(n.Name)
			glog.Info("Node die")
			return
		}

		// Check size
		sizeBuf, err := reader.Peek(2)
		if err != nil {
			return
		}
		var size uint16
		binary.Read(bytes.NewBuffer(sizeBuf), binary.BigEndian, &size)
		if _, err = reader.Peek(2+int(size)); err != nil {
			return
		}

		// Read Body
		packBuf := make([]byte, 2 + int(size))
		reader.Read(packBuf)
		bodyBuf := packBuf[2:]

		// Parse packet id from body
		var packetId uint16
		binary.Read(bytes.NewBuffer(bodyBuf), binary.BigEndian, &packetId)
		handlerFunc, ok := handlers[packetId]
		if !ok {
			glog.Errorf("packet id error, pid=%d", packetId)
			return
		}

		// Decode message from body
		handler := handlerFunc()
		msgBuf := bytes.NewBuffer(bodyBuf[2:])
		err = gob.NewDecoder(msgBuf).Decode(handler)
		if err != nil {
			glog.Errorf("packet decode error, pid=%d err=%s", packetId, err.Error())
			return
		}

		glog.Infof("packet is %v", handler)

		handler.Handle(n)
	}

	myNet.RecvLoop(n.conn, cb)
}

func (n *Node) Send(msg []byte) {
	var headBuffer bytes.Buffer
	binary.Write(&headBuffer, binary.BigEndian, uint16(len(msg)))
	_, err := n.conn.Write(headBuffer.Bytes())
	if err != nil {
		//TODO
	}
	_, err = n.conn.Write(msg)
	if err != nil {
		//TODO
	}
}

