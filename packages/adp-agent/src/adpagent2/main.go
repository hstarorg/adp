package main

import (
	"gopkg.in/baa.v1"
)

func main() {
	app := baa.New()
	Router(app)
	app.Run(":7777")
}
