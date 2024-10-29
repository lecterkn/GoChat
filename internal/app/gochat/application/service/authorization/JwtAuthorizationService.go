package authorization

import (
	"lecter/goserver/internal/app/gochat/common"
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtAuthorizationService struct {
	UserRepository repository.UserRepository
}

/*
 * JWT認証を行うミドルウェア
 */
func (JwtAuthorizationService) JwtAuthorization(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		ctx.Header("WWW-Authenticate", `Bearer realm="Authorization Required"`)
		ctx.JSON(response.Unauthorized("authorization required").ToResponse())
		ctx.Abort()
		return
	}
	jwtSecretKey := os.Getenv("SECRET_KEY")
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(response.Unauthorized("invalid token").ToResponse())
		ctx.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(response.Unauthorized("invalid claims").ToResponse())
		ctx.Abort()
		return
	}
	userIdString, ok := claims["sub"].(string)
	if !ok {
		ctx.JSON(response.Unauthorized("invalid claims sub").ToResponse())
		ctx.Abort()
		return

	}
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		ctx.JSON(response.Unauthorized("invalid claims sub format").ToResponse())
		ctx.Abort()
		return
	}
	ctx.Set("userId", userId)

	userName, ok := claims["name"].(string)
	if !ok {
		ctx.JSON(response.Unauthorized("invalid claims name").ToResponse())
		ctx.Abort()
		return
	}
	ctx.Set("username", userName)

	ctx.Next()
}

func (jas JwtAuthorizationService) Login(ctx *gin.Context) {
	var requset Credentials
	if err := ctx.ShouldBindJSON(&requset); err != nil {
		ctx.JSON(response.ValidationError("invalid request").ToResponse())
		return
	}
	jwtToken, error := jas.generateJwt(requset)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

type Credentials struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

func (jas JwtAuthorizationService) generateJwt(credentials Credentials) (string, *response.ErrorResponse) {
	jwtSecretKey := os.Getenv("SECRET_KEY")
	if jwtSecretKey == "" {
		return "", response.InternalError("secret key error")
	}
	exists, userModel := jas.userAuthorize(credentials.Username, credentials.Password)
	if !exists || userModel == nil {
		return "", response.Unauthorized("authorization error")
	}
	claims := jwt.MapClaims{
		"sub":  userModel.Id,
		"name": userModel.Name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", response.InternalError("jwt signing error")
	}
	return jwtToken, nil
}

func (jas JwtAuthorizationService) userAuthorize(username, password string) (bool, *entity.UserEntity) {
	userModel, err := jas.UserRepository.SelectByName(username)
	if err != nil {
		return false, nil
	}
	return common.HashEquals(password, userModel.Password), userModel
}
