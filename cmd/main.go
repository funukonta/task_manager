package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
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

	r.GET("/google/login", GoogleLogin)
	r.GET("/google/callback", GoogleCallback)
	routers.Routrer(r, db)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	r.Run(port)

}

func GoogleLogin(c *gin.Context) {
	googleConf := pkg.LoadConfig()

	oauthState := pkg.GenerateStateOauthCookie(c.Writer)
	url := googleConf.AuthCodeURL(oauthState)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	// get oauth state from cookie for this user
	oauthState, err := c.Cookie("oauthstate")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	state := c.Query("state")
	code := c.Query("code")
	c.Header("Content-Type", "application/json")

	// ERROR : Invalid OAuth State
	if state != oauthState {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		fmt.Fprintf(c.Writer, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := pkg.AppConfig.GoogleLogConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(c.Writer, "failed code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from google server
	response, err := http.Get(pkg.OauthGoogleUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(c.Writer, "failed getting user info: %s", err.Error())
		return
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(c.Writer, "failed read response: %s", err.Error())
		return
	}

	// send back response to browser
	c.Writer.Write(contents)
}

// migrate create -ext sql -dir migrate crate_table_task
// migrate create -ext sql -dir migrate crate_table_user

// docker postges
// docker run --name task-manager -e POSTGRES_PASSWORD=task-manager -p 5432:5432 -d postgres && sleep 3 && docker exec -it task-manager psql -U postgres -d postgres -c "CREATE DATABASE task_manager;"
