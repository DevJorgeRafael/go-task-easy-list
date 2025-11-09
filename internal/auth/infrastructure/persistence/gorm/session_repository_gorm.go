package gorm

import (
	"go-task-easy-list/internal/auth/domain/model"
	"go-task-easy-list/internal/auth/domain/repository"

	"gorm.io/gorm"
)

type SessionRepositoryGorm struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &SessionRepositoryGorm{db: db}
}

func (r *SessionRepositoryGorm) Create(session *model.Session) error {
	sessionModel := &SessionModel{
		ID: session.ID,
		UserID: session.UserID,
		RefreshToken: session.RefreshToken,
		ExpiresAt: session.ExpiresAt,
		CreatedAt: session.CreatedAt,
	}

	return r.db.Create(sessionModel).Error
}

func (r *SessionRepositoryGorm) FindByRefreshToken(token string) (*model.Session, error) {
	sessionModel := &SessionModel{}
	if err := r.db.Where("refresh_token = ?", token).First(sessionModel).Error; err != nil {
		return nil ,err
	}

	session := &model.Session{
		ID: sessionModel.ID,
		UserID: sessionModel.UserID,
		RefreshToken: sessionModel.RefreshToken,
		ExpiresAt: sessionModel.ExpiresAt,
		CreatedAt: sessionModel.CreatedAt,
	}

	return session, nil
}

func (r *SessionRepositoryGorm) DeleteByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&SessionModel{}).Error
}

func (r *SessionRepositoryGorm) DeleteExpired() error {
	return r.db.Where("expires_at < ?", gorm.Expr("NOW()")).Delete(&SessionModel{}).Error
}