package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"efficaxcounter/cmd/database"
	"efficaxcounter/cmd/router"

	"github.com/gin-contrib/cors"
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

	//* Config Cors
	serve.Use(cors.Default())
	serve.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Key"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	serve.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"title":  "Efficax Counter",
			"status": "ok",
			"code":   200,
		})
	})
	
	//* Routes
	router.Users(&serve.RouterGroup, db)
	router.Counters(&serve.RouterGroup, db)

	serve.Run(":8000")
	fmt.Println("Servidor na Porta 8000")
}
