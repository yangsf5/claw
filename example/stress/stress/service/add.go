// Author: sheppard(ysf1026@gmail.com) 2014-03-12

package service

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/yangsf5/claw/center"
	"github.com/yangsf5/claw/example/stress/proto"
)

type Add struct {
	testing bool
}

func (s *Add) ClawCallback(session int, source string, msgType int, msg interface{}) {
	fmt.Printf("Service.Add recv msgType=%v msg=%v\n", msgType, string(msg.([]byte)))
	switch msgType {
	case center.MsgTypeHarbor:
		switch string(msg.([]byte)) {
		case "start":
			s.start()
		case "stop":
			s.stop()
		}
	}
}

func (s *Add) ClawStart() {
	go s.stressTest()
}

func (s *Add) start() {
	s.testing = true
}

func (s *Add) stop() {
	s.testing = false
}

func (s *Add) stressTest() {
	for {
		time.Sleep(1 * time.Second)

		if !s.testing {
			continue
		}

		client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
		if err != nil {
			panic(err)
		}

		req := &proto.ReqAdd{27, 26}
		var res int
		err = client.Call("Math.Add", req, &res)
		if err != nil {
			panic(err)
		}
		fmt.Println("Service.Add Res", res)
	}
}
