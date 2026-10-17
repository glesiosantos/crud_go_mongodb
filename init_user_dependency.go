package main

import (
	"github.com/glesiosantos/crud_go_mongodb/src/controller"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository"
	"github.com/glesiosantos/crud_go_mongodb/src/model/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependecies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := services.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
