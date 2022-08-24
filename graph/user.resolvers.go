package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/renaldiaddison/tpa-web-backend/auth"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/tools"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.UserCredentials) (interface{}, error) {
	return auth.UserLogin(ctx, input.Email, input.Password)
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (interface{}, error) {
	return auth.UserRegister(ctx, input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.NewUser) (*model.User, error) {
	model := new(model.User)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	model.Email = input.Email
	model.FirstName = input.FirstName
	model.LastName = input.LastName
	model.Password = input.Password
	return model, r.DB.Save(model).Error
}

// UpdateProfilePicture is the resolver for the updateProfilePicture field.
func (r *mutationResolver) UpdateProfilePicture(ctx context.Context, id string, imageURL string) (interface{}, error) {
	model := new(model.User)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	model.ProfilePicture = imageURL
	return model, r.DB.Save(model).Error
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	model := new(model.User)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	return model, r.DB.Delete(model).Error
}

// ActivateUser is the resolver for the activateUser field.
func (r *mutationResolver) ActivateUser(ctx context.Context, id string) (interface{}, error) {
	user := new(model.User)
	link := new(model.ActivationLink)
	if err := r.DB.First(user, "id = ?", id).Error; err != nil {
		panic(err)
	}
	if is_active := user.IsActive; is_active == false {
		user.IsActive = true
	} else {
		user.IsActive = false
	}

	if err := r.DB.Delete(link, "user_id = ?", id).Error; err != nil {
		panic(err)
	}

	return user, r.DB.Save(user).Error
}

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, email string, newPassword string) (interface{}, error) {
	user := new(model.User)
	link := new(model.ResetPasswordLink)

	if err := r.DB.First(user, "email = ?", email).Error; err != nil {
		panic(err)
	}

	if err := r.DB.Delete(link, "email = ?", email).Error; err != nil {
		panic(err)
	}

	user.Password = tools.HashPassword(newPassword)

	return user, r.DB.Save(user).Error
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	model := new(model.User)
	return model, r.DB.First(model, "id = ?", id).Error
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var models []*model.User
	return models, r.DB.Find(&models).Error
}

// Protected is the resolver for the protected field.
func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "Success", nil
}

// FollowedUser is the resolver for the followed_user field.
func (r *userResolver) FollowedUser(ctx context.Context, obj *model.User) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// ConnectedUser is the resolver for the connected_user field.
func (r *userResolver) ConnectedUser(ctx context.Context, obj *model.User) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// RequestConnect is the resolver for the request_connect field.
func (r *userResolver) RequestConnect(ctx context.Context, obj *model.User) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
