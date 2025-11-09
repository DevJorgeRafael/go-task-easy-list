package repository

import "go-task-easy-list/internal/auth/domain/model"

type SessionRepository interface {
	Create(session *model.Session) error
	FindByRefreshToken(token string) (*model.Session, error)
	DeleteByUserID(userID string) error
	DeleteExpired() error
}