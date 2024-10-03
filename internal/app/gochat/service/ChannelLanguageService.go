package service

import (
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/enum/language"
	"lecter/goserver/internal/app/gochat/model"
	"lecter/goserver/internal/app/gochat/repository"

	"github.com/google/uuid"
)

type ChannelLanguageService struct{}

var channelLanguageRepository = repository.ChannelLanguageRepository{}

func (ChannelLanguageService) GetChannelLanguages(channelId uuid.UUID) ([]model.ChannelLanguageModel, *response.ErrorResponse) {
	models, err := channelLanguageRepository.Index(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	return models, nil
}

func (ChannelLanguageService) SaveChannelLanguages(userId, channelId uuid.UUID, langs []language.Language) ([]model.ChannelLanguageModel, *response.ErrorResponse) {
	var models []model.ChannelLanguageModel
	// チャンネルの存在確認
	channel, err := channelRepository.Select(channelId)
	if err != nil {
		return nil, response.NotFoundError("the channel does not exist")
	}
	// 権限確認
	if !IsChannelOwner(userId, *channel) {
		return nil, response.ForbiddenError("permission error")
	}
	// チャンネル言語モデルを作成
	for _, lang := range langs {
		models = append(models, model.ChannelLanguageModel{
			ChannelId: channelId,
			Language:  lang,
		})
	}
	// チャンネル一括削除・挿入
	err = channelLanguageRepository.Delete(channelId)
	if err != nil {
		return nil, response.InternalError("failed to update channel languages")
	}
	models, err = channelLanguageRepository.InsertAll(models)
	if err != nil {
		return nil, response.InternalError("failed to update channel languages")
	}
	return models, nil
}
