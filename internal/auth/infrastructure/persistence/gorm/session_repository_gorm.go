package gorm

import (
	"go-task-easy-list/internal/auth/domain/model"
	"go-task-easy-list/internal/auth/domain/repository"
	"time"

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

func (r *SessionRepositoryGorm) CountByUserID(userID string) (int64, error) {
	var count int64
	err := r.db.Model(&SessionModel{}).Where("user_id = ? AND expires_at > ?", userID, time.Now()).Count(&count).Error
	return count, err
}

func (r *SessionRepositoryGorm) DeleteOldestByUserID(userID string) error {
	var oldestSession SessionModel
	if err := r.db.Where("user_id = ?", userID).Order("created_at ASC").First(&oldestSession).Error; err != nil{
		return err
	}
	return r.db.Delete(&oldestSession).Error
}

func (r *SessionRepositoryGorm) DeleteExpiredByUserID(userID string) error {
	return r.db.Where("user_id = ? AND expires_at < ?", userID, time.Now()).Delete(&SessionModel{}).Error
}