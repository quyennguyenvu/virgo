package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"virgo/storage"
)

// TodoService ..
type TodoService interface {
	List(strStatus string) Response
	ByID(strID string) Response
	Create(readClose io.ReadCloser) Response
	Update(strID string, readCloser io.ReadCloser) Response
	Destroy(strID string) Response
}

// todoHandlerImpl ..
type todoServiceImpl struct {
	todoDAO storage.TodoStorage
}

// NewTodoService ..
func NewTodoService() TodoService {
	return &todoServiceImpl{
		todoDAO: storage.NewTodoStorage(),
	}
}

func (sc *todoServiceImpl) List(strStatus string) Response {
	status, _ := strconv.Atoi(strStatus)
	todos, err := sc.todoDAO.List(status)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	res, err := json.Marshal(todos)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: res,
		Code: http.StatusOK,
		Err:  nil,
	}
}

func (sc *todoServiceImpl) ByID(strID string) Response {
	id, _ := strconv.Atoi(strID)
	todo, err := sc.todoDAO.ByID(id)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	res, err := json.Marshal(todo)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: res,
		Code: http.StatusOK,
		Err:  nil,
	}
}

func (sc *todoServiceImpl) Create(readCloser io.ReadCloser) Response {
	var todo storage.Todo
	getBody(&todo, readCloser)

	err := sc.todoDAO.Create(&todo)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: []byte("Successfully create task"),
		Code: http.StatusOK,
		Err:  nil,
	}
}

func (sc *todoServiceImpl) Update(strID string, readCloser io.ReadCloser) Response {
	id, _ := strconv.Atoi(strID)

	todo, err := sc.todoDAO.ByID(id)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	getBody(todo, readCloser)

	err = sc.todoDAO.Update(todo)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	return Response{
		Data: []byte("Successfully update task: " + todo.Name),
		Code: http.StatusOK,
		Err:  nil,
	}
}

func (sc *todoServiceImpl) Destroy(strID string) Response {
	id, _ := strconv.Atoi(strID)

	todo, err := sc.todoDAO.ByID(id)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	err = sc.todoDAO.Destroy(todo)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	return Response{
		Data: []byte("Successfully deleted task: " + todo.Name),
		Code: http.StatusOK,
		Err:  nil,
	}
}
