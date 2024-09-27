package handlers

import (
	"encoding/json"
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
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}
