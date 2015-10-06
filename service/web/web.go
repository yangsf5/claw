package web

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"runtime/debug"

	"github.com/golang/glog"
	"golang.org/x/net/websocket"

	"github.com/yangsf5/claw/center"
)

const (
	TEMPLATE_DIR = "./view"
	TEMPLATE_COMMON_DIR = "./view/common"
)

var (
	templates = make(map[string] *template.Template)
	mux *http.ServeMux
)

func Start() {
	if mux == nil {
		panic("Claw.Web please register http handler.")
	}

	cacheTemplate()

	type HttpConfig struct {
		ListenAddr string `xml:"listenAddr,attr"`
	}
	type ConfigPack struct {
		XMLName xml.Name `xml:"clawconfig"`
		Http HttpConfig `xml:"http"`
	}
	var cfg ConfigPack
	center.GetConfig(&cfg)

	err := http.ListenAndServe(cfg.Http.ListenAddr, mux)
	checkError(err)
}

func RegisterHttpHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if mux == nil {
		mux = http.NewServeMux()
	}
	mux.HandleFunc(pattern, safeHandler(handler))
}

func RegisterWebSocketHandler(pattern string, handler func(*websocket.Conn)) {
	if mux == nil {
		mux = http.NewServeMux()
	}
	mux.Handle(pattern, websocket.Handler(handler))
}

func StaticDirHandler(prefix string, staticDir string) {
	if mux == nil {
		mux = http.NewServeMux()
	}
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		http.ServeFile(w, r, file)
	})
}

func cacheTemplate() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	checkError(err)

	header := TEMPLATE_COMMON_DIR + "/header.html"
	footer := TEMPLATE_COMMON_DIR + "/footer.html"
	glog.Infof("Common template [%s %s]", header, footer)

	var tplName, tplPath string
	for _, fileInfo := range fileInfoArr {
		tplName = fileInfo.Name()
		if ext := path.Ext(tplName); ext != ".html" {
			continue
		}
		tplPath = TEMPLATE_DIR + "/" + tplName
		glog.Infof("Loading template %s", tplPath)
		t := template.Must(template.ParseFiles(tplPath, header, footer))
		templates[tplPath] = t
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func RenderHtml(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = TEMPLATE_DIR + "/" + tmpl
	tpl, ok := templates[tmpl];
	if !ok {
		glog.Errorf("Render html, but template is nil, name=%s", tmpl)
		return
	}

	err := tpl.Execute(w, data)
	checkError(err)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				glog.Errorf("WARN: panic in %v - %v", fn, err)
				glog.Errorf(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
