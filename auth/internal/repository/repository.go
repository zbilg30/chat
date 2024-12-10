package repository

import (
	"auth/internal/pkg"
)

type RepositoryLayer struct {
    UserRepository    IUserRepository
}

func NewRepositoryLayer(db pkg.IMongoClient) *RepositoryLayer {
    return &RepositoryLayer{
        UserRepository:    NewUserRepository(db),
    }
}
