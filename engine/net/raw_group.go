// Author: sheppard(ysf1026@gmail.com) 2014-04-12

package net

import (
	"sync"
)

type RawGroup struct {
	peers map[int] Peer
	mutex sync.RWMutex

	broadcast chan []byte
}

func NewRawGroup() *RawGroup {
	g := &RawGroup{}
	g.peers = make(map[int] Peer)
	g.broadcast = make(chan []byte)

	go g.tick()
	return g
}

func (g *RawGroup) AddPeer(peerId int, peer Peer) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.peers[peerId]; ok {
		return false
	}
	g.peers[peerId] = peer
	return true
}

func (g *RawGroup) GetPeer(peerId int) Peer {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	peer, ok := g.peers[peerId]
	if ok {
		return peer
	}
	return nil
}

func (g *RawGroup) DelPeer(peerId int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	delete(g.peers, peerId)
}

type RawWalkFunc func(peerId int, peer Peer)

func (g *RawGroup) Walk(walkFn RawWalkFunc) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.peers {
		walkFn(k, v)
	}
}

type RawCondFunc func(peerId int, peer Peer) bool

func (g *RawGroup) Find(condFn RawCondFunc) (peerId int, peer Peer) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.peers {
		if condFn(k, v) {
			return k, v
		}
	}

	return -1, nil
}

func (g *RawGroup) Broadcast(msg []byte) {
	g.broadcast <- msg
}

func (g *RawGroup) Close() {
	close(g.broadcast)
}

func (g *RawGroup) tick() {
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

