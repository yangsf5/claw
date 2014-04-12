// Author: sheppard(ysf1026@gmail.com) 2014-04-12

package net

import (
	"sync"
)

type Group2 struct {
	peers map[int] Peer
	mutex sync.RWMutex

	broadcast chan []byte
}

func NewGroup2() *Group2 {
	g := &Group2{}
	g.peers = make(map[int] Peer)
	g.broadcast = make(chan []byte)

	go g.tick()
	return g
}

func (g *Group2) AddPeer(peerId int, peer Peer) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.peers[peerId]; ok {
		return false
	}
	g.peers[peerId] = peer
	return true
}

func (g *Group2) GetPeer(peerId int) Peer {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	peer, ok := g.peers[peerId]
	if ok {
		return peer
	}
	return nil
}

func (g *Group2) DelPeer(peerId int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	delete(g.peers, peerId)
}

type WalkFunc2 func(peerId int, peer Peer)

func (g *Group2) Walk(walkFn WalkFunc2) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.peers {
		walkFn(k, v)
	}
}

type CondFunc2 func(peerId int, peer Peer) bool

func (g *Group2) Find(condFn CondFunc2) (peerId int, peer Peer) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.peers {
		if condFn(k, v) {
			return k, v
		}
	}

	return -1, nil
}

func (g *Group2) Broadcast(msg []byte) {
	g.broadcast <- msg
}

func (g *Group2) Close() {
	close(g.broadcast)
}

func (g *Group2) tick() {
	for {
		select {
		case msg, ok := <-g.broadcast:
			if !ok {
				return
			}

			g.Walk(func(peerId int, peer Peer) {
				peer.Send(msg)
			})
		}
	}
}

