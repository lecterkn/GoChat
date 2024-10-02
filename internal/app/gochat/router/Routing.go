package router

import (
	"lecter/goserver/internal/app/gochat/controller"
	"lecter/goserver/internal/app/gochat/service"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	// Basic認証のグループを設定
	auth := service.AuthenticationService{}
	userApi := r.Group("/api/v1/users")
	userApi.Use(auth.BasicAuthorization)

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
}
