// Author: sheppard(ysf1026@gmail.com) 2014-03-17

package center

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
)

var (
	configFile = flag.String("config", "./config.xml", "config file name")
	IsMaster = flag.Bool("isMaster", false, "this server is master?")
)

func init() {
	flag.Parse()
}

func GetConfig(v interface{}) {
	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(content, v)
	if err != nil {
		panic(err)
	}
}
