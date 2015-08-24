package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/yangsf5/claw/center"
	clawNet "github.com/yangsf5/claw/engine/net"
)

var (
	conn net.Conn
)

func main() {
	center.InitConfig()
	connect()

	recv()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		send([]byte(msg))
	}

	if err := scanner.Err(); err !=nil {
		panic(err)
	}
}

func connect() {
	var err error
	addr := center.BaseConfig.Gate.ListenAddr
	conn, err = net.Dial("tcp", addr)
	checkError(err)
	fmt.Printf("connect to %s ok\n", addr)
}

func send(msg []byte) {
	_, err := conn.Write(msg)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func recv() {
	cb := func(reader *bufio.Reader, err error) {
		if err != nil {
			fmt.Println("Recv err", err.Error())
			os.Exit(2)
		}

		packBuf := make([]byte, 512)
		reader.Read(packBuf)
		fmt.Printf("%s say: %s\n", conn.RemoteAddr(), string(packBuf))
	}

	go clawNet.RecvLoop(conn, cb)
}
