package controller

import (
	"fmt"
	"lecter/goserver/internal/app/gochat/controller/request"
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/enum/language"
	"lecter/goserver/internal/app/gochat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct{}

var userService = service.UserService{}

/*
 * リクエスト送信者のユーザー情報を取得
 */
func (uc UserController) Select(ctx *gin.Context) {
	// ユーザー名取得
	name, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(response.ValidationError("Invalid username").ToResponse())
		return
	}
	// ユーザーモデルを取得
	model, error := userService.GetUserByName(name.(string))
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model.ToResponse())
}

/*
 * ユーザーを作成
 */
func (uc UserController) Create(ctx *gin.Context) {
	// 作成リクエストを取得
	var request request.UserCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid request body").ToResponse())
		return
	}

	// ユーザー作成
	model, error := userService.CreateUser(request.Name, request.Password)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model.ToResponse())
}

/*
 * ユーザーの更新
 */
func (uc UserController) Update(ctx *gin.Context) {
	// ユーザーID取得
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}

	// 更新リクエスト取得
	var request request.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(response.ValidationError("invalid request").ToResponse())
		return
	}

	langCode, err := language.GetLanguageFromCode(request.Language)
	if err != nil {
		ctx.JSON(response.ValidationError("invalid languageCode").ToResponse())
		return
	}

	// ユーザー更新
	model, error := userService.UpdateUser(userId, request.Name, langCode)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model.ToResponse())
}
