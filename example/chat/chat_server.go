// Author: sheppard(ysf1026@gmail.com) 2014-03-25

package main

import (
	"bufio"
	"net"
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	clawNet "github.com/yangsf5/claw/engine/net"
	"github.com/yangsf5/claw/service"
)

var (
	clients *clawNet.Group
)

type Client struct {
	name string
	conn net.Conn
}

func (c *Client) Send(msg []byte) {
	c.conn.Write(msg)
}

func init() {
	clients = clawNet.NewGroup()
}

func main() {
	service.Register()
	service.GateRegisterConnHandler(connHandle)
	center.Use([]string{"Error", "Gate"})

	glog.Info("Chat start!")

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Chat exit!")
	glog.Flush()
}

func connHandle(conn net.Conn) {
	client := &Client{conn.RemoteAddr().String(), conn}
	clients.AddPeer(client.name, client)

	cb := func(reader *bufio.Reader, err error) {
		if err != nil {
			defer conn.Close()
			clients.DelPeer(client.name)
			glog.Infof("[Event]: %s leave", client.name)
			return
		}

		packBuf := make([]byte, 512)
		reader.Read(packBuf)
		glog.Infof("%s say: %s", client.name, string(packBuf))
		clients.Broadcast(packBuf)
	}

	go clawNet.RecvLoop(conn, cb)
}
