package model

import "github.com/google/uuid"

type UserModel struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type UserTable struct {
	Id []byte
	Name string
	Url string
}

func (userModel UserModel) ToTable() *UserTable {
	return &UserTable{
		Id: userModel.Id[:],
		Name: userModel.Name,
		Url: userModel.Url,
	}
}

func (userTable UserTable) ToModel() *UserModel {
	id, err := uuid.FromBytes(userTable.Id)
	if err != nil {
		return nil
	}
	return &UserModel{
		Id: id,
		Name: userTable.Name,
		Url: userTable.Url,
	}
}