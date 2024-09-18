package controller

import (
	"lecter/goserver/controller/request"
	"lecter/goserver/controller/response"
	"lecter/goserver/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileController struct{}

var userProfileService = service.UserProfileService{}
var authenicateService = service.RelationAuthenticationService{}

func (upc UserProfileController) Select(ctx *gin.Context) {
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(response.InternalError("failed to get username").ToResponse())
		return
	}
	if error := authenicateService.IsUserRelated(userId, username.(string)); error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	model, error := userProfileService.SelectUserProfile(userId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}

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

	model, error := userProfileService.UpdateUserProfile(userId, request.DisplayName, request.Url, request.Description)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}