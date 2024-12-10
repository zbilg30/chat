package resolver

import "auth/internal/repository"

type Resolver struct {
    UserResolver    *UserResolver
}

func NewResolver(repositories *repository.RepositoryLayer) *Resolver {
    return &Resolver{
        UserResolver:    &UserResolver{UserRepo: repositories.UserRepository},
    }
}