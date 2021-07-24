package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"github.com/kons16/graphql-todo/repository/redis"
	"log"

	"github.com/kons16/graphql-todo/graph/generated"
	"github.com/kons16/graphql-todo/graph/model"
	"github.com/kons16/graphql-todo/models"
)

type Resolver struct {
	TaskRepo *redis.TaskRepository
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	log.Printf("queryResolver.Todos")
	tasks, err := r.TaskRepo.GetAllTasks()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return tasks, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
