# adp-agent

利用 `Go` 来实现代理程序，用于执行 center 中设定的控制命令。

# 操作手册

1、在 `golang` (官方网站)[https://golang.org/]下载 go 的安装包，并进行安装

2、使用 `glide` 作为包管理工具，在 [glide 下载地址](https://github.com/Masterminds/glide/releases) 下载好二进制包，解压后放到环境变量 `PATH` 能够指向的一个目录中（为了能够全局使用 `glide` 命令），通过新开一个控制台，输入 `glide` 进行校验，如果能正确输出 `glide` 帮助说明即是安装成功。

**注意：glide如果是在Windows10下请使用 v0.12.3，如果使用最新版本，很大几率不能正常使用**

3、代码编辑器采用 `VS Code`，安装 `Go` 插件。

4、在 `VS Code 中的 Workspace Settings` 设置如下参数

```json
{
  "go.toolsGopath": "C:\\Go\\tools",
  "go.gopath": "E:\\MyProjects\\GitProjects\\adp\\packages\\adp-agent"
}
```

其中 `go.toolsGopath` 用来安装 GO 开发的一系列工具（VSC 会自动提示安装）；`go.gopath` 一般设置为当前项目的根目录。

5、在环境变量中设置 `GOPATH`，可以设置多个值，如：`GOPATH=xxx;E:\MyCodeLibrary\adp\packages\adp-agent`
