package repository

import (
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	USER_COLLECTION = "MONGO_DB_USERS_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	SaveUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errorsc.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *errorsc.RestErr)
	FindUserById(Id string) (model.UserDomainInterface, *errorsc.RestErr)
}
