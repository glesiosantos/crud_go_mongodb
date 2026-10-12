package model

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) UpdateUser(id string) *errorsc.RestErr {
	logger.Info("Init updateUser model", zap.String("journey", "updateUser"))
	return nil
}
