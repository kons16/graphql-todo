package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/kons16/graphql-todo/models"
	"log"
)

type TaskRepository struct {
	rdMap *redis.Client
}

func NewTaskRepository(rdMap *redis.Client) *TaskRepository {
	return &TaskRepository{rdMap: rdMap}
}

func (tr *TaskRepository) CreateTask(text string) error {
	return nil
}

func (tr *TaskRepository) GetAllTasks() ([]*models.Todo, error) {
	ctx := context.Background()

	var cursor uint64
	var keys []string
	keys, cursor, err := tr.rdMap.Scan(ctx, cursor, "keys*", 0).Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var todos []map[string]string
	for _, key := range keys {
		val, err := tr.rdMap.HGetAll(ctx, key).Result()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		todos = append(todos, val)
	}

	return nil, nil
}

func (tr *TaskRepository) GetTask(id string) (*models.Todo, error) {
	return nil, nil
}
