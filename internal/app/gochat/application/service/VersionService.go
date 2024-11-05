package service

import (
	"lecter/goserver/internal/app/gochat/common"
	"lecter/goserver/internal/app/gochat/infrastructure/model"
)

type VersionService struct{}

func NewVersionService() VersionService {
	return VersionService{}
}

/*
 * アプリケーションコンフィグのバージョン情報を取得
 */
func (vc VersionService) GetVersion() *model.VersionModel {
	return &model.VersionModel{
		Name:    common.ApplicationConfig.Name,
		Version: common.ApplicationConfig.Version,
	}
}
