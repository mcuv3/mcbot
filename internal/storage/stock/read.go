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

type FindOneParams struct {
	Symbol   string `bson:"symbol"`
	Exchange string `bson:"exchange"`
}

func (r Reader) FindOne(ctx context.Context, params FindOneParams) (Stock, error) {
	collection := r.client.Collection(emptyStock.CollectionName())
	var trend Stock
	res := collection.FindOne(ctx, params)

	err := res.Decode(&trend)
	if err != nil {
		return trend, err
	}

	return trend, nil
}
