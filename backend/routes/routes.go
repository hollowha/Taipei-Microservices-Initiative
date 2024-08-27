package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 建立 RouterGroup
	api := r.Group("/api")
	{
		api.GET("/example", controllers.Example)
		// 可以在這裡添加更多的路由

		// user routes
		user := api.Group("/users")
		{
			user.GET("/", controllers.GetAllUsers)
			user.GET("/:id", controllers.GetUserByID)
			user.POST("/", controllers.CreateUser)
			user.PUT("/:id", controllers.UpdateUserByID)
			user.DELETE("/:id", controllers.DeleteUserByID)
		}

		// activity routes
		activity := api.Group("/activitys")
		{
			activity.GET("/", controllers.GetAllActivitys)
		}
	}
}
