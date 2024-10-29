package controller

import (
	"lecter/goserver/internal/app/gochat/application/service"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/presentation/controller/request"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChannelLanguageController struct {
	ChannelLanguageService service.ChannelLanguageService
}

/*
 * チャンネルの言語一覧を取得
 */
func (clc ChannelLanguageController) Index(ctx *gin.Context) {
	// チャンネルIDを取得
	channelId, err := uuid.Parse(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid channelId").ToResponse())
		return
	}
	// チャンネルの言語一覧を取得
	models, error := clc.ChannelLanguageService.GetChannelLanguages(channelId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	// モデルをレスポンスに変換
	response := response.ChannelLanguageResponse{
		ChannelId: channelId,
		Languages: []string{},
	}
	for _, model := range models {
		response.Languages = append(response.Languages, model.Language.ToCode())
	}
	ctx.JSON(http.StatusOK, response)
}

func (clc ChannelLanguageController) Save(ctx *gin.Context) {
	// ユーザーIDを取得
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
	var request request.ChannelLanguageRequest
	if err = ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid requestBody").ToResponse())
		return
	}
	var langs []language.Language
	for _, langCode := range request.Languages {
		lang, err := language.GetLanguageFromCode(langCode)
		if err != nil {
			ctx.JSON(response.ValidationError("invalid languageCode").ToResponse())
			return
		}
		langs = append(langs, lang)
	}
	models, error := clc.ChannelLanguageService.SaveChannelLanguages(
		userId.(uuid.UUID),
		channelId,
		langs)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	response := response.ChannelLanguageResponse{
		ChannelId: channelId,
		Languages: []string{},
	}
	for _, lang := range models {
		response.Languages = append(response.Languages, lang.Language.ToCode())
	}
	ctx.JSON(http.StatusOK, response)
}
