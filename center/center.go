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
	dataType int
	data interface{}
}

type Service interface {
	ClawCallback(session int, source string, msgType int, msg interface{})

	//RelyServices() []string

	ClawStart()
}

var (
	services map[string]Service
	channels map[string]chan<- message
)

func init() {
	services = make(map[string]Service)
	channels = make(map[string]chan<- message)
}

func Register(name string, service Service) {
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
			serv.ClawStart()

			channel := make(chan message)

			go func() {
				for {
					select {
					case msg, ok := <-channel:
						if !ok {
							return
						}

						serv.ClawCallback(msg.session, msg.source, msg.dataType, msg.data)
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

func Send(source, destination string, session int, msgType int, msg interface{}) {
	channel, ok := channels[destination]
	if !ok {
		Error("Center", fmt.Sprintf("[Send] destination is not found, source=%s destination=%s", source, destination))
		return
	}
	channel <- message{
		source: source,
		session: session,
		dataType: msgType,
		data: msg,
	}
}

func Error(source, msg string) {
	Send(source, "Error", 0, MsgTypeText, msg)
}

