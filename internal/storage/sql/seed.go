package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/MauricioAntonioMartinez/mcbot/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type config struct {
	connStr string
	dbName  string
}

type seeder struct {
	client *mongo.Client
	stores storage.Store
	dbName string
}

var cfg config

func init() {
	flag.StringVar(&cfg.connStr, "connStr", "", "Connection string to connect to the database")
	flag.StringVar(&cfg.dbName, "dbName", "mcbot", "Database name")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()
	ctx := context.Background()

	client, err := storage.ConnectDB(cfg.connStr)
	if err != nil {
		log.Fatal("ERROR CONNECTING TO DB", err)
	}

	stores := storage.NewStore(storage.Params{
		Connection: client,
		Database:   cfg.dbName,
	})

	sd := seeder{client, stores, cfg.dbName}
	fmt.Println("Seeding database ....")
	if err := sd.seed(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding database finished")
}

func (s seeder) seed(ctx context.Context) error {
	return s.createIndexes(ctx)
}

func (s seeder) createIndexes(ctx context.Context) error {
	for coll, idxs := range indexes {
		if err := s.addIndex(ctx, coll, idxs); err != nil {
			return err
		}
	}
	return nil
}

func (s seeder) addIndex(ctx context.Context, collection string, idx mongo.IndexModel) error {
	serviceCollection := s.client.Database(s.dbName).Collection(collection)
	indexName, err := serviceCollection.Indexes().CreateOne(ctx, idx)
	if err != nil {
		return err
	}
	fmt.Printf("Index created: %s\n", indexName)
	return nil
}
