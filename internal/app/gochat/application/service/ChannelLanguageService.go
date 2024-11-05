package service

import (
	"lecter/goserver/internal/app/gochat/domain/entity"
	"lecter/goserver/internal/app/gochat/domain/enum/language"
	"lecter/goserver/internal/app/gochat/domain/repository"
	"lecter/goserver/internal/app/gochat/presentation/controller/response"

	"github.com/google/uuid"
)

type ChannelLanguageService struct {
	ChannelRepository         repository.ChannelRepository
	ChannelLanguageRepository repository.ChannelLanguageRepository
}

func NewChannelLanguageService (
	channelRepository repository.ChannelRepository,
	channelLanguageRepository repository.ChannelLanguageRepository,
	) ChannelLanguageService {
	return ChannelLanguageService{
		ChannelRepository: channelRepository,
		ChannelLanguageRepository: channelLanguageRepository,
	}
}

func (cls ChannelLanguageService) GetChannelLanguages(channelId uuid.UUID) ([]entity.ChannelLanguageEntity, *response.ErrorResponse) {
	models, err := cls.ChannelLanguageRepository.Index(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	return models, nil
}

func (cls ChannelLanguageService) SaveChannelLanguages(userId, channelId uuid.UUID, langs []language.Language) ([]entity.ChannelLanguageEntity, *response.ErrorResponse) {
	var models []entity.ChannelLanguageEntity
	// チャンネルの存在確認
	channel, err := cls.ChannelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	// 権限確認
	if !IsChannelOwner(userId, *channel) {
		return nil, response.ForbiddenError("permission error")
	}
	// チャンネル言語モデルを作成
	for _, lang := range langs {
		models = append(models, entity.ChannelLanguageEntity{
			ChannelId: channelId,
			Language:  lang,
		})
	}
	// チャンネル一括削除・挿入
	err = cls.ChannelLanguageRepository.Delete(channelId)
	if err != nil {
		return nil, response.InternalError("failed to update channel languages")
	}
	models, err = cls.ChannelLanguageRepository.InsertAll(models)
	if err != nil {
		return nil, response.InternalError("failed to update channel languages")
	}
	return models, nil
}
