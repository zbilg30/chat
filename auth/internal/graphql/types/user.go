package types

import (
	"auth/internal/models"

	"github.com/graphql-go/graphql"
)


var UserType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.User:
                    return source.ID.Hex(), nil
                case map[string]interface{}:
                    return source["id"], nil
                default:
                    return nil, nil
                }
            },
        },
        "firstName": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.User:
                    return source.FirstName, nil
                case map[string]interface{}:
                    return source["firstName"], nil
                default:
                    return nil, nil
                }
            },
        },
        "lastName": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.User:
                    return source.LastName, nil
                case map[string]interface{}:
                    return source["lastName"], nil
                default:
                    return nil, nil
                }
            },
        },
        "email": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.User:
                    return source.Email, nil
                case map[string]interface{}:
                    return source["email"], nil
                default:
                    return nil, nil
                }
            },
        },
        "password": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.User:
                    return source.Password, nil
                case map[string]interface{}:
                    return source["password"], nil
                default:
                    return nil, nil
                }
            },
        },
    },
})


var CreateUserInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateUserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"firstName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"lastName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String, 
		},
		"password": &graphql.InputObjectFieldConfig{
			Type: graphql.String, 
		},
	},
})