package main

import (
	_ "demo/internal/pkg/config"

	"demo/internal/pkg/db"
	_ "demo/internal/pkg/db"
	"fmt"
)

import "demo/internal/app"

func main() {
	app.Run()

	name := db.RdRedis.Get(db.RdRedisContext, "name")
	fmt.Printf("name: %s\n", name)
}
