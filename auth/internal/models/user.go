package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	ID       	 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FirstName    string             `bson:"firstName" json:"firstName"`
    LastName     string             `bson:"lastName" json:"lastName"`
    Email   	 string             `bson:"email" json:"email"`
    Password 	 string            	`bson:"password" json:"password"`
}

func (u *User) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "id":        u.ID.Hex(),  // Convert ObjectID to string
        "firstName": u.FirstName,
        "lastName":  u.LastName,
        "email":     u.Email,
    }
}