package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// AddJob is the resolver for the addJob field.
func (r *mutationResolver) AddJob(ctx context.Context, title string, companyName string, workplace string, city string, country string, employmentType string, description string) (*model.Job, error) {
	modelJobs := &model.Job{
		ID:             uuid.NewString(),
		Title:          title,
		CompanyName:    companyName,
		Workplace:      workplace,
		City:           city,
		Country:        country,
		EmploymentType: employmentType,
		Description:    description,
		CreatedAt:      time.Now(),
	}

	return modelJobs, r.DB.Create(modelJobs).Error
}

// Jobs is the resolver for the Jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.Job, error) {
	var modelJobs []*model.Job

	if err := r.DB.Order("created_at desc").Find(&modelJobs).Error; err != nil {
		return nil, err
	}

	return modelJobs, nil
}
