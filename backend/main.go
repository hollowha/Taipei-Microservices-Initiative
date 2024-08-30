package main

import (
	"backend/controllers"
	"backend/middleware"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// func main() {

// 	// Initialize GORM with SQLite
// 	db, err := gorm.Open(sqlite.Open("service.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("failed to connect database")
// 	}
// 	// Auto-migrate the User model
// 	db.AutoMigrate(&controllers.User{})
// 	// 設定資料庫實例到控制器
// 	controllers.SetDB(db)

// 	r := gin.Default()

// 	// 設定全域的 CORS middleware
// 	r.Use(middleware.CORSMiddleware())

// 	// 初始化所有路由
// 	routes.InitRoutes(r)

// 	// 啟動伺服器
// 	r.Run(":8081") // 預設會在本機的 8080 port 啟動
// }

func main() {
	// Set Gin to Release Mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize GORM with SQLite
	db, err := gorm.Open(sqlite.Open("service.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Check if the 'users' table exists, and create it if not
	if !db.Migrator().HasTable(&controllers.User{}) {
		if err := db.Migrator().CreateTable(&controllers.User{}); err != nil {
			log.Fatal("failed to create table: ", err)
		}
	}

	// Set database instance to controllers
	controllers.SetDB(db)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) // Set global CORS middleware
	routes.InitRoutes(r)               // Initialize all routes
	r.Run(":8081")                     // Default runs on localhost port 8081
}
