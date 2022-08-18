package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/auth"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.UserCredentials) (interface{}, error) {
	return auth.UserLogin(ctx, input.Email, input.Password)
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (interface{}, error) {
	return auth.UserRegister(ctx, input)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	model := &model.User{
		ID:        uuid.NewString(),
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
		IsActive:  false,
	}

	err := r.DB.Create(model).Error
	return model, err
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
	model := new(model.User)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	if is_active := model.IsActive; is_active == false {
		model.IsActive = true
	} else {
		model.IsActive = false
	}
	return model, r.DB.Save(model).Error
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
