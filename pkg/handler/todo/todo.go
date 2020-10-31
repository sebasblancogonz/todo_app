package todo

import (
	"log"

	"github.com/sebasblancogonz/todo_app/config"
	"gopkg.in/mgo.v2"
)

//TodoCollection static collection
const TodoCollection = "todo"

//MongoConfig returns db
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		log.Println(err)
	}
	return db
}
