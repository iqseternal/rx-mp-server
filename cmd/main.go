package main

import (
	_ "rx-mp/config"

	"fmt"
	"rx-mp/internal/app"
	"rx-mp/internal/pkg/storage"
)

func main() {
	app.Run()

	name := storage.RdRedis.Get(storage.RdRedisContext, "name")
	fmt.Printf("name: %s\n", name)
}
