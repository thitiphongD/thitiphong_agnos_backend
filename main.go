package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	db "github.com/thitiphongD/thitiphong_agnos_backend/db"
	"github.com/thitiphongD/thitiphong_agnos_backend/middlewares"
	"github.com/thitiphongD/thitiphong_agnos_backend/modules/password"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}

	defer func() {
		sqlDB, _ := database.DB()
		_ = sqlDB.Close()
	}()

	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.CustomRecoveryMiddleware())
	r.NoRoute(middlewares.NotFoundMiddleware())

	password.NewHTTPPassword(r)

	r.Run()
}
