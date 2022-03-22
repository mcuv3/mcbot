package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/MauricioAntonioMartinez/mcbot/internal/collectors/lykke/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("cert.crt", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := grpc.Dial("hft-apiv2-grpc.lykke.com:443", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal("CONNECTION FAILED", err)
	}

	svc := service.NewPublicServiceClient(conn)
	defer conn.Close()

	streamPrices(svc)

}

func streamPrices(client service.PublicServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := client.GetPriceUpdates(ctx, &service.PriceUpdatesRequest{
		AssetPairIds: []string{"BTCUSD"},
	})

	if err != nil {
		log.Fatal("ERROR GET ASSETS: ", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("EOF")
			break
		}

		if err != nil {
			fmt.Println("ERROR GET ASSETS: ", err)
			streamPrices(client)
		}

		fmt.Printf("%s BID: %s ASK: %s\n", res.GetAssetPairId(), res.GetBid(), res.GetAsk())
	}

}
