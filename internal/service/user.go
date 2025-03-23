package service

import (
	"context"
	"fmt"
	"github.com/aliskhannn/goals-api/internal/hash"
	"github.com/aliskhannn/goals-api/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, user *model.User) error {
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to find user by username: %w", err)
	}

	if isValidPassword := hash.VerifyPassword(user.Password, password); !isValidPassword {
		return "", fmt.Errorf("invalid password")
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	return token, nil
}
