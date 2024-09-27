package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"mongo-api/services"
	"net/http"
)

var todo services.Todo

func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func createTodo(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	err = todo.InsertTodo(todo)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 304,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := Response{
		Msg:  "Success! Created Todo",
		Code: 201,
		Data: todo,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todo.GetAllTodos()
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 500,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	res := Response{
		Msg:  "Success! Retrieved Todos",
		Code: 200,
		Data: todos,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo, err := todo.GetTodoById(id)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 500,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := Response{
		Msg:  "Success! Retrieved Todo",
		Code: 200,
		Data: todo,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func updateTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var updatedTodo services.Todo
	err := json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		log.Fatal(err)
	}

	todo, err := todo.UpdateById(id, updatedTodo)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 500,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := Response{
		Msg:  "Success! Updated Todo",
		Code: 200,
		Data: todo,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func deleteTodoById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := todo.DeleteById(id)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 500,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := Response{
		Msg:  "Success! Deleted Todo",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}
