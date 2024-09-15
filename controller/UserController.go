package controller

import (
	"fmt"
	"lecter/hello/common"
	"lecter/hello/controller/request"
	"lecter/hello/model"
	"lecter/hello/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct{}

/*
 * ユーザー一覧を取得
 */
func (uc UserController) Index(ctx *gin.Context) {
	userRepository := repository.UserRepository{}
	models := userRepository.Index()
	if models == nil {
		ctx.JSON(common.InternalErrorResponse("failed to get models"))
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
		ctx.JSON(common.ValidationErrorResponse("invalid id"))
		return
	}
	userRepository := repository.UserRepository{}
	model := userRepository.Select(userId)
	if model == nil {
		ctx.JSON(common.NotFoundErrorResponse("user not found"))
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
		ctx.JSON(common.ValidationErrorResponse("validation error"))
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(common.InternalErrorResponse("uuid error"))
		return
	}
	model := &model.UserModel{
		Id: id,
		Name: request.Name,
		Url: request.Url,
	}

	userRepository := repository.UserRepository{}
	model = userRepository.Insert(*model)

	if model == nil {
		ctx.JSON(common.InternalErrorResponse("db connection error"))
		fmt.Println("failed to insert UserModel")
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
		ctx.JSON(common.ValidationErrorResponse("invalid userId"))
		return
	}

	var request request.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(common.ValidationErrorResponse("invalid request"))
		return
	}

	userRepository := repository.UserRepository{}
	model := userRepository.Select(userId)

	if model == nil {
		ctx.JSON(common.NotFoundErrorResponse("user not found"))
		return
	}

	model.Name = request.Name
	model.Url = request.Url

	model = userRepository.Update(*model)
	ctx.JSON(http.StatusOK, *model)
}