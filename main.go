package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/config/database"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()

	context := context.Background()
	database, err := database.NewMongoDBConnection(context)

	if err != nil {
		log.Fatal("Error ao carregar as variaveis de ambiente")
	}

	// init dependencias
	userController := initDependecies(database)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
