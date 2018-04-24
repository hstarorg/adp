package god 

import (
	"fmt"
	"net/http"
)

type God struct {
	debug           bool
	name            string
	middleware      []HandlerFunc
}

// 定义中间件
type Middleware interface{}

// 定义处理函数
type HandlerFunc func(*Context)

func New() *God {
	god := new(God)
	god.middleware = make([]HandlerFunc, 0)
	return god
}

func (god *God) Server(addr string) *http.Server {
	server := &http.Server{Addr: addr}
	return server
}

func (god *God) Listen(addr string) {
	god.run(god.Server(addr))
}

func (b *God) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("xxx"))
}

func (god *God) run(server *http.Server, files ...string) {
	server.Handler = god
	if len(files) == 0 {
		println("Listen %s", server.Addr)
	}else {
		panic("invalid TLS configuration")
	}
}
