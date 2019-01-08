package storage

import (
	"virgo/helper"

	"github.com/jinzhu/gorm"
)

// Todo ..
type Todo struct {
	gorm.Model
	Name   string
	Task   string
	Status int
}

// TableName ...
func (m Todo) TableName() string {
	return "todo"
}

// TodoStorage ..
type TodoStorage interface {
	ByID(id int) (*Todo, error)
	Create(data *Todo) error
	List(status int) (*[]*Todo, error)
	Update(data *Todo) error
	Destroy(data *Todo) error
}

type todoImpl struct {
	db *gorm.DB
}

// NewTodoStorage ..
func NewTodoStorage() TodoStorage {
	return &todoImpl{db: db}
}

func (s *todoImpl) ByID(id int) (*Todo, error) {
	todo := &Todo{}
	result := s.db.First(&todo, id)
	if result.Error != nil {
		helper.Logging("Todo", "ByID", result.Error.Error())
		return nil, result.Error
	}
	return todo, nil
}

func (s *todoImpl) Create(data *Todo) error {
	result := s.db.Create(data)
	if result.Error != nil {
		helper.Logging("Todo", "Create", result.Error.Error())
		return result.Error
	}
	return nil
}

func (s *todoImpl) List(status int) (*[]*Todo, error) {
	todos := &[]*Todo{}
	result := s.db.Model(&Todo{}).Where("status = ?", status).Find(todos)
	if result.Error != nil {
		helper.Logging("Todo", "List", result.Error.Error())
		return nil, result.Error
	}
	return todos, nil
}

func (s *todoImpl) Update(data *Todo) error {
	result := s.db.Save(data)
	if result.Error != nil {
		helper.Logging("Todo", "Update", result.Error.Error())
		return result.Error
	}
	return nil
}

func (s *todoImpl) Destroy(data *Todo) error {
	result := s.db.Delete(data)
	if result.Error != nil {
		helper.Logging("Todo", "Destroy", result.Error.Error())
		return result.Error
	}
	return nil
}
