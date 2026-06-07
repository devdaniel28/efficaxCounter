package router

import (
	"database/sql"
	"efficaxcounter/cmd/middleware"
	"efficaxcounter/cmd/models"

	"github.com/gin-gonic/gin"
)


func Counters(rtr *gin.RouterGroup, db *sql.DB) {
	counterGroup := rtr.Group("/counter")
	counterGroup.Use(middleware.Authrequired())

	counterGroup.POST("/", func(ctx *gin.Context) {
		var counter models.Counter

		if err := ctx.ShouldBindJSON(&counter); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		var exist bool
		
		queryStr := `SELECT EXISTS(SELECT 1 FROM ecounter.users WHERE id = $1)`
		db.QueryRow(queryStr, counter.UserId).Scan(&exist)

		if !exist {
			ctx.JSON(400, gin.H{
				"error": "Count not found",
			})
			return
		}

		var counterId int
		query := `INSERT INTO ecounter.quotes (name, userid) VALUES ($1, $2) RETURNING id`
		err := db.QueryRow(query, counter.Name, counter.UserId).Scan(&counterId)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Failed create counter",
			})
		}

		ctx.JSON(200, gin.H{
			"id": counterId,
			"counters": counter,
		})
	})
	
	counterGroup.GET("/list", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT id, name, quote, userid FROM ecounter.quotes")

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Error retrieving data: " + err.Error(),
			})
			return
		}

		var counters []models.Counter

		for rows.Next() {
			var counter models.Counter
			err := rows.Scan(&counter.Id, &counter.Name, &counter.Quote, &counter.UserId)
			if err != nil {
				ctx.JSON(500, gin.H{
					"error": "Error retrieving data: " + err.Error(),
				})
				return
			}

			counters = append(counters, counter)
		}

		ctx.JSON(200, gin.H{
			"counters": counters,
		})
	})

}