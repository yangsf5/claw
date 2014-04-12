// Author: sheppard(ysf1026@gmail.com) 2014-04-11

package gate

import (
	"bufio"
	"net"

	"github.com/golang/glog"
	clawNet "github.com/yangsf5/claw/engine/net"
)

var (
	clients *clawNet.Group2
	sessionIdGenerator int
)

type Client struct {
	id int
	conn net.Conn
}

func (c *Client) Send(msg []byte) {
	c.conn.Write(msg)
}

func init() {
	clients = clawNet.NewGroup2()
}

func ConnHandle(conn net.Conn) {
	sessionIdGenerator++
	client := &Client{ sessionIdGenerator, conn}
	clients.AddPeer(client.id, client)

	cb := func(reader *bufio.Reader, err error) {
		if err != nil {
			defer conn.Close()
			clients.DelPeer(client.id)
			glog.Infof("[Event]: %d leave", client.id)
			return
		}

		packBuf := make([]byte, 512)
		reader.Read(packBuf)
		glog.Infof("%d say: %s", client.id, string(packBuf))
		clients.Broadcast(packBuf)
	}

	go clawNet.RecvLoop(conn, cb)
}

func Broadcast(msg []byte) {
	clients.Broadcast(msg)
}
