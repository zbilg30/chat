package graphql

import (
	"auth/internal/graphql/types"
	"auth/internal/models"
	"auth/internal/resolver"
	"context"

	"github.com/graphql-go/graphql"
)

func NewMutation(resolver *resolver.Resolver) *graphql.Object {
	RootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: types.UserType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(types.CreateUserInput),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					user := &models.User{
						FirstName:  input["firstName"].(string),
						LastName:  input["lastName"].(string),
						Email: input["email"].(string),
						Password: input["password"].(string),
					}

					return resolver.UserResolver.CreateUser(context.Background(), user);
				},
			},
		},
	})

	return RootMutation
}

	