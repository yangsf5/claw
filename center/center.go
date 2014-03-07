// Author: sheppard(ysf1026@gmail.com) 2014-03-03

package center

import (
	"fmt"
	"os"
)

const (
	MsgTypeText = iota
	MsgTypeResponse
	MsgTypeMulticast
	MsgTypeClient
	MsgTypeSystem
	MsgTypeHarbor
)

type message struct {
	source string
	session int
	data []byte
}

type ClawCallback func(session int, source string, msg []byte)

var (
	services map[string]ClawCallback
	channels map[string]chan<- message
)

func init() {
	services = make(map[string]ClawCallback)
	channels = make(map[string]chan<- message)
}

func Register(name string, service ClawCallback) {
	_, ok := services[name]
	if ok {
		fmt.Println("[Center.Register] service name is existed, name=" + name)
		os.Exit(2)
	}
	services[name] = service
}

func Use(names []string) {
	for _, name := range names {
		serv, ok := services[name]
		if ok {
			channel := make(chan message)

			go func() {
				for {
					select {
					case msg, ok := <-channel:
						if !ok {
							return
						}

						serv(msg.session, msg.source, msg.data)
					}
				}
			}()

			channels[name] = channel
		} else {
			fmt.Println("[Center.Use] service is not found, name=" + name)
			os.Exit(2)
		}
	}

	check()
}

func check() {
	needs := []string{"Error"}
	for _, need := range needs {
		_, ok := channels[need]
		if !ok {
			fmt.Println("[Center.Check] lack a service, name=" + need)
			os.Exit(2)
		}
	}
}

func Send(source, destination string, session int, msg []byte) {
	channel, ok := channels[destination]
	if !ok {
		Error("Center", fmt.Sprintf("[Send] destination is not found, source=%s destination=%s", source, destination))
		return
	}
	channel <- message{
		source: source,
		session: session,
		data: msg,
	}
}

func Error(source, msg string) {
	Send(source, "Error", 0, []byte(msg))
}

