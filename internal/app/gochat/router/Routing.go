package router

import (
	"lecter/goserver/internal/app/gochat/controller"
	"lecter/goserver/internal/app/gochat/service/authorization"

	_ "lecter/goserver/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routing(r *gin.Engine) {
	// cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// JWT認証のグループを設定
	r.POST("/api/v1/login", authorization.Login)
	userApi := r.Group("/api/v1/users")
	userApi.Use(authorization.JwtAuthorization)

	// Version
	vc := controller.VersionController{}
	r.GET("/api/v1/version", vc.Index)

	// User
	uc := controller.UserController{}
	r.POST("api/v1/register", uc.Create)
	userApi.GET("/", uc.Select)
	userApi.PATCH("/:userId", uc.Update)

	// User Profile
	upc := controller.UserProfileController{}
	userApi.GET("/:userId/profiles", upc.Select)
	userApi.PUT("/:userId/profiles", upc.Update)

	cc := controller.ChannelController{}
	userApi.GET("/channels", cc.Index)
	userApi.POST("/channels", cc.Create)
	userApi.GET("/channels/:channelId", cc.Select)
	userApi.PATCH("/channels/:channelId", cc.Update)
	userApi.DELETE("/channels/:channelId", cc.Delete)

	mc := controller.MessageController{}
	userApi.GET("/channels/:channelId/messages", mc.Index)
	userApi.POST("/channels/:channelId/messages", mc.Create)
	userApi.PATCH("/channels/:channelId/messages/:messageId", mc.Update)
	userApi.DELETE("/channels/:channelId/messages/:messageId", mc.Delete)

	clc := controller.ChannelLanguageController{}
	userApi.GET("/channels/:channelId/languages", clc.Index)
	userApi.PUT("/channels/:channelId/languages", clc.Save)
}
