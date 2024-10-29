package controller

import (
	"lecter/goserver/internal/app/gochat/application/service"
	"lecter/goserver/internal/app/gochat/presentation/controller/request"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileController struct{
	UserProfileService service.UserProfileService
}

/*
 *	ユーザープロフィールを取得する
 */
func (upc UserProfileController) Select(ctx *gin.Context) {
	// ユーザーID取得
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// ユーザーモデル取得
	model, error := upc.UserProfileService.SelectUserProfile(userId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}

/*
 * ユーザープロフィールを新規作成・更新を行う
 */
func (upc UserProfileController) Update(ctx *gin.Context) {
	// 更新対象のユーザーID取得
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// リクエスト送信者のユーザー名取得
	requestUserId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(response.InternalError("failed to get username").ToResponse())
		return
	}
	// リクエスト送信者と対象のIDの一致確認
	if requestUserId != userId {
		ctx.JSON(response.ValidationError("permission error").ToResponse())
		return
	}

	// 更新リクエストボディ取得
	var request request.UserProfileUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid request").ToResponse())
		return
	}

	// ユーザープロフィールを更新
	model, error := upc.UserProfileService.UpdateUserProfile(userId, request.DisplayName, request.Url, request.Description)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}
