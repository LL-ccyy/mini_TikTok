package main

import (
	"Minimalist_TikTok/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.Init()
	//go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	http.ListenAndServe(":39090", nil)
}
