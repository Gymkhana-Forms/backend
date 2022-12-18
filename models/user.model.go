package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          *string            `json:"Name" validate:"required,min=2,max=100"`
	RollNo        *string            `json:"RollNo" validate:"required"`
	Email         *string            `json:"Email" validate:"email,required"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Token         *string            `json:"Token"`
	User_type     *string            `json:"User_type" validate:"required,eq=ADMIN|eq=STUDENT"`
	Refresh_token *string            `json:"Refresh_token"`
}
