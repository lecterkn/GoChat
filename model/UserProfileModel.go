package model

import (
	"github.com/google/uuid"
)

type UserProfileModel struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID  `json:"userId"`
	DisplayName string  `json:"displayName"`
	Url         string  `json:"url"`
	Description string  `json:"description"`
}

type UserProfileTable struct {
	Id          []byte `json:"id"`
	UserId      []byte `json:"userId"`
	DisplayName string `json:"displayName"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func (upm UserProfileModel) ToTable() *UserProfileTable {
	return &UserProfileTable{
		Id: upm.Id[:],
		UserId: upm.UserId[:],
		DisplayName: upm.DisplayName,
		Url: upm.Url,
		Description: upm.Description,
	}
}

func (upm UserProfileModel) ToArray() (uuid.UUID, uuid.UUID, string, string, string) {
	return upm.Id, upm.UserId, upm.DisplayName, upm.Url, upm.Description
}

func (upt UserProfileTable) ToModel() *UserProfileModel {
	var id, userId uuid.UUID
	var err error
	id, err = uuid.FromBytes(upt.Id)
	if err != nil {
		return nil
	}
	userId, err = uuid.FromBytes(upt.UserId)
	if err != nil {
		return nil
	}
	return &UserProfileModel{
		Id: id,
		UserId: userId,
		DisplayName: upt.DisplayName,
		Url: upt.Url,
		Description: upt.Description,
	}
}

func (upt UserProfileTable) ToArray() (*[]byte, *[]byte, *string, *string, *string) {
	return &upt.Id, &upt.UserId, &upt.DisplayName, &upt.Url, &upt.Description
}

func (upt UserProfileTable) ToValues() ([]byte, []byte, string, string, string) {
	return upt.Id, upt.UserId, upt.DisplayName, upt.Url, upt.Description
}