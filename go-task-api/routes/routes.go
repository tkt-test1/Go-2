package routes

import (
	"go-task-api/controllers"
	"go-task-api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		protected := api.Group("/tasks")
		protected.Use(utils.AuthMiddleware())
		{
			protected.POST("", controllers.CreateTask)
			protected.GET("", controllers.ListTasks)
			protected.PUT("/:id", controllers.UpdateTask)
			protected.DELETE("/:id", controllers.DeleteTask)
		}
	}
}
