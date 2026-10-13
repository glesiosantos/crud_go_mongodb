package services

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct{}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *errorsc.RestErr
	UpdateUser(string, model.UserDomainInterface) *errorsc.RestErr
	FindUser(string) (*model.UserDomainInterface, *errorsc.RestErr)
	DeleteUser(string) *errorsc.RestErr
}
