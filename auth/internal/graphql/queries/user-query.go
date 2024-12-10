package queires

import (
	"auth/internal/graphql/types"
	"auth/internal/resolver"

	"errors"

	"github.com/graphql-go/graphql"
)

func GetUserQueries(resolver *resolver.Resolver) graphql.Fields {
	return graphql.Fields{
		"getUserById": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, ok := params.Args["id"].(string)
				if !ok {
					return nil, errors.New("invalid ID argument")
				}
				return resolver.UserResolver.GetUserByID(params.Context, id)
			},
		},
	}
}