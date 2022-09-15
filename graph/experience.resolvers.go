package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// CreateExperience is the resolver for the createExperience field.
func (r *mutationResolver) CreateExperience(ctx context.Context, input model.NewExperience) (interface{}, error) {
	var endYear string
	if input.Active {
		endYear = "Present"
	} else {
		endYear = input.EndYear
	}
	model := &model.Experience{
		ID:             uuid.NewString(),
		UserID:         input.UserID,
		Title:          input.Title,
		EmploymentType: input.EmploymentType,
		CompanyName:    input.CompanyName,
		Location:       input.Location,
		Active:         input.Active,
		StartYear:      input.StartYear,
		EndYear:        endYear,
		Industry:       input.Industry,
		Description:    input.Description,
	}
	err := r.DB.Create(model).Error
	if err != nil {
		panic(err)
	}
	return model, nil
}

// UpdateExperience is the resolver for the updateExperience field.
func (r *mutationResolver) UpdateExperience(ctx context.Context, id string, input model.NewExperience) (interface{}, error) {
	var model *model.Experience

	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		return "Error", err
	}
	model.Title = input.Title
	model.EmploymentType = input.EmploymentType
	model.CompanyName = input.CompanyName
	model.Location = input.Location
	model.Active = input.Active
	model.StartYear = input.StartYear
	model.EndYear = input.EndYear
	model.Industry = input.Industry
	model.Description = input.Description

	return model, r.DB.Save(model).Error
}

// DeleteExperience is the resolver for the deleteExperience field.
func (r *mutationResolver) DeleteExperience(ctx context.Context, id string) (interface{}, error) {
	experience := new(model.Experience)
	if err := r.DB.First(experience, "id=?", id).Error; err != nil {
		panic(err)
	}
	return experience, nil
}

// UserExperience is the resolver for the userExperience field.
func (r *queryResolver) UserExperience(ctx context.Context, userID string) ([]*model.Experience, error) {
	var model []*model.Experience
	return model, r.DB.Where("user_id = ?", userID).Find(&model).Error
}
