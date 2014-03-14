// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
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
			break
		}
		var size uint16
		binary.Read(bytes.NewBuffer(sizeBuf[:]), binary.BigEndian, &size)

		bodyBuf := make([]byte, size)
		_, err = n.conn.Read(bodyBuf[:])
		if err != nil {
			break
		}

		var packetId uint16
		binary.Read(bytes.NewBuffer(bodyBuf), binary.BigEndian, &packetId)
		packetFunc, ok := packets[packetId]
		if !ok {
			fmt.Printf("packet id error, pid=%d\n", packetId)
			continue
		}
		packet := packetFunc()
		msgBuf := bytes.NewBuffer(bodyBuf[2:])
		err = gob.NewDecoder(msgBuf).Decode(packet)
		if err != nil {
			fmt.Printf("packet decode error, pid=%d\n", packetId)
			continue
		}

		fmt.Println("packet is", packet)
	}
}


var (
	nodes map[string]*Node
)

func init() {
	nodes = make(map[string]*Node)
}

func HandleConnection(conn net.Conn) {
	node := NewNode(conn)
	go node.Handle()
}
