package resolver

import (
	"auth/internal/models"
	"auth/internal/repository"
	"context"
	"fmt"
)


type UserResolver struct {
    UserRepo repository.IUserRepository
}

func (r *UserResolver) GetUserByID(ctx context.Context, id string) (interface{}, error) {
	user, err := r.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user.ToMap(), nil
}

func (r *UserResolver) CreateUser(ctx context.Context, user *models.User) (interface{} , error) {
	user, err := r.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return user.ToMap(), nil
}