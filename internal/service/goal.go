package service

import (
	"context"
	"github.com/aliskhanx/goals-api/internal/model"
)

type Repository interface {
	Create(ctx context.Context, goal *model.Goal) error
	GetAll(ctx context.Context) ([]*model.Goal, error)
	GetById(ctx context.Context, id int) (*model.Goal, error)
	Update(ctx context.Context, goal *model.Goal, id int) error
	Delete(ctx context.Context, id int) error
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

func (s *GoalService) GetAllGoals(ctx context.Context) ([]*model.Goal, error) {
	return s.repo.GetAll(ctx)
}

func (s *GoalService) GetGoalById(ctx context.Context, id int) (*model.Goal, error) {
	return s.repo.GetById(ctx, id)
}

func (s *GoalService) UpdateGoal(ctx context.Context, goal *model.Goal, id int) error {
	return s.repo.Update(ctx, goal, id)
}

func (s *GoalService) DeleteGoal(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
