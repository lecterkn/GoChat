package service

import (
	"lecter/goserver/controller/response"
	"lecter/goserver/repository"

	"github.com/google/uuid"
)

type RelationAuthenticationService struct{}

func (ras RelationAuthenticationService) IsUserRelated(id uuid.UUID, name string) (*response.ErrorResponse) {
	var userRepository = repository.UserRepository{}
	model, err := userRepository.Select(id)
	if err != nil {
		return response.NotFoundError("user not found")
	}
	if model.Name != name {
		return response.ValidationError("permission denied")
	}
	return nil
}