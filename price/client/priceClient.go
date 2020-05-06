package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	pb "gpProto/price"
	"log"
	"time"
)

const address = "localhost:50051"

const dateLayoutISO = "2006-01-02"

func getPrice(priceClient pb.PriceSrvClient, ctx context.Context) {
	_ = ctx
	searchParms := &pb.SearchParams{
		Id:         "price1a",
		ProductId:  "prod1a",
		CustomerId: "cust1",
		Type:       pb.Type_LIST,
	}
	foundPrice, err := priceClient.GetPrice(ctx, searchParms)
	if err != nil {
		log.Fatalf("Unable to find price. Error: %v", err)
	}
	fmt.Printf("Pulled price %v\n", foundPrice)
}

func CreatePrice(priceClient pb.PriceSrvClient, ctx context.Context) {
	_ = ctx

	priceVF := ptypes.TimestampNow()

	priceVTstr := "2021-06-03"
	err, priceVT := timeStringToTimestamp(priceVTstr)

	price := &pb.Price{
		Id:         "price1a",
		ProductId:  "prod1a",
		Value:      2.99,
		ValidFrom:  priceVF,
		ValidThru:  priceVT,
		Type:       pb.Type_GOGS,
		CustomerId: "cust1",
	}

	newPrice, err := priceClient.CreatePrice(ctx, price)
	if err != nil {
		log.Fatal("Unable to create Price")
	}
	fmt.Printf("Received price %v\n", newPrice)
}

func timeStringToTimestamp(priceVTstr string) (error, *timestamp.Timestamp) {
	priceVTtime, err := time.Parse(dateLayoutISO, priceVTstr)
	if err != nil {
		log.Fatalf("Unable to Format date %v", priceVTstr)
	}
	priceVT, err := ptypes.TimestampProto(priceVTtime)
	if err != nil {
		log.Fatalf("Unable to convert time to timestamp %v", priceVTtime)
	}
	return err, priceVT
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Unable to connect to %s. Error: %v", address, err)
	}
	defer conn.Close()
	priceClient := pb.NewPriceSrvClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	CreatePrice(priceClient, ctx)
	getPrice(priceClient, ctx)
}
