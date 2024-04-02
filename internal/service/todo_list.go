package service

import (
	todo "github.com/Jereyji/FQW.git"
	"github.com/Jereyji/FQW.git/internal/repository"
)

type todoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *todoListService {
	return &todoListService{repo : repo}
}

func (s *todoListService) Create(userId int, list todo.TodoList) (int,error) {
	return s.repo.Create(userId, list)
}

func (s *todoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *todoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}