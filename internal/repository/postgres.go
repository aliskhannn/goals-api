package repository

import (
	"context"
	"fmt"
	"github.com/aliskhanx/goals-api/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) Create(ctx context.Context, goal *model.Goal) error {
	query := `INSERT INTO goals (title, description, completed) VALUES ($1, $2, $3)`

	_, err := r.pool.Exec(ctx, query, goal.Title, goal.Description, goal.Completed)
	if err != nil {
		return fmt.Errorf("error creating goal: %w", err)
	}

	return nil
}
