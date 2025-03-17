package service

import (
	"context"
	"github.com/aliskhanx/goals-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, goal *model.Goal) error
}

type GoalService struct {
	repo Repository
}

func NewGoalService(repo Repository) *GoalService {
	return &GoalService{
		repo: repo,
	}
}

func (s *GoalService) CreateGoal(ctx context.Context, goal *model.Goal) error {
	return s.repo.Create(ctx, goal)
}
