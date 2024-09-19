package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Password []byte `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserTable struct {
	Id []byte
	Name string
	Password []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (userModel UserModel) ToTable() *UserTable {
	return &UserTable{
		Id: userModel.Id[:],
		Name: userModel.Name,
		Password: userModel.Password,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
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
		CreatedAt: userTable.CreatedAt,
		UpdatedAt: userTable.UpdatedAt,
	}
}