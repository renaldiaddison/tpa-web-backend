package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/mail"
	"github.com/renaldiaddison/tpa-web-backend/service"
)

// CreateAndSendResetLink is the resolver for the createAndSendResetLink field.
func (r *mutationResolver) CreateAndSendResetLink(ctx context.Context, email string) (string, error) {
	_, err := service.ResetLinkGetByEmail(ctx, email)

	if err == nil {
		return "", errors.New("You have already requested, please check your email!")
	}

	link := model.ResetPasswordLink{
		ID:    uuid.NewString(),
		Email: email,
	}

	if err := r.DB.Model(link).Create(&link).Error; err != nil {
		return "", err
	}

	resetLink := "http://localhost:5173/reset-password/" + link.ID

	mail.SendEmail("This is your password reset link!! ", "linkHEdIn Reset Password", email, resetLink)

	return resetLink, nil
}

// GetResetLink is the resolver for the getResetLink field.
func (r *queryResolver) GetResetLink(ctx context.Context, id string) (*model.ResetPasswordLink, error) {
	model := new(model.ResetPasswordLink)
	return model, r.DB.First(model, "id = ?", id).Error
}
