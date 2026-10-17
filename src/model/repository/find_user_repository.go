package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository/entity"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Iniciando busca de usuario pelo email no repositório")

	collectionName := os.Getenv(USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)

			logger.Error(errorMessage, err, zap.String("journed", "findUserByEmail"))
			return nil, errorsc.NewBadRequestError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journed", "findUserByEmail"))
		return nil, errorsc.NewInternalError(errorMessage)
	}

	return converter.ConverterEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *errorsc.RestErr) {
	logger.Info("Iniciando busca de usuario pelo Id no repositório")

	collectionName := os.Getenv(USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	objectId, _ := bson.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this Id: %s", id)

			logger.Error(errorMessage, err, zap.String("journed", "findUserById"))
			return nil, errorsc.NewBadRequestError(errorMessage)
		}

		errorMessage := "Error trying to find user by Id"
		logger.Error(errorMessage, err, zap.String("journed", "findUserById"))
		return nil, errorsc.NewInternalError(errorMessage)
	}

	return converter.ConverterEntityToDomain(*userEntity), nil
}
