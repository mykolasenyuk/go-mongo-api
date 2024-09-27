package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Todo struct {
	Id        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo

	return Todo{}
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("todos_db").Collection(collection)
}

func (t *Todo) InsertTodo(entry Todo) error {
	collection := returnCollectionPointer("todos")

	_, err := collection.InsertOne(context.TODO(), Todo{
		Task:      entry.Task,
		Completed: entry.Completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}
