package config

import (
	"go-task-easy-list/internal/auth/application/service"
	"go-task-easy-list/internal/auth/infrastructure/http/handler"
	gormRepo "go-task-easy-list/internal/auth/infrastructure/persistence/gorm"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type AuthModule struct {
	Handler *handler.AuthHandler
}

func NewAuthModule(db *gorm.DB, jwtSecre string) *AuthModule {
	// Repositories
	userRepo := gormRepo.NewUserRepository(db)
	sessionRepo := gormRepo.NewSessionRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, sessionRepo, jwtSecre)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)

	return &AuthModule{
		Handler: authHandler,
	}
}

// RegisterRoutes registra las rutas del módulo auth
func (m *AuthModule) RegisterRoutes(r chi.Router) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", m.Handler.Register)
		r.Post("/login", m.Handler.Login)
		r.Post("/logout", m.Handler.Logout)
	})
}