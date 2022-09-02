package auth

import (
	"context"
	"errors"

	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/mail"
	"github.com/renaldiaddison/tpa-web-backend/service"
	"github.com/renaldiaddison/tpa-web-backend/tools"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, input model.NewUser) (*model.User, error) {
	_, err := service.UserGetByEmail(ctx, input.Email)

	if err == nil {
		if err != gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "Email is already taken",
			}
		}
	}

	createdUser, err := service.UserCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	link, err := service.ActivationLinkCreate(ctx, createdUser.ID)
	if err != nil {
		return nil, err
	}

	mail.SendEmail("This is your linkHEdIn's account activation link!! ", "linkHEdIn Account Activation", createdUser.Email, link)
	return createdUser, nil
}

func UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := service.UserGetByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "Email not found",
			}
		}
		return nil, err
	}

	if getUser.IsActive == false {
		return nil, errors.New("Your account hasn't been authenticated")
	}

	if err := tools.ComparePassword(getUser.Password, password); err != nil {
		return nil, errors.New("Wrong Credentials")
	}

	token, err := service.JwtGenerate(ctx, getUser.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token":              token,
		"id":                 getUser.ID,
		"email":              getUser.Email,
		"name":               getUser.FirstName + " " + getUser.LastName + " " + getUser.AdditionalName,
		"profile_picture":    getUser.ProfilePicture,
		"background_picture": getUser.BackgroundPicture,
		"headline":           getUser.Headline,
		"about":              getUser.About,
		"location":           getUser.Location,
		"followed_user":      getUser.FollowedUser,
		"profile_views":      getUser.ProfileViews,
		"request_connect":    getUser.RequestConnect,
		"connected_user":     getUser.ConnectedUser,
	}, nil
}
