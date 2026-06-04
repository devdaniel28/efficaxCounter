package main

import (
	"database/sql"
	"fmt"
	"time"

	"efficaxcounter/cmd/database"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	CreatAt time.Time `json:"created_at"`
}

var db *sql.DB

func main() {
	var err error
	db, err = database.ConnectDB()
	if err != nil {

	}

	defer db.Close()

	serve := gin.Default()
	serve.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"title":  "Efficax Counter",
			"status": "ok",
			"code":   200,
		})
	})

	serve.Run(":8000")
	fmt.Println("Servidor na Porta 8000")
}
