package services

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *errorsc.RestErr {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))
	return nil
}
