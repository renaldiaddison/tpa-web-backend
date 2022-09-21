package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/middlewares"
	"github.com/samber/lo"
)

// Search is the resolver for the Search field.
func (r *queryResolver) Search(ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	search := new(model.Search)

	userID := middlewares.GetJwtValueData(ctx).Userid
	var modelUsers []*model.User
	var modelPosts []*model.Post

	// SEARCH USER BY KEYWORD
	if err := r.DB.Limit(limit).Offset(offset).Not("id = ?", userID).Find(&modelUsers, "concat(first_name, last_name) like ?", "%"+keyword+"%").Error; err != nil {
		return nil, err
	}

	// SEARCH POSTS BY KEYWOARD

	if err := r.DB.Limit(limit).Offset(offset).Find(&modelPosts, "text like ? ", "%"+keyword+"%").Error; err != nil {
		return nil, err
	}

	search.Users = modelUsers
	search.Posts = modelPosts

	return search, nil
}

// SearchHastag is the resolver for the SearchHastag field.
func (r *queryResolver) SearchHastag(ctx context.Context, keyword string, limit int, offset int) (*model.Search, error) {
	search := new(model.Search)

	var modelPosts []*model.Post

	// SEARCH POSTS BY KEYWOARD

	if err := r.DB.Limit(limit).Offset(offset).Find(&modelPosts, "text like ? ", "%#"+keyword+"%").Error; err != nil {
		return nil, err
	}

	search.Posts = modelPosts

	return search, nil
}

// Users is the resolver for the Users field.
func (r *searchResolver) Users(ctx context.Context, obj *model.Search) ([]*model.User, error) {
	var users []*model.User

	userIds := lo.Map(obj.Users, func(user *model.User, _ int) string {
		return user.ID
	})
	if len(userIds) == 0 {
		return users, nil
	}

	if err := r.DB.Find(&users, userIds).Error; err != nil {
		return nil, err
	}

	// log.Print(users)
	return users, nil
}

// Posts is the resolver for the Posts field.
func (r *searchResolver) Posts(ctx context.Context, obj *model.Search) ([]*model.Post, error) {
	var posts []*model.Post

	postIds := lo.Map(obj.Posts, func(post *model.Post, _ int) string {
		return post.ID
	})

	if len(postIds) == 0 {
		return posts, nil
	}

	if err := r.DB.Find(&posts, postIds).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// Search returns generated.SearchResolver implementation.
func (r *Resolver) Search() generated.SearchResolver { return &searchResolver{r} }

type searchResolver struct{ *Resolver }
