package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
	"github.com/renaldiaddison/tpa-web-backend/middlewares"
)

// CreateEducation is the resolver for the createEducation field.
func (r *mutationResolver) CreateEducation(ctx context.Context, input model.NewEducation) (interface{}, error) {
	model := &model.Education{
		ID:           uuid.NewString(),
		UserID:       input.UserID,
		School:       input.School,
		Degree:       input.Degree,
		FieldOfStudy: input.FieldOfStudy,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
		Grade:        input.Grade,
		Activities:   input.Activities,
		Description:  input.Description,
	}
	err := r.DB.Create(model).Error
	if err != nil {
		panic(err)
	}
	return model, nil
}

// UpdateEducation is the resolver for the updateEducation field.
func (r *mutationResolver) UpdateEducation(ctx context.Context, id string, input model.NewEducation) (interface{}, error) {
	var model *model.Education

	if err := r.DB.First(&model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	model.School = input.School
	model.Degree = input.Degree
	model.FieldOfStudy = input.FieldOfStudy
	model.StartDate = input.StartDate
	model.EndDate = input.EndDate
	model.Grade = input.Grade
	model.Activities = input.Activities
	model.Description = input.Description

	return model, r.DB.Save(model).Error
}

// DeleteEducation is the resolver for the deleteEducation field.
func (r *mutationResolver) DeleteEducation(ctx context.Context, id string) (interface{}, error) {
	model := new(model.Education)
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		panic(err)
	}
	return model, r.DB.Delete(model).Error
}

// UserEducation is the resolver for the userEducation field.
func (r *queryResolver) UserEducation(ctx context.Context, userID string) ([]*model.Education, error) {
	var model []*model.Education
	return model, r.DB.Where("user_id = ?", userID).Find(&model).Error
}

// MyEducation is the resolver for the myEducation field.
func (r *queryResolver) MyEducation(ctx context.Context) ([]*model.Education, error) {
	val := *middlewares.GetJwtValueData(ctx)
	var models []*model.Education
	return models, r.DB.Where("user_id = ?", val.Userid).Find(&models).Error
}
