package adpagent

import (
	"god"
)

func main(){
	app := god.New()
	app.Listen(":7777")
}
