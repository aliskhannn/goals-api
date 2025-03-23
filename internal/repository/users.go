package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/aliskhannn/goals-api/internal/hash"
	"github.com/aliskhannn/goals-api/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) CreateUser(ctx context.Context, user *model.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return fmt.Errorf("username, email and password cannot be empty")
	}

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	user.Password = hashedPassword

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`

	err = r.pool.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("user with this email or username already exists")
		}
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	query := `SELECT id, username, email, password FROM users WHERE username = $1`

	row := r.pool.QueryRow(ctx, query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}

	return &user, nil
}
