package service

import (
	"encoding/base64"
	"lecter/goserver/common"
	"lecter/goserver/controller/response"
	"lecter/goserver/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticationService struct{}

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

	if (!basicAuthorize(name, password)) {
		c.JSON(response.Unauthorized("invalid name or password").ToResponse())
		c.Abort()
		return 
	}
	c.Next()
}

func basicAuthorize(username, password string) bool {
	userRepository := repository.UserRepository{}
	userModel, err := userRepository.SelectByName(username)
	if err != nil {
		return false
	}
	return common.HashEquals(password, userModel.Password)
}