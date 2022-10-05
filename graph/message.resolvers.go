package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/renaldiaddison/tpa-web-backend/graph/generated"
	"github.com/renaldiaddison/tpa-web-backend/graph/model"
)

// Sender is the resolver for the sender field.
func (r *messageResolver) Sender(ctx context.Context, obj *model.Message) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.Find(&modelUser, "id = ?", obj.SenderID).Error
}

// SharePost is the resolver for the SharePost field.
func (r *messageResolver) SharePost(ctx context.Context, obj *model.Message) (*model.Post, error) {
	modelPost := new(model.Post)

	if err := r.DB.Find(&modelPost, "id = ?", obj.SharePostID).Error; err != nil {
		return nil, err
	}

	return modelPost, nil
}

// ShareProfile is the resolver for the ShareProfile field.
func (r *messageResolver) ShareProfile(ctx context.Context, obj *model.Message) (*model.User, error) {
	modelUser := new(model.User)

	if obj.ShareProfileID == nil {
		return modelUser, nil
	}

	if err := r.DB.Find(&modelUser, "id = ?", obj.ShareProfileID).Error; err != nil {
		return nil, err
	}

	return modelUser, nil
}

// AddRoom is the resolver for the addRoom field.
func (r *mutationResolver) AddRoom(ctx context.Context, userID1 string, userID2 string) (*model.Room, error) {
	modelRoom := &model.Room{
		ID:        uuid.NewString(),
		User1ID:   userID1,
		User2ID:   userID2,
		CreatedAt: time.Now(),
	}

	return modelRoom, r.DB.Create(modelRoom).Error
}

// AddMessage is the resolver for the addMessage field.
func (r *mutationResolver) AddMessage(ctx context.Context, senderID string, text string, imageURL string, roomID string) (*model.Message, error) {
	modelMessage := &model.Message{
		ID:        uuid.NewString(),
		Text:      text,
		ImageURL:  imageURL,
		SenderID:  senderID,
		RoomID:    roomID,
		CreatedAt: time.Now(),
	}

	return modelMessage, r.DB.Create(modelMessage).Error
}

// AddMessageSharePost is the resolver for the addMessageSharePost field.
func (r *mutationResolver) AddMessageSharePost(ctx context.Context, senderID string, roomID string, sharePostID string) (*model.Message, error) {
	modelMessage := &model.Message{
		ID:          uuid.NewString(),
		SenderID:    senderID,
		RoomID:      roomID,
		SharePostID: &sharePostID,
		CreatedAt:   time.Now(),
	}

	return modelMessage, r.DB.Create(modelMessage).Error
}

// AddMessageShareProfile is the resolver for the addMessageShareProfile field.
func (r *mutationResolver) AddMessageShareProfile(ctx context.Context, senderID string, roomID string, shareProfileID string) (*model.Message, error) {
	modelMessage := &model.Message{
		ID:             uuid.NewString(),
		SenderID:       senderID,
		RoomID:         roomID,
		ShareProfileID: &shareProfileID,
		CreatedAt:      time.Now(),
	}

	return modelMessage, r.DB.Create(modelMessage).Error
}

// Room is the resolver for the room field.
func (r *queryResolver) Room(ctx context.Context, roomID string) (*model.Room, error) {
	modelRoom := new(model.Room)

	return modelRoom, r.DB.Order("created_at desc").Find(modelRoom, "id = ?", roomID).Error
}

// Rooms is the resolver for the rooms field.
func (r *queryResolver) Rooms(ctx context.Context, userID string) ([]*model.Room, error) {
	var modelRooms []*model.Room

	if err := r.DB.Order("created_at desc").Where("user1_id = ?", userID).Or("user2_id = ?", userID).Find(&modelRooms).Error; err != nil {
		return nil, err
	}

	return modelRooms, nil
}

// User1 is the resolver for the user1 field.
func (r *roomResolver) User1(ctx context.Context, obj *model.Room) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.Find(&modelUser, "id = ?", obj.User1ID).Error
}

// User2 is the resolver for the user2 field.
func (r *roomResolver) User2(ctx context.Context, obj *model.Room) (*model.User, error) {
	modelUser := new(model.User)
	return modelUser, r.DB.Find(&modelUser, "id = ?", obj.User2ID).Error
}

// LastMessage is the resolver for the lastMessage field.
func (r *roomResolver) LastMessage(ctx context.Context, obj *model.Room) (*model.Message, error) {
	modelMessage := new(model.Message)

	if err := r.DB.Order("created_at desc").Limit(1).Find(&modelMessage, "room_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelMessage, nil
}

// Messages is the resolver for the messages field.
func (r *roomResolver) Messages(ctx context.Context, obj *model.Room) ([]*model.Message, error) {
	var modelMessages []*model.Message

	if err := r.DB.Order("created_at asc").Find(&modelMessages, "room_id = ?", obj.ID).Error; err != nil {
		return nil, err
	}

	return modelMessages, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type messageResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
