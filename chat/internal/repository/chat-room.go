package repository

import (
	"auth/internal/models"
	"auth/internal/pkg"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRoomRepository struct {
    mongoClient pkg.IMongoClient 
}

type IChatRoomRepository interface {
	CreateChatRoom(ctx context.Context, chatRoom *models.ChatRoom) (*models.ChatRoom, error)
}

func NewChatRoomRepository(db pkg.IMongoClient) IChatRoomRepository {
    return &ChatRoomRepository{
        mongoClient: db,
    }
}

func(c *ChatRoomRepository) CreateChatRoom(ctx context.Context, chatRoom *models.ChatRoom) (*models.ChatRoom, error) {
    chatRoom.ID = primitive.NewObjectID()
    _, err := c.mongoClient.GetCollection("chatRooms").InsertOne(ctx, chatRoom)
    
    if err != nil {
        return nil, err
    }
    return chatRoom, nil
}