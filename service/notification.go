package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"gorm.io/gorm"
)

func AddNotification(db *gorm.DB, ctx context.Context, toUserID string, fromUserID string, message string) (*model.Notification, error) {
	modelNotification := &model.Notification{
		ID:         uuid.NewString(),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Message:    message,
	}

	return modelNotification, db.Create(modelNotification).Error
}
