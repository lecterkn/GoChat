package service

import (
	"encoding/base64"
	"lecter/goserver/internal/app/gochat/common"
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/model"
	"lecter/goserver/internal/app/gochat/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthenticationService struct{}

/*
 * Basic認証を行う
 */
func (as AuthenticationService) BasicAuthorization(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Basic ") {
		c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
		c.JSON(response.Unauthorized("authorization required").ToResponse())
		c.Abort()
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[len("Basic "):])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		c.JSON(response.Unauthorized("invalid format").ToResponse())
		c.Abort()
		return
	}

	name := pair[0]
	password := pair[1]

	authorized, userModel := basicAuthorize(name, password)
	if !authorized {
		c.JSON(response.Unauthorized("invalid name or password").ToResponse())
		c.Abort()
		return
	}
	c.Set("username", userModel.Name)
	c.Set("userId", userModel.Id)
	c.Next()
}

/*
 * ユーザー名とユーザーIDの関連性があるか確認する
 */
func (as AuthenticationService) IsUserRelated(id uuid.UUID, name string) *response.ErrorResponse {
	var userRepository = repository.UserRepository{}
	model, err := userRepository.Select(id)
	if err != nil {
		return response.NotFoundError("user not found")
	}
	if model.Name != name {
		return response.ValidationError("permission denied")
	}
	return nil
}

func basicAuthorize(username, password string) (bool, *model.UserModel) {
	userRepository := repository.UserRepository{}
	userModel, err := userRepository.SelectByName(username)
	if err != nil {
		return false, nil
	}
	return common.HashEquals(password, userModel.Password), userModel
}
