package controller

import (
	"lecter/goserver/internal/app/gochat/application/service"
	"lecter/goserver/internal/app/gochat/presentation/controller/request"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageController struct {
	MessageService service.MessageService
}

/*
 * メッセージ一覧を取得する
 */
func (mc MessageController) Index(ctx *gin.Context) {
	// ユーザーID取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// チャンネルID取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}
	// 一覧取得リクエストを取得
	var params request.MessageListRequestParam
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(response.ValidationError("invalid query parameters").ToResponse())
		return
	}
	// メッセージリストを取得
	models, error := mc.MessageService.GetMessages(userId.(uuid.UUID), channelId, params.LastId, params.Limit, params.Language)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, models)
}

/*
 * メッセージを作成する
 */
func (mc MessageController) Create(ctx *gin.Context) {
	// ユーザーID取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// チャンネルID取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}
	// 一覧取得リクエストを取得
	var request request.MessageCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid requestBody").ToResponse())
		return
	}
	// メッセージを作成
	model, error := mc.MessageService.CreateMessage(userId.(uuid.UUID), channelId, request.Message)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

func (mc MessageController) Update(ctx *gin.Context) {
	// ユーザーID取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// チャンネルID取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}
	// メッセージID取得
	messageId, err := uuid.Parse(ctx.Param("messageId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid messageId").ToResponse())
		return
	}
	// 一覧取得リクエストを取得
	var request request.MessageUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid requestBody").ToResponse())
		return
	}
	// メッセージを作成
	model, error := mc.MessageService.UpdateMessage(userId.(uuid.UUID), channelId, messageId, request.Message)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

func (mc MessageController) Delete(ctx *gin.Context) {
	// ユーザーID取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// チャンネルID取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}
	// メッセージID取得
	messageId, err := uuid.Parse(ctx.Param("messageId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid messageId").ToResponse())
		return
	}
	// メッセージを作成
	error := mc.MessageService.DeleteMessage(userId.(uuid.UUID), channelId, messageId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.Status(http.StatusNoContent)
}
