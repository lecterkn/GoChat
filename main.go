package main

import (
	"fmt"
	"lecter/hello/config"
	"lecter/hello/db"
	"lecter/hello/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// configs.jsonから読み込んだConfig
	appConfig := config.LoadConfig()
	if (appConfig != nil) {
		config.ApplicationConfig = *appConfig
	}

	// DB接続
	err := db.Connect()
	if (err != nil) {
		fmt.Println("DB Connection ERROR")
		return
	}
	defer db.DB().Close()

	server := gin.Default()
	router.Routing(server)

	// サーバー起動
	server.Run(":" + strconv.Itoa(config.ApplicationConfig.Port))
}