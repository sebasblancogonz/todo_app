package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/todo_app/pkg/handler/task"
)

//Routes struct
type Routes struct{}

//StartGin will start the server
func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", welcome)
		api.GET("/tasks", task.GetAllTasks)
		api.GET("/tasks/task")
		api.POST("/tasks/task", task.CreateTask)
	}

	r.Run(":8000")
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Resource not found",
	})
}
