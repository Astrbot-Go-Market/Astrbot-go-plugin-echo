// Echo 插件示例（子进程插件运行时）。
//
// 演示钩子写法：命令不在 main 里内联，而是：
//   - handlers.go 定义命令处理函数与 setup()
//   - main.go 只负责 sdk.Serve，并挂上 OnLoad 钩子
//
// 也可以改用另一种写法：直接在 sdk.Plugin{Commands: []sdk.Command{...}}
// 结构体里声明，或用 init() 调用 sdk.RegisterCommand 注册。两者等价。
package main

import (
	sdk "github.com/WaterGodFurina/Astrbot-go-plugin-sdk"
)

func main() {
	sdk.Serve(&sdk.Plugin{
		Name:        "echo",
		Version:     "1.0.0",
		Description: "Echoes your message back",
		Author:      "AstrBot Devs",
		OnLoad:      setup, // 服务启动时先跑 setup()，注册命令
	})
}
