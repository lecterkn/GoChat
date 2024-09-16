package service

import (
	"lecter/goserver/common"
	"lecter/goserver/model"
)

type VersionService struct{}

/*
 * アプリケーションコンフィグのバージョン情報を取得
 */
func (vc VersionService) GetVersion() *model.VersionModel {
	return &model.VersionModel{
		Name: common.ApplicationConfig.Name,
		Version: common.ApplicationConfig.Version,
	}
}