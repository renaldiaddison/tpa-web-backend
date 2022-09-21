package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// Commenter is the resolver for the Commenter field.
func (r *commentResolver) Commenter(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Replies is the resolver for the Replies field.
func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Likes is the resolver for the Likes field.
func (r *commentResolver) Likes(ctx context.Context, obj *model.Comment) ([]*model.LikeComment, error) {
	panic(fmt.Errorf("not implemented"))
}

// User is the resolver for the User field.
func (r *likeCommentResolver) User(ctx context.Context, obj *model.LikeComment) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, postID string, commenterID string, comment string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddLikeComment is the resolver for the addLikeComment field.
func (r *mutationResolver) AddLikeComment(ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	panic(fmt.Errorf("not implemented"))
}

// DeleteLikeComment is the resolver for the deleteLikeComment field.
func (r *mutationResolver) DeleteLikeComment(ctx context.Context, commentID string, userID string) (*model.LikeComment, error) {
	panic(fmt.Errorf("not implemented"))
}

// AddReply is the resolver for the addReply field.
func (r *mutationResolver) AddReply(ctx context.Context, commenterID string, postID string, replyToCommentID string, comment string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// PostComment is the resolver for the postComment field.
func (r *queryResolver) PostComment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// RepliedToComments is the resolver for the repliedToComments field.
func (r *queryResolver) RepliedToComments(ctx context.Context, limit int, offset int, commentID string) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// PostComments is the resolver for the postComments field.
func (r *queryResolver) PostComments(ctx context.Context, limit int, offset int, postID string) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// LikeComment returns generated.LikeCommentResolver implementation.
func (r *Resolver) LikeComment() generated.LikeCommentResolver { return &likeCommentResolver{r} }

type commentResolver struct{ *Resolver }
type likeCommentResolver struct{ *Resolver }
