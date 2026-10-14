package repository

import (
	"context"
	"os"

	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	USER_COLLECTION = "MONGO_DB_USERS_COLLECTION"
)

func (ur *userRepository) SaveUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Init created user repository")

	userCollection := os.Getenv(USER_COLLECTION)
	collection := ur.databaseConnection.Collection(userCollection)

	value, err := userDomain.GetJsonValue()
	if err != nil {
		return nil, errorsc.NewInternalError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, errorsc.NewInternalError(err.Error())
	}

	id, ok := result.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, errorsc.NewInternalError("Tipo inesperado para o ID inserido")
	}

	userDomain.SetID(id.Hex())

	return userDomain, nil
}
