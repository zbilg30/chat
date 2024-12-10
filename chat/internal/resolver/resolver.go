package resolver

import "auth/internal/repository"

type Resolver struct {
    ChatResolver    *ChatResolver
    ChatRoomResolver *ChatRoomResolver
}

func NewResolver(repositories *repository.RepositoryLayer) *Resolver {
    return &Resolver{
        ChatResolver:    &ChatResolver{ChatRepo: repositories.ChatRepository},
        ChatRoomResolver: &ChatRoomResolver{ChatRoomRepo: repositories.ChatRoomRepository},
    }
}