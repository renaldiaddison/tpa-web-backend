package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/database"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/tools"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := database.GetDatabase()

	input.Password = tools.HashPassword(input.Password)
	user := model.User{
		ID:             uuid.NewString(),
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Email:          input.Email,
		Password:       input.Password,
		ProfilePicture: "https://firebasestorage.googleapis.com/v0/b/linkhedin-2b334.appspot.com/o/default.jpg?alt=media&token=bc68161a-27b9-47f1-a07a-eaa2e2f1a757",
	}

	if err := db.Model(user).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetByID(ctx context.Context, id string) (*model.User, error) {
	db := database.GetDatabase()

	var user model.User
	if err := db.Model(user).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	db := database.GetDatabase()

	var user model.User
	if err := db.Model(user).Where("email LIKE ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
