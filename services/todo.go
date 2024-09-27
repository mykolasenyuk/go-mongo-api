package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (t *Todo) GetAllTodos() ([]Todo, error) {
	collection := returnCollectionPointer("todos")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var todos []Todo
	for cursor.Next(context.TODO()) {
		var todo Todo
		err := cursor.Decode(&todo)
		if err != nil {
			log.Println("Error:", err)
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *Todo) GetTodoById(id string) (Todo, error) {
	collection := returnCollectionPointer("todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error: Invalid ID format")
		return Todo{}, err
	}

	var todo Todo
	err = collection.FindOne(context.TODO(), bson.M{"_id": mongoID}).Decode(&todo)
	if err != nil {
		log.Println("Error:", err)
		return Todo{}, err
	}

	return todo, nil
}

func (t *Todo) UpdateById(id string, entry Todo) (Todo, error) {
	collection := returnCollectionPointer("todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error: Invalid ID format")
		return Todo{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"task":       entry.Task,
			"completed":  entry.Completed,
			"updated_at": time.Now(),
		},
	}

	var updatedTodo Todo
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": mongoID}, update, opts).Decode(&updatedTodo)
	if err != nil {
		log.Println("Error:", err)
		return Todo{}, err
	}

	return updatedTodo, nil
}

func (t *Todo) DeleteById(id string) error {
	collection := returnCollectionPointer("todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Error: Invalid ID format")
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": mongoID})
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}
