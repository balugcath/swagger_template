package middleware

import (
	"context"
	"swagger_template/src/main_service/swagger/models_server"
)

type storage interface {
	Add(context.Context, models_server.Item) (models_server.Item, error)
	Destroy(context.Context, string) error
	Find(context.Context, models_server.Item) ([]models_server.Item, error)
	Update(context.Context, string, models_server.Item) (models_server.Item, error)
}

type client interface {
	Delete(context.Context, string) error
}

// Middleware ...
type Middleware struct {
	storage storage
	client  client
}

// NewMiddleware ...
func NewMiddleware(storage storage, client client) *Middleware {
	s := Middleware{
		storage: storage,
	}
	return &s
}

// Add ...
func (s *Middleware) Add(ctx context.Context, item models_server.Item) (models_server.Item, error) {
	return s.storage.Add(ctx, item)
}

// Destroy ...
func (s *Middleware) Destroy(ctx context.Context, id string) error {
	s.client.Delete(ctx, id)
	return s.storage.Destroy(ctx, id)
}

// Find ...
func (s *Middleware) Find(ctx context.Context, item models_server.Item) ([]models_server.Item, error) {
	return s.storage.Find(ctx, item)
}

// Update ...
func (s *Middleware) Update(ctx context.Context, id string, item models_server.Item) (models_server.Item, error) {
	return s.storage.Update(ctx, id, item)
}
