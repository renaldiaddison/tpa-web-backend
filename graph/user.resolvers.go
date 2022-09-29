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
	"github.com/samber/lo"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUser) (*model.User, error) {
	model := new(model.User)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	model.FirstName = input.FirstName
	model.LastName = input.LastName
	model.AdditionalName = input.AdditionalName
	model.About = input.About
	model.Location = input.Location
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

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, id1 string, id2 string) (interface{}, error) {
	modelFollow := new(model.Follow)

	modelFollow.UserID = id1
	modelFollow.FollowID = id2

	r.DB.Table("user_follows").Create(modelFollow)

	var modelFollows []*model.Follow
	r.DB.Table("user_follows").Find(&modelFollows, "follow_id = ?", id2)

	return map[string]interface{}{
		"length": len(modelFollows),
	}, nil
}

// UnFollowUser is the resolver for the unFollowUser field.
func (r *mutationResolver) UnFollowUser(ctx context.Context, id1 string, id2 string) (interface{}, error) {
	modelFollow := new(model.Follow)

	if err := r.DB.Table("user_follows").First(&modelFollow, "user_id = ? AND follow_id = ?", id1, id2).Error; err != nil {
		return nil, err
	}

	if modelFollow.UserID == "" {
		var modelFollows []*model.Follow
		r.DB.Table("user_follows").Find(&modelFollows, "follow_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelFollows),
		}, nil
	} else {
		r.DB.Table("user_follows").Delete(&modelFollow, "user_id = ? AND follow_id = ?", id1, id2)

		var modelFollows []*model.Follow
		r.DB.Table("user_follows").Find(&modelFollows, "follow_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelFollows),
		}, nil
	}
}

// VisitUser is the resolver for the visitUser field.
func (r *mutationResolver) VisitUser(ctx context.Context, id1 string, id2 string) (interface{}, error) {
	modelVisit := new(model.Visit)

	r.DB.Table("user_visits").First(&modelVisit, "user_id = ? AND visit_id = ?", id1, id2)

	if modelVisit.UserID != "" {
		var modelVisits []*model.Visit
		r.DB.Table("user_visits").Find(&modelVisits, "visit_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelVisits),
		}, nil
	} else {
		modelVisit.UserID = id1
		modelVisit.VisitID = id2

		if err := r.DB.Table("user_visits").Create(modelVisit).Error; err == nil {
			service.AddNotification(r.DB, ctx, id2, id1, "Visit Your Profile")
		}

		var modelVisits []*model.Visit
		r.DB.Table("user_visits").Find(&modelVisits, "visit_id = ?", id2)

		return map[string]interface{}{
			"length": len(modelVisits),
		}, nil
	}
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

// UserSuggestion is the resolver for the UserSuggestion field.
func (r *queryResolver) UserSuggestion(ctx context.Context, userID string) ([]*model.User, error) {
	var modelUsers []*model.User
	var userIdList []string
	var userSuggestionId []string

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

	var friendConnection1 []*model.Connection
	var friendConnection2 []*model.Connection

	if err := r.DB.Where("user1_id IN ?", userIdList).Not("user2_id = ?", userID).Find(&friendConnection1).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Where("user2_id IN ?", userIdList).Not("user1_id = ?", userID).Find(&friendConnection2).Error; err != nil {
		return nil, err
	}

	userSuggestion1Ids := lo.Map(friendConnection1, func(connectionData *model.Connection, _ int) string {
		return connectionData.User2ID
	})

	userSuggestion2Ids := lo.Map(friendConnection2, func(connectionData *model.Connection, _ int) string {
		return connectionData.User1ID
	})

	userSuggestionId = append(userSuggestionId, userSuggestion1Ids...)
	userSuggestionId = append(userSuggestionId, userSuggestion2Ids...)
	userSuggestionId = lo.Uniq(userSuggestionId)

	var finalUserSuggestionId []string
	for _, suggestionIdUser := range userSuggestionId {
		checkSame := false
		for _, userConnectionId := range userIdList {
			if suggestionIdUser == userConnectionId {
				checkSame = true
			}
		}

		if !checkSame {
			finalUserSuggestionId = append(finalUserSuggestionId, suggestionIdUser)
		}
	}

	if len(finalUserSuggestionId) == 0 {
		return nil, gqlerror.Errorf("No Connection User Data")
	}

	if err := r.DB.Find(&modelUsers, finalUserSuggestionId).Error; err != nil {
		return nil, err
	}

	return modelUsers, nil
}

// Visits is the resolver for the visits field.
func (r *userResolver) Visits(ctx context.Context, obj *model.User) ([]*model.Visit, error) {
	var modelVisits []*model.Visit

	return modelVisits, r.DB.Table("user_visits").Find(&modelVisits, "visit_id = ?", obj.ID).Error
}

// Follows is the resolver for the follows field.
func (r *userResolver) Follows(ctx context.Context, obj *model.User) ([]*model.Follow, error) {
	var modelFollow []*model.Follow

	return modelFollow, r.DB.Table("user_follows").Find(&modelFollow, "follow_id = ? ", obj.ID).Error
}

// Block is the resolver for the Block field.
func (r *userResolver) Block(ctx context.Context, obj *model.User) ([]*model.Block, error) {
	var modelBlocks []*model.Block

	if err := r.DB.Table("user_blocks").Find(&modelBlocks, "user_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelBlocks, nil
}

// Connection is the resolver for the Connection field.
func (r *userResolver) Connection(ctx context.Context, obj *model.User) ([]*model.Connection, error) {
	var modelConnections []*model.Connection

	if err := r.DB.Where("user1_id = ?", obj.ID).Or("user2_id = ?", obj.ID).Find(&modelConnections).Error; err != nil {
		return nil, err
	}

	return modelConnections, nil
}

// ConnectRequest is the resolver for the ConnectRequest field.
func (r *userResolver) ConnectRequest(ctx context.Context, obj *model.User) ([]*model.ConnectRequest, error) {
	var modelConnectionRequests []*model.ConnectRequest

	if err := r.DB.Find(&modelConnectionRequests, "to_user_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelConnectionRequests, nil
}

// Experiences is the resolver for the experiences field.
func (r *userResolver) Experiences(ctx context.Context, obj *model.User) ([]*model.Experience, error) {
	var modelExperiences []*model.Experience

	return modelExperiences, r.DB.Where("user_id = ?", obj.ID).Find(&modelExperiences).Error
}

// Educations is the resolver for the educations field.
func (r *userResolver) Educations(ctx context.Context, obj *model.User) ([]*model.Education, error) {
	var modelEducations []*model.Education

	return modelEducations, r.DB.Where("user_id = ? ", obj.ID).Find(&modelEducations).Error
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
