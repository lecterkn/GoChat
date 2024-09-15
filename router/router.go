package router

import (
	"lecter/hello/controller"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine)  {
	// Version
	vc := controller.VersionController{}
	r.GET("/version", vc.Index)

	// User
	uc := controller.UserController{}
	r.GET("/users", uc.Index)
	r.GET("/users/:user_id", uc.Select)
	r.POST("/users", uc.Create)
}