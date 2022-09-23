package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Commenter is the resolver for the Commenter field.
func (r *commentResolver) Commenter(ctx context.Context, obj *model.Comment) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.CommenterID).Error
}

// Replies is the resolver for the Replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	var modelComments []*model.Comment

	if err := r.DB.Find(&modelComments, "reply_to_comment_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelComments, nil
}

// Likes is the resolver for the Likes field.
func (r *commentResolver) Likes(ctx context.Context, obj *model.Comment) ([]*model.LikeComment, error) {
	var modelLikes []*model.LikeComment

	if err := r.DB.Find(&modelLikes, "comment_id", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelLikes, nil
}

// User is the resolver for the User field.
func (r *likeCommentResolver) User(ctx context.Context, obj *model.LikeComment) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.First(modelUser, "id = ?", obj.UserID).Error
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, postID string, commenterID string, comment string) (*model.Comment, error) {
	modelComment := &model.Comment{
		ID:          uuid.NewString(),
		PostID:      postID,
		CommenterID: commenterID,
		Comment:     comment,
		CreatedAt:   time.Now(),
	}

	return modelComment, r.DB.Create(modelComment).Error
}

// AddLikeComment is the resolver for the addLikeComment field.
func (r *mutationResolver) AddLikeComment(ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	modelLikeComment := &model.LikeComment{
		ID:        uuid.NewString(),
		CommentID: commentID,
		UserID:    userID,
	}

	modelLikeCommentDb := new(model.LikeComment)
	if err := r.DB.Find(modelLikeComment, "user_id = ? AND comment_id = ?", userID, commentID).Error; err != nil {
		return nil, err
	}

	if modelLikeCommentDb.ID != "" {
		return nil, gqlerror.Errorf("You cannot like twice")
	}

	return modelLikeComment, r.DB.Create(modelLikeComment).Error
}

// DeleteLikeComment is the resolver for the deleteLikeComment field.
func (r *mutationResolver) DeleteLikeComment(ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	modelLikeComment := new(model.LikeComment)

	if err := r.DB.Find(modelLikeComment, "comment_id = ? AND user_id = ?", commentID, userID).Error; err != nil {
		return nil, err
	}

	return modelLikeComment, r.DB.Delete(modelLikeComment).Error
}

// AddReply is the resolver for the addReply field.
func (r *mutationResolver) AddReply(ctx context.Context, commenterID string, postID string, replyToCommentID string, comment string) (*model.Comment, error) {
	modelComment := &model.Comment{
		ID:               uuid.NewString(),
		PostID:           postID,
		CommenterID:      commenterID,
		Comment:          comment,
		ReplyToCommentID: &replyToCommentID,
		CreatedAt:        time.Now(),
	}

	return modelComment, r.DB.Create(modelComment).Error
}

// PostComment is the resolver for the postComment field.
func (r *queryResolver) PostComment(ctx context.Context, id string) (*model.Comment, error) {
	modelComment := new(model.Comment)

	if err := r.DB.Find(modelComment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return modelComment, nil
}

// RepliedToComments is the resolver for the repliedToComments field.
func (r *queryResolver) RepliedToComments(ctx context.Context, limit int, offset int, commentID string) ([]*model.Comment, error) {
	var modelComments []*model.Comment

	if err := r.DB.Limit(limit).Offset(offset).Order("created_at desc").Find(&modelComments, "reply_to_comment_id = ?", commentID).Error; err != nil {
		return nil, err
	}

	return modelComments, nil
}

// PostComments is the resolver for the postComments field.
func (r *queryResolver) PostComments(ctx context.Context, limit int, offset int, postID string) ([]*model.Comment, error) {
	var modelComments []*model.Comment

	if err := r.DB.Limit(limit).Offset(offset).Order("created_at desc").Find(&modelComments, "post_id = ? AND reply_to_comment_id IS NULL", postID).Error; err != nil {
		return nil, err
	}

	if len(modelComments) == 0 {
		return nil, gqlerror.Errorf("No More Data")
	}

	return modelComments, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// LikeComment returns generated.LikeCommentResolver implementation.
func (r *Resolver) LikeComment() generated.LikeCommentResolver { return &likeCommentResolver{r} }

type commentResolver struct{ *Resolver }
type likeCommentResolver struct{ *Resolver }
