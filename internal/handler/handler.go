package handler

import (
	"context"
	"encoding/json"
	"github.com/aliskhanx/goals-api/internal/model"
	"github.com/aliskhanx/goals-api/internal/service"
	"net/http"
)

type Handler struct {
	goalService *service.GoalService
}

func NewHandler(goalService *service.GoalService) *Handler {
	return &Handler{
		goalService: goalService,
	}
}

func (h *Handler) HandleCreateGoal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var goal model.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.goalService.CreateGoal(ctx, &goal)
	if err != nil {
		http.Error(w, "Error creating goal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(goal)
}
