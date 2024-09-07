package main

import (
	"backend/controllers"
	"backend/routes"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	db, err := gorm.Open(sqlite.Open("service.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err := db.AutoMigrate(&controllers.User{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	// load api key from .env file
	/*
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
 	*/

	controllers.SetDB(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.InitRoutes(r)

	r.Run(":8081")
}
