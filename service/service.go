package service

import (
	"test_project/model"
	"test_project/repository"
)

type TodoList interface {
	GetAll() ([]model.Todo, error)
	GetListById(id string) (model.Todo, error)
	UpdateList(listId int, input model.UpdateListInput) error
	DeleteList(id string) error
	CreateList(input model.Todo) (model.Todo, error)
}

type Service struct {
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(repos.TodoList),
	}
}

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) GetAll() ([]model.Todo, error) {
	return s.repo.GetAll()
}
func (s *TodoListService) GetListById(id string) (model.Todo, error) {
	return s.repo.GetListById(id)
}
func (s *TodoListService) UpdateList(listId int, input model.UpdateListInput) error {
	return s.repo.UpdateList(listId, input)
}



func (s *TodoListService) DeleteList(id string) error {
	return s.repo.DeleteList(id)
}

func (s *TodoListService) CreateList(input model.Todo) (model.Todo, error) {
	return s.repo.CreateList(input)
}
