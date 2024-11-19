package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "lecter/goserver/docs"
	"lecter/goserver/internal/app/gochat"
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
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// JWT認証のグループを設定
	controllerSets := gochat.InitializeControllerSet()
	jwtAuthorizationService := gochat.InitializeJwtAuthorizationService()
	r.POST("/api/v1/login", jwtAuthorizationService.Login)
	userApi := r.Group("/api/v1/users")
	userApi.Use(jwtAuthorizationService.JwtAuthorization)

	// Version
	vc := controllerSets.VersionController
	r.GET("/api/v1/version", vc.Index)

	// User
	uc := controllerSets.UserController
	r.POST("api/v1/register", uc.Create)
	userApi.GET("/", uc.Select)
	userApi.PATCH("/:userId", uc.Update)

	// User Profile
	upc := controllerSets.UserProfileController
	userApi.GET("/:userId/profiles", upc.Select)
	userApi.PUT("/:userId/profiles", upc.Update)

	// Channel
	cc := controllerSets.ChannelController
	userApi.GET("/channels", cc.Index)
	userApi.POST("/channels", cc.Create)
	userApi.GET("/channels/:channelId", cc.Select)
	userApi.PATCH("/channels/:channelId", cc.Update)
	userApi.DELETE("/channels/:channelId", cc.Delete)

	// Message
	mc := controllerSets.MessageController
	userApi.GET("/channels/:channelId/messages", mc.Index)
	userApi.POST("/channels/:channelId/messages", mc.Create)
	userApi.PATCH("/channels/:channelId/messages/:messageId", mc.Update)
	userApi.DELETE("/channels/:channelId/messages/:messageId", mc.Delete)

	// Channel Language
	clc := controllerSets.ChannelLanguageController
	userApi.GET("/channels/:channelId/languages", clc.Index)
	userApi.PUT("/channels/:channelId/languages", clc.Save)
}
