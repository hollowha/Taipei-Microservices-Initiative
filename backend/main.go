package main

import (
	"backend/middleware"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 設定全域的 CORS middleware
	r.Use(middleware.CORSMiddleware())

	// 初始化所有路由
	routes.InitRoutes(r)

	// 啟動伺服器
	r.Run(":8080") // 預設會在本機的 8080 port 啟動
}
