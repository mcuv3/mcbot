// run it on your machine as a web server instead of lambdas, create orders,feed data with grpc calls
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/MauricioAntonioMartinez/mcbot/api/onprem/handlers"
	"github.com/MauricioAntonioMartinez/mcbot/internal/storage"
)

type config struct {
	port string
}

func main() {

	ctx := context.Background()

	client, err := storage.ConnectDB("")
	if err != nil {
		log.Fatal("Error connecting with the database read the error: ", err)
	}

	defer client.Disconnect(ctx)

	stores := storage.NewStore(storage.Params{
		Connection: client,
		Database:   "mcbot",
	})

	s := newSever(handlers.NewLogic(handlers.Params{
		Stores: stores,
	}))

	fmt.Println("Listenning server on port 8080")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func newSever(handlers handlers.Handler) http.Server {
	return http.Server{
		Addr:    ":8080",
		Handler: routes(handlers),
	}
}
