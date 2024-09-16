package controller

import (
	"fmt"
	"lecter/goserver/controller/request"
	"lecter/goserver/controller/response"
	"lecter/goserver/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct{}

var userService = service.UserService{}

/*
 * ユーザー一覧を取得
 */
func (uc UserController) Index(ctx *gin.Context) {
	models, error := userService.GetUsers()
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, models)
}

/*
 * 特定のユーザーを取得
 */
func (uc UserController) Select(ctx *gin.Context) {
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}

	model, error := userService.GetUser(userId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, model)
}

/*
 * ユーザーを作成
 */
func (uc UserController) Create(ctx *gin.Context) {
	var request request.UserCreateRequest
	// バリデーションチェック
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid request body").ToResponse())
		return
	}

	model, error := userService.CreateUser(request.Name, request.Url)

	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}

/*
 * ユーザーの更新
 */
func (uc UserController) Update(ctx *gin.Context) {
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}

	var request request.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(response.ValidationError("invalid request").ToResponse())
		return
	}

	model, error:= userService.UpdateUser(userId, request.Name, request.Url)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}