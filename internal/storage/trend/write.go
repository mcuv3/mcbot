package trend

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Writer struct {
	client *mongo.Database
}

func NewWriter(client *mongo.Database) Writer {
	return Writer{
		client: client,
	}
}

func (w Writer) Save(trend Trend) error {
	collection := w.client.Collection("trends")
	_, err := collection.InsertOne(context.TODO(), trend)
	return err
}
