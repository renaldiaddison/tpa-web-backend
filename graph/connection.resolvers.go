package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// User1 is the resolver for the user1 field.
func (r *connectionResolver) User1(ctx context.Context, obj *model.Connection) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.User1ID).Error
}

// User2 is the resolver for the user2 field.
func (r *connectionResolver) User2(ctx context.Context, obj *model.Connection) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.User2ID).Error
}

// AddConnection is the resolver for the addConnection field.
func (r *mutationResolver) AddConnection(ctx context.Context, user1id string, user2id string) (*model.Connection, error) {
	modelConnection := &model.Connection{
		ID:      uuid.NewString(),
		User1ID: user1id,
		User2ID: user2id,
	}

	return modelConnection, r.DB.Create(modelConnection).Error
}

// Connection returns generated.ConnectionResolver implementation.
func (r *Resolver) Connection() generated.ConnectionResolver { return &connectionResolver{r} }

type connectionResolver struct{ *Resolver }
