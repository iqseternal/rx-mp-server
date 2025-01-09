package main

import (
	"demo/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

type config struct {
	PORT string
	HOST string
	URL  string
}

var (
	Config = config{
		PORT: "8080",
		HOST: "localhost",
		URL:  "localhost:8080",
	}
)

func main() {
	engine := gin.Default()

	router.InitRouter(engine)

	err := engine.Run(Config.URL)

	if err != nil {
		fmt.Println("启动失败")
	}
}
