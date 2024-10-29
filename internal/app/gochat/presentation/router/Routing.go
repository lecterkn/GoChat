package router

import (
	"context"
	_ "lecter/goserver/docs"
	"lecter/goserver/internal/app/gochat/application/service"
	"lecter/goserver/internal/app/gochat/application/service/authorization"
	implemenst "lecter/goserver/internal/app/gochat/infrastructure/repository/implements"
	"lecter/goserver/internal/app/gochat/presentation/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getRedis() service.RedisService {
	return service.RedisService{
		Context: context.TODO(),
		Client: *redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func Routing(r *gin.Engine) {
	// cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	// Swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// JWT認証のグループを設定
	ur := implemenst.UserRepositoryImpl{}
	jwtAuthorizationService := authorization.JwtAuthorizationService{
		UserRepository: ur,
	}
	r.POST("/api/v1/login", jwtAuthorizationService.Login)
	userApi := r.Group("/api/v1/users")
	userApi.Use(jwtAuthorizationService.JwtAuthorization)

	// Version
	vc := controller.VersionController{
		VersionService: service.VersionService{},
	}
	r.GET("/api/v1/version", vc.Index)

	// User
	ups := service.UserProfileService{
		UserProfileRepository: implemenst.UserProfileRepositoryImpl{},
	}
	uc := controller.UserController{
		UserService: service.UserService{
			UserRepository: ur,
		},
		UserProfileService: ups,
	}
	r.POST("api/v1/register", uc.Create)
	userApi.GET("/", uc.Select)
	userApi.PATCH("/:userId", uc.Update)

	// User Profile
	upc := controller.UserProfileController{
		UserProfileService: ups,
	}
	userApi.GET("/:userId/profiles", upc.Select)
	userApi.PUT("/:userId/profiles", upc.Update)

	cr := implemenst.ChannelRepositoryImpl{}
	cc := controller.ChannelController{
		ChannelService: service.ChannelService{
			ChannelRepository: cr,
		},
	}
	userApi.GET("/channels", cc.Index)
	userApi.POST("/channels", cc.Create)
	userApi.GET("/channels/:channelId", cc.Select)
	userApi.PATCH("/channels/:channelId", cc.Update)
	userApi.DELETE("/channels/:channelId", cc.Delete)

	mr := implemenst.MessageRepositoryImpl{}
	mc := controller.MessageController{
		MessageService: service.MessageService{
			MessageRepository: mr,
			ChannelRepository: cr,
			MessageDomainService: service.MessageDomainService{
				MessageRepository: mr,
			},
			RedisService: getRedis(),
		},
	}
	userApi.GET("/channels/:channelId/messages", mc.Index)
	userApi.POST("/channels/:channelId/messages", mc.Create)
	userApi.PATCH("/channels/:channelId/messages/:messageId", mc.Update)
	userApi.DELETE("/channels/:channelId/messages/:messageId", mc.Delete)

	clc := controller.ChannelLanguageController{
		ChannelLanguageService: service.ChannelLanguageService{
			ChannelRepository: cr,
			ChannelLanguageRepository: implemenst.ChannelLanguageRepositoryImpl{},
		},
	}
	userApi.GET("/channels/:channelId/languages", clc.Index)
	userApi.PUT("/channels/:channelId/languages", clc.Save)
}
