package service

import "github.com/Jereyji/FQW.git/pkg/repository"

type Authorization interface {

}

type TodoList interface {

}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Reposizitory) *Service {
	return &Service{}
}