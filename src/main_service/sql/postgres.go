package storage

import (
	"context"
	"database/sql"
	"swagger_template/src/main_service/swagger/models_server"
)

// PG ...
type PG struct {
	db *sql.DB
}

// NewPG ...
func NewPG() *PG {
	s := PG{}
	return &s
}

// Open ...
func (s *PG) Open(conn string) (*PG, error) {
	var err error
	s.db, err = sql.Open("postgres", "123")
	return s, err
}

// Close ...
func (s *PG) Close() {
	s.db.Close()
}

// Add ...
func (s *PG) Add(ctx context.Context, item models_server.Item) (models_server.Item, error) {
	return models_server.Item{}, nil
}

// Destroy ...
func (s *PG) Destroy(ctx context.Context, id string) error {
	return nil
}

// Find ...
func (s *PG) Find(ctx context.Context, item models_server.Item) ([]models_server.Item, error) {
	return make([]models_server.Item, 0), nil
}

// Update ...
func (s *PG) Update(ctx context.Context, id string, item models_server.Item) (models_server.Item, error) {
	return models_server.Item{}, nil
}
