package kline

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type reader struct {
	client *mongo.Database
}

func NewReader(client *mongo.Database) reader {
	return reader{
		client: client,
	}
}

func (r reader) List(ctx context.Context, filters map[string]interface{}) ([]Model, error) {
	collection := r.client.Collection(collectionName)
	var trends []Model
	cur, err := collection.Find(ctx, filters, options.Find().SetSort(bson.D{{Key: "created_at", Value: 1}}))
	if err != nil {
		return trends, err
	}
	for cur.Next(context.TODO()) {
		var trend Model
		err := cur.Decode(&trend)
		if err != nil {
			return trends, err
		}
		trends = append(trends, trend)
	}

	return trends, nil
}

func (r reader) GetLast(ctx context.Context, symbol string) (Model, error) {
	collection := r.client.Collection(collectionName)
	var trend Model

	res := collection.FindOne(ctx, bson.D{{Key: "symbol", Value: symbol}}, options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}}))
	if res.Err() != nil {
		return trend, res.Err()
	}
	if err := res.Decode(&trend); err != nil {
		return trend, err
	}

	return trend, nil
}
