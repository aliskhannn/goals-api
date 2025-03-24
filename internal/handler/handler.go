package handler

import (
	"encoding/json"
	"github.com/aliskhannn/goals-api/internal/model"
	"github.com/aliskhannn/goals-api/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type Handler struct {
	goalService *service.GoalService
	userService *service.UserService
}

func NewHandler(goalService *service.GoalService, userService *service.UserService) *Handler {
	return &Handler{
		goalService: goalService,
		userService: userService,
	}
}

var validate = validator.New()

func (h *Handler) HandleCreateGoal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only method POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userId").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var goal model.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err = validate.Struct(goal); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.goalService.CreateGoal(r.Context(), &goal, userID)
	if err != nil {
		http.Error(w, "Error creating goal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(goal)
}

func (h *Handler) HandleGetAllGoals(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only method GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	goals, err := h.goalService.GetAllGoals(r.Context())
	if err != nil {
		http.Error(w, "Error getting goals", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goals)
}

func (h *Handler) HandleGetGoalByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only method GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	goal, err := h.goalService.GetGoalById(r.Context(), id)
	if err != nil {
		http.Error(w, "Error getting goal by id", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(goal)
}

func (h *Handler) HandleUpdateGoal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only method PUT is allowed", http.StatusMethodNotAllowed)
		return
	}

	var goal model.Goal
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err = validate.Struct(goal); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	err = h.goalService.UpdateGoal(r.Context(), &goal, id)
	if err != nil {
		http.Error(w, "Error updating goal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Goal updated successfully"})
}

func (h *Handler) HandleDeleteGoal(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only method DELETE is allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid goal ID", http.StatusBadRequest)
		return
	}

	err = h.goalService.DeleteGoal(r.Context(), id)
	if err != nil {
		http.Error(w, "Error deleting post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Goal deleted successfully"})
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only method POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err = validate.Struct(user); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.Register(r.Context(), user)
	if err != nil {
		http.Error(w, "Registration error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only method POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.userService.Login(r.Context(), user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error logging in "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
