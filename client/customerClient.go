package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "gpProto/customer"
	"log"
	"time"
)

const address = "localhost:50051"

func createCustomer(custClient pb.CustomerSrvClient, ctx context.Context) {
	loc := &pb.Location{
		Street:  "123 Street",
		ZipCode: "5555",
		City:    "Life City",
		Country: "Duckland",
	}
	newCust := &pb.Customer{
		Id:   "cust1",
		Name: "Mighty Duck",
		Type: pb.Customer_INDIVIDUAL,
	}
	newCust.Address = append(newCust.Address, loc)

	cust, err := custClient.CreateCustomer(ctx, newCust)
	if err != nil {
		log.Fatal("Unable to create customer")
	}
	fmt.Printf("Received customer %v\n", cust)
}

func GetCusstomer(custClient pb.CustomerSrvClient, ctx context.Context) {
	searchParams := &pb.SearchParams{
		Id:   "cust1",
		Name: "Mighty Duck Jr.",
	}
	cust, err := custClient.GetCustomer(ctx, searchParams)
	if err != nil {
		log.Fatalf("Unable to find cutomer. Error: %v", err)
	}
	fmt.Printf("Pulled customer %v\n", cust)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Unable to connect to %s. Error: %v", address, err)
	}
	defer conn.Close()
	custClient := pb.NewCustomerSrvClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createCustomer(custClient, ctx)
	GetCusstomer(custClient, ctx)

}
