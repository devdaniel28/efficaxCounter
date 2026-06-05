package main

import (
	"database/sql"
	"fmt"
	"log"

	"efficaxcounter/cmd/database"
	"efficaxcounter/cmd/router"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatal("Failed connect to database ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Ping Db falied ", err)
	}

	serve := gin.Default()
	serve.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"title":  "Efficax Counter",
			"status": "ok",
			"code":   200,
		})
	})

	router.Users(&serve.RouterGroup, db)

	serve.Run(":8000")
	fmt.Println("Servidor na Porta 8000")
}
