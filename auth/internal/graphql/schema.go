package graphql

import (
	"auth/internal/resolver"

	"github.com/graphql-go/graphql"
)

func NewSchema(resolver *resolver.Resolver) *graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: NewQuery(resolver),
        Mutation: NewMutation(resolver),
	})
	if err != nil {
		return nil
	}
	return &schema
}


