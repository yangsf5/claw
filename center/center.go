// Author: sheppard(ysf1026@gmail.com) 2014-03-03

package center

import (
	"fmt"
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

var (
	services map[string]chan<- message
)

func init() {
	services = make(map[string]chan<- message)
}

type cb func(session int, source string, msg []byte)

func Register(name string, service cb) {
	//TODO check repeated name
	channel := make(chan message)

	go func() {
		for {
			select {
			case msg, ok := <-channel:
				if !ok {
					return
				}

				service(msg.session, msg.source, msg.data)
			}
		}
	}()

	services[name] = channel
}

func Cancel(name string) {
	// Some services cannot be cancelled
	needs := []string{"Error"}
	for _, need := range needs {
		if name == need {
			Error("Center", "[Cancel] this service cannot be cancelled, name=" + name)
			return
		}
	}

	serv, ok := services[name]
	if ok {
		delete(services, name)
		close(serv)
	}
}

func Send(source, destination string, session int, msg []byte) {
	channel, ok := services[destination]
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

