package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/config/database"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/controller"
	"github.com/glesiosantos/crud_go_mongodb/src/controller/routes"
	"github.com/glesiosantos/crud_go_mongodb/src/model/services"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()

	context := context.Background()

	database.NewMongoDBConnection(context)

	if err != nil {
		log.Fatal("Error ao carregar as variaveis de ambiente")
	}

	// init dependencias
	service := services.NewUserDomainService()
	controller := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
