package apihandler

import (
	"context"

	resp "github.com/go-openapi/runtime/middleware"
	"swagger_template/src/main_service/swagger/models_server"
	"swagger_template/src/main_service/swagger/restapi/operations/todos"
)

type middleware interface {
	Add(context.Context, models_server.Item) (models_server.Item, error)
	Destroy(context.Context, string) error
	Find(context.Context, models_server.Item) ([]models_server.Item, error)
	Update(context.Context, string, models_server.Item) (models_server.Item, error)
}

// Todo ...
type Todo struct {
	middleware middleware
}

// NewTodo ...
func NewTodo(middleware middleware) *Todo {
	s := Todo{middleware: middleware}
	return &s
}

// Destroy ...
func (s *Todo) Destroy(ctx context.Context, params todos.DestroyParams) resp.Responder {
	if err := s.middleware.Destroy(ctx, params.ID); err != nil {
		msg := err.Error()
		return todos.NewDestroyDefault(-1).WithPayload(&models_server.Error{Code: 500, Message: &msg})
	}
	return todos.NewDestroyNoContent()
}

// Find ...
func (s *Todo) Find(ctx context.Context, params todos.FindParams) resp.Responder {
	return resp.Error(500, "operation todos.FindTodos has not yet been implemented")
}

// Update ...
func (s *Todo) Update(ctx context.Context, params todos.UpdateParams) resp.Responder {
	return resp.NotImplemented("operation todos.UpdateOne has not yet been implemented")
}

// Add ...
func (s *Todo) Add(ctx context.Context, params todos.AddParams) resp.Responder {
	reply := models_server.Item{Description: nil, ID: 909, Completed: true}
	return todos.NewAddCreated().WithPayload(&reply)
}
