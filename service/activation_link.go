package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/database"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

func ActivationLinkCreate(ctx context.Context, input string) (string, error) {
	db := database.GetDatabase()

	link := model.ActivationLink{
		ID:     uuid.NewString(),
		UserID: input,
	}

	if err := db.Model(link).Create(&link).Error; err != nil {
		return "", err
	}

	return "http://localhost:5173/activate-account/" + link.ID, nil
}
