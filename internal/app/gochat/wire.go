//go:build wireinject
// +build wireinject

package gochat

import (
	"github.com/google/wire"
	"lecter/goserver/internal/app/gochat/application/service"
	"lecter/goserver/internal/app/gochat/application/service/authorization"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/infrastructure/db"
	"lecter/goserver/internal/app/gochat/infrastructure/repository/implements"
	"lecter/goserver/internal/app/gochat/presentation/controller"
)

var databaseSet = wire.NewSet(
	db.Database,
)

var repositorySet = wire.NewSet(
	implements.NewUserRepositoryImpl,
	implements.NewUserProfileRepositoryImpl,
	implements.NewMessageRepositoryImpl,
	implements.NewChannelRepositoryImpl,
	implements.NewChannelLanguageRepositoryImpl,
)

var serviceSet = wire.NewSet(
	service.NewChannelLanguageService,
	service.NewChannelService,
	service.NewMessageDomainService,
	service.NewMessageService,
	service.NewRedisService,
	service.NewUserProfileService,
	service.NewUserService,
	service.NewVersionService,
	authorization.NewJwtAuthorizationService,
)

var controllerSet = wire.NewSet(
	controller.NewChannelController,
	controller.NewChannelLanguageController,
	controller.NewMessageController,
	controller.NewUserController,
	controller.NewUserProfileController,
	controller.NewVersionController,
)

type ControllerSet struct {
	ChannelController         controller.ChannelController
	ChannelLanguageController controller.ChannelLanguageController
	MessageController         controller.MessageController
	UserController            controller.UserController
	UserProfileController     controller.UserProfileController
	VersionController         controller.VersionController
}

func InitializeControllerSet() *ControllerSet {
	wire.Build(
		databaseSet,
		repositorySet,
		serviceSet,
		controllerSet,
		wire.Bind(new(repository.UserRepository), new(implements.UserRepositoryImpl)),
		wire.Bind(new(repository.UserProfileRepository), new(implements.UserProfileRepositoryImpl)),
		wire.Bind(new(repository.MessageRepository), new(implements.MessageRepositoryImpl)),
		wire.Bind(new(repository.ChannelRepository), new(implements.ChannelRepositoryImpl)),
		wire.Bind(new(repository.ChannelLanguageRepository), new(implements.ChannelLanguageRepositoryImpl)),
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil
}

func InitializeJwtAuthorizationService() *authorization.JwtAuthorizationService {
	wire.Build(
		databaseSet,
		repositorySet,
		wire.Bind(new(repository.UserRepository), new(implements.UserRepositoryImpl)),
		serviceSet,
	)
	return nil
}
