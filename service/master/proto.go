// Author: sheppard(ysf1026@gmail.com) 2014-03-14

package master

const (
	LOGIN = iota
)

type Login struct {
	Name string
}

type Command struct {
	Name string
}

var (
	packets map[uint16]func()interface{}
)

func init() {
	packets = make(map[uint16]func()interface{})

	packets[LOGIN] = func() interface{} {
		return &Login{}
	}
}
