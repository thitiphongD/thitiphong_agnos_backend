package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/thitiphongD/thitiphong_agnos_backend/db"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	db := database.InitDB()
    defer func() {
        sqlDB, err := db.DB()
        if err != nil {
            log.Fatalf("Error getting underlying database object: %s", err)
        }
        sqlDB.Close()
    }()

	if err := db.Error; err != nil {
        log.Fatalf("Error initializing database: %s", err)
    }

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(500, gin.H{
			"error_message": recovered,
		})
	}))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error_message": "404 page not found",
		})
	})
	r.Run()
}
