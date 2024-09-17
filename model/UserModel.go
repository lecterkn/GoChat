package model

import "github.com/google/uuid"

type UserModel struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Password []byte `json:"password"`
}

type UserTable struct {
	Id []byte
	Name string
	Password []byte
}

func (userModel UserModel) ToTable() *UserTable {
	return &UserTable{
		Id: userModel.Id[:],
		Name: userModel.Name,
		Password: userModel.Password,
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
		Password: userTable.Password,
	}
}