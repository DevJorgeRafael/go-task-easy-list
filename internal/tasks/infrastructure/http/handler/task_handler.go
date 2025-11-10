package handler

import (
	"encoding/json"
	sharedhttp "go-task-easy-list/internal/shared/http"
	sharedContext "go-task-easy-list/internal/shared/context"
	format "go-task-easy-list/internal/shared/http/utils"
	"go-task-easy-list/internal/tasks/application/service"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type TaskHandler struct {
	taskService *service.TaskService
	validator   *validator.Validate
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		validator:   validator.New(),
	}
}

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	StatusId    string `json:"statusId"`
	PriorityId  string `json:"priorityId"`
	StartsAt    string `json:"startsAt"`
	DueDate     string `json:"dueDate"`
}

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StatusId    string `json:"statusId"`
	PriorityId  string `json:"priorityId"`
	StartsAt    string `json:"startsAt"`
	DueDate     string `json:"dueDate"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// CreateTask - POST /api/tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(sharedContext.UserIdKey).(string)
	if !ok {
		sharedhttp.ErrorResponse(w, http.StatusUnauthorized, "Usuario no autenticado")
		return
	}

	var req TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, format.FormatValidationError(err))
		return
	}

	statusId, err := strconv.Atoi(req.StatusId)
	if err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, "StatusId inválido")
		return
	}

	priorityId, err := strconv.Atoi(req.PriorityId)
	if err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, "PriorityId inválido")
		return
	}

	startsAt, err := time.Parse(time.RFC3339, req.StartsAt)
	if err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, "StartsAt inválido")
		return
	}

	dueDate, err := time.Parse(time.RFC3339, req.DueDate)
	if err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, "DueDate inválido")
		return
	}

	task, err := h.taskService.CreateTask(req.Title, req.Description, statusId, priorityId, startsAt, dueDate, userID)
	if err != nil {
		sharedhttp.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	resp := TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		StatusId:    strconv.Itoa(task.StatusID.ID),
		PriorityId:  strconv.Itoa(task.PriorityID.ID),
		StartsAt:    task.StartsAt.Format(time.RFC3339),
		DueDate:     task.DueDate.Format(time.RFC3339),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}

	sharedhttp.SuccessResponse(w, http.StatusCreated, resp)
}