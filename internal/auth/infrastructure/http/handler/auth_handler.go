package handler

import (
	"encoding/json"
	"fmt"
	"go-task-easy-list/internal/auth/application/service"
	"go-task-easy-list/internal/shared/http"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *service.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	User interface{} `json:"user"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// Register - POST /api/auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		infrastructure.ErrorResponse(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := h.validator.Struct(req); err != nil {
		infrastructure.ErrorResponse(w, http.StatusBadRequest, formatValidationError(err))
		return
	}

	user, err := h.authService.Register(req.Email, req.Password, req.Name)
	if err != nil {
		status := http.StatusBadRequest
		if err == service.ErrEmailExists {
			status = http.StatusConflict
		}
		infrastructure.ErrorResponse(w, status, err.Error())
		return
	}

	infrastructure.SuccessResponse(w, http.StatusCreated, user)
}

// Login - POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		infrastructure.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		infrastructure.ErrorResponse(w, http.StatusUnauthorized, "Credenciales inválidas")
		return
	}

	response := AuthResponse{
		User:        nil,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}

	infrastructure.SuccessResponse(w, http.StatusOK, response)
}

func formatValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := e.Field()
			switch e.Tag() {
			case "required":
				return fmt.Sprintf("El campo '%s' es requerido", field)
			case "email":
				return "Formato de email inválido"
			case "min":
				return fmt.Sprintf("El campo '%s' debe tener al menos %s caracteres", field, e.Param())
			}
		}
	}
	return "Datos inválidos"
}