package queires

import (
	"auth/internal/graphql/types"
	"auth/internal/resolver"

	"github.com/graphql-go/graphql"
)

func ChatQueries(resolver *resolver.Resolver) graphql.Fields {
	return graphql.Fields{
		"getChats": &graphql.Field{
			Type: types.ChatType,
			Args: graphql.FieldConfigArgument{
				"chatRoomId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
		},
	}
}