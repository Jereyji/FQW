package repository

type Authorization interface {

}

type TodoList interface {

}

type TodoItem interface {

}

type Reposizitory struct {
	Authorization
	TodoList
	TodoItem
}

func NewReposizitory() *Reposizitory {
	return &Reposizitory{}
}