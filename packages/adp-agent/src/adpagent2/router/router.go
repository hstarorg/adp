package router

import (
	"adpagent2/bizs"

	"gopkg.in/baa.v1"
)

// AttachToApp 挂载Router
func AttachToApp(b *baa.Baa) {
	b.Get("/", bizs.AgentBiz.ExecuteBuild)
	b.Get("/ping", bizs.AgentBiz.ProcessPing)
}
