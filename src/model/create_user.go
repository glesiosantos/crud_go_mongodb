package model

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *errorsc.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()
	return nil
}
