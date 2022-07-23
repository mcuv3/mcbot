package kline

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type writer struct {
	client *mongo.Database
}

func NewWriter(client *mongo.Database) writer {
	return writer{
		client: client,
	}
}

func (w writer) Save(trend Model) error {
	trend.CreatedAt = time.Now().Unix()
	collection := w.client.Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), trend)
	return err
}
