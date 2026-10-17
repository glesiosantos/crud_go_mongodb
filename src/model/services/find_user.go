package services

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Init findUserByEmail service", zap.String("journey", "findUserByEmail"))
	return ud.repository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserById(id string) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Init findUserById service", zap.String("journey", "findUserByEmail"))
	return ud.repository.FindUserById(id)
}
