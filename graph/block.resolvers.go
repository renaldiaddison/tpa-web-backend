package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// AddBlock is the resolver for the addBlock field.
func (r *mutationResolver) AddBlock(ctx context.Context, userID string, blockID string) (*model.Block, error) {
	modelBlock := &model.Block{
		UserID:  userID,
		BlockID: blockID,
	}

	return modelBlock, r.DB.Table("user_blocks").Create(modelBlock).Error
}

// DeleteBlock is the resolver for the deleteBlock field.
func (r *mutationResolver) DeleteBlock(ctx context.Context, userID string, blockID string) (*model.Block, error) {
	modelBlock := new(model.Block)

	return modelBlock, r.DB.Table("user_blocks").Delete(modelBlock, "user_id = ? AND block_id = ?", userID, blockID).Error
}

// Blocks is the resolver for the blocks field.
func (r *queryResolver) Blocks(ctx context.Context, userID string) ([]*model.Block, error) {
	var modelBlocks []*model.Block

	if err := r.DB.Table("user_blocks").Where("user_id = ?", userID).Or("block_id = ?", userID).Find(&modelBlocks).Error; err != nil {
		return nil, err
	}

	return modelBlocks, nil
}
