package routers

import (
	"task_manager/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", taskController.GetAllTasks)
	r.GET("/tasks/:id", taskController.GetTaskByID)
	r.POST("/tasks", taskController.CreateTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	r.POST("/register", userController.RegisterUser)

	return r
}
