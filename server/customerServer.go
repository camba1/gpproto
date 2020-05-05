package main

import (
	"context"
	"google.golang.org/grpc"
	pb "gpProto/customer"
	"log"
	"net"
)

const port = ":50051"

type custServer struct {
	pb.UnimplementedCustomerSrvServer
}

func (s *custServer) GetCustomer(ctx context.Context, saerchParms *pb.SearchParams) (*pb.Customer, error) {
	_ = ctx
	cust := pb.Customer{
		Id:      saerchParms.Id,
		Name:    saerchParms.Name,
		Address: nil,
		Type:    pb.Customer_BUSINESS,
	}
	return &cust, nil
}

func (s *custServer) CreateCustomer(ctx context.Context, customer *pb.Customer) (*pb.Customer, error) {
	_ = ctx
	return customer, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot start listening on port %s. Error: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterCustomerSrvServer(s, &custServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start server. Error: %v", err)
	}
}
