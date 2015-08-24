package net

import (
	"sync"
)

type Peer interface {
	Send([]byte)
}

type Group struct {
	peers map[string] Peer
	mutex sync.RWMutex

	broadcast chan []byte
}

func NewGroup() *Group {
	g := &Group{}
	g.peers = make(map[string] Peer)
	g.broadcast = make(chan []byte)

	go g.tick()
	return g
}

func (g *Group) AddPeer(peerId string, peer Peer) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.peers[peerId]; ok {
		return false
	}
	g.peers[peerId] = peer
	return true
}

func (g *Group) GetPeer(peerId string) Peer {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	peer, ok := g.peers[peerId]
	if ok {
		return peer
	}
	return nil
}

func (g *Group) DelPeer(peerId string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	delete(g.peers, peerId)
}

type WalkFunc func(peerId string, peer Peer)

func (g *Group) Walk(walkFn WalkFunc) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.peers {
		walkFn(k, v)
	}
}

type CondFunc func(peerId string, peer Peer) bool

func (g *Group) Find(condFn CondFunc) (peerId string, peer Peer) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.peers {
		if condFn(k, v) {
			return k, v
		}
	}

	return "", nil
}

func (g *Group) Broadcast(msg []byte) {
	g.broadcast <- msg
}

func (g *Group) Close() {
	close(g.broadcast)
}

func (g *Group) tick() {
	for {
		select {
		case msg, ok := <-g.broadcast:
			if !ok {
				return
			}

			g.Walk(func(peerId string, peer Peer) {
				peer.Send(msg)
			})
		}
	}
}

