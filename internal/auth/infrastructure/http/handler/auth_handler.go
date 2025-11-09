package handler

import (
	"go-task-easy-list/internal/auth/application/service"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register - POST /api/auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// 1. Parsear JSON del body
	// 2. Validar campos (usar go-playground/validator)
	// 3. Llamar authService.Register()
	// 4. Retornar respuesta JSON
}

// Login - POST /api/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Similar a Register
}