package router

import (
	"lecter/goserver/controller"
	"lecter/goserver/service"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine)  {
	auth := service.AuthenticationService{}
	userApi := r.Group("/api/v1/user")
	userApi.Use(auth.BasicAuthorization)
	// Version
	vc := controller.VersionController{}
	r.GET("/api/v1/version", vc.Index)

	// User
	uc := controller.UserController{}
	//r.GET("/api/v1/users", uc.Index)
	//r.GET("/api/v1/users/:userId", uc.Select)
	r.POST("api/v1/register", uc.Create)
	//r.PATCH("/api/v1/users/:userId", uc.Update)
}