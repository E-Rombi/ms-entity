package repository

import (
	"context"
	"log"
	"ms-entity/internal/app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
)

type SaveProductPort interface {
	Save(product model.Product) (*model.Product, error)
}

type FindProductByIdPort interface {
	FindById(id string) (*model.Product, error)
}

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository() *ProductRepository {
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &ProductRepository{
		collection: client.Database("entity").Collection("products"),
	}
}

func (pr ProductRepository) Save(product model.Product) (*model.Product, error) {
	id, err := pr.collection.InsertOne(context.TODO(), product)
	if err != nil {
		return nil, err
	}
	product.ID = id.InsertedID.(primitive.ObjectID)

	return &product, nil
}

func (pr ProductRepository) FindById(id string) (*model.Product, error) {
	product := model.Product{}
	objId, _ := primitive.ObjectIDFromHex(id)
	if err := pr.collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}
