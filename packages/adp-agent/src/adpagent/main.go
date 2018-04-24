package main

import (
	"god"
)

func XXX(ctx *god.Context){

}

func main(){
	app := god.New()
	// 定义Router
	router:= new(god.Router)
	// 加载路由
	app.Use(router.Routes())
	// 启动程序
	app.Listen("127.0.0.1:7777")
}
