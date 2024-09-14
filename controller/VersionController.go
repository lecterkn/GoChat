package controller

import (
	"lecter/hello/config"
	"lecter/hello/model"

	"github.com/gin-gonic/gin"
)

type VersionController struct {
}

func (vc VersionController) Index(ctx *gin.Context) {
	model := model.VersionModel{
		Name: config.ApplicationConfig.Name,
		Version: config.ApplicationConfig.Version,
	}
	ctx.JSON(200, model)
}