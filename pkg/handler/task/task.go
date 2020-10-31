package task

import (
	"log"
	"net/http"

	model_task "github.com/sebasblancogonz/model/task"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/todo_app/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TaskCollection static collection
const TaskCollection = "tasks"

//MongoConfig returns db
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		log.Println(err)
	}
	return db
}

//GetAllTasks will return all tasks
func GetAllTasks(c *gin.Context) {
	db := *MongoConfig()

	tasks := model_task.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{}).All(&tasks)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching all tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": &tasks,
	})
}

//GetTasksByStatus will return tasks given a status
func GetTasksByStatus(c *gin.Context) {
	db := *MongoConfig()

	status := c.Param("status")

	tasks := model_task.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{"status": &status})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching tasks by status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": &tasks,
	})
}
