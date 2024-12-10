package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ChatRoom struct {
	ID       	 primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
    UserIds      []string             `bson:"userIds" json:"userIds"`
    Name         string  `bson:"name" json:"name"`
}

func (u *ChatRoom) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":        u.ID.Hex(),  // Convert ObjectID to string
        "name":     u.Name,
    }
}