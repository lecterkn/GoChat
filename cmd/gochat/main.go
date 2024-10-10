package main

import (
	"fmt"
	"lecter/goserver/internal/app/gochat/common"
	"lecter/goserver/internal/app/gochat/db"
	"lecter/goserver/internal/app/gochat/router"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数を読み込む
	if err := loadEnv(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// configs.jsonから読み込んだConfig
	appConfig := common.LoadConfig()
	if appConfig != nil {
		common.ApplicationConfig = *appConfig
	}

	// DB接続
	err := db.Connect()
	if err != nil {
		fmt.Println("DB Connection ERROR")
		return
	}
	defer db.Close()

	server := gin.Default()
	router.Routing(server)

	// サーバー起動
	err = server.Run(":" + strconv.Itoa(common.ApplicationConfig.Port))
	if err != nil {
		fmt.Print(err.Error())
	}
}

func loadEnv() error {
	err := godotenv.Load() // .envファイルの読み込み
	if err != nil {
		return err
	}
	return nil
}
