// Author: sheppard(ysf1026@gmail.com) 2014-03-12

package service

import (
	"fmt"
	"net/rpc"
	"github.com/yangsf5/claw/example/stress/proto"
)

func init() {
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
	fmt.Println("Res", res)
}

func addCallback(session int, source string, msg []byte) {
}
