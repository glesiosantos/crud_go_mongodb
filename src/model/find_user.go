package model

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) FindUser(id string) (*userDomain, *errorsc.RestErr) {
	logger.Info("Init findUser model", zap.String("journey", "findUser"))
	return nil, nil
}
