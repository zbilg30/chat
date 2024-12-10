package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID       	 primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
    Message      string               `bson:"message" json:"message"`
    UserId       string               `bson:"userId" json:"userId"`
    ChatRoomId   string   `bson:"chatRoomId,omitempty" json:"chatRoomId"`
}

func (u *Chat) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":        u.ID.Hex(),  // Convert ObjectID to string
        "message":   u.Message,
        "userId":    u.UserId,
        "chatRoomId":     u.ChatRoomId,
    }
}