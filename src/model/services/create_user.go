package services

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	repository, err := ud.repository.SaveUser(userDomain)

	if err != nil {
		logger.Error("Init create user model", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", repository.GetId()),
		zap.String("journey", "createUser"))

	return repository, nil
}
