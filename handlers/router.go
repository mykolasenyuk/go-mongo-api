package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Response struct {
	Msg  string
	Code int
	Data interface{} `json:"data"`
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Route("/api", func(router chi.Router) {

		// version 1
		router.Route("/v1", func(router chi.Router) {

			router.Get("/healthcheck", healthCheck)
			router.Post("/todos", createTodo)
			router.Get("/todos", getAllTodos)
			router.Get("/todos/{id}", getTodoById)
			router.Put("/todos/{id}", updateTodoById)
			router.Delete("/todos/{id}", deleteTodoById)

		})

		// version 2 - add it if you want
		// router.Route("/v2", func(router chi.Router) {
		// })

	})

	return router

}
