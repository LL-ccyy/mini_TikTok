package main

import (
	"Minimalist_TikTok/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	//go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
