package config

import (
	gormModels "go-task-easy-list/internal/auth/infrastructure/persistence/gorm"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	db.Exec("PRAGMA foreign_keys = ON")

	// AutoMigrate: Crear tablas automáticamente
	if err := db.AutoMigrate(
		&gormModels.UserModel{},
		&gormModels.SessionModel{},
	); err != nil {
		return nil, err
	}

	return db, nil
}