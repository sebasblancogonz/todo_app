package model

import (
	"time"

	status "github.com/sebasblancogonz/todo_app/pkg/status"
)

//Task struct
type Task struct {
	ID          string        `bson:"id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Status      status.Status `bson:"status"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

//Tasks list
type Tasks []Task
