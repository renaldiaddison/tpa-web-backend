package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/renaldiaddison/tpa-web-backend/auth"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/service"
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

// UpdateBackgroundPicture is the resolver for the updateBackgroundPicture field.
func (r *mutationResolver) UpdateBackgroundPicture(ctx context.Context, id string, imageURL string) (interface{}, error) {
	model := new(model.User)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	model.BackgroundPicture = imageURL
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

// RequestConnect is the resolver for the requestConnect field.
func (r *mutationResolver) RequestConnect(ctx context.Context, id string, recipientID string) (interface{}, error) {
	user := new(model.User)
	err := r.DB.First(user, "id = ?", recipientID).Error
	if err != nil {
		panic(err)
	}
	user.RequestConnect = append(user.RequestConnect, id)
	return map[string]interface{}{}, r.DB.Save(user).Error
}

// AcceptConnect is the resolver for the acceptConnect field.
func (r *mutationResolver) AcceptConnect(ctx context.Context, id string, senderID string) (interface{}, error) {
	recepient := new(model.User)
	sender := new(model.User)
	if err := r.DB.First(recepient, "id = ?", id).Error; err != nil {
		panic(err)
	}
	if err := r.DB.First(sender, "id = ?", senderID).Error; err != nil {
		panic(err)
	}

	new_arr := make([]string, (len(recepient.RequestConnect) - 1))
	k := 0
	for i := 0; i < (len(recepient.RequestConnect) - 1); {
		if recepient.RequestConnect[i] != senderID {
			new_arr[i] = recepient.RequestConnect[k]
			k++
			i++
		} else {
			k++
		}
	}
	recepient.RequestConnect = new_arr
	sender.ConnectedUser = append(sender.ConnectedUser, id)
	recepient.ConnectedUser = append(recepient.ConnectedUser, senderID)
	if err := r.DB.Save(recepient).Error; err != nil {
		panic(err)
	}
	if err := r.DB.Save(sender).Error; err != nil {
		panic(err)
	}
	return map[string]interface{}{}, nil
}

// IgnoreConnect is the resolver for the ignoreConnect field.
func (r *mutationResolver) IgnoreConnect(ctx context.Context, id string, senderID string) (interface{}, error) {
	recepient := new(model.User)
	if err := r.DB.First(recepient, "id=?", id).Error; err != nil {
		panic(err)
	}
	new_arr := make([]string, (len(recepient.RequestConnect) - 1))
	k := 0
	for i := 0; i < (len(recepient.RequestConnect) - 1); {
		if recepient.RequestConnect[i] != senderID {
			new_arr[i] = recepient.RequestConnect[k]
			k++
			i++
		} else {
			k++
		}
	}
	recepient.RequestConnect = new_arr
	if err := r.DB.Save(recepient).Error; err != nil {
		panic(err)
	}
	return map[string]interface{}{}, nil
}

// Follow is the resolver for the follow field.
func (r *mutationResolver) Follow(ctx context.Context, id string, followedID string) (interface{}, error) {
	user := new(model.User)
	if err := r.DB.First(user, "id=?", id).Error; err != nil {
		return "failed", err
	}
	user.FollowedUser = append(user.FollowedUser, followedID)
	if err := r.DB.Save(user).Error; err != nil {
		return "failed", err
	}
	return map[string]interface{}{}, nil
}

// Unfollow is the resolver for the unfollow field.
func (r *mutationResolver) Unfollow(ctx context.Context, id string, unfollowedID string) (interface{}, error) {
	user := new(model.User)
	if err := r.DB.First(user, "id=?", id).Error; err != nil {
		panic(err)
	}
	new_arr := make([]string, (len(user.FollowedUser) - 1))
	k := 0
	for i := 0; i < (len(user.FollowedUser) - 1); {
		if user.FollowedUser[i] != unfollowedID {
			new_arr[i] = user.FollowedUser[k]
			k++
			i++
		} else {
			k++
		}
	}
	user.FollowedUser = new_arr
	if err := r.DB.Save(user).Error; err != nil {
		panic(err)
	}
	return map[string]interface{}{}, nil
}

// GetUserByID is the resolver for the getUserById field.
func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	model, err := service.UserGetByID(ctx, id)
	return model, err
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	var models []*model.User
	return models, r.DB.Find(&models).Error
}

// FollowedUser is the resolver for the followed_user field.
func (r *userResolver) FollowedUser(ctx context.Context, obj *model.User) ([]string, error) {
	return obj.FollowedUser, nil
}

// ConnectedUser is the resolver for the connected_user field.
func (r *userResolver) ConnectedUser(ctx context.Context, obj *model.User) ([]string, error) {
	return obj.ConnectedUser, nil
}

// RequestConnect is the resolver for the request_connect field.
func (r *userResolver) RequestConnect(ctx context.Context, obj *model.User) ([]string, error) {
	return obj.RequestConnect, nil
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
