package router

import (
	"lecter/hello/controller"

	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine)  {
	// controllers
	vc := controller.VersionController{}
	r.GET("/version", vc.Index)

	uc := controller.UserController{}
	r.GET("/users", uc.Index)
	r.POST("/users", uc.Create)
}