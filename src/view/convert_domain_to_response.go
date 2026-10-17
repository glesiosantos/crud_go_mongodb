package view

import (
	"github.com/glesiosantos/crud_go_mongodb/src/controller/dto/response"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
