package storage

import (
	"context"
	"time"

	"github.com/mcuv3/mcbot/internal/storage/kline"
	"github.com/mcuv3/mcbot/internal/storage/stock"
	"github.com/mcuv3/mcbot/internal/storage/trend"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Params struct {
	Connection *mongo.Client
	Database   string
}

type Store struct {
	Trend trend.Unit
	Stock stock.Unit
	Kline kline.Unit
}

func NewStore(params Params) Store {
	db := params.Connection.Database(params.Database)

	return Store{
		Trend: trend.NewUnit(db),
		Stock: stock.NewUnit(db),
		Kline: kline.NewUnit(db),
	}
}

func ConnectDB(str string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(str))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
