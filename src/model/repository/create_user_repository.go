package repository

import (
	"context"
	"os"

	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ur *userRepository) SaveUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Iniciando salvamento do usuário no repositório")

	collectionName := os.Getenv(USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConverterDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, errorsc.NewInternalError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)
	return converter.ConverterEntityToDomain(*value), nil

}
