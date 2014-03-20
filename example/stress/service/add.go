// Author: sheppard(ysf1026@gmail.com) 2014-03-12

package service

import (
	"fmt"
	"net/rpc"
	"github.com/yangsf5/claw/example/stress/proto"
)

type Add struct {
}

func (s *Add) ClawCallback(session int, source string, msgType int, msg interface{}) {
}

func (s *Add) ClawStart() {
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

