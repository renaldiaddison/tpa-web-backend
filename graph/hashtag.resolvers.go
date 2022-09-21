package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// AddHashtag is the resolver for the addHashtag field.
func (r *mutationResolver) AddHashtag(ctx context.Context, hashtag string) (*model.Hashtag, error) {
	modelHashtag := new(model.Hashtag)
	modelInputHashtag := &model.Hashtag{
		ID:      uuid.NewString(),
		Hashtag: hashtag,
	}

	if err := r.DB.Find(modelHashtag, "hashtag = ?", hashtag).Error; err != nil {
		return nil, err
	}

	if modelHashtag.ID != "" {
		return nil, gqlerror.Errorf("Hashtag already exists")
	}

	return modelInputHashtag, r.DB.Create(modelInputHashtag).Error
}

// Hashtags is the resolver for the Hashtags field.
func (r *queryResolver) Hashtags(ctx context.Context) ([]*model.Hashtag, error) {
	var modelHashtags []*model.Hashtag

	return modelHashtags, r.DB.Find(&modelHashtags).Error
}
