package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MauricioAntonioMartinez/mcbot/storage/stock"
	"github.com/MauricioAntonioMartinez/mcbot/storage/trend"
	"go.mongodb.org/mongo-driver/bson"
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
}

func NewStore(params Params) Store {
	db := params.Connection.Database(params.Database)
	return Store{
		Trend: trend.NewUnit(db),
		Stock: stock.NewUnit(db),
	}
}

func ConnectDB(str string) {
	client, err := mongo.NewClient(options.Client().ApplyURI(str))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

}
