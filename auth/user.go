package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/database"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := database.GetDatabase()

	input.Password = HashPassword(input.Password)
	user := model.User{
		ID:        uuid.NewString(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
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
