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

func getAllTaks(c *gin.Context, db *mgo.Database) (model_task.Tasks, error) {
	tasks := model_task.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{}).All(&tasks)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func getTasksByStatus(c *gin.Context, status string, db *mgo.Database) (model_task.Tasks, error) {

	tasks := model_task.Tasks{}

	err := db.C(TaskCollection).Find(bson.M{"status": &status}).All(&tasks)

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func errorResponse(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something happened",
			"error":   err,
		})
		return
	}
}

//GetTasks will return all tasks
func GetTasks(c *gin.Context) {
	db := *MongoConfig()

	status := strings.ToUpper(c.Query("status"))

	tasks := model_task.Tasks{}

	var err error

	if status != "" {
		tasks, err = getTasksByStatus(c, status, &db)
	} else {
		tasks, err = getAllTaks(c, &db)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching tasks",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": &tasks,
	})
}

//GetTask will return an specific task
func GetTask(c *gin.Context) {
	db := *MongoConfig()

	taskID := c.Query("taskId")

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query string \"taskId\" is missing on the url",
		})
		return
	}

	task := model_task.Task{}

	err := db.C(TaskCollection).FindId(bson.ObjectIdHex(taskID)).One(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something happened",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, &task)
}

//UpdateTaskStatus will update a task
func UpdateTaskStatus(c *gin.Context) {
	db := *MongoConfig()

	taskID := c.Query("taskId")

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query string \"taskId\" is missing on the url",
		})
		return
	}

	task := model_task.Task{}

	task.ID = bson.ObjectIdHex(taskID)

	err := c.BindJSON(&task)

	newData := bson.M{
		"$set": bson.M{
			"status":     task.Status,
			"updated_at": time.Now(),
		},
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body format",
		})
		return
	}

	err = db.C(TaskCollection).UpdateId(task.ID, newData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something happened",
			"error":   err,
		})
		return
	}

	err = db.C(TaskCollection).FindId(task.ID).One(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something happened",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfuly",
		"task":    &task,
	})
}

//CreateTask will create a new task
func CreateTask(c *gin.Context) {
	db := *MongoConfig()

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

//DeleteTask will delete an specifica task
func DeleteTask(c *gin.Context) {
	db := *MongoConfig()

	taskID := c.Query("taskId")

	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Query string \"taskId\" is missing on the url",
		})
	}

	err := db.C(TaskCollection).RemoveId(bson.ObjectIdHex(taskID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something happened",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfuly",
	})

}
