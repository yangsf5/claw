// Author: sheppard(ysf1026@gmail.com) 2014-01-19

package net

import (
	"sync"
)

type User interface {
	Send(string)
}

type Group struct {
	users map[string] User
	mutex sync.RWMutex

	broadcast chan string

	//TODO check
	cmd chan string
}

func NewGroup() *Group {
	g := &Group{}
	g.users = make(map[string] User)
	g.broadcast = make(chan string)
	g.cmd = make(chan string)

	go g.tick()
	return g
}

func (g *Group) AddUser(uid string, u User) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.users[uid]; ok {
		return false
	}
	g.users[uid] = u
	return true
}

func (g *Group) GetUser(uid string) User {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	user, ok := g.users[uid]
	if ok {
		return user
	}
	return nil
}

func (g *Group) DelUser(uid string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	delete(g.users, uid)
}

type WalkFunc func(uid string, u User)

func (g *Group) Walk(walkFn WalkFunc) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.users {
		walkFn(k, v)
	}
}

type CondFunc func(uid string, u User) bool

func (g *Group) Find(condFn CondFunc) (uid string, u User) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	for k, v := range g.users {
		if condFn(k, v) {
			return k, v
		}
	}

	return "", nil
}

func (g *Group) Broadcast(msg string) {
	g.broadcast <- msg
}

func (g *Group) SendCommand(cmd string) {
	g.cmd <- cmd
}

func (g *Group) Close() {
	close(g.broadcast)
	close(g.cmd)
}

func (g *Group) tick() {
	for {
		select {
		case msg, ok := <-g.broadcast:
			if !ok {
				return
			}

			g.Walk(func(uid string, u User) {
				u.Send(msg)
			})

		case _, ok := <-g.cmd:
			if !ok {
				return
			}
		}
	}
}

