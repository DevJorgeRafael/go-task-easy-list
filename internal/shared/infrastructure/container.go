package infrastructure

import (
	authConfig "go-task-easy-list/internal/auth/infrastructure/config"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Container struct {
	AuthModule *authConfig.AuthModule
	// TaskModule *taskConfig.TaskModule
}

func NewContainer(db *gorm.DB, jwtSecret string) *Container {
	return &Container {
		AuthModule: authConfig.NewAuthModule(db, jwtSecret),
		// TaskModule: taskConfig.NewTaskModule(db),
	}
}

// RegisterRoutes registra las rutas de todos los módulos
func (c *Container) RegisterRoutes(r chi.Router) {
	c.AuthModule.RegisterRoutes(r)
	// c.TaskModule.RegisterRoutes(r)
}