package main

import (
	"runtime"
	"rx-mp/internal/app"
)

// check64Runtime 检查运行环境是否处于 64 位下
func check64Runtime() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}

func main() {
	is64Runtime := check64Runtime()
	if !is64Runtime {
		panic("请在64位 操作系统架构下运行本程序")
	}

	app.Run()
}
