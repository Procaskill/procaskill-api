package routes

import (
	"procaskill/controllers"
	"procaskill/middlewares"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	protectedTaskRoutes := router.Group("/api/tasks")
	protectedTaskRoutes.Use(middlewares.JwtMiddleware())
	{
		protectedTaskRoutes.GET("/list", func(c *gin.Context) {
			controllers.ListTasksController(c)
		});
		protectedTaskRoutes.GET("/:id", func(c *gin.Context) {
			controllers.GetTaskController(c)
		});
		protectedTaskRoutes.POST("/create", func(c *gin.Context) {
			controllers.CreateTaskController(c)
		});
		protectedTaskRoutes.PUT("/update/:id", func(c *gin.Context) {
			controllers.UpdateTaskController(c)
		});
		protectedTaskRoutes.DELETE("/delete/:id", func(c *gin.Context) {
			controllers.DeleteTaskController(c)
		});
	}
}