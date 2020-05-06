package main

import (
	"context"
	"google.golang.org/grpc"
	pb "gpProto/price"
	"log"
	"net"
)

const port = ":50051"

type priceServer struct {
	pb.UnimplementedPriceSrvServer
}

func (p *priceServer) GetPrice(ctx context.Context, saerchParms *pb.SearchParams) (*pb.Price, error) {
	_ = ctx
	prod := &pb.Price{
		Id:         saerchParms.Id,
		ProductId:  saerchParms.ProductId,
		CustomerId: saerchParms.CustomerId,
		Type:       saerchParms.Type,
		Value:      1.99,
	}
	return prod, nil
}

func (p *priceServer) CreatePrice(ctx context.Context, price *pb.Price) (*pb.Price, error) {
	_ = ctx
	return price, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot start listening on port %s. Error: %v", port, err)
	}
	s := grpc.NewServer()
	pb.RegisterPriceSrvServer(s, &priceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start product server. Error: %v", err)
	}
}
