package main

import (
	"Minimalist_TikTok/config"
	//"Minimalist_TikTok/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	//go service.Manager.Start()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
