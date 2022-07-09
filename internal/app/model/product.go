package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description"`
	Coast       float64            `bson:"coast"`
	Price       float64            `bson:"price"`
}

func NewProduct(product *NewProductRequest) *Product {
	return &Product{
		Description: product.Description,
		Coast:       product.Coast,
		Price:       product.Price,
	}
}

type NewProductRequest struct {
	Description string  `json:"description"`
	Coast       float64 `json:"coast"`
	Price       float64 `json:"price"`
}
