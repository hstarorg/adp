package god

import (
	"fmt"
	"net/http"
)

// God 定义God类型
type God struct {
	debug       bool   `json: "debug" xml:"debug"`
	name        string `json: "name" xml:"name"`
	middlewares []HandlerFunc
}

// Middleware 定义中间件类型
type Middleware interface{}

// HandlerFunc 定义处理函数
type HandlerFunc func(*Context)

// New 创建God实例
func New() *God {
	god := new(God)
	god.middlewares = make([]HandlerFunc, 0)
	return god
}

// Server 创建一个WebServer
func (god *God) Server(addr string) *http.Server {
	server := &http.Server{Addr: addr}
	return server
}

// Listen 启动Web Server
func (god *God) Listen(addr string) {
	god.run(god.Server(addr))
}

func parseRequest(req *http.Request) *Request {
	return new(Request)
}

func (god *God) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	ctx := new(Context)
	ctx.handlers = god.middlewares
	ctx.req = req
	ctx.res = res
	ctx.request = parseRequest(req)
	ctx.Next()
	// res.Write([]byte("xxx"))
}

// Use 方法是为了加载中间件
func (god *God) Use(middleware HandlerFunc) *God {
	god.middlewares = append(god.middlewares, middleware)
	return god
}

func (god *God) run(server *http.Server) {
	server.Handler = god
	fmt.Printf("Listen %s", server.Addr)
	server.ListenAndServe()
}
