package auth

import (
	"context"
	"errors"

	"github.com/renaldiaddison/tpa-web-backend/email"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, input model.NewUser) (interface{}, error) {
	_, err := UserGetByEmail(ctx, input.Email)

	if err == nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	createdUser, err := UserCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	link, err := ActivationLinkCreate(ctx, createdUser.ID)

	if err != nil {
		return nil, err
	}

	email.SendEmail(createdUser.Email, link)
	return map[string]interface{}{}, nil
}

func UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := UserGetByEmail(ctx, email)
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

	if err := ComparePassword(getUser.Password, password); err != nil {
		return nil, err
	}

	token, err := JwtGenerate(ctx, getUser.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":    getUser.ID,
		"token": token,
		"name":  getUser.FirstName + " " + getUser.LastName,
		"email": getUser.Email,
	}, nil
}
