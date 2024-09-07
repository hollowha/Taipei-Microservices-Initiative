package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Create RouterGroup for API
	api := r.Group("/api")
	{
		api.GET("/example", controllers.Example)
		api.POST("/upload/image", controllers.UploadImage)
		// Additional routes can be added here

		// User routes
		user := api.Group("/users")
		{
			user.GET("/", controllers.GetAllUsers)
			user.GET("/image/:filename", controllers.GetUserIMGByFileName)
			user.GET("/:id", controllers.GetUserByID)
			user.POST("/", controllers.CreateUser)
			user.PUT("/:id", controllers.UpdateUserByID)
			user.DELETE("/:id", controllers.DeleteUserByID)
		}

		// Activity routes
		activity := api.Group("/activitys")
		{
			activity.GET("/all", controllers.GetAllActivitys)
			activity.POST("/", controllers.CreateEvent)
			activity.GET("/:title", controllers.GetEvents)
			activity.GET("/image/:imageTitle", controllers.ServeImage)
		}

		autofillform := api.Group("/autofillform")
		{

			autofillform.GET("/:formname", controllers.FillFormbyID)
			autofillform.GET("/detail/:formname", controllers.GetFormDetail)

		}

		report := api.Group("/report")
		{
			report.GET("/all", controllers.GetAllReports)
			report.POST("/", controllers.CreateReport)
			report.DELETE("/:id", controllers.DeleteReportByID)
			//report.GET("/:id", controllers.GetReportByID)
		}

		vote := api.Group("/vote")
		{
			vote.GET("/activities", controllers.GetVoteActivities) // 獲取所有投票活動及其選項
			vote.POST("/submit", controllers.SubmitVote)           // 提交投票
			vote.POST("/create", controllers.CreateVoteActivities)
		}
	}
}
