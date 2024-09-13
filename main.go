package main

import (
	"lecter/hello/config"
	"lecter/hello/db"
	"lecter/hello/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// configs.jsonから読み込んだConfig
	appConfig := *config.LoadConfig()

	// DB接続
	db.Connect()

	server := gin.Default()
	router.Routing(server)

	// サーバー起動
	server.Run(":" + strconv.Itoa(appConfig.Port))
}