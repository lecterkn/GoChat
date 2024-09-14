package controller

import (
	"fmt"
	"lecter/hello/common"
	"lecter/hello/model"
	"lecter/hello/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct{}

type UserRequest struct {
	Name string `json:"name" binding:"required,min=1,max=20"`
	Url string `json:"url" binding:"required"`
}

func (uc UserController) Index(ctx *gin.Context) {
	userRepository := repository.UserRepository{}
	models := userRepository.Index()
	if models == nil {
		ctx.JSON(http.StatusInternalServerError, common.InternalErrorResponse("failed to get models"))
		return
	}
	ctx.JSON(http.StatusOK, models)
}

func (uc UserController) Create(ctx *gin.Context) {
	var request UserRequest
	// バリデーションチェック
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ValidationErrorResponse("validation error"))
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.InternalErrorResponse("uuid error"))
		return
	}
	model := model.UserModel{
		Id: id,
		Name: request.Name,
		Url: request.Url,
	}

	userRepository := repository.UserRepository{}
	result := userRepository.Insert(&model)

	if result == nil {
		ctx.JSON(http.StatusInternalServerError, common.InternalErrorResponse("db connection error"))
		fmt.Println("failed to insert UserModel")
		return
	}
	ctx.JSON(http.StatusOK, *result)
}