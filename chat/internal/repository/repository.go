package repository

import (
	"auth/internal/pkg"
)

type RepositoryLayer struct {
    ChatRepository    IChatRepository
    ChatRoomRepository IChatRoomRepository
}

func NewRepositoryLayer(db pkg.IMongoClient) *RepositoryLayer {
    return &RepositoryLayer{
        ChatRepository:    NewChatRepository(db),
        ChatRoomRepository: NewChatRoomRepository(db),
    }
}
