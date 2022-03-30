package stock

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func (w Writer) Save(trend Stock) error {
	collection := w.client.Collection(emptyStock.CollectionName())
	_, err := collection.InsertOne(context.TODO(), trend)
	return err
}

type DeleteParams struct {
	StockID uuid.UUID
}

func (w Writer) Delete(ctx context.Context, params DeleteParams) error {
	collection := w.client.Collection(emptyStock.CollectionName())
	_, err := collection.DeleteOne(ctx, bson.M{"_id": params.StockID.String()})
	return err
}
