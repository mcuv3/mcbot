package stock

import "go.mongodb.org/mongo-driver/mongo"

type Unit struct {
	Reader
	Writer
}

func NewUnit(client *mongo.Database) Unit {
	return Unit{
		Reader: NewReader(client),
		Writer: NewWrite(client),
	}
}
