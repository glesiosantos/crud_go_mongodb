package services

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository"
)

func NewUserDomainService(
	repository repository.UserRepository,
) UserDomainService {
	return &userDomainService{repository}
}

type userDomainService struct {
	repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *errorsc.RestErr)
	UpdateUser(string, model.UserDomainInterface) *errorsc.RestErr
	FindUserByEmail(email string) (model.UserDomainInterface, *errorsc.RestErr)
	FindUserById(id string) (model.UserDomainInterface, *errorsc.RestErr)
	DeleteUser(string) *errorsc.RestErr
}
