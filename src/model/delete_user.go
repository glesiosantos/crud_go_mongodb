package model

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) DeleteUser(id string) (*userDomain, *errorsc.RestErr) {
	logger.Info("Init deleteUser model", zap.String("journey", "deleteUser"))
	return nil, nil
}
