package bizs

import (
	"gopkg.in/baa.v1"
)

// agentBiz 业务类
type agentBiz struct{}

// AgentBiz 暴露业务类实例
var AgentBiz = agentBiz{}

// ExecuteBuild 执行构建
func (a agentBiz) ExecuteBuild(c *baa.Context) {

}

// ProcessPing 处理ping请求，校验存活
func (a agentBiz) ProcessPing(c *baa.Context) {
	c.Text(200, []byte("PONG"))
}
