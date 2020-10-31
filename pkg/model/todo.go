package model

//TODO struct
type TODO struct {
	ID          string `bson:"id"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Status      Status `bson:"status"`
}

//TODOS list
type TODOS []TODO
