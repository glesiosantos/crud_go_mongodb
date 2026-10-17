package repository

import (
	"context"
	"encoding/json"
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
	logger.Info("Iniciando salvamento do usuário no repositório")

	collectionName := os.Getenv(USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	jsonValue, err := userDomain.GetJsonValue()
	if err != nil {
		return nil, errorsc.NewInternalError(err.Error())
	}

	var document bson.M
	if err := json.Unmarshal([]byte(jsonValue), &document); err != nil {
		return nil, errorsc.NewInternalError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), document)
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
