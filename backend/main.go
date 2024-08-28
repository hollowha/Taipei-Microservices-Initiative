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

func main() {

	// Initialize GORM with SQLite
	db, err := gorm.Open(sqlite.Open("service.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	// Auto-migrate the User model
	db.AutoMigrate(&controllers.User{})
	// 設定資料庫實例到控制器
	controllers.SetDB(db)

	r := gin.Default()

	// 設定全域的 CORS middleware
	r.Use(middleware.CORSMiddleware())

	// 初始化所有路由
	routes.InitRoutes(r)

	// 啟動伺服器
	r.Run(":8081") // 預設會在本機的 8080 port 啟動
}
