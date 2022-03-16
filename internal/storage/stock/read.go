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

func (r Reader) Find(ctx context.Context, filters map[string]interface{}) ([]Stock, error) {
	collection := r.client.Collection(emptyStock.Table())
	var trends []Stock
	cur, err := collection.Find(ctx, filters)
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
