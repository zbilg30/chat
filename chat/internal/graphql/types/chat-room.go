package types

import (
	"auth/internal/models"

	"github.com/graphql-go/graphql"
)

var ChatRoomType = graphql.NewObject(graphql.ObjectConfig{
    Name: "ChatRoom",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.ChatRoom:
                    return source.ID.Hex(), nil
                case map[string]interface{}:
                    return source["id"], nil
                default:
                    return nil, nil
                }
            },
        },
        "userIds": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.ChatRoom:
                    return source.UserIds, nil
                case map[string]interface{}:
                    return source["userIds"], nil
                default:
                    return nil, nil
                }
            },
        },
      
    },
})

var CreateChatRoomInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateChatRoomInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"userIds": &graphql.InputObjectFieldConfig{
            Type: graphql.NewList(graphql.String),  // Changed from graphql.String to NewList(graphql.String)
		},
        "name": &graphql.InputObjectFieldConfig{
            Type: graphql.String,
        },
	},
})