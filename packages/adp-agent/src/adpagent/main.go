package main

import (
	"god"
)

// XXX 没什么卵用
func XXX(ctx *god.Context) {
	ctx.Next()
}

func main() {
	app := god.New()
	// 定义Router
	router := new(god.Router)
	app.Use(XXX)
	// 加载路由
	app.Use(router.Routes())
	// 启动程序
	app.Listen(":7777")
}
