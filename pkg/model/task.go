package model

//Task struct
type Task struct {
	ID          string `bson:"id"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Status      Status `bson:"status"`
}

//Tasks list
type Tasks []Task
