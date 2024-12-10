package graphql

import (
	"auth/internal/graphql/types"
	"auth/internal/kafka"
	"auth/internal/models"
	"auth/internal/resolver"
	"context"

	"github.com/graphql-go/graphql"
)

func NewMutation(resolver *resolver.Resolver, kafkaProducer *kafka.Producer) *graphql.Object {
	RootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createChat": &graphql.Field{
				Type: types.ChatType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(types.CreateChatInput),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					user := &models.Chat{
						UserId:  input["userId"].(string),
						Message:  input["message"].(string),
						ChatRoomId: input["chatRoomId"].(string),
					}

					return resolver.ChatResolver.CreateChat(context.Background(), user, kafkaProducer);
				},
			},
			"createChatRoom": &graphql.Field{
				Type: types.ChatRoomType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(types.CreateChatRoomInput),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					interfaceIds := input["userIds"].([]interface{})

					userIds := make([]string, len(interfaceIds))
					for i, v := range interfaceIds {
						userIds[i] = v.(string)
					}

					chat := &models.ChatRoom{
						UserIds:  userIds,
						Name:  input["name"].(string),
					}

					return resolver.ChatRoomResolver.CreateChatRoom(context.Background(), chat);
				},
			},
		},
	})

	return RootMutation
}

	