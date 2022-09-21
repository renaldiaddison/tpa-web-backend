package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// User1 is the resolver for the user1 field.
func (r *connectionResolver) User1(ctx context.Context, obj *model.Connection) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// User2 is the resolver for the user2 field.
func (r *connectionResolver) User2(ctx context.Context, obj *model.Connection) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddConnection is the resolver for the addConnection field.
func (r *mutationResolver) AddConnection(ctx context.Context, user1id string, user2id string) (*model.Connection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Connection returns generated.ConnectionResolver implementation.
func (r *Resolver) Connection() generated.ConnectionResolver { return &connectionResolver{r} }

type connectionResolver struct{ *Resolver }
