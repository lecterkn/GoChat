package main

import (
	"lecter/hello/router"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	router.Routing(server)
	server.Run(":8080")
}