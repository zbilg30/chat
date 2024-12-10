package repository

import (
	"auth/internal/models"
	"auth/internal/pkg"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRepository struct {
    mongoClient pkg.IMongoClient 
}

type IChatRepository interface {
	CreateChat(ctx context.Context, chatRoom *models.Chat) (*models.Chat, error)
}

func NewChatRepository(db pkg.IMongoClient) IChatRepository {
    return &ChatRepository{
        mongoClient: db,
    }
}

func(c *ChatRepository) CreateChat(ctx context.Context, chat *models.Chat) (*models.Chat, error) {
	chat.ID = primitive.NewObjectID()
    _, err := c.mongoClient.GetCollection("chats").InsertOne(ctx, chat)

    if err != nil {
        return nil, err
    }
    return chat, nil
}

