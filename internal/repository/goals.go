package repository

import (
	"context"
	"fmt"
	"github.com/aliskhannn/goals-api/internal/model"
)

func (r *Repository) CreateGoal(ctx context.Context, goal *model.Goal, userID int) error {
	query := `INSERT INTO goals (title, description, completed, user_id) VALUES ($1, $2, $3, $4)`

	_, err := r.pool.Exec(ctx, query, goal.Title, goal.Description, goal.Completed, userID)
	if err != nil {
		return fmt.Errorf("error creating goal: %w", err)
	}

	return nil
}

func (r *Repository) GetAllGoals(ctx context.Context) ([]*model.Goal, error) {
	query := `SELECT id, title, description, completed FROM goals`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting goals: %w", err)
	}
	defer rows.Close()

	var goals []*model.Goal

	for rows.Next() {
		var goal model.Goal
		if err := rows.Scan(&goal.ID, &goal.Title, &goal.Description, &goal.Completed); err != nil {
			return nil, fmt.Errorf("error scanning goal: %w", err)
		}

		goals = append(goals, &goal)
	}

	return goals, nil
}

func (r *Repository) GetGoalById(ctx context.Context, id int) (*model.Goal, error) {
	query := `SELECT id, title, description, completed FROM goals WHERE id = $1`

	var goal model.Goal
	err := r.pool.QueryRow(ctx, query, id).Scan(&goal.ID, &goal.Title, &goal.Description, &goal.Completed)
	if err != nil {
		return nil, fmt.Errorf("error getting goal by id: %w", err)
	}

	return &goal, nil
}

func (r *Repository) UpdateGoal(ctx context.Context, goal *model.Goal, id int) error {
	query := `UPDATE goals SET title = $1, description = $2, completed = $3 WHERE id = $4`

	result, err := r.pool.Exec(
		ctx,
		query,
		goal.Title,
		goal.Description,
		goal.Completed,
		id,
	)
	if err != nil {
		return fmt.Errorf("error updating goal: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("goal not found")
	}

	return nil
}

func (r *Repository) DeleteGoal(ctx context.Context, id int) error {
	query := `DELETE FROM goals WHERE id = $1`

	result, err := r.pool.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting goal: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("goal not found")
	}

	return nil
}
