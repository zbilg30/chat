package resolver

import (
	"auth/internal/models"
	"auth/internal/repository"
	"context"
)


type ChatRoomResolver struct {
    ChatRoomRepo repository.IChatRoomRepository
}


func (r *ChatRoomResolver) CreateChatRoom(ctx context.Context, chatRoom *models.ChatRoom) (interface{} , error) {
	chatRoom, err := r.ChatRoomRepo.CreateChatRoom(ctx, chatRoom)
	if err != nil {
		return nil, err
	}
	
	return chatRoom.ToMap(), nil
}