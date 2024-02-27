package main

import (
	"go-web-ssh/controller"
	"go-web-ssh/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	server.Static("/static", "./client")
	server.GET("/ws", controller.WSHandler())
	server.GET("/", gin.WrapH(http.FileServer(http.Dir("./client"))))
	server.Run("0.0.0.0:3000")
}
