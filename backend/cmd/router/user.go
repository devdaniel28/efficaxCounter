package router

import (
	"database/sql"
	"efficaxcounter/cmd/middleware"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	CreatAt time.Time `json:"created_at"`
}

func Users(rtr *gin.RouterGroup, db *sql.DB) {
	userGroup := rtr.Group("/user")
	userGroup.Use(middleware.Authrequired())

	userGroup.POST("/", func(ctx *gin.Context) {
		typeUser := User{}
		err := json.NewDecoder(ctx.Request.Body).Decode(&typeUser)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid Request",
			})
			return
		}

		if typeUser.Name == "" || typeUser.Email == "" {
			ctx.JSON(400, gin.H{
				"error": "Name and email are required",
			})
			return
		}

		var userId int
		err = db.QueryRow(
			"INSERT INTO ecounter.users (name, email) VALUES ($1, $2) RETURNING id",
			typeUser.Name, typeUser.Email,
		).Scan(&userId)

		if err != nil {
			fmt.Println("Serve error: ", err)
			ctx.JSON(500, gin.H{
				"error": "Failed to create user",
				"code": 500,
			})
			return
		}

		typeUser.Id = userId
		typeUser.CreatAt = time.Now()
		
		ctx.JSON(201, gin.H{
			"user": typeUser,
		})
	
	})

	userGroup.GET("/list", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT id, name, email, created_at FROM ecounter.users")

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Error retrieving data: " + err.Error(),
			})
			return
		}

		var user []User

		for rows.Next() {
			var users User
			err := rows.Scan(&users.Id, &users.Name, &users.Email, &users.CreatAt)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": "Error retrieving data: " + err.Error(),
				})
				return
			}
			user = append(user, users)
		}

		if err = rows.Err(); err != nil {
			ctx.JSON(500, gin.H{
				"error": "Error retrieving data: " + err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"users": user,
		})

	})
}