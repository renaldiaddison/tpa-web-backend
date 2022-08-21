package service

import (
	"context"

	"github.com/renaldiaddison/tpa-web-backend/database"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

func ResetLinkGetByEmail(ctx context.Context, email string) (*model.ResetPasswordLink, error) {
	db := database.GetDatabase()

	var resetLink model.ResetPasswordLink
	if err := db.Model(resetLink).Where("email LIKE ?", email).Take(&resetLink).Error; err != nil {
		return nil, err
	}

	return &resetLink, nil
}
