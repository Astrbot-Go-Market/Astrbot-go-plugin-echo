// Example AstrBot Go plugin: "echo" (gRPC subprocess or Native in-process).
//
// This module is independent of the AstrBot host. The same source supports both
// runtimes; the runtime is chosen per plugin on the host side:
//
//   - gRPC (default): the host builds an executable and launches it as a child
//     process, talking over gRPC (go-plugin). `func main` runs `sdk.Serve`.
//   - Native: the host builds an in-process shared library (.so / .dll) and
//     injects a generated entry that hands the package-level `plugin` variable
//     to the SDK Native runtime. `func main` is not executed for Native builds.
//
// The `plugin` variable is therefore package-level (not inlined in main), which
// is the only source-shape requirement shared by both runtimes. This plugin
// registers an "echo" command that echoes the user's message back.
package main

import (
	"strings"

	sdk "github.com/WaterGodFurina/Astrbot-go-plugin-sdk/v2"
)

// plugin 是插件定义。提升为包级变量后，同一份源码既可用于 gRPC 运行方式
// （main 内 sdk.Serve(plugin)），也可用于 Native 运行方式（构建期注入的
// native_entry.go 直接引用本变量并交给 SDK Native 入口）。
var plugin = &sdk.Plugin{
	Name:        "echo",
	Version:     "2.0.0",
	Description: "Echoes your message back",
	Author:      "AstrBot Devs",
	ConfigSchema: map[string]any{
		"description": "Echo 插件配置",
		"type":        "object",
		"items": map[string]any{
			"prefix": map[string]any{
				"description": "回复前缀",
				"type":        "string",
				"default":     "[Echo] ",
			},
			"upper": map[string]any{
				"description": "转大写",
				"type":        "bool",
				"default":     false,
			},
		},
	},
	Commands: []sdk.Command{
		{
			Name:        "echo",
			Aliases:     []string{"repeat"},
			Description: "Echoes your message",
			Usage:       "echo <text>",
			Permission:  "everyone",
			Handler: func(e *sdk.Event, args []string) (string, error) {
				if len(args) == 0 {
					return "Usage: echo <text>", nil
				}
				return strings.Join(args, " "), nil
			},
		},
	},
	Hooks: []sdk.Hook{
		{
			Name:  "echo_on_start",
			Event: "startup",
			Handler: func(e *sdk.Event) error {
				return nil
			},
		},
	},
}

func main() {
	sdk.Serve(plugin)
}
