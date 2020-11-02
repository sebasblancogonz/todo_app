package task

import (
	"log"
	"net/http"
	"strings"
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
func GetTasks(c *gin.Context) {
	db := *MongoConfig()

	status := strings.ToUpper(c.Query("status"))

	tasks := model_task.Tasks{}

	err := error

	if status != "" {
		println(status)
		err, tasks = getTasksByStatus(c, status, &db)
	} else {
		err, tasks = getAllTaks(c, &db)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "Something went wrong",
		})
	} 

	c.JSON(http.StatusOK, gin.H{
		"tasks": &tasks,
	})

}

func getAllTaks(c *gin.Context, db *mgo.Database) err Error, model_task.Tasks {
	tasks := model_task.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{}).All(&tasks)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching all tasks",
		})
		return tasks
	}

	return tasks
}

func getTasksByStatus(c *gin.Context, status string, db *mgo.Database) err Error, model_task.Tasks {

	tasks := model_task.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{"status": &status}).All(&tasks)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching tasks by status",
		})
		return tasks
	}

	return tasks
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

//UpdateTaskStatus will update a task
func UpdateTaskStatus(c *gin.Context) {

}
