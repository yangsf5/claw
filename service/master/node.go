// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"net"
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
	for {
		var sizeBuf [2]byte
		_, err := n.conn.Read(sizeBuf[:])
		if err != nil {
			fmt.Printf("packet id error, err=%s\n", err.Error())
			break
		}
		var size uint16
		binary.Read(bytes.NewBuffer(sizeBuf[:]), binary.BigEndian, &size)

		bodyBuf := make([]byte, size)
		_, err = n.conn.Read(bodyBuf[:])
		if err != nil && err != io.EOF {
			break
		}

		var packetId uint16
		binary.Read(bytes.NewBuffer(bodyBuf), binary.BigEndian, &packetId)
		handler, ok := handlers[packetId]
		if !ok {
			fmt.Printf("packet id error, pid=%d\n", packetId)
			continue
		}
		msgBuf := bytes.NewBuffer(bodyBuf[2:])
		err = gob.NewDecoder(msgBuf).Decode(handler)
		if err != nil {
			fmt.Printf("packet decode error, pid=%d\n", packetId)
			continue
		}

		fmt.Println("packet is", handler)

		handler.Handle(n)
	}

	nodes.DelPeer(n.Name)
	fmt.Println("Node die")
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

