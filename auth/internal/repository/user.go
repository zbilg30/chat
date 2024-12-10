package repository

import (
	"auth/internal/models"
	"auth/internal/pkg"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
    mongoClient pkg.IMongoClient 
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

func NewUserRepository(db pkg.IMongoClient) IUserRepository {
    return &UserRepository{
        mongoClient: db,
    }
}

func(c *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = primitive.NewObjectID()
    _, err := c.mongoClient.GetCollection("users").InsertOne(ctx, user)

    if err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid ID format")
    }

    var user models.User
    err = r.mongoClient.GetCollection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}