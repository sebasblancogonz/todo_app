package task

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/todo_app/config"
	model_task "github.com/sebasblancogonz/todo_app/pkg/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	todo       = "TODO"
	inProgress = "IN_PROGRESS"
	done       = "DONE"
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

	println(status)

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

//CreateTask will create a new task
func CreateTask(c *gin.Context) {
	db := *MongoConfig()

	println(db.Name)

	task := model_task.Task{}

	err := c.Bind(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting request body",
			"error":   err,
		})
		return
	}

	task.Status = todo
	task.ID = bson.NewObjectId()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err = db.C(TaskCollection).Insert(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong creating a new task",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task created successfuly",
		"task":    &task,
	})
}

func UpdateTaskStatus(c *gin.Context) {

}
