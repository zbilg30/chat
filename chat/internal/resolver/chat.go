package resolver

import (
	"auth/internal/kafka"
	"auth/internal/models"
	"auth/internal/repository"
	"context"
	"fmt"
	"log"
)


type ChatResolver struct {
    ChatRepo repository.IChatRepository
}


func (r *ChatResolver) CreateChat(ctx context.Context, chat *models.Chat, kafkaProducer *kafka.Producer) (interface{} , error) {
	chat, err := r.ChatRepo.CreateChat(ctx, chat)
	if err != nil {
		return nil, err
	}

    message := &kafka.Message{
        SenderID: chat.UserId,
        Content:  chat.Message,
		ChatRoomId: chat.ChatRoomId,
    }

    if err := kafkaProducer.SendMessage("chat.messages", message); err != nil {
        log.Fatalf("Failed to send message: %v", err)
    }
	
	fmt.Println("Send to kafka")
	return chat.ToMap(), nil
}