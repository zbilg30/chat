package graphql

import (
	queires "auth/internal/graphql/queries"
	"auth/internal/resolver"

	"github.com/graphql-go/graphql"
)

func NewQuery(resolver *resolver.Resolver) *graphql.Object {
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: mergeFields(
			queires.GetUserQueries(resolver),
		),
	})
	return query;
}

func mergeFields(fieldsList ...graphql.Fields) graphql.Fields {
	merged := graphql.Fields{}
	for _, fields := range fieldsList {
		for k, v := range fields {
			merged[k] = v
		}
	}
	return merged
}