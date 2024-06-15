package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID           primitive.ObjectID `bson:"_id"`
	Title        *string            `json:"title" validate:"required,min=2,max=100"`
	Author       *string            `json:"author" validate:"required,min=2,max=100"`
	ISBN         *string            `json:"isbn" validate:"required"`
	Publisher    *string            `json:"publisher" validate:"required,min=2,max=100"`
	Publish_Date time.Time          `json:"publish_date"`
}
