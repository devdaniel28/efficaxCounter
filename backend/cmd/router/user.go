package router

import (
	"database/sql"
	"efficaxcounter/cmd/middleware"
	"efficaxcounter/cmd/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Users(rtr *gin.RouterGroup, db *sql.DB) {
	userGroup := rtr.Group("/user")
	userGroup.Use(middleware.Authrequired())

	userGroup.POST("/", func(ctx *gin.Context) {
		typeUser := models.User{}
		err := json.NewDecoder(ctx.Request.Body).Decode(&typeUser)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid Request" + err.Error(),
			})
			return
		}

		if typeUser.Name == "" || typeUser.Email == "" || typeUser.Password == ""{
			ctx.JSON(400, gin.H{
				"error": "Name, email and Password are required",
			})
			return
		}

		hashedPassword, err := middleware.HashPassword(typeUser.Password)
		if err != nil {
			ctx.JSON(500, gin.H{
				 "error": "Failed to process password",
			})
			return
		}

		var userId int
		err = db.QueryRow(
			"INSERT INTO ecounter.users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
			typeUser.Name, typeUser.Email, hashedPassword,
		).Scan(&userId)

		if err != nil {
			fmt.Println("Serve error: ", err)
			ctx.JSON(500, gin.H{
				"error": "Failed to create user" + err.Error(),
			})
			return
		}

		token, err := middleware.GerenerateJwt(userId, typeUser.Email)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Failed to generate authentication token" + err.Error(),
			})
			return
		}

		typeUser.Id = userId
		typeUser.CreatAt = time.Now()
		
		ctx.JSON(201, gin.H{
			"user": typeUser,
			"token": token,
		})
	
	})

	userGroup.GET("/list", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT id, name, email, password, created_at FROM ecounter.users")

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Error retrieving data: " + err.Error(),
			})
			return
		}

		var user []models.User

		for rows.Next() {
			var users models.User
			err := rows.Scan(&users.Id, &users.Name, &users.Email, &users.Password, &users.CreatAt)
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