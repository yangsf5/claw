package main

import (
	"bufio"
	"net"
	"time"

	"github.com/golang/glog"
	"github.com/yangsf5/claw/center"
	clawNet "github.com/yangsf5/claw/engine/net"
	"github.com/yangsf5/claw/service"
	"github.com/yangsf5/claw/service/gate"
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
	gate.RegisterReader(regReader)
	center.Use([]string{"Error", "Gate"})

	glog.Info("Chat start!")

	for {
		time.Sleep(100 * time.Millisecond)
	}

	glog.Info("Chat exit!")
	glog.Flush()
}

func regReader(session int, reader *bufio.Reader, err error) {
	if err != nil {
		glog.Infof("[Event]: %d leave", session)
		return
	}

	packBuf := make([]byte, 512)
	n, _ := reader.Read(packBuf)
	msg := packBuf[:n]
	glog.Infof("%d say: %s", session, string(msg))
	center.Send("main", "Gate", 0, center.MsgTypeText, msg)
}
