package main

import (
	"runtime"
	"rx-mp/internal/app"
)

// check64bRuntime 检查运行环境是否处于 64 位下
func check64bRuntime() bool {
	return runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64"
}

func main() {
	is64bRuntime := check64bRuntime()
	if !is64bRuntime {
		panic("请在64位 操作系统架构下运行本程序")
	}

	app.Run()
}
