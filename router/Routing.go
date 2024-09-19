package router

import (
	"lecter/goserver/controller"
	"lecter/goserver/service"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine)  {
	// BasicAuthorization by User
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
}