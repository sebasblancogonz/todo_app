package task

import (
	"log"
	"net/http"

	model_todo "github.com/sebasblancogonz/model/task"

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

	tasks := model_todo.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{}).All(&tasks)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching all taskss",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": &tasks,
	})
}

func GetPen
