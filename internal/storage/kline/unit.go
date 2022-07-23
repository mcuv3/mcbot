package kline

import "go.mongodb.org/mongo-driver/mongo"

type Unit struct {
	reader
	writer
}

func NewUnit(client *mongo.Database) Unit {
	return Unit{
		reader: NewReader(client),
		writer: NewWriter(client),
	}
}
