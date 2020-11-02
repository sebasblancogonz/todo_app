package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/todo_app/pkg/handler/task"
)

//Routes struct
type Routes struct{}

const (
	tasksResource = "/tasks"
	taskResource  = tasksResource + "/task"
)

//StartGin will start the server
func (c Routes) StartGin() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/", welcome)
		api.GET(tasksResource, task.GetAllTasks)
		api.POST(taskResource, task.CreateTask)
		api.PATCH(taskResource)
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
