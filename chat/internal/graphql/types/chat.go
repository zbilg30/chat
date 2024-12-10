package types

import (
	"auth/internal/models"

	"github.com/graphql-go/graphql"
)


var ChatType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Chat",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.Chat:
                    return source.ID.Hex(), nil
                case map[string]interface{}:
                    return source["id"], nil
                default:
                    return nil, nil
                }
            },
        },
        "message": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.Chat:
                    return source.Message, nil
                case map[string]interface{}:
                    return source["message"], nil
                default:
                    return nil, nil
                }
            },
        },
        "userId": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.Chat:
                    return source.UserId, nil
                case map[string]interface{}:
                    return source["userId"], nil
                default:
                    return nil, nil
                }
            },
        },
        "chatRoomId": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                switch source := p.Source.(type) {
                case models.Chat:
                    return source.ChatRoomId, nil
                case map[string]interface{}:
                    return source["chatRoomId"], nil
                default:
                    return nil, nil
                }
            },
        },
    },
})


var CreateChatInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateChatInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"chatRoomId": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"userId": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"message": &graphql.InputObjectFieldConfig{
			Type: graphql.String, 
		},
		
	},
})

