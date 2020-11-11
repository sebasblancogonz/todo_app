package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	api := r.Group("/api")
	{
		api.GET("/", welcome)
		api.GET(tasksResource, task.GetTasks)
		api.GET(taskResource, task.GetTask)
		api.POST(taskResource, task.CreateTask)
		api.PATCH(taskResource, task.UpdateTaskStatus)
		api.DELETE(taskResource, task.DeleteTask)
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
