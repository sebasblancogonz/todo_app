package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Task struct
type Task struct {
	ID          bson.ObjectId `bson:"_id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Status      string        `bson:"status"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

//Tasks list
type Tasks []Task
