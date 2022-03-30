package stock

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Reader struct {
	client *mongo.Database
}

func NewReader(client *mongo.Database) Reader {
	return Reader{
		client: client,
	}
}

type ListParams struct{}

func (r Reader) Find(ctx context.Context, params ListParams) ([]Stock, error) {
	collection := r.client.Collection(emptyStock.CollectionName())
	var trends []Stock
	cur, err := collection.Find(ctx, params)
	if err != nil {
		return trends, err
	}
	for cur.Next(context.TODO()) {
		var trend Stock
		err := cur.Decode(&trend)
		if err != nil {
			return trends, err
		}
		trends = append(trends, trend)
	}

	return trends, nil
}
