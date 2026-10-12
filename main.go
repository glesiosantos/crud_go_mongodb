package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error ao carregar as variaveis de ambiente")
	}

	fmt.Println(os.Getenv("TEST"))
}
