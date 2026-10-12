package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error ao carregar as variaveis de ambiente")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
