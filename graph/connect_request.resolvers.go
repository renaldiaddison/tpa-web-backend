package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// FromUser is the resolver for the fromUser field.
func (r *connectRequestResolver) FromUser(ctx context.Context, obj *model.ConnectRequest) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// ToUser is the resolver for the toUser field.
func (r *connectRequestResolver) ToUser(ctx context.Context, obj *model.ConnectRequest) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddConnectRequest is the resolver for the addConnectRequest field.
func (r *mutationResolver) AddConnectRequest(ctx context.Context, fromUserID string, toUserID string, message string) (*model.ConnectRequest, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteConnectRequest is the resolver for the deleteConnectRequest field.
func (r *mutationResolver) DeleteConnectRequest(ctx context.Context, fromUserID string, toUserID string) (*model.ConnectRequest, error) {
	panic(fmt.Errorf("not implemented"))
}

// ConnectRequest returns generated.ConnectRequestResolver implementation.
func (r *Resolver) ConnectRequest() generated.ConnectRequestResolver {
	return &connectRequestResolver{r}
}

type connectRequestResolver struct{ *Resolver }
