package controller

import (
	"lecter/goserver/internal/app/gochat/service"

	"github.com/gin-gonic/gin"
)

type VersionController struct {
}

var versionService = service.VersionService{}

// @Summary バージョンを取得
// @Accept json
// @Produce json
// @Router /version [get]
func (vc VersionController) Index(ctx *gin.Context) {
	ctx.JSON(200, versionService.GetVersion())
}
