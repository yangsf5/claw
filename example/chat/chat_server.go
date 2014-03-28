// Author: sheppard(ysf1026@gmail.com) 2014-03-25

package main

import (
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
	conn net.Conn
}

func (c *Client) Send(msg []byte) {
	c.conn.Write(msg)
}

func init() {
	clients = clawNet.NewGroup
}

func main() {
	service.Register()
	service.GateRegisterConnHandler(connHandle)
	center.Use([]string{"Error", "Master", "Harbor", "Gate"})

	glog.Info("Chat start!")

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Chat exit!")
	glog.Flush()
}

func connHandle(conn net.Conn) {
	defer conn.Close()
	clients.AddPeer(conn.RemoteAddr(), client)

	cb := func(reader *bufio.Reader, err error) {
		if err != nil {
			clients.DelPeer(conn.RemoteAddr())
			glog.Infof("Event: %s leave", conn.RemoteAddr())
			return
		}

		n := reader.Buffered()
		packBuf := make([]byte, n)
		reader.Read(packBuf)
		glog.Infof("%s say: %s", conn.RemoteAddr(), string(packBuf))
	}

	clawNet.RecvLoop(conn, cb)
}
