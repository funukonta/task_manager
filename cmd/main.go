package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Env
	evnPath := filepath.Join("..", ".env")
	err := godotenv.Load(evnPath)
	if err != nil {
		log.Fatal("Error load .env" + err.Error())
	}

	// gin init
	r := gin.Default()

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	r.Run(port)

}
