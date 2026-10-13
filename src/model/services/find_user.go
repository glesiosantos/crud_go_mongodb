package services

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(id string) (*model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))
	return nil, nil
}
