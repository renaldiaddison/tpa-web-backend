package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/service"
)

// AddNotification is the resolver for the addNotification field.
func (r *mutationResolver) AddNotification(ctx context.Context, toUserID string, fromUserID string, message string) (*model.Notification, error) {
	return service.AddNotification(r.DB, ctx, toUserID, fromUserID, message)
}

// FromUser is the resolver for the fromUser field.
func (r *notificationResolver) FromUser(ctx context.Context, obj *model.Notification) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.FromUserID).Error
}

// ToUser is the resolver for the toUser field.
func (r *notificationResolver) ToUser(ctx context.Context, obj *model.Notification) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.ToUserID).Error
}

// UserNotification is the resolver for the userNotification field.
func (r *queryResolver) UserNotification(ctx context.Context, toUserID string) ([]*model.Notification, error) {
	var modelNotifications []*model.Notification

	if err := r.DB.Find(&modelNotifications, "to_user_id = ?", toUserID).Error; err != nil {
		return nil, err
	}

	return modelNotifications, nil
}

// Notification returns generated.NotificationResolver implementation.
func (r *Resolver) Notification() generated.NotificationResolver { return &notificationResolver{r} }

type notificationResolver struct{ *Resolver }
