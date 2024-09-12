package controller

import (
	"lecter/hello/model"

	"github.com/gin-gonic/gin"
)

type VersionController struct {
}

func (vc VersionController) Index(ctx *gin.Context) {
	model := model.VersionModel{
		Name: "goGin",
		Version: "0.1",
	}
	ctx.JSON(200, model)
}