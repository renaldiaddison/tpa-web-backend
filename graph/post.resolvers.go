package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/middlewares"
	"github.com/renaldiaddison/tpa-web-backend/service"
	"github.com/samber/lo"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.InputPost) (*model.Post, error) {
	modelPost := &model.Post{
		ID:        uuid.NewString(),
		Text:      input.Text,
		PhotoUrl:  input.PhotoURL,
		VideoUrl:  input.VideoURL,
		CreatedAt: time.Now(),
		SenderId:  input.SenderID,
	}

	// CONNECTED USER POST
	var userIdList []string
	var connections1 []*model.Connection
	var connections2 []*model.Connection

	if err := r.DB.Find(&connections1, "user1_id", input.SenderID).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Find(&connections2, "user2_id", input.SenderID).Error; err != nil {
		return nil, err
	}

	connetions1Ids := lo.Map(connections1, func(connectionData *model.Connection, _ int) string {
		return connectionData.User2ID
	})

	connetions2Ids := lo.Map(connections2, func(connectionData *model.Connection, _ int) string {
		return connectionData.User1ID
	})

	userIdList = append(userIdList, connetions1Ids...)
	userIdList = append(userIdList, connetions2Ids...)
	userIdList = lo.Uniq(userIdList)

	for _, userId := range userIdList {
		service.AddNotification(r.DB, ctx, userId, input.SenderID, "Create A New Post")
	}

	return modelPost, r.DB.Create(modelPost).Error
}

// LikePost is the resolver for the LikePost field.
func (r *mutationResolver) LikePost(ctx context.Context, postID string, userID string) (*model.LikePosts, error) {
	modelLikePost := &model.LikePosts{
		PostId: postID,
		UserId: userID,
	}

	return modelLikePost, r.DB.Create(modelLikePost).Error
}

// UnLikePost is the resolver for the UnLikePost field.
func (r *mutationResolver) UnLikePost(ctx context.Context, postID string, userID string) (*model.LikePosts, error) {
	modelLikePost := new(model.LikePosts)

	if err := r.DB.Find(modelLikePost, "post_id = ? AND user_id = ?", postID, userID).Error; err != nil {
		return nil, err
	}

	return modelLikePost, r.DB.Delete(modelLikePost, "post_id = ? AND user_id = ?", postID, userID).Error
}

// Sender is the resolver for the Sender field.
func (r *postResolver) Sender(ctx context.Context, obj *model.Post) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.SenderId).Error
}

// Likes is the resolver for the Likes field.
func (r *postResolver) Likes(ctx context.Context, obj *model.Post) ([]*model.LikePosts, error) {
	var modelLikePost []*model.LikePosts

	if err := r.DB.Find(&modelLikePost, "post_id", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelLikePost, nil
}

// Comments is the resolver for the Comments field.
func (r *postResolver) Comments(ctx context.Context, obj *model.Post) ([]*model.Comment, error) {
	var modelComment []*model.Comment

	if err := r.DB.Find(&modelComment, "post_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelComment, nil
}

// Posts is the resolver for the Posts field.
func (r *queryResolver) Posts(ctx context.Context, limit int, offset int) ([]*model.Post, error) {
	var userIdList []string
	userID := middlewares.GetJwtValueData(ctx).Userid
	userIdList = append(userIdList, userID)

	// FOLLOWED USER POST
	var follows []*model.Follow

	if err := r.DB.Table("user_follows").Find(&follows, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	followIds := lo.Map(follows, func(x *model.Follow, _ int) string {
		return x.FollowID
	})

	userIdList = append(userIdList, followIds...)

	// CONNECTED USER POST
	var connections1 []*model.Connection
	var connections2 []*model.Connection

	if err := r.DB.Find(&connections1, "user1_id", userID).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Find(&connections2, "user2_id", userID).Error; err != nil {
		return nil, err
	}

	connetions1Ids := lo.Map(connections1, func(connectionData *model.Connection, _ int) string {
		return connectionData.User2ID
	})

	connetions2Ids := lo.Map(connections2, func(connectionData *model.Connection, _ int) string {
		return connectionData.User1ID
	})

	userIdList = append(userIdList, connetions1Ids...)
	userIdList = append(userIdList, connetions2Ids...)
	userIdList = lo.Uniq(userIdList)

	var posts []*model.Post
	if err := r.DB.Limit(limit).Offset(offset).Order("created_at desc").Find(&posts, "sender_id IN ?", userIdList).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
