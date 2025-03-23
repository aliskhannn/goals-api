package service

import (
	"context"
	"github.com/aliskhannn/goals-api/internal/model"
)

type GoalRepository interface {
	CreateGoal(ctx context.Context, goal *model.Goal) error
	GetAllGoals(ctx context.Context) ([]*model.Goal, error)
	GetGoalById(ctx context.Context, id int) (*model.Goal, error)
	UpdateGoal(ctx context.Context, goal *model.Goal, id int) error
	DeleteGoal(ctx context.Context, id int) error
}

type GoalService struct {
	repo GoalRepository
}

func NewGoalService(repo GoalRepository) *GoalService {
	return &GoalService{
		repo: repo,
	}
}

func (s *GoalService) CreateGoal(ctx context.Context, goal *model.Goal) error {
	return s.repo.CreateGoal(ctx, goal)
}

func (s *GoalService) GetAllGoals(ctx context.Context) ([]*model.Goal, error) {
	return s.repo.GetAllGoals(ctx)
}

func (s *GoalService) GetGoalById(ctx context.Context, id int) (*model.Goal, error) {
	return s.repo.GetGoalById(ctx, id)
}

func (s *GoalService) UpdateGoal(ctx context.Context, goal *model.Goal, id int) error {
	return s.repo.UpdateGoal(ctx, goal, id)
}

func (s *GoalService) DeleteGoal(ctx context.Context, id int) error {
	return s.repo.DeleteGoal(ctx, id)
}
