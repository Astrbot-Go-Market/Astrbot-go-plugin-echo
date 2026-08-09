package main

import (
	"strings"

	sdk "github.com/WaterGodFurina/Astrbot-go-plugin-sdk"
)

// echoHandler 把用户消息原样返回。
func echoHandler(e *sdk.Event, args []string) (string, error) {
	if len(args) == 0 {
		return "Usage: echo <text>", nil
	}
	return strings.Join(args, " "), nil
}

// setup 由 main 的 OnLoad 钩子调用，用于注册命令（复杂插件可在这里
// 读取配置文件后动态注册，也可拆多个函数/文件）。
func setup() error {
	sdk.RegisterCommand(sdk.Command{
		Name:        "echo",
		Aliases:     []string{"repeat"},
		Description: "Echoes your message",
		Usage:       "echo <text>",
		Permission:  "everyone",
		Handler:     echoHandler,
	})
	return nil
}
