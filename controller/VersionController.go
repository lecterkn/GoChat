package controller

import (
	"lecter/hello/model"
	"lecter/hello/common"
	"github.com/gin-gonic/gin"
)

type VersionController struct {
}

/*
 * バージョンを取得
 */
func (vc VersionController) Index(ctx *gin.Context) {
	model := model.VersionModel{
		Name: common.ApplicationConfig.Name,
		Version: common.ApplicationConfig.Version,
	}
	ctx.JSON(200, model)
}