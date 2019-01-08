package router

import (
	"time"
	"virgo/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// NewRouter ...
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	// Todo handler
	todo := handler.NewTodoHandler()
	r.Get("/todo", todo.Index())
	r.Post("/todo", todo.Create())
	r.Get("/todo/{id:[1-9]+}", todo.Show())
	r.Put("/todo/{id:[1-9]+}", todo.Update())
	r.Delete("/todo/{id:[1-9]+}", todo.Destroy())

	return r
}
