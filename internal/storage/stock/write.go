package stock

import (
	"context"
	"fmt"

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
	res, err := collection.InsertOne(context.Background(), trend)
	fmt.Println(res)
	return err
}

func (w Writer) SaveMany(trend interface{}) error {
	collection := w.client.Collection(emptyStock.CollectionName())

	_, err := collection.InsertMany(context.Background(), nil)

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
