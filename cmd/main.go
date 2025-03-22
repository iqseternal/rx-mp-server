package main

import (
	"runtime"
	_ "rx-mp/config"

	"fmt"
	"rx-mp/internal/app"
	"rx-mp/internal/pkg/storage"
)

func main() {
	if runtime.GOARCH != "amd64" && runtime.GOARCH != "arm64" {
		panic("unsupported 32-bit architecture")
	}

	app.Run()

	name := storage.RdRedis.Get(storage.RdRedisContext, "name")
	fmt.Printf("name: %s\n", name)
}
