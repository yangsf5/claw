package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
	"github.com/yangsf5/claw/example/stress/proto"
)

type Math int

func (m *Math) Add(req *proto.ReqAdd, res *int) error {
	fmt.Printf("Server accept Math.Add, req=%v\n", req)
	*res = req.A + req.B
	return nil
}

func main() {
	fmt.Println("Stress mock server start!")
	math := new(Math)
	rpc.Register(math)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	go http.Serve(l, nil)

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
