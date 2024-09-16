package controller

import (
	"lecter/goserver/service"

	"github.com/gin-gonic/gin"
)

type VersionController struct {
}

var versionService = service.VersionService{}

/*
 * バージョンを取得
 */
func (vc VersionController) Index(ctx *gin.Context) {
	ctx.JSON(200, versionService.GetVersion())
}