package controller

import (
	"lecter/goserver/controller/request"
	"lecter/goserver/controller/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileController struct{}

func (upc UserProfileController) Update(ctx *gin.Context) {
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}

	var request request.UserProfileUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid request").ToResponse())
		return
	}
}