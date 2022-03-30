// run it on your machine as a web server instead of lambdas, create orders,feed data with grpc calls
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/api/onprem/handlers"
	"github.com/MauricioAntonioMartinez/mcbot/internal/storage"
)

type config struct {
	port    string
	dbURL   string
	dbName  string
	apikeys struct {
		cryptoCompare string
		coinAPI       string
	}
}

var cfg config

func init() {
	flag.StringVar(&cfg.dbURL, "dbURL", "", "Database URL")
	flag.StringVar(&cfg.port, "port", "8080", "Port to listen")
	flag.StringVar(&cfg.apikeys.cryptoCompare, "cryptoCompAPI", "", "CryptoCompare API Key")
	flag.StringVar(&cfg.apikeys.coinAPI, "coinAPI", "", "CoinAPI API Key")
	flag.StringVar(&cfg.dbName, "dbName", "mcbot", "Database name")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	client, err := storage.ConnectDB(cfg.dbURL)
	if err != nil {
		log.Fatal("Error connecting with the database read the error: ", err)
	}

	defer client.Disconnect(ctx)

	stores := storage.NewStore(storage.Params{
		Connection: client,
		Database:   cfg.dbName,
	})

	s := newSever(handlers.NewLogic(handlers.Params{
		Stores: stores,
	}), cfg.port)

	fmt.Printf("Listenning server on port %s\n", cfg.port)
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func newSever(handlers handlers.Handler, port string) http.Server {
	return http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes(handlers),
	}
}
