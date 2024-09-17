package service

import (
	"lecter/goserver/model"

	"github.com/google/uuid"
)

type UserProfileService struct{}

func (ups UserProfileService) UpdateUserProfile(userId uuid.UUID, DisplayName, Url, Description string) *model.UserProfileModel {
	
}