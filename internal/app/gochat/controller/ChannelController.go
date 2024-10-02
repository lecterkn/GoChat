package controller

import (
	"lecter/goserver/internal/app/gochat/controller/request"
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChannelController struct{}

var channelService = service.ChannelService{}

/*
 * チャンネル一覧を取得する
 */
func (ChannelController) Index(ctx *gin.Context) {
	// チャンネル一覧取得
	models, error := channelService.GetChannels()
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, models)
}

/*
 * チャンネルを取得する
 */
func (ChannelController) Select(ctx *gin.Context) {
	// チャンネルID取得
	id, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}
	// チャンネル取得
	model, error := channelService.GetChannel(id)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

/*
 * チャンネルを作成する
 */
func (ChannelController) Create(ctx *gin.Context) {
	// 作成リクエスト取得
	var request request.ChannelCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid requestBody").ToResponse())
		return
	}

	// 作成者ユーザーID取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.InternalError("failed to get userId").ToResponse())
		return
	}

	// チャンネル作成
	model, error := channelService.CreateChannel(userId.(uuid.UUID), request.Name, *request.Permission)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

/*
 * チャンネル更新する
 */
func (ChannelController) Update(ctx *gin.Context) {
	// 送信者のユーザーIDを取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.InternalError("failed to get userId").ToResponse())
		return
	}

	// チャンネルIDを取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}

	// チャンネル更新リクエストを取得
	var request request.ChannelUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid requestBody").ToResponse())
		return
	}

	// チャンネルを更新
	model, error := channelService.UpdateChannel(channelId, userId.(uuid.UUID), request.Name, *request.Permission)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

/*
 * チャンネルを削除する
 */
func (ChannelController) Delete(ctx *gin.Context) {
	// 送信者のユーザーIDを取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.InternalError("failed to get userId").ToResponse())
		return
	}

	// チャンネルID取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}

	// チャンネル削除
	error := channelService.DeleteChannel(userId.(uuid.UUID), channelId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.Status(http.StatusNoContent)
}
