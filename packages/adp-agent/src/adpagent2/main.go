package main

import (
	"adpagent2/router"

	"gopkg.in/baa.v1"
)

func main() {
	app := baa.New()
	router.AttachToApp(app)
	app.Run(":7777")
}
