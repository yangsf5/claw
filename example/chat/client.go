// Author: sheppard(ysf1026@gmail.com) 2014-03-26

package main

import (
	"net"
	"io/ioutil"

	clawNet "github.com/yangsf5/claw/engine/net"
)

var (
	conn net.Conn
)

func main() {
}

func connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:11001")
	checkError(err)
	recv()
}

func send(msg []byte) {
	_, err := conn.Write(buf)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func recv() {
	//TODO clawNet.RecvLoop
}
