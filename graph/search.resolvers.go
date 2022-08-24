package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	panic(fmt.Errorf("not implemented"))
}

// Users is the resolver for the users field.
func (r *searchResolver) Users(ctx context.Context, obj *model.Search) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Search returns generated.SearchResolver implementation.
func (r *Resolver) Search() generated.SearchResolver { return &searchResolver{r} }

type searchResolver struct{ *Resolver }
