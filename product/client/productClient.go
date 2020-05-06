package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "gpProto/product"
	"log"
	"time"
)

const address = "localhost:50051"

func getProduct(prodClient pb.ProductSrvClient, ctx context.Context) {
	_ = ctx
	searchParms := &pb.SearchParams{
		Id:   "prod1a",
		Name: "Awesome Cream",
	}
	foundProd, err := prodClient.GetProduct(ctx, searchParms)
	if err != nil {
		log.Fatalf("Unable to find product. Error: %v", err)
	}
	fmt.Printf("Pulled product %v\n", foundProd)
}

func CreateProduct(prodClient pb.ProductSrvClient, ctx context.Context) {
	_ = ctx
	prod := &pb.Product{
		Id:          "prod1a",
		Name:        "Awesome Cream",
		Description: "Really awesome cream",
	}
	prod.Dimensions = append(prod.Dimensions, &pb.Dimension{
		Height: 10.2,
		Length: 5.3,
		Width:  2.1,
		Type:   pb.Dimension_PRODUCT,
	},
	)
	prod.Dimensions = append(prod.Dimensions, &pb.Dimension{
		Height: 12.2,
		Length: 6.3,
		Width:  3.1,
		Type:   pb.Dimension_SHIPPING,
	},
	)
	newProd, err := prodClient.CreateProduct(ctx, prod)
	if err != nil {
		log.Fatal("Unable to create Product")
	}
	fmt.Printf("Received product %v\n", newProd)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Unable to connect to %s. Error: %v", address, err)
	}
	defer conn.Close()
	prodClient := pb.NewProductSrvClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	CreateProduct(prodClient, ctx)
	getProduct(prodClient, ctx)
}
