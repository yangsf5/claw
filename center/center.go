// Author: sheppard(ysf1026@gmail.com) 2014-03-03

package center

type cb func(session int, source string, msg []byte)

var (
	services map[string]cb
)

func init() {
	services = make(map[string]cb)
}

func Register(name string, service cb) {
	//TODO check repeated name
	services[name] = service
}

func Send(source, destination string, session int, msg []byte) {
	services[destination](session, source, msg)
}
