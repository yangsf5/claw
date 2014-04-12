// Author: sheppard(ysf1026@gmail.com) 2014-04-11

package gate

import (
	"bufio"
	"net"

	"github.com/golang/glog"
	clawNet "github.com/yangsf5/claw/engine/net"
)

type regReaderFunc func(session int, reader *bufio.Reader, err error)

var (
	clients *clawNet.Group2
	sessionIdGenerator int

	regReader regReaderFunc
)

type Client struct {
	session int
	conn net.Conn
}

func (c *Client) Send(msg []byte) {
	c.conn.Write(msg)
}

func init() {
	clients = clawNet.NewGroup2()
}

func RegisterReader(reader regReaderFunc) {
	regReader = reader
}

func ConnHandle(conn net.Conn) {
	sessionIdGenerator++
	client := &Client{ sessionIdGenerator, conn}
	clients.AddPeer(client.session, client)

	cb := func(reader *bufio.Reader, err error) {
		if err != nil {
			defer conn.Close()
			clients.DelPeer(client.session)
			glog.Infof("Gate lose peer, session=%d", client.session)
		}

		if regReader != nil {
			regReader(client.session, reader, err)
		}
	}

	go clawNet.RecvLoop(conn, cb)
}

func Broadcast(msg []byte) {
	clients.Broadcast(msg)
}

func SendSingle(session int, msg []byte) {
	if client := clients.GetPeer(session); client != nil {
		client.Send(msg)
	}
}

