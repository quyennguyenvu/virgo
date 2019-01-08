package handler

import (
	"net/http"
	"virgo/service"

	"github.com/go-chi/chi"
)

// TodoHandler ..
type TodoHandler interface {
	Index() http.HandlerFunc
	Create() http.HandlerFunc
	Show() http.HandlerFunc
	Update() http.HandlerFunc
	Destroy() http.HandlerFunc
}

// todoHandlerImpl ..
type todoHandlerImpl struct {
	todoSC service.TodoService
}

// NewTodoHandler ..
func NewTodoHandler() TodoHandler {
	return &todoHandlerImpl{
		todoSC: service.NewTodoService(),
	}
}

// Index ..
func (h *todoHandlerImpl) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strStatus := r.URL.Query().Get("status")
		response := h.todoSC.List(strStatus)

		respond(w, response)
	}
}

// Create ..
func (h *todoHandlerImpl) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := h.todoSC.Create(r.Body)
		respond(w, response)
	}
}

// Show ..
func (h *todoHandlerImpl) Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := chi.URLParam(r, "id")
		response := h.todoSC.ByID(strID)

		respond(w, response)
	}
}

// Update ..
func (h *todoHandlerImpl) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := chi.URLParam(r, "id")
		response := h.todoSC.Update(strID, r.Body)

		respond(w, response)
	}
}

// Destroy ..
func (h *todoHandlerImpl) Destroy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strID := chi.URLParam(r, "id")
		response := h.todoSC.Destroy(strID)

		respond(w, response)
	}
}
