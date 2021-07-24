package redis

import "github.com/kons16/graphql-todo/models"

type Todo interface {
	CreateTask(text string) error
	GetTask(id string) (*models.Todo, error)
	GetAllTasks() ([]*models.Todo, error)
}
