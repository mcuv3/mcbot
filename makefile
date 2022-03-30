# Makefile
include .env

hello:
	echo "Hello"

seed:
	go build -o dist/seed ./internal/storage/sql
	./dist/seed -connStr=$(MONGO_URI)

backend:
	MONGO_URI=$(MONGO_URI) COIN_API=$(COIN_API) CRYPTO_COMPATE_KEY=$(CRYPTO_COMPATE_KEY) air 
