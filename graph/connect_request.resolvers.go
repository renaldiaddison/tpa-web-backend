package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// FromUser is the resolver for the fromUser field.
func (r *connectRequestResolver) FromUser(ctx context.Context, obj *model.ConnectRequest) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.FromUserID).Error
}

// ToUser is the resolver for the toUser field.
func (r *connectRequestResolver) ToUser(ctx context.Context, obj *model.ConnectRequest) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.ToUserID).Error
}

// AddConnectRequest is the resolver for the addConnectRequest field.
func (r *mutationResolver) AddConnectRequest(ctx context.Context, fromUserID string, toUserID string, message string) (*model.ConnectRequest, error) {
	modelConnectRequest := &model.ConnectRequest{
		ID:         uuid.NewString(),
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Message:    message,
	}

	return modelConnectRequest, r.DB.Create(modelConnectRequest).Error
}

// DeleteConnectRequest is the resolver for the deleteConnectRequest field.
func (r *mutationResolver) DeleteConnectRequest(ctx context.Context, fromUserID string, toUserID string) (*model.ConnectRequest, error) {
	modelConnectRequest := new(model.ConnectRequest)

	if err := r.DB.Find(&modelConnectRequest, "from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).Error; err != nil {
		return nil, err
	}

	return modelConnectRequest, r.DB.Delete(modelConnectRequest).Error
}

// ConnectRequest returns generated.ConnectRequestResolver implementation.
func (r *Resolver) ConnectRequest() generated.ConnectRequestResolver {
	return &connectRequestResolver{r}
}

type connectRequestResolver struct{ *Resolver }
