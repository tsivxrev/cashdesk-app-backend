package models

type Entry struct {
	Id          string  `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name" validate:"required"`
	Description *string `json:"description" bson:"description"`
	Image       *string `json:"image" bson:"image"`
	Price       float64 `json:"price" bson:"price" validate:"required"`
	Hidden      bool    `json:"hidden" bson:"hidden"`

	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
}
