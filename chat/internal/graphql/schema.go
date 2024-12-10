package graphql

import (
	"auth/internal/kafka"
	"auth/internal/resolver"

	"github.com/graphql-go/graphql"
)

func NewSchema(resolver *resolver.Resolver, kafkaProduceer *kafka.Producer) *graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: NewQuery(resolver),
        Mutation: NewMutation(resolver, kafkaProduceer),
	})
	if err != nil {
		return nil
	}
	return &schema
}


