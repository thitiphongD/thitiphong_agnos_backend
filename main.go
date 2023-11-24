package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/thitiphongD/thitiphong_agnos_backend/db"
	"github.com/thitiphongD/thitiphong_agnos_backend/middlewares"
	"github.com/thitiphongD/thitiphong_agnos_backend/modules/password"
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
	// r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CustomRecoveryMiddleware())
    r.NoRoute(middlewares.NotFoundMiddleware())
	
	password.InitPassword(r)

	r.Run()
}
