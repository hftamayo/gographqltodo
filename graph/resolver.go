package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"

	"github.com/hftamayo/gographqltodo/api/v1/models"
	"gorm.io/gorm"
)

type Resolver struct {
	DB *gorm.DB
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, title string, body string) (*models.Todo, error) {
	newTodo := models.Todo{
		Title: title,
		Body:  body,
		Done:  new(bool), // default to false
	}
	result := r.DB.Create(&newTodo) // Assuming r.DB is your GORM database connection
	if result.Error != nil {
		return nil, result.Error
	}
	return &newTodo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, title string, body string) (*models.Todo, error) {
	var todo models.Todo
	result := r.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	todo.Title = title
	todo.Body = body
	result = r.DB.Save(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (r *mutationResolver) MarkTodoAsDone(ctx context.Context, id int, done bool) (*models.Todo, error) {
	var todo models.Todo
	result := r.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	todo.Done = &done
	result = r.DB.Save(&todo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	var todos []*models.Todo
	result := r.DB.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
func (r *queryResolver) Todo(ctx context.Context, id int) (*models.Todo, error) {
	var todo models.Todo
	result := r.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int) (bool, error) {
	result := r.DB.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
