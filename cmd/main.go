package main

import (
	"runtime"
	"rx-mp/internal/app"
)

func main() {
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		panic("请在64位 操作系统架构下运行本程序")
	}

	app.Run()
}
