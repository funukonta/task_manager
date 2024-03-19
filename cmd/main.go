package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/funukonta/task_manager/internal/routers"
	"github.com/funukonta/task_manager/pkg"
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

	db, err := pkg.ConnectPostgres()
	if err != nil {
		log.Fatal("Connect DB err:", err.Error())
	}

	// gin init
	r := gin.Default()

	routers.Routrer(r, db)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	r.Run(port)

}

// migrate create -ext sql -dir migrate crate_table_task
// migrate create -ext sql -dir migrate crate_table_user

// docker postges
// docker run --name task-manager -e POSTGRES_PASSWORD=task-manager -p 5432:5432 -d postgres && sleep 3 && docker exec -it task-manager psql -U postgres -d postgres -c "CREATE DATABASE task_manager;"
