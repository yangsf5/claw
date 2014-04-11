// Author: sheppard(ysf1026@gmail.com) 2014-04-11

package gate

import (
	"bufio"
	"net"

	"github.com/golang/glog"
	clawNet "github.com/yangsf5/claw/engine/net"
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

func ConnHandle(conn net.Conn) {
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

func Broadcast(msg []byte) {
	clients.Broadcast(msg)
}
