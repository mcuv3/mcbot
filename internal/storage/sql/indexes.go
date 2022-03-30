package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var indexes = map[string]mongo.IndexModel{
	"stock": {
		Keys: bsonx.Doc{{Key: "symbol", Value: bsonx.String("text")},
			{Key: "exchange", Value: bsonx.String("text")}},
		Options: options.Index().SetUnique(true),
	},
}
