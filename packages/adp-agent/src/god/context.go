package god

import (
	"net/http"
)

type Context struct {
	req          *http.Request
	res          http.ResponseWriter
	handlers     []HandlerFunc
	request      *Request
	response     *Response
	app          *God
	state        map[string]interface{}
	currentIndex int
}

func (ctx *Context) Next() {
	if ctx.currentIndex >= len(ctx.handlers) {
		return
	}
	i := ctx.currentIndex
	ctx.currentIndex++
	if ctx.handlers[i] != nil {
		ctx.handlers[i](ctx)
	} else {
		ctx.Next()
	}
}
